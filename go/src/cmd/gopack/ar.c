<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gopack/ar.c</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gopack/ar.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/iar/ar.c
// http://code.google.com/p/inferno-os/source/browse/utils/iar/ar.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

/*
 * ar - portable (ascii) format version
 */

/* protect a couple of our names */
#define select your_select
#define rcmd your_rcmd

#include &lt;u.h&gt;
#include &lt;time.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;
#include &lt;ar.h&gt;

#undef select
#undef rcmd

/*
 *	The algorithm uses up to 3 temp files.  The &#34;pivot member&#34; is the
 *	archive member specified by and a, b, or i option.  The temp files are
 *	astart - contains existing members up to and including the pivot member.
 *	amiddle - contains new files moved or inserted behind the pivot.
 *	aend - contains the existing members that follow the pivot member.
 *	When all members have been processed, function &#39;install&#39; streams the
 * 	temp files, in order, back into the archive.
 */

typedef struct	Arsymref
{
	char	*name;
	char *file;
	int	type;
	int	len;
	vlong	offset;
	struct	Arsymref *next;
} Arsymref;

typedef struct	Armember	/* Temp file entry - one per archive member */
{
	struct Armember	*next;
	struct ar_hdr	hdr;
	long		size;
	long		date;
	void		*member;
} Armember;

typedef	struct Arfile		/* Temp file control block - one per tempfile */
{
	int	paged;		/* set when some data paged to disk */
	char	*fname;		/* paging file name */
	int	fd;		/* paging file descriptor */
	vlong	size;
	Armember *head;		/* head of member chain */
	Armember *tail;		/* tail of member chain */
	Arsymref *sym;		/* head of defined symbol chain */
} Arfile;

typedef struct Hashchain
{
	char	*name;
	char *file;
	struct Hashchain *next;
} Hashchain;

#define	NHASH	1024

/*
 *	macro to portably read/write archive header.
 *	&#39;cmd&#39; is read/write/Bread/Bwrite, etc.
 */
#define	HEADER_IO(cmd, f, h)	cmd(f, h.name, sizeof(h.name)) != sizeof(h.name)\
				|| cmd(f, h.date, sizeof(h.date)) != sizeof(h.date)\
				|| cmd(f, h.uid, sizeof(h.uid)) != sizeof(h.uid)\
				|| cmd(f, h.gid, sizeof(h.gid)) != sizeof(h.gid)\
				|| cmd(f, h.mode, sizeof(h.mode)) != sizeof(h.mode)\
				|| cmd(f, h.size, sizeof(h.size)) != sizeof(h.size)\
				|| cmd(f, h.fmag, sizeof(h.fmag)) != sizeof(h.fmag)

		/* constants and flags */
char	*man =		&#34;mrxtdpq&#34;;
char	*opt =		&#34;uvnbailo&#34;;
char	artemp[] =	&#34;/tmp/vXXXXX&#34;;
char	movtemp[] =	&#34;/tmp/v1XXXXX&#34;;
char	tailtemp[] =	&#34;/tmp/v2XXXXX&#34;;
char	symdef[] =	&#34;__.SYMDEF&#34;;
char	pkgdef[] =	&#34;__.PKGDEF&#34;;

int	aflag;				/* command line flags */
int	bflag;
int	cflag;
int	gflag;
int	oflag;
int	uflag;
int	vflag;

int	errors;

Arfile *astart, *amiddle, *aend;	/* Temp file control block pointers */
int	allobj = 1;			/* set when all members are object files of the same type */
int	symdefsize;			/* size of symdef file */
char	*pkgstmt;		/* string &#34;package foo&#34; */
int	dupfound;			/* flag for duplicate symbol */
Hashchain	*hash[NHASH];		/* hash table of text symbols */

#define	ARNAMESIZE	sizeof(astart-&gt;tail-&gt;hdr.name)

char	poname[ARNAMESIZE+1];		/* name of pivot member */
char	*file;				/* current file or member being worked on */
Biobuf	bout;
Biobuf bar;

void	arcopy(Biobuf*, Arfile*, Armember*);
int	arcreate(char*);
void	arfree(Arfile*);
void	arinsert(Arfile*, Armember*);
void	*armalloc(int);
char *arstrdup(char*);
void	armove(Biobuf*, Arfile*, Armember*);
void	arread(Biobuf*, Armember*, int);
void	arstream(int, Arfile*);
int	arwrite(int, Armember*);
int	bamatch(char*, char*);
int	duplicate(char*, char**);
Armember *getdir(Biobuf*);
void	getpkgdef(char**, int*);
int	getspace(void);
void	install(char*, Arfile*, Arfile*, Arfile*, int);
void	loadpkgdata(char*, int);
void	longt(Armember*);
int	match(int, char**);
void	mesg(int, char*);
Arfile	*newtempfile(char*);
Armember *newmember(void);
void	objsym(Sym*, void*);
int	openar(char*, int, int);
int	page(Arfile*);
void	pmode(long);
void	rl(int);
void	scanobj(Biobuf*, Arfile*, long);
void	scanpkg(Biobuf*, long);
void	select(int*, long);
void	setcom(void(*)(char*, int, char**));
void	skip(Biobuf*, vlong);
int	symcomp(void*, void*);
void	trim(char*, char*, int);
void	usage(void);
void	wrerr(void);
void	wrsym(Biobuf*, long, Arsymref*);

void	rcmd(char*, int, char**);		/* command processing */
void	dcmd(char*, int, char**);
void	xcmd(char*, int, char**);
void	tcmd(char*, int, char**);
void	pcmd(char*, int, char**);
void	mcmd(char*, int, char**);
void	qcmd(char*, int, char**);
void	(*comfun)(char*, int, char**);

