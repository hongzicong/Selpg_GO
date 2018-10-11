package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "pflag"
)

const INBUFSIZE = 16 * 1024
const INT_MAX = int(^uint(0) >> 1)
const BUFSIZ = 1024

/*================================= globals =======================*/

var progname string

/*================================= types =========================*/

type sp_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int
	page_type   bool
	print_dest  string
}

/*================================= main()=== =====================*/

func main() {
	progname = os.Args[0]

	sa := sp_args{-1, -1, "", 72, false, ""}

	sa.process_args()

	sa.process_input()
}

/*================================= process_args() ================*/

func (sa *sp_args) process_args() {

	startpage := flag.IntP("start_page", "s", 0, "The page should start printing")

	endpage := flag.IntP("end_page", "e", 0, "The page to end printing at [Necessary, no less than startpage]")

	pagelen := flag.IntP("page_len", "l", 72, "Default value is 72, can be overriden by \"-l number\" on command line ")

	pagetype := flag.BoolP("page_type", "f", false, "We set false for lines-delimited, and true for form-feed-delimited")

	printdest := flag.StringP("print_destination", "d", "", "Choose a printer to accept the result as a task")

	flag.Parse()

	sa.start_page = *startpage
	sa.end_page = *endpage
	sa.page_len = *pagelen
	sa.page_type = *pagetype
	sa.print_dest = *printdest

	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "%s: too much arguments\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if sa.start_page == 0 {
		fmt.Fprintf(os.Stderr, "%s: 1st arg should be given\n", progname)
		flag.Usage()
		os.Exit(2)
	}

	if sa.start_page < 1 || sa.start_page > INT_MAX {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n", progname, sa.start_page)
		flag.Usage()
		os.Exit(3)

	}

	if sa.end_page == 0 {
		fmt.Fprintf(os.Stderr, "%s: 2nd arg should be given\n", progname)
		flag.Usage()
		os.Exit(4)
	}

	if sa.end_page < 1 || sa.end_page > INT_MAX || sa.end_page < sa.start_page {
		fmt.Fprintf(os.Stderr, "%s: invalid end page %s\n", progname, sa.end_page)
		flag.Usage()
		os.Exit(5)
	}

	if sa.page_len != 72 && sa.page_type {
		fmt.Fprintf(os.Stderr, "%s: Line number and force paging should not be set at the same time!", progname)
		flag.Usage()
		os.Exit(6)
	}

	if sa.page_len < 1 || sa.page_len > INT_MAX-1 {
		fmt.Fprintf(os.Stderr, "%s: invalid page length %s\n", progname, sa.page_len)
		flag.Usage()
		os.Exit(7)
	}

	if flag.NArg() != 0 {
		sa.in_filename = flag.Args()[0]
		/* check if file exists */
		_, err := os.Stat(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n", progname, sa.in_filename)
			os.Exit(8)
		}
	}
}

/*================================= process_input() ===============*/

func (sa sp_args) process_input() {

	var fin *os.File  /* input stream */
	var fout *os.File /* output stream */
	var line_ctr int  /* line counter */
	var page_ctr int  /* page counter */
	var cmd *exec.Cmd

	/* set the input source */
	if len(sa.in_filename) == 0 {
		fin = os.Stdin
	} else {
		_, err := os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", progname, sa.in_filename)
			os.Exit(12)
		}
	}

	/* set the output destination */
	if len(sa.print_dest) == 0 {
		fout = os.Stdout
	} else {
		cmd = exec.Command("lp", "-d"+sa.print_dest)
		_, err := cmd.StdinPipe()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open pipe to \"%s\"\n", progname, "lp"+"-d"+sa.print_dest)
			os.Exit(13)
		}
	}
	/* begin one of two main loops based on page type */
	if sa.page_type {
		line_ctr = 0
		page_ctr = 1
		line := make([]byte, BUFSIZ)
		for true {
			_, err := fin.Read(line)
			/* error or EOF */
			if err != nil || err == io.EOF {
				break
			}
			line_ctr++
			if line_ctr > sa.page_len {
				page_ctr++
				line_ctr = 1
			}
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
				fout.Write(line)
			}
		}
	} else {
		page_ctr = 1
		c := make([]byte, 1)

		fmt.Fprintf(os.Stderr, "s\n")
		for true {
			_, err := fin.Read(c)
			if err == io.EOF {
				break
			}
			if c[0] == '\f' {
				page_ctr++
			}
			if page_ctr >= sa.start_page && page_ctr <= sa.end_page {
				fout.Write(c)
			}
		}
	}

	/* end main loop */

	if page_ctr < sa.start_page {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", progname, sa.start_page, page_ctr)
	} else if page_ctr < sa.end_page {
		fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, sa.end_page, page_ctr)
	}

	fin.Close()
	if len(sa.print_dest) != 0 {
		cmd.CombinedOutput()
	}
	fmt.Fprintf(os.Stderr, "%s: done\n", progname)

}