void
main(int argc, char *argv[])
{
	char *cp;

	Binit(&amp;bout, 1, OWRITE);
	if(argc &lt; 3)
		usage();
	for (cp = argv[1]; *cp; cp++) {
		switch(*cp) {
		case &#39;a&#39;:	aflag = 1;	break;
		case &#39;b&#39;:	bflag = 1;	break;
		case &#39;c&#39;:	cflag = 1;	break;
		case &#39;d&#39;:	setcom(dcmd);	break;
		case &#39;g&#39;:	gflag = 1; break;
		case &#39;i&#39;:	bflag = 1;	break;
		case &#39;l&#39;:
				strcpy(artemp, &#34;vXXXXX&#34;);
				strcpy(movtemp, &#34;v1XXXXX&#34;);
				strcpy(tailtemp, &#34;v2XXXXX&#34;);
				break;
		case &#39;m&#39;:	setcom(mcmd);	break;
		case &#39;o&#39;:	oflag = 1;	break;
		case &#39;p&#39;:	setcom(pcmd);	break;
		case &#39;q&#39;:	setcom(qcmd);	break;
		case &#39;r&#39;:	setcom(rcmd);	break;
		case &#39;t&#39;:	setcom(tcmd);	break;
		case &#39;u&#39;:	uflag = 1;	break;
		case &#39;v&#39;:	vflag = 1;	break;
		case &#39;x&#39;:	setcom(xcmd);	break;
		default:
			fprint(2, &#34;gopack: bad option `%c&#39;\n&#34;, *cp);
			exits(&#34;error&#34;);
		}
	}
	if (aflag &amp;&amp; bflag) {
		fprint(2, &#34;gopack: only one of &#39;a&#39; and &#39;b&#39; can be specified\n&#34;);
		usage();
	}
	if(aflag || bflag) {
		trim(argv[2], poname, sizeof(poname));
		argv++;
		argc--;
		if(argc &lt; 3)
			usage();
	}
	if(comfun == 0) {
		if(uflag == 0) {
			fprint(2, &#34;gopack: one of [%s] must be specified\n&#34;, man);
			usage();
		}
		setcom(rcmd);
	}
	cp = argv[2];
	argc -= 3;
	argv += 3;
	(*comfun)(cp, argc, argv);	/* do the command */
	cp = 0;
	while (argc--) {
		if (*argv) {
			fprint(2, &#34;gopack: %s not found\n&#34;, *argv);
			cp = &#34;error&#34;;
		}
		argv++;
	}
	if (errors)
		cp = &#34;error&#34;;
	exits(cp);
}
/*
 *	select a command
 */
void
setcom(void (*fun)(char *, int, char**))
{

	if(comfun != 0) {
		fprint(2, &#34;gopack: only one of [%s] allowed\n&#34;, man);
		usage();
	}
	comfun = fun;
}
/*
 *	perform the &#39;r&#39; and &#39;u&#39; commands
 */
void
rcmd(char *arname, int count, char **files)
{
	int fd;
	int i;
	Arfile *ap;
	Armember *bp;
	Dir *d;
	Biobuf *bfile;

	fd = openar(arname, ORDWR, 1);
	if (fd &gt;= 0) {
		Binit(&amp;bar, fd, OREAD);
		Bseek(&amp;bar,seek(fd,0,1), 1);
	}
	astart = newtempfile(artemp);
	ap = astart;
	aend = 0;
	for(i = 0; fd &gt;= 0; i++) {
		bp = getdir(&amp;bar);
		if (!bp)
			break;
		if (bamatch(file, poname)) {		/* check for pivot */
			aend = newtempfile(tailtemp);
			ap = aend;
		}
			/* pitch symdef file */
		if (i == 0 &amp;&amp; strcmp(file, symdef) == 0) {
			skip(&amp;bar, bp-&gt;size);
			continue;
		}
			/* pitch pkgdef file */
		if (gflag &amp;&amp; strcmp(file, pkgdef) == 0) {
			skip(&amp;bar, bp-&gt;size);
			continue;
		}
		if (count &amp;&amp; !match(count, files)) {
			scanobj(&amp;bar, ap, bp-&gt;size);
			arcopy(&amp;bar, ap, bp);
			continue;
		}
		bfile = Bopen(file, OREAD);
		if (!bfile) {
			if (count != 0) {
				fprint(2, &#34;gopack: cannot open %s\n&#34;, file);
				errors++;
			}
			scanobj(&amp;bar, ap, bp-&gt;size);
			arcopy(&amp;bar, ap, bp);
			continue;
		}
		d = dirfstat(Bfildes(bfile));
		if(d == nil)
			fprint(2, &#34;gopack: cannot stat %s: %r\n&#34;, file);
		if (uflag &amp;&amp; (d==nil || d-&gt;mtime &lt;= bp-&gt;date)) {
			scanobj(&amp;bar, ap, bp-&gt;size);
			arcopy(&amp;bar, ap, bp);
			Bterm(bfile);
			free(d);
			continue;
		}
		mesg(&#39;r&#39;, file);
		skip(&amp;bar, bp-&gt;size);
		scanobj(bfile, ap, d-&gt;length);
		free(d);
		armove(bfile, ap, bp);
		Bterm(bfile);
	}
	if(fd &gt;= 0)
		close(fd);
		/* copy in remaining files named on command line */
	for (i = 0; i &lt; count; i++) {
		file = files[i];
		if(file == 0)
			continue;
		files[i] = 0;
		bfile = Bopen(file, OREAD);
		if (!bfile) {
			fprint(2, &#34;gopack: cannot open %s\n&#34;, file);
			errors++;
		} else {
			mesg(&#39;a&#39;, file);
			d = dirfstat(Bfildes(bfile));
			if (d == nil)
				fprint(2, &#34;can&#39;t stat %s\n&#34;, file);
			else {
				scanobj(bfile, astart, d-&gt;length);
				armove(bfile, astart, newmember());
				free(d);
			}
			Bterm(bfile);
		}
	}
	if(fd &lt; 0 &amp;&amp; !cflag)
		install(arname, astart, 0, aend, 1);	/* issue &#39;creating&#39; msg */
	else
		install(arname, astart, 0, aend, 0);
}

void
dcmd(char *arname, int count, char **files)
{
	Armember *bp;
	int fd, i;

	if (!count)
		return;
	fd = openar(arname, ORDWR, 0);
	Binit(&amp;bar, fd, OREAD);
	Bseek(&amp;bar,seek(fd,0,1), 1);
	astart = newtempfile(artemp);
	for (i = 0; bp = getdir(&amp;bar); i++) {
		if(match(count, files)) {
			mesg(&#39;d&#39;, file);
			skip(&amp;bar, bp-&gt;size);
			if (strcmp(file, symdef) == 0)
				allobj = 0;
		} else if (i == 0 &amp;&amp; strcmp(file, symdef) == 0) {
			skip(&amp;bar, bp-&gt;size);
		} else if (gflag &amp;&amp; strcmp(file, pkgdef) == 0) {
			skip(&amp;bar, bp-&gt;size);
		} else {
			scanobj(&amp;bar, astart, bp-&gt;size);
			arcopy(&amp;bar, astart, bp);
		}
	}
	close(fd);
	install(arname, astart, 0, 0, 0);
}

void
xcmd(char *arname, int count, char **files)
{
	int fd, f, mode, i;
	Armember *bp;
	Dir dx;

	fd = openar(arname, OREAD, 0);
	Binit(&amp;bar, fd, OREAD);
	Bseek(&amp;bar,seek(fd,0,1), 1);
	i = 0;
	while (bp = getdir(&amp;bar)) {
		if(count == 0 || match(count, files)) {
			mode = strtoul(bp-&gt;hdr.mode, 0, 8) &amp; 0777;
			f = create(file, OWRITE, mode);
			if(f &lt; 0) {
				fprint(2, &#34;gopack: %s cannot create\n&#34;, file);
				skip(&amp;bar, bp-&gt;size);
			} else {
				mesg(&#39;x&#39;, file);
				arcopy(&amp;bar, 0, bp);
				if (write(f, bp-&gt;member, bp-&gt;size) &lt; 0)
					wrerr();
				if(oflag) {
					nulldir(&amp;dx);
					dx.atime = bp-&gt;date;
					dx.mtime = bp-&gt;date;
					if(dirwstat(file, &amp;dx) &lt; 0)
						perror(file);
				}
				free(bp-&gt;member);
				close(f);
			}
			free(bp);
			if (count &amp;&amp; ++i &gt;= count)
				break;
		} else {
			skip(&amp;bar, bp-&gt;size);
			free(bp);
		}
	}
	close(fd);
}
void
pcmd(char *arname, int count, char **files)
{
	int fd;
	Armember *bp;

	fd = openar(arname, OREAD, 0);
	Binit(&amp;bar, fd, OREAD);
	Bseek(&amp;bar,seek(fd,0,1), 1);
	while(bp = getdir(&amp;bar)) {
		if(count == 0 || match(count, files)) {
			if(vflag)
				print(&#34;\n&lt;%s&gt;\n\n&#34;, file);
			arcopy(&amp;bar, 0, bp);
			if (write(1, bp-&gt;member, bp-&gt;size) &lt; 0)
				wrerr();
		} else
			skip(&amp;bar, bp-&gt;size);
		free(bp);
	}
	close(fd);
}
void
mcmd(char *arname, int count, char **files)
{
	int fd, i;
	Arfile *ap;
	Armember *bp;

	if (count == 0)
		return;
	fd = openar(arname, ORDWR, 0);
	Binit(&amp;bar, fd, OREAD);
	Bseek(&amp;bar,seek(fd,0,1), 1);
	astart = newtempfile(artemp);
	amiddle = newtempfile(movtemp);
	aend = 0;
	ap = astart;
	for (i = 0; bp = getdir(&amp;bar); i++) {
		if (bamatch(file, poname)) {
			aend = newtempfile(tailtemp);
			ap = aend;
		}
		if(match(count, files)) {
			mesg(&#39;m&#39;, file);
			scanobj(&amp;bar, amiddle, bp-&gt;size);
			arcopy(&amp;bar, amiddle, bp);
		} else if (ap == astart &amp;&amp; i == 0 &amp;&amp; strcmp(file, symdef) == 0) {
			/*
			 * pitch the symdef file if it is at the beginning
			 * of the archive and we aren&#39;t inserting in front
			 * of it (ap == astart).
			 */
			skip(&amp;bar, bp-&gt;size);
		} else if (ap == astart &amp;&amp; gflag &amp;&amp; strcmp(file, pkgdef) == 0) {
			/*
			 * pitch the pkgdef file if we aren&#39;t inserting in front
			 * of it (ap == astart).
			 */
			skip(&amp;bar, bp-&gt;size);
		} else {
			scanobj(&amp;bar, ap, bp-&gt;size);
			arcopy(&amp;bar, ap, bp);
		}
	}
	close(fd);
	if (poname[0] &amp;&amp; aend == 0)
		fprint(2, &#34;gopack: %s not found - files moved to end.\n&#34;, poname);
	install(arname, astart, amiddle, aend, 0);
}
void
tcmd(char *arname, int count, char **files)
{
	int fd;
	Armember *bp;
	char name[ARNAMESIZE+1];

	fd = openar(arname, OREAD, 0);
	Binit(&amp;bar, fd, OREAD);
	Bseek(&amp;bar,seek(fd,0,1), 1);
	while(bp = getdir(&amp;bar)) {
		if(count == 0 || match(count, files)) {
			if(vflag)
				longt(bp);
			trim(file, name, ARNAMESIZE);
			Bprint(&amp;bout, &#34;%s\n&#34;, name);
		}
		skip(&amp;bar, bp-&gt;size);
		free(bp);
	}
	close(fd);
}
void
qcmd(char *arname, int count, char **files)
{
	int fd, i;
	Armember *bp;
	Biobuf *bfile;

	if(aflag || bflag) {
		fprint(2, &#34;gopack: abi not allowed with q\n&#34;);
		exits(&#34;error&#34;);
	}
	fd = openar(arname, ORDWR, 1);
	if (fd &lt; 0) {
		if(!cflag)
			fprint(2, &#34;gopack: creating %s\n&#34;, arname);
		fd = arcreate(arname);
	}
	Binit(&amp;bar, fd, OREAD);
	Bseek(&amp;bar,seek(fd,0,1), 1);
	/* leave note group behind when writing archive; i.e. sidestep interrupts */
	rfork(RFNOTEG);
	Bseek(&amp;bar, 0, 2);
	bp = newmember();
	for(i=0; i&lt;count &amp;&amp; files[i]; i++) {
		file = files[i];
		files[i] = 0;
		bfile = Bopen(file, OREAD);
		if(!bfile) {
			fprint(2, &#34;gopack: cannot open %s\n&#34;, file);
			errors++;
		} else {
			mesg(&#39;q&#39;, file);
			armove(bfile, 0, bp);
			if (!arwrite(fd, bp))
				wrerr();
			free(bp-&gt;member);
			bp-&gt;member = 0;
			Bterm(bfile);
		}
	}
	free(bp);
	close(fd);
}

/*
 *	extract the symbol references from an object file
 */
void
scanobj(Biobuf *b, Arfile *ap, long size)
{
	int obj;
	vlong offset;
	Dir *d;
	static int lastobj = -1;

	if (!allobj)			/* non-object file encountered */
		return;
	offset = Boffset(b);
	obj = objtype(b, 0);
	if (obj &lt; 0) {			/* not an object file */
		if (!gflag || strcmp(file, pkgdef) != 0) {  /* don&#39;t clear allobj if it&#39;s pkg defs */
			fprint(2, &#34;gopack: non-object file %s\n&#34;, file);
			allobj = 0;
		}
		d = dirfstat(Bfildes(b));
		if (d != nil &amp;&amp; d-&gt;length == 0)
			fprint(2, &#34;gopack: zero length file %s\n&#34;, file);
		free(d);
		Bseek(b, offset, 0);
		return;
	}
	if (lastobj &gt;= 0 &amp;&amp; obj != lastobj) {
		fprint(2, &#34;gopack: inconsistent object file %s\n&#34;, file);
		allobj = 0;
		Bseek(b, offset, 0);
		return;
	}
	lastobj = obj;
	if (!readar(b, obj, offset+size, 0)) {
		fprint(2, &#34;gopack: invalid symbol reference in file %s\n&#34;, file);
		allobj = 0;
		Bseek(b, offset, 0);
		return;
	}
	Bseek(b, offset, 0);
	objtraverse(objsym, ap);
	if (gflag) {
		scanpkg(b, size);
		Bseek(b, offset, 0);
	}
}

/*
 * does line contain substring (length-limited)
 */
int
strstrn(char *line, int len, char *sub)
{
	int i;
	int sublen;

	sublen = strlen(sub);
	for (i = 0; i &lt; len - sublen; i++)
		if (memcmp(line+i, sub, sublen) == 0)
			return 1;
	return 0;
}

/*
 * Extract the package definition data from an object file
 */
void
scanpkg(Biobuf *b, long size)
{
	long n;
	int c;
	long start, end, pkgsize;
	char *data, *line, pkgbuf[1024], *pkg;
	int first;

	/*
	 * scan until $$
	 */
	for (n=0; n&lt;size; ) {
		c = Bgetc(b);
		if(c == Beof)
			break;
		n++;
		if(c != &#39;$&#39;)
			continue;
		c = Bgetc(b);
		if(c == Beof)
			break;
		n++;
		if(c != &#39;$&#39;)
			continue;
		goto foundstart;
	}
	// fprint(2, &#34;gopack: warning: no package import section in %s\n&#34;, file);
	return;

foundstart:
	/* found $$; skip rest of line */
	while((c = Bgetc(b)) != &#39;\n&#39;)
		if(c == Beof)
			goto bad;

	/* how big is it? */
	pkg = nil;
	first = 1;
	start = end = 0;
	for (n=0; n&lt;size; n+=Blinelen(b)) {
		line = Brdline(b, &#39;\n&#39;);
		if (line == 0)
			goto bad;
		if (first &amp;&amp; strstrn(line, Blinelen(b), &#34;package &#34;)) {
			if (Blinelen(b) &gt; sizeof(pkgbuf)-1)
				goto bad;
			memmove(pkgbuf, line, Blinelen(b));
			pkgbuf[Blinelen(b)] = &#39;\0&#39;;
			pkg = pkgbuf;
			while(*pkg == &#39; &#39; || *pkg == &#39;\t&#39;)
				pkg++;
			if(strncmp(pkg, &#34;package &#34;, 8) != 0)
				goto bad;
			start = Boffset(b);  // after package statement
			first = 0;
			continue;
		}
		if(line[0] == &#39;$&#39; &amp;&amp; line[1] == &#39;$&#39;)
			goto foundend;
		end = Boffset(b);  // before closing $$
	}
bad:
	fprint(2, &#34;gopack: bad package import section in %s\n&#34;, file);
	return;

foundend:
	if (start == 0)
		return;
	if (end == 0)
		goto bad;
	if (pkgstmt == nil) {
		/* this is the first package */
		pkgstmt = arstrdup(pkg);
	} else {
		if (strcmp(pkg, pkgstmt) != 0) {
			fprint(2, &#34;gopack: inconsistent package name\n&#34;);
			return;
		}
	}

	pkgsize = end-start;
	data = armalloc(pkgsize);
	Bseek(b, start, 0);
	if (Bread(b, data, pkgsize) != pkgsize) {
		fprint(2, &#34;gopack: error reading package import section in %s\n&#34;, file);
		return;
	}
	loadpkgdata(data, pkgsize);
}

/*
 *	add text and data symbols to the symbol list
 */
void
objsym(Sym *s, void *p)
{
	int n;
	Arsymref *as;
	Arfile *ap;
	char *ofile;

	if (s-&gt;type != &#39;T&#39; &amp;&amp;  s-&gt;type != &#39;D&#39;)
		return;
	ap = (Arfile*)p;
	as = armalloc(sizeof(Arsymref));
	as-&gt;offset = ap-&gt;size;
	as-&gt;name = arstrdup(s-&gt;name);
	as-&gt;file = arstrdup(file);
	if(s-&gt;type == &#39;T&#39; &amp;&amp; duplicate(as-&gt;name, &amp;ofile)) {
		dupfound = 1;
		fprint(2, &#34;duplicate text symbol: %s and %s: %s\n&#34;, as-&gt;file, ofile, as-&gt;name);
		free(as-&gt;name);
		free(as);
		return;
	}
	as-&gt;type = s-&gt;type;
	n = strlen(s-&gt;name);
	symdefsize += 4+(n+1)+1;
	as-&gt;len = n;
	as-&gt;next = ap-&gt;sym;
	ap-&gt;sym = as;
}

/*
 *	Check the symbol table for duplicate text symbols
 */
int
hashstr(char *name)
{
	int h;
	char *cp;

	h = 0;
	for(cp = name; *cp; h += *cp++)
		h *= 1119;
	
	// the code used to say
	//	if(h &lt; 0)
	//		h = ~h;
	// but on gcc 4.3 with -O2 on some systems,
	// the if(h &lt; 0) gets compiled away as not possible.
	// use a mask instead, leaving plenty of bits but
	// definitely not the sign bit.

	return h &amp; 0xfffffff;
}

int
duplicate(char *name, char **ofile)
{
	Hashchain *p;
	int h;

	h = hashstr(name) % NHASH;

	for(p = hash[h]; p; p = p-&gt;next)
		if(strcmp(p-&gt;name, name) == 0) {
			*ofile = p-&gt;file;
			return 1;
		}
	p = armalloc(sizeof(Hashchain));
	p-&gt;next = hash[h];
	p-&gt;name = name;
	p-&gt;file = file;
	hash[h] = p;
	*ofile = nil;
	return 0;
}

/*
 *	open an archive and validate its header
 */
int
openar(char *arname, int mode, int errok)
{
	int fd;
	char mbuf[SARMAG];

	fd = open(arname, mode);
	if(fd &gt;= 0){
		if(read(fd, mbuf, SARMAG) != SARMAG || strncmp(mbuf, ARMAG, SARMAG)) {
			fprint(2, &#34;gopack: %s not in archive format\n&#34;, arname);
			exits(&#34;error&#34;);
		}
	}else if(!errok){
		fprint(2, &#34;gopack: cannot open %s: %r\n&#34;, arname);
		exits(&#34;error&#34;);
	}
	return fd;
}

/*
 *	create an archive and set its header
 */
int
arcreate(char *arname)
{
	int fd;

	fd = create(arname, OWRITE, 0664);
	if(fd &lt; 0){
		fprint(2, &#34;gopack: cannot create %s: %r\n&#34;, arname);
		exits(&#34;error&#34;);
	}
	if(write(fd, ARMAG, SARMAG) != SARMAG)
		wrerr();
	return fd;
}

/*
 *		error handling
 */
void
wrerr(void)
{
	perror(&#34;gopack: write error&#34;);
	exits(&#34;error&#34;);
}

void
rderr(void)
{
	perror(&#34;gopack: read error&#34;);
	exits(&#34;error&#34;);
}

void
phaseerr(int offset)
{
	fprint(2, &#34;gopack: phase error at offset %d\n&#34;, offset);
	exits(&#34;error&#34;);
}

void
usage(void)
{
	fprint(2, &#34;usage: gopack [%s][%s] archive files ...\n&#34;, opt, man);
	exits(&#34;error&#34;);
}

/*
 *	read the header for the next archive member
 */
Armember *
getdir(Biobuf *b)
{
	Armember *bp;
	char *cp;
	static char name[ARNAMESIZE+1];

	bp = newmember();
	if(HEADER_IO(Bread, b, bp-&gt;hdr)) {
		free(bp);
		return 0;
	}
	if(strncmp(bp-&gt;hdr.fmag, ARFMAG, sizeof(bp-&gt;hdr.fmag)))
		phaseerr(Boffset(b));
	strncpy(name, bp-&gt;hdr.name, sizeof(bp-&gt;hdr.name));
	cp = name+sizeof(name)-1;
	while(*--cp==&#39; &#39;)
		;
	cp[1] = &#39;\0&#39;;
	file = arstrdup(name);
	bp-&gt;date = strtol(bp-&gt;hdr.date, 0, 0);
	bp-&gt;size = strtol(bp-&gt;hdr.size, 0, 0);
	return bp;
}

/*
 *	Copy the file referenced by fd to the temp file
 */
void
armove(Biobuf *b, Arfile *ap, Armember *bp)
{
	char *cp;
	Dir *d;

	d = dirfstat(Bfildes(b));
	if (d == nil) {
		fprint(2, &#34;gopack: cannot stat %s\n&#34;, file);
		return;
	}
	trim(file, bp-&gt;hdr.name, sizeof(bp-&gt;hdr.name));
	for (cp = strchr(bp-&gt;hdr.name, 0);		/* blank pad on right */
		cp &lt; bp-&gt;hdr.name+sizeof(bp-&gt;hdr.name); cp++)
			*cp = &#39; &#39;;
	sprint(bp-&gt;hdr.date, &#34;%-12ld&#34;, d-&gt;mtime);
	sprint(bp-&gt;hdr.uid, &#34;%-6d&#34;, 0);
	sprint(bp-&gt;hdr.gid, &#34;%-6d&#34;, 0);
	sprint(bp-&gt;hdr.mode, &#34;%-8lo&#34;, d-&gt;mode);
	sprint(bp-&gt;hdr.size, &#34;%-10lld&#34;, d-&gt;length);
	strncpy(bp-&gt;hdr.fmag, ARFMAG, 2);
	bp-&gt;size = d-&gt;length;
	arread(b, bp, bp-&gt;size);
	if (d-&gt;length&amp;0x01)
		d-&gt;length++;
	if (ap) {
		arinsert(ap, bp);
		ap-&gt;size += d-&gt;length+SAR_HDR;
	}
	free(d);
}

/*
 *	Copy the archive member at the current offset into the temp file.
 */
void
arcopy(Biobuf *b, Arfile *ap, Armember *bp)
{
	long n;

	n = bp-&gt;size;
	if (n &amp; 01)
		n++;
	arread(b, bp, n);
	if (ap) {
		arinsert(ap, bp);
		ap-&gt;size += n+SAR_HDR;
	}
}

/*
 *	Skip an archive member
 */
void
skip(Biobuf *bp, vlong len)
{
	if (len &amp; 01)
		len++;
	Bseek(bp, len, 1);
}

/*
 *	Stream the three temp files to an archive
 */
void
install(char *arname, Arfile *astart, Arfile *amiddle, Arfile *aend, int createflag)
{
	int fd;

	if(allobj &amp;&amp; dupfound) {
		fprint(2, &#34;%s not changed\n&#34;, arname);
		return;
	}
	/* leave note group behind when copying back; i.e. sidestep interrupts */
	rfork(RFNOTEG);

	if(createflag)
		fprint(2, &#34;gopack: creating %s\n&#34;, arname);
	fd = arcreate(arname);

	if(allobj)
		rl(fd);

	if (astart) {
		arstream(fd, astart);
		arfree(astart);
	}
	if (amiddle) {
		arstream(fd, amiddle);
		arfree(amiddle);
	}
	if (aend) {
		arstream(fd, aend);
		arfree(aend);
	}
	close(fd);
}

void
rl(int fd)
{
	Biobuf b;
	char *cp;
	struct ar_hdr a;
	long len;
	int headlen;
	char *pkgdefdata;
	int pkgdefsize;

	pkgdefdata = nil;
	pkgdefsize = 0;

	Binit(&amp;b, fd, OWRITE);
	Bseek(&amp;b,seek(fd,0,1), 0);

	len = symdefsize;
	if(len&amp;01)
		len++;
	sprint(a.date, &#34;%-12ld&#34;, time(0));
	sprint(a.uid, &#34;%-6d&#34;, 0);
	sprint(a.gid, &#34;%-6d&#34;, 0);
	sprint(a.mode, &#34;%-8lo&#34;, 0644L);
	sprint(a.size, &#34;%-10ld&#34;, len);
	strncpy(a.fmag, ARFMAG, 2);
	strcpy(a.name, symdef);
	for (cp = strchr(a.name, 0);		/* blank pad on right */
		cp &lt; a.name+sizeof(a.name); cp++)
			*cp = &#39; &#39;;
	if(HEADER_IO(Bwrite, &amp;b, a))
			wrerr();

	headlen = Boffset(&amp;b);
	len += headlen;
	if (gflag) {
		getpkgdef(&amp;pkgdefdata, &amp;pkgdefsize);
		len += SAR_HDR + pkgdefsize;
		if (len &amp; 1)
			len++;
	}
	if (astart) {
		wrsym(&amp;b, len, astart-&gt;sym);
		len += astart-&gt;size;
	}
	if(amiddle) {
		wrsym(&amp;b, len, amiddle-&gt;sym);
		len += amiddle-&gt;size;
	}
	if(aend)
		wrsym(&amp;b, len, aend-&gt;sym);

	if(symdefsize&amp;0x01)
		Bputc(&amp;b, 0);

	if (gflag) {
		len = pkgdefsize;
		sprint(a.date, &#34;%-12ld&#34;, time(0));
		sprint(a.uid, &#34;%-6d&#34;, 0);
		sprint(a.gid, &#34;%-6d&#34;, 0);
		sprint(a.mode, &#34;%-8lo&#34;, 0644L);
		sprint(a.size, &#34;%-10ld&#34;, (len + 1) &amp; ~1);
		strncpy(a.fmag, ARFMAG, 2);
		strcpy(a.name, pkgdef);
		for (cp = strchr(a.name, 0);		/* blank pad on right */
			cp &lt; a.name+sizeof(a.name); cp++)
				*cp = &#39; &#39;;
		if(HEADER_IO(Bwrite, &amp;b, a))
				wrerr();

		if (Bwrite(&amp;b, pkgdefdata, pkgdefsize) != pkgdefsize)
			wrerr();
		if(len&amp;0x01)
			Bputc(&amp;b, 0);
	}
	Bterm(&amp;b);
}

/*
 *	Write the defined symbols to the symdef file
 */
void
wrsym(Biobuf *bp, long offset, Arsymref *as)
{
	int off;

	while(as) {
		Bputc(bp, as-&gt;type);
		off = as-&gt;offset+offset;
		Bputc(bp, off);
		Bputc(bp, off&gt;&gt;8);
		Bputc(bp, off&gt;&gt;16);
		Bputc(bp, off&gt;&gt;24);
		if (Bwrite(bp, as-&gt;name, as-&gt;len+1) != as-&gt;len+1)
			wrerr();
		as = as-&gt;next;
	}
}

/*
 *	Check if the archive member matches an entry on the command line.
 */
int
match(int count, char **files)
{
	int i;
	char name[ARNAMESIZE+1];

	for(i=0; i&lt;count; i++) {
		if(files[i] == 0)
			continue;
		trim(files[i], name, ARNAMESIZE);
		if(strncmp(name, file, ARNAMESIZE) == 0) {
			file = files[i];
			files[i] = 0;
			return 1;
		}
	}
	return 0;
}

/*
 *	compare the current member to the name of the pivot member
 */
int
bamatch(char *file, char *pivot)
{
	static int state = 0;

	switch(state)
	{
	case 0:			/* looking for position file */
		if (aflag) {
			if (strncmp(file, pivot, ARNAMESIZE) == 0)
				state = 1;
		} else if (bflag) {
			if (strncmp(file, pivot, ARNAMESIZE) == 0) {
				state = 2;	/* found */
				return 1;
			}
		}
		break;
	case 1:			/* found - after previous file */
		state = 2;
		return 1;
	case 2:			/* already found position file */
		break;
	}
	return 0;
}

/*
 *	output a message, if &#39;v&#39; option was specified
 */
void
mesg(int c, char *file)
{

	if(vflag)
		Bprint(&amp;bout, &#34;%c - %s\n&#34;, c, file);
}

/*
 *	isolate file name by stripping leading directories and trailing slashes
 */
void
trim(char *s, char *buf, int n)
{
	char *p;

	for(;;) {
		p = strrchr(s, &#39;/&#39;);
		if (!p) {		/* no slash in name */
			strncpy(buf, s, n);
			return;
		}
		if (p[1] != 0) {	/* p+1 is first char of file name */
			strncpy(buf, p+1, n);
			return;
		}
		*p = 0;			/* strip trailing slash */
	}
}

/*
 *	utilities for printing long form of &#39;t&#39; command
 */
#define	SUID	04000
#define	SGID	02000
#define	ROWN	0400
#define	WOWN	0200
#define	XOWN	0100
#define	RGRP	040
#define	WGRP	020
#define	XGRP	010
#define	ROTH	04
#define	WOTH	02
#define	XOTH	01
#define	STXT	01000

void
longt(Armember *bp)
{
	char *cp;
	time_t date;

	pmode(strtoul(bp-&gt;hdr.mode, 0, 8));
	Bprint(&amp;bout, &#34;%3ld/%1ld&#34;, strtol(bp-&gt;hdr.uid, 0, 0), strtol(bp-&gt;hdr.gid, 0, 0));
	Bprint(&amp;bout, &#34;%7ld&#34;, bp-&gt;size);
	date = bp-&gt;date;
	cp = ctime(&amp;date);
	Bprint(&amp;bout, &#34; %-12.12s %-4.4s &#34;, cp+4, cp+24);
}

int	m1[] = { 1, ROWN, &#39;r&#39;, &#39;-&#39; };
int	m2[] = { 1, WOWN, &#39;w&#39;, &#39;-&#39; };
int	m3[] = { 2, SUID, &#39;s&#39;, XOWN, &#39;x&#39;, &#39;-&#39; };
int	m4[] = { 1, RGRP, &#39;r&#39;, &#39;-&#39; };
int	m5[] = { 1, WGRP, &#39;w&#39;, &#39;-&#39; };
int	m6[] = { 2, SGID, &#39;s&#39;, XGRP, &#39;x&#39;, &#39;-&#39; };
int	m7[] = { 1, ROTH, &#39;r&#39;, &#39;-&#39; };
int	m8[] = { 1, WOTH, &#39;w&#39;, &#39;-&#39; };
int	m9[] = { 2, STXT, &#39;t&#39;, XOTH, &#39;x&#39;, &#39;-&#39; };

int	*m[] = { m1, m2, m3, m4, m5, m6, m7, m8, m9};

void
pmode(long mode)
{
	int **mp;

	for(mp = &amp;m[0]; mp &lt; &amp;m[9];)
		select(*mp++, mode);
}

void
select(int *ap, long mode)
{
	int n;

	n = *ap++;
	while(--n&gt;=0 &amp;&amp; (mode&amp;*ap++)==0)
		ap++;
	Bputc(&amp;bout, *ap);
}

/*
 *	Temp file I/O subsystem.  We attempt to cache all three temp files in
 *	core.  When we run out of memory we spill to disk.
 *	The I/O model assumes that temp files:
 *		1) are only written on the end
 *		2) are only read from the beginning
 *		3) are only read after all writing is complete.
 *	The architecture uses one control block per temp file.  Each control
 *	block anchors a chain of buffers, each containing an archive member.
 */
Arfile *
newtempfile(char *name)		/* allocate a file control block */
{
	Arfile *ap;

	ap = armalloc(sizeof(Arfile));
	ap-&gt;fname = name;
	return ap;
}

Armember *
newmember(void)			/* allocate a member buffer */
{
	return armalloc(sizeof(Armember));
}

void
arread(Biobuf *b, Armember *bp, int n)	/* read an image into a member buffer */
{
	int i;

	bp-&gt;member = armalloc(n);
	i = Bread(b, bp-&gt;member, n);
	if (i &lt; 0) {
		free(bp-&gt;member);
		bp-&gt;member = 0;
		rderr();
	}
}

/*
 * insert a member buffer into the member chain
 */
void
arinsert(Arfile *ap, Armember *bp)
{
	bp-&gt;next = 0;
	if (!ap-&gt;tail)
		ap-&gt;head = bp;
	else
		ap-&gt;tail-&gt;next = bp;
	ap-&gt;tail = bp;
}

/*
 *	stream the members in a temp file to the file referenced by &#39;fd&#39;.
 */
void
arstream(int fd, Arfile *ap)
{
	Armember *bp;
	int i;
	char buf[8192];

	if (ap-&gt;paged) {		/* copy from disk */
		seek(ap-&gt;fd, 0, 0);
		for (;;) {
			i = read(ap-&gt;fd, buf, sizeof(buf));
			if (i &lt; 0)
				rderr();
			if (i == 0)
				break;
			if (write(fd, buf, i) != i)
				wrerr();
		}
		close(ap-&gt;fd);
		ap-&gt;paged = 0;
	}
		/* dump the in-core buffers */
	for (bp = ap-&gt;head; bp; bp = bp-&gt;next) {
		if (!arwrite(fd, bp))
			wrerr();
	}
}

/*
 *	write a member to &#39;fd&#39;.
 */
int
arwrite(int fd, Armember *bp)
{
	int len;

	if(HEADER_IO(write, fd, bp-&gt;hdr))
		return 0;
	len = bp-&gt;size;
	if (len &amp; 01)
		len++;
	if (write(fd, bp-&gt;member, len) != len)
		return 0;
	return 1;
}

/*
 *	Spill a member to a disk copy of a temp file
 */
int
page(Arfile *ap)
{
	Armember *bp;

	bp = ap-&gt;head;
	if (!ap-&gt;paged) {		/* not yet paged - create file */
		ap-&gt;fname = mktemp(ap-&gt;fname);
		ap-&gt;fd = create(ap-&gt;fname, ORDWR|ORCLOSE, 0600);
		if (ap-&gt;fd &lt; 0) {
			fprint(2,&#34;gopack: can&#39;t create temp file\n&#34;);
			return 0;
		}
		ap-&gt;paged = 1;
	}
	if (!arwrite(ap-&gt;fd, bp))	/* write member and free buffer block */
		return 0;
	ap-&gt;head = bp-&gt;next;
	if (ap-&gt;tail == bp)
		ap-&gt;tail = bp-&gt;next;
	free(bp-&gt;member);
	free(bp);
	return 1;
}

/*
 *	try to reclaim space by paging.  we try to spill the start, middle,
 *	and end files, in that order.  there is no particular reason for the
 *	ordering.
 */
int
getspace(void)
{
fprint(2, &#34;IN GETSPACE\n&#34;);
	if (astart &amp;&amp; astart-&gt;head &amp;&amp; page(astart))
		return 1;
	if (amiddle &amp;&amp; amiddle-&gt;head &amp;&amp; page(amiddle))
		return 1;
	if (aend &amp;&amp; aend-&gt;head &amp;&amp; page(aend))
		return 1;
	return 0;
}

void
arfree(Arfile *ap)		/* free a member buffer */
{
	Armember *bp, *next;

	for (bp = ap-&gt;head; bp; bp = next) {
		next = bp-&gt;next;
		if (bp-&gt;member)
			free(bp-&gt;member);
		free(bp);
	}
	free(ap);
}

/*
 *	allocate space for a control block or member buffer.  if the malloc
 *	fails we try to reclaim space by spilling previously allocated
 *	member buffers.
 */
void *
armalloc(int n)
{
	char *cp;

	// bump so that arwrite can do the same
	if(n&amp;1)
		n++;

	do {
		cp = malloc(n);
		if (cp) {
			memset(cp, 0, n);
			return cp;
		}
	} while (getspace());
	fprint(2, &#34;gopack: out of memory\n&#34;);
	exits(&#34;malloc&#34;);
	return 0;
}

char *
arstrdup(char *s)
{
	char *t;

	t = armalloc(strlen(s) + 1);
	strcpy(t, s);
	return t;
}


/*
 *	package import data
 */
typedef struct Import Import;
struct Import
{
	Import *hash;	// next in hash table
	char *prefix;	// &#34;type&#34;, &#34;var&#34;, &#34;func&#34;, &#34;const&#34;
	char *name;
	char *def;
	char *file;
};
enum {
	NIHASH = 1024
};
Import *ihash[NIHASH];
int nimport;

Import *
ilookup(char *name)
{
	int h;
	Import *x;

	h = hashstr(name) % NIHASH;
	for(x=ihash[h]; x; x=x-&gt;hash)
		if(x-&gt;name[0] == name[0] &amp;&amp; strcmp(x-&gt;name, name) == 0)
			return x;
	x = armalloc(sizeof *x);
	x-&gt;name = name;
	x-&gt;hash = ihash[h];
	ihash[h] = x;
	nimport++;
	return x;
}

/*
 * a and b don&#39;t match.
 * is one a forward declaration and the other a valid completion?
 * if so, return the one to keep.
 */
char*
forwardfix(char *a, char *b)
{
	char *t;

	if(strlen(a) &gt; strlen(b)) {
		t = a;
		a = b;
		b = t;
	}
	if(strcmp(a, &#34;struct&#34;) == 0 &amp;&amp; strncmp(b, &#34;struct &#34;, 7) == 0)
		return b;
	if(strcmp(a, &#34;interface&#34;) == 0 &amp;&amp; strncmp(b, &#34;interface &#34;, 10) == 0)
		return b;
	return nil;
}

int parsemethod(char**, char*, char**);
int parsepkgdata(char**, char*, char**, char**, char**);

void
loadpkgdata(char *data, int len)
{
	char *p, *ep, *prefix, *name, *def, *ndef;
	Import *x;

	p = data;
	ep = data + len;
	while(parsepkgdata(&amp;p, ep, &amp;prefix, &amp;name, &amp;def) &gt; 0) {
		x = ilookup(name);
		if(x-&gt;prefix == nil) {
			x-&gt;prefix = prefix;
			x-&gt;def = def;
			x-&gt;file = file;
		} else if(strcmp(x-&gt;prefix, prefix) != 0) {
			fprint(2, &#34;gopack: conflicting definitions for %s\n&#34;, name);
			fprint(2, &#34;%s:\t%s %s ...\n&#34;, x-&gt;file, x-&gt;prefix, name);
			fprint(2, &#34;%s:\t%s %s ...\n&#34;, file, prefix, name);
			errors++;
		} else if(strcmp(x-&gt;def, def) == 0) {
			// fine
		} else if((ndef = forwardfix(x-&gt;def, def)) != nil) {
			x-&gt;def = ndef;
		} else {
			fprint(2, &#34;gopack: conflicting definitions for %s\n&#34;, name);
			fprint(2, &#34;%s:\t%s %s %s\n&#34;, x-&gt;file, x-&gt;prefix, name, x-&gt;def);
			fprint(2, &#34;%s:\t%s %s %s\n&#34;, file, prefix, name, def);
			errors++;
		}
	}
}

int
parsepkgdata(char **pp, char *ep, char **prefixp, char **namep, char **defp)
{
	char *p, *prefix, *name, *def, *edef, *meth;
	int n;

	// skip white space
	p = *pp;
	while(p &lt; ep &amp;&amp; (*p == &#39; &#39; || *p == &#39;\t&#39;))
		p++;
	if(p == ep)
		return 0;

	// prefix: (var|type|func|const)
	prefix = p;

	prefix = p;
	if(p + 6 &gt; ep)
		return -1;
	if(strncmp(p, &#34;var &#34;, 4) == 0)
		p += 4;
	else if(strncmp(p, &#34;type &#34;, 5) == 0)
		p += 5;
	else if(strncmp(p, &#34;func &#34;, 5) == 0)
		p += 5;
	else if(strncmp(p, &#34;const &#34;, 6) == 0)
		p += 6;
	else {
		fprint(2, &#34;gopack: confused in pkg data near &lt;&lt;%.20s&gt;&gt;\n&#34;, p);
		errors++;
		return -1;
	}
	p[-1] = &#39;\0&#39;;

	// name: a.b followed by space
	name = p;
	while(p &lt; ep &amp;&amp; *p != &#39; &#39;)
		p++;
	if(p &gt;= ep)
		return -1;
	*p++ = &#39;\0&#39;;

	// def: free form to new line
	def = p;
	while(p &lt; ep &amp;&amp; *p != &#39;\n&#39;)
		p++;
	if(p &gt;= ep)
		return -1;
	edef = p;
	*p++ = &#39;\0&#39;;

	// include methods on successive lines in def of named type
	while(parsemethod(&amp;p, ep, &amp;meth) &gt; 0) {
		*edef++ = &#39;\n&#39;;	// overwrites &#39;\0&#39;
		if(edef+1 &gt; meth) {
			// We want to indent methods with a single \t.
			// 6g puts at least one char of indent before all method defs,
			// so there will be room for the \t.  If the method def wasn&#39;t
			// indented we could do something more complicated,
			// but for now just diagnose the problem and assume
			// 6g will keep indenting for us.
			fprint(2, &#34;gopack: %s: expected methods to be indented %p %p %.10s\n&#34;, file, edef, meth, meth);
			errors++;
			return -1;
		}
		*edef++ = &#39;\t&#39;;
		n = strlen(meth);
		memmove(edef, meth, n);
		edef += n;
	}

	// done
	*pp = p;
	*prefixp = prefix;
	*namep = name;
	*defp = def;
	return 1;
}

int
parsemethod(char **pp, char *ep, char **methp)
{
	char *p;

	// skip white space
	p = *pp;
	while(p &lt; ep &amp;&amp; (*p == &#39; &#39; || *p == &#39;\t&#39;))
		p++;
	if(p == ep)
		return 0;

	// if it says &#34;func (&#34;, it&#39;s a method
	if(p + 6 &gt;= ep || strncmp(p, &#34;func (&#34;, 6) != 0)
		return 0;

	// definition to end of line
	*methp = p;
	while(p &lt; ep &amp;&amp; *p != &#39;\n&#39;)
		p++;
	if(p &gt;= ep) {
		fprint(2, &#34;gopack: lost end of line in method definition\n&#34;);
		*pp = ep;
		return -1;
	}
	*p++ = &#39;\0&#39;;
	*pp = p;
	return 1;
}

int
importcmp(const void *va, const void *vb)
{
	Import *a, *b;
	int i;

	a = *(Import**)va;
	b = *(Import**)vb;

	i = strcmp(a-&gt;prefix, b-&gt;prefix);
	if(i != 0) {
		// rewrite so &#34;type&#34; comes first
		if(strcmp(a-&gt;prefix, &#34;type&#34;) == 0)
			return -1;
		if(strcmp(b-&gt;prefix, &#34;type&#34;) == 0)
			return 1;
		return i;
	}
	return strcmp(a-&gt;name, b-&gt;name);
}

char*
strappend(char *s, char *t)
{
	int n;

	n = strlen(t);
	memmove(s, t, n);
	return s+n;
}

void
getpkgdef(char **datap, int *lenp)
{
	int i, j, len;
	char *data, *p;
	Import **all, *x;

	if(pkgstmt == nil) {
		// Write out non-empty, parseable __.PKGDEF,
		// so that import of an empty archive works.
		*datap = &#34;import\n$$\npackage __emptypackage__\n$$\n&#34;;
		*lenp = strlen(*datap);
		return;
	}

	// make a list of all the exports and count string sizes
	all = armalloc(nimport*sizeof all[0]);
	j = 0;
	len = 7 + 3 + strlen(pkgstmt) + 1;	// import\n$$\npkgstmt\n
	for(i=0; i&lt;NIHASH; i++) {
		for(x=ihash[i]; x; x=x-&gt;hash) {
			all[j++] = x;
			len += strlen(x-&gt;prefix) + 1
				+ strlen(x-&gt;name) + 1
				+ strlen(x-&gt;def) + 1;
		}
	}
	if(j != nimport) {
		fprint(2, &#34;gopack: import count mismatch (internal error)\n&#34;);
		exits(&#34;oops&#34;);
	}
	len += 3;	// $$\n

	// sort exports (unnecessary but nicer to look at)
	qsort(all, nimport, sizeof all[0], importcmp);

	// print them into buffer
	data = armalloc(len);

	// import\n
	// $$\n
	// pkgstmt\n
	p = data;
	p = strappend(p, &#34;import\n$$\n&#34;);
	p = strappend(p, pkgstmt);
	p = strappend(p, &#34;\n&#34;);
	for(i=0; i&lt;nimport; i++) {
		x = all[i];
		// prefix name def\n
		p = strappend(p, x-&gt;prefix);
		p = strappend(p, &#34; &#34;);
		p = strappend(p, x-&gt;name);
		p = strappend(p, &#34; &#34;);
		p = strappend(p, x-&gt;def);
		p = strappend(p, &#34;\n&#34;);
	}
	p = strappend(p, &#34;$$\n&#34;);
	if(p != data+len) {
		fprint(2, &#34;gopack: internal math error\n&#34;);
		exits(&#34;oops&#34;);
	}

	*datap = data;
	*lenp = len;
}

</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
