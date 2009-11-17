<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/sym.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/sym.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/sym.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/sym.c
//
// 	Copyright © 1994-1999 Lucent Technologies Inc.
// 	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
// 	Portions Copyright © 1997-1999 Vita Nuova Limited.
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
// 	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
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

#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;

#define	HUGEINT	0x7fffffff
#define	NNAME	20		/* a relic of the past */

typedef	struct txtsym Txtsym;
typedef	struct file File;
typedef	struct hist Hist;

struct txtsym {				/* Text Symbol table */
	int 	n;			/* number of local vars */
	Sym	**locals;		/* array of ptrs to autos */
	Sym	*sym;			/* function symbol entry */
};

struct hist {				/* Stack of include files &amp; #line directives */
	char	*name;			/* Assumes names Null terminated in file */
	int32	line;			/* line # where it was included */
	int32	offset;			/* line # of #line directive */
};

struct file {				/* Per input file header to history stack */
	uvlong	addr;			/* address of first text sym */
	union {
		Txtsym	*txt;		/* first text symbol */
		Sym	*sym;		/* only during initilization */
	};
	int	n;			/* size of history stack */
	Hist	*hist;			/* history stack */
};

static	int	debug = 0;

static	Sym	**autos;		/* Base of auto variables */
static	File	*files;			/* Base of file arena */
static	int	fmaxi;			/* largest file path index */
static	Sym	**fnames;		/* file names path component table */
static	Sym	**globals;		/* globals by addr table */
static	Hist	*hist;			/* base of history stack */
static	int	isbuilt;		/* internal table init flag */
static	int32	nauto;			/* number of automatics */
static	int32	nfiles;			/* number of files */
static	int32	nglob;			/* number of globals */
static	int32	nhist;			/* number of history stack entries */
static	int32	nsym;			/* number of symbols */
static	int	ntxt;			/* number of text symbols */
static	uchar	*pcline;		/* start of pc-line state table */
static	uchar 	*pclineend;		/* end of pc-line table */
static	uchar	*spoff;			/* start of pc-sp state table */
static	uchar	*spoffend;		/* end of pc-sp offset table */
static	Sym	*symbols;		/* symbol table */
static	Txtsym	*txt;			/* Base of text symbol table */
static	uvlong	txtstart;		/* start of text segment */
static	uvlong	txtend;			/* end of text segment */
static	uvlong	firstinstr;		/* as found from symtab; needed for amd64 */

static void	cleansyms(void);
static int32	decodename(Biobuf*, Sym*);
static short	*encfname(char*);
static int 	fline(char*, int, int32, Hist*, Hist**);
static void	fillsym(Sym*, Symbol*);
static int	findglobal(char*, Symbol*);
static int	findlocvar(Symbol*, char *, Symbol*);
static int	findtext(char*, Symbol*);
static int	hcomp(Hist*, short*);
static int	hline(File*, short*, int32*);
static void	printhist(char*, Hist*, int);
static int	buildtbls(void);
static int	symcomp(const void*, const void*);
static int	symerrmsg(int, char*);
static int	txtcomp(const void*, const void*);
static int	filecomp(const void*, const void*);

/*
 *	initialize the symbol tables
 */
int
syminit(int fd, Fhdr *fp)
{
	Sym *p;
	int32 i, l, size;
	vlong vl;
	Biobuf b;
	int svalsz;

	if(fp-&gt;symsz == 0)
		return 0;
	if(fp-&gt;type == FNONE)
		return 0;

	cleansyms();
	textseg(fp-&gt;txtaddr, fp);
		/* minimum symbol record size = 4+1+2 bytes */
	symbols = malloc((fp-&gt;symsz/(4+1+2)+1)*sizeof(Sym));
	if(symbols == 0) {
		werrstr(&#34;can&#39;t malloc %ld bytes&#34;, fp-&gt;symsz);
		return -1;
	}
	Binit(&amp;b, fd, OREAD);
	Bseek(&amp;b, fp-&gt;symoff, 0);
	nsym = 0;
	size = 0;
	for(p = symbols; size &lt; fp-&gt;symsz; p++, nsym++) {
		if(fp-&gt;_magic &amp;&amp; (fp-&gt;magic &amp; HDR_MAGIC)){
			svalsz = 8;
			if(Bread(&amp;b, &amp;vl, 8) != 8)
				return symerrmsg(8, &#34;symbol&#34;);
			p-&gt;value = beswav(vl);
		}
		else{
			svalsz = 4;
			if(Bread(&amp;b, &amp;l, 4) != 4)
				return symerrmsg(4, &#34;symbol&#34;);
			p-&gt;value = (u32int)beswal(l);
		}
		if(Bread(&amp;b, &amp;p-&gt;type, sizeof(p-&gt;type)) != sizeof(p-&gt;type))
			return symerrmsg(sizeof(p-&gt;value), &#34;symbol&#34;);

		i = decodename(&amp;b, p);
		if(i &lt; 0)
			return -1;
		size += i+svalsz+sizeof(p-&gt;type);

		if(svalsz == 8){
			if(Bread(&amp;b, &amp;vl, 8) != 8)
				return symerrmsg(8, &#34;symbol&#34;);
			p-&gt;gotype = beswav(vl);
		}
		else{
			if(Bread(&amp;b, &amp;l, 4) != 4)
				return symerrmsg(4, &#34;symbol&#34;);
			p-&gt;gotype = (u32int)beswal(l);
		}
		size += svalsz;

		/* count global &amp; auto vars, text symbols, and file names */
		switch (p-&gt;type) {
		case &#39;l&#39;:
		case &#39;L&#39;:
		case &#39;t&#39;:
		case &#39;T&#39;:
			ntxt++;
			break;
		case &#39;d&#39;:
		case &#39;D&#39;:
		case &#39;b&#39;:
		case &#39;B&#39;:
			nglob++;
			break;
		case &#39;f&#39;:
			if(strcmp(p-&gt;name, &#34;.frame&#34;) == 0) {
				p-&gt;type = &#39;m&#39;;
				nauto++;
			}
			else if(p-&gt;value &gt; fmaxi)
				fmaxi = p-&gt;value;	/* highest path index */
			break;
		case &#39;a&#39;:
		case &#39;p&#39;:
		case &#39;m&#39;:
			nauto++;
			break;
		case &#39;z&#39;:
			if(p-&gt;value == 1) {		/* one extra per file */
				nhist++;
				nfiles++;
			}
			nhist++;
			break;
		default:
			break;
		}
	}
	if (debug)
		print(&#34;NG: %ld NT: %d NF: %d\n&#34;, nglob, ntxt, fmaxi);
	if (fp-&gt;sppcsz) {			/* pc-sp offset table */
		spoff = (uchar *)malloc(fp-&gt;sppcsz);
		if(spoff == 0) {
			werrstr(&#34;can&#39;t malloc %ld bytes&#34;, fp-&gt;sppcsz);
			return -1;
		}
		Bseek(&amp;b, fp-&gt;sppcoff, 0);
		if(Bread(&amp;b, spoff, fp-&gt;sppcsz) != fp-&gt;sppcsz){
			spoff = 0;
			return symerrmsg(fp-&gt;sppcsz, &#34;sp-pc&#34;);
		}
		spoffend = spoff+fp-&gt;sppcsz;
	}
	if (fp-&gt;lnpcsz) {			/* pc-line number table */
		pcline = (uchar *)malloc(fp-&gt;lnpcsz);
		if(pcline == 0) {
			werrstr(&#34;can&#39;t malloc %ld bytes&#34;, fp-&gt;lnpcsz);
			return -1;
		}
		Bseek(&amp;b, fp-&gt;lnpcoff, 0);
		if(Bread(&amp;b, pcline, fp-&gt;lnpcsz) != fp-&gt;lnpcsz){
			pcline = 0;
			return symerrmsg(fp-&gt;lnpcsz, &#34;pc-line&#34;);
		}
		pclineend = pcline+fp-&gt;lnpcsz;
	}
	return nsym;
}

static int
symerrmsg(int n, char *table)
{
	werrstr(&#34;can&#39;t read %d bytes of %s table&#34;, n, table);
	return -1;
}

static int32
decodename(Biobuf *bp, Sym *p)
{
	char *cp;
	int c1, c2;
	int32 n;
	vlong o;

	if((p-&gt;type &amp; 0x80) == 0) {		/* old-style, fixed length names */
		p-&gt;name = malloc(NNAME);
		if(p-&gt;name == 0) {
			werrstr(&#34;can&#39;t malloc %d bytes&#34;, NNAME);
			return -1;
		}
		if(Bread(bp, p-&gt;name, NNAME) != NNAME)
			return symerrmsg(NNAME, &#34;symbol&#34;);
		Bseek(bp, 3, 1);
		return NNAME+3;
	}

	p-&gt;type &amp;= ~0x80;
	if(p-&gt;type == &#39;z&#39; || p-&gt;type == &#39;Z&#39;) {
		o = Bseek(bp, 0, 1);
		if(Bgetc(bp) &lt; 0) {
			werrstr(&#34;can&#39;t read symbol name&#34;);
			return -1;
		}
		for(;;) {
			c1 = Bgetc(bp);
			c2 = Bgetc(bp);
			if(c1 &lt; 0 || c2 &lt; 0) {
				werrstr(&#34;can&#39;t read symbol name&#34;);
				return -1;
			}
			if(c1 == 0 &amp;&amp; c2 == 0)
				break;
		}
		n = Bseek(bp, 0, 1)-o;
		p-&gt;name = malloc(n);
		if(p-&gt;name == 0) {
			werrstr(&#34;can&#39;t malloc %ld bytes&#34;, n);
			return -1;
		}
		Bseek(bp, -n, 1);
		if(Bread(bp, p-&gt;name, n) != n) {
			werrstr(&#34;can&#39;t read %ld bytes of symbol name&#34;, n);
			return -1;
		}
	} else {
		cp = Brdline(bp, &#39;\0&#39;);
		if(cp == 0) {
			werrstr(&#34;can&#39;t read symbol name&#34;);
			return -1;
		}
		n = Blinelen(bp);
		p-&gt;name = malloc(n);
		if(p-&gt;name == 0) {
			werrstr(&#34;can&#39;t malloc %ld bytes&#34;, n);
			return -1;
		}
		strcpy(p-&gt;name, cp);
	}
	return n;
}

/*
 *	free any previously loaded symbol tables
 */
static void
cleansyms(void)
{
	if(globals)
		free(globals);
	globals = 0;
	nglob = 0;
	if(txt)
		free(txt);
	txt = 0;
	ntxt = 0;
	if(fnames)
		free(fnames);
	fnames = 0;
	fmaxi = 0;

	if(files)
		free(files);
	files = 0;
	nfiles = 0;
	if(hist)
		free(hist);
	hist = 0;
	nhist = 0;
	if(autos)
		free(autos);
	autos = 0;
	nauto = 0;
	isbuilt = 0;
	if(symbols)
		free(symbols);
	symbols = 0;
	nsym = 0;
	if(spoff)
		free(spoff);
	spoff = 0;
	if(pcline)
		free(pcline);
	pcline = 0;
}

/*
 *	delimit the text segment
 */
void
textseg(uvlong base, Fhdr *fp)
{
	txtstart = base;
	txtend = base+fp-&gt;txtsz;
}

/*
 *	symbase: return base and size of raw symbol table
 *		(special hack for high access rate operations)
 */
Sym *
symbase(int32 *n)
{
	*n = nsym;
	return symbols;
}

/*
 *	Get the ith symbol table entry
 */
Sym *
getsym(int index)
{
	if(index &gt;= 0 &amp;&amp; index &lt; nsym)
		return &amp;symbols[index];
	return 0;
}

/*
 *	initialize internal symbol tables
 */
static int
buildtbls(void)
{
	int32 i;
	int j, nh, ng, nt;
	File *f;
	Txtsym *tp;
	Hist *hp;
	Sym *p, **ap;

	if(isbuilt)
		return 1;
	isbuilt = 1;
			/* allocate the tables */
	firstinstr = 0;
	if(nglob) {
		globals = malloc(nglob*sizeof(*globals));
		if(!globals) {
			werrstr(&#34;can&#39;t malloc global symbol table&#34;);
			return 0;
		}
	}
	if(ntxt) {
		txt = malloc(ntxt*sizeof(*txt));
		if (!txt) {
			werrstr(&#34;can&#39;t malloc text symbol table&#34;);
			return 0;
		}
	}
	fnames = malloc((fmaxi+1)*sizeof(*fnames));
	if (!fnames) {
		werrstr(&#34;can&#39;t malloc file name table&#34;);
		return 0;
	}
	memset(fnames, 0, (fmaxi+1)*sizeof(*fnames));
	files = malloc(nfiles*sizeof(*files));
	if(!files) {
		werrstr(&#34;can&#39;t malloc file table&#34;);
		return 0;
	}
	hist = malloc(nhist*sizeof(Hist));
	if(hist == 0) {
		werrstr(&#34;can&#39;t malloc history stack&#34;);
		return 0;
	}
	autos = malloc(nauto*sizeof(Sym*));
	if(autos == 0) {
		werrstr(&#34;can&#39;t malloc auto symbol table&#34;);
		return 0;
	}
		/* load the tables */
	ng = nt = nh = 0;
	f = 0;
	tp = 0;
	i = nsym;
	hp = hist;
	ap = autos;
	for(p = symbols; i-- &gt; 0; p++) {
//print(&#34;sym %d type %c name %s value %llux\n&#34;, p-symbols, p-&gt;type, p-&gt;name, p-&gt;value);
		switch(p-&gt;type) {
		case &#39;D&#39;:
		case &#39;d&#39;:
		case &#39;B&#39;:
		case &#39;b&#39;:
			if(debug)
				print(&#34;Global: %s %llux\n&#34;, p-&gt;name, p-&gt;value);
			globals[ng++] = p;
			break;
		case &#39;z&#39;:
			if(p-&gt;value == 1) {		/* New file */
				if(f) {
					f-&gt;n = nh;
					f-&gt;hist[nh].name = 0;	/* one extra */
					hp += nh+1;
					f++;
				}
				else
					f = files;
				f-&gt;hist = hp;
				f-&gt;sym = 0;
				f-&gt;addr = 0;
				nh = 0;
			}
				/* alloc one slot extra as terminator */
			f-&gt;hist[nh].name = p-&gt;name;
			f-&gt;hist[nh].line = p-&gt;value;
			f-&gt;hist[nh].offset = 0;
			if(debug)
				printhist(&#34;-&gt; &#34;, &amp;f-&gt;hist[nh], 1);
			nh++;
			break;
		case &#39;Z&#39;:
			if(f &amp;&amp; nh &gt; 0)
				f-&gt;hist[nh-1].offset = p-&gt;value;
			break;
		case &#39;T&#39;:
		case &#39;t&#39;:	/* Text: terminate history if first in file */
		case &#39;L&#39;:
		case &#39;l&#39;:
			tp = &amp;txt[nt++];
			tp-&gt;n = 0;
			tp-&gt;sym = p;
			tp-&gt;locals = ap;
			if(debug)
				print(&#34;TEXT: %s at %llux\n&#34;, p-&gt;name, p-&gt;value);
			if (firstinstr == 0 || p-&gt;value &lt; firstinstr)
				firstinstr = p-&gt;value;
			if(f &amp;&amp; !f-&gt;sym) {			/* first  */
				f-&gt;sym = p;
				f-&gt;addr = p-&gt;value;
			}
			break;
		case &#39;a&#39;:
		case &#39;p&#39;:
		case &#39;m&#39;:		/* Local Vars */
			if(!tp)
				print(&#34;Warning: Free floating local var: %s\n&#34;,
					p-&gt;name);
			else {
				if(debug)
					print(&#34;Local: %s %llux\n&#34;, p-&gt;name, p-&gt;value);
				tp-&gt;locals[tp-&gt;n] = p;
				tp-&gt;n++;
				ap++;
			}
			break;
		case &#39;f&#39;:		/* File names */
			if(debug)
				print(&#34;Fname: %s\n&#34;, p-&gt;name);
			fnames[p-&gt;value] = p;
			break;
		default:
			break;
		}
	}
		/* sort global and text tables into ascending address order */
	qsort(globals, nglob, sizeof(Sym*), symcomp);
	qsort(txt, ntxt, sizeof(Txtsym), txtcomp);
	qsort(files, nfiles, sizeof(File), filecomp);
	tp = txt;
	for(i = 0, f = files; i &lt; nfiles; i++, f++) {
		for(j = 0; j &lt; ntxt; j++) {
			if(f-&gt;sym == tp-&gt;sym) {
				if(debug) {
					print(&#34;LINK: %s to at %llux&#34;, f-&gt;sym-&gt;name, f-&gt;addr);
					printhist(&#34;... &#34;, f-&gt;hist, 1);
				}
				f-&gt;txt = tp++;
				break;
			}
			if(++tp &gt;= txt+ntxt)	/* wrap around */
				tp = txt;
		}
	}
	return 1;
}

/*
 * find symbol function.var by name.
 *	fn != 0 &amp;&amp; var != 0	=&gt; look for fn in text, var in data
 *	fn != 0 &amp;&amp; var == 0	=&gt; look for fn in text
 *	fn == 0 &amp;&amp; var != 0	=&gt; look for var first in text then in data space.
 */
int
lookup(char *fn, char *var, Symbol *s)
{
	int found;

	if(buildtbls() == 0)
		return 0;
	if(fn) {
		found = findtext(fn, s);
		if(var == 0)		/* case 2: fn not in text */
			return found;
		else if(!found)		/* case 1: fn not found */
			return 0;
	} else if(var) {
		found = findtext(var, s);
		if(found)
			return 1;	/* case 3: var found in text */
	} else return 0;		/* case 4: fn &amp; var == zero */

	if(found)
		return findlocal(s, var, s);	/* case 1: fn found */
	return findglobal(var, s);		/* case 3: var not found */

}

/*
 * strcmp, but allow &#39;_&#39; to match center dot (rune 00b7 == bytes c2 b7)
 */
int
cdotstrcmp(char *sym, char *user) {
	for (;;) {
		while (*sym == *user) {
			if (*sym++ == &#39;\0&#39;)
				return 0;
			user++;
		}
		/* unequal - but maybe &#39;_&#39; matches center dot */
		if (user[0] == &#39;_&#39; &amp;&amp; (sym[0]&amp;0xFF) == 0xc2 &amp;&amp; (sym[1]&amp;0xFF) == 0xb7) {
			/* &#39;_&#39; matches center dot - advance and continue */
			user++;
			sym += 2;
			continue;
		}
		break;
	}
	return *user - *sym;
}

/*
 * find a function by name
 */
static int
findtext(char *name, Symbol *s)
{
	int i;

	for(i = 0; i &lt; ntxt; i++) {
		if(cdotstrcmp(txt[i].sym-&gt;name, name) == 0) {
			fillsym(txt[i].sym, s);
			s-&gt;handle = (void *) &amp;txt[i];
			s-&gt;index = i;
			return 1;
		}
	}
	return 0;
}
/*
 * find global variable by name
 */
static int
findglobal(char *name, Symbol *s)
{
	int32 i;

	for(i = 0; i &lt; nglob; i++) {
		if(cdotstrcmp(globals[i]-&gt;name, name) == 0) {
			fillsym(globals[i], s);
			s-&gt;index = i;
			return 1;
		}
	}
	return 0;
}

/*
 *	find the local variable by name within a given function
 */
int
findlocal(Symbol *s1, char *name, Symbol *s2)
{
	if(s1 == 0)
		return 0;
	if(buildtbls() == 0)
		return 0;
	return findlocvar(s1, name, s2);
}

/*
 *	find the local variable by name within a given function
 *		(internal function - does no parameter validation)
 */
static int
findlocvar(Symbol *s1, char *name, Symbol *s2)
{
	Txtsym *tp;
	int i;

	tp = (Txtsym *)s1-&gt;handle;
	if(tp &amp;&amp; tp-&gt;locals) {
		for(i = 0; i &lt; tp-&gt;n; i++)
			if (cdotstrcmp(tp-&gt;locals[i]-&gt;name, name) == 0) {
				fillsym(tp-&gt;locals[i], s2);
				s2-&gt;handle = (void *)tp;
				s2-&gt;index = tp-&gt;n-1 - i;
				return 1;
			}
	}
	return 0;
}

/*
 *	Get ith text symbol
 */
int
textsym(Symbol *s, int index)
{

	if(buildtbls() == 0)
		return 0;
	if(index &lt; 0 || index &gt;= ntxt)
		return 0;
	fillsym(txt[index].sym, s);
	s-&gt;handle = (void *)&amp;txt[index];
	s-&gt;index = index;
	return 1;
}

/*
 *	Get ith file name
 */
int
filesym(int index, char *buf, int n)
{
	Hist *hp;

	if(buildtbls() == 0)
		return 0;
	if(index &lt; 0 || index &gt;= nfiles)
		return 0;
	hp = files[index].hist;
	if(!hp || !hp-&gt;name)
		return 0;
	return fileelem(fnames, (uchar*)hp-&gt;name, buf, n);
}

/*
 *	Lookup name of local variable located at an offset into the frame.
 *	The type selects either a parameter or automatic.
 */
int
getauto(Symbol *s1, int off, int type, Symbol *s2)
{
	Txtsym *tp;
	Sym *p;
	int i, t;

	if(s1 == 0)
		return 0;
	if(type == CPARAM)
		t = &#39;p&#39;;
	else if(type == CAUTO)
		t = &#39;a&#39;;
	else
		return 0;
	if(buildtbls() == 0)
		return 0;
	tp = (Txtsym *)s1-&gt;handle;
	if(tp == 0)
		return 0;
	for(i = 0; i &lt; tp-&gt;n; i++) {
		p = tp-&gt;locals[i];
		if(p-&gt;type == t &amp;&amp; p-&gt;value == off) {
			fillsym(p, s2);
			s2-&gt;handle = s1-&gt;handle;
			s2-&gt;index = tp-&gt;n-1 - i;
			return 1;
		}
	}
	return 0;
}

/*
 * Find text symbol containing addr; binary search assumes text array is sorted by addr
 */
static int
srchtext(uvlong addr)
{
	uvlong val;
	int top, bot, mid;
	Sym *sp;

	val = addr;
	bot = 0;
	top = ntxt;
	for (mid = (bot+top)/2; mid &lt; top; mid = (bot+top)/2) {
		sp = txt[mid].sym;
		if(val &lt; sp-&gt;value)
			top = mid;
		else if(mid != ntxt-1 &amp;&amp; val &gt;= txt[mid+1].sym-&gt;value)
			bot = mid;
		else
			return mid;
	}
	return -1;
}

/*
 * Find data symbol containing addr; binary search assumes data array is sorted by addr
 */
static int
srchdata(uvlong addr)
{
	uvlong val;
	int top, bot, mid;
	Sym *sp;

	bot = 0;
	top = nglob;
	val = addr;
	for(mid = (bot+top)/2; mid &lt; top; mid = (bot+top)/2) {
		sp = globals[mid];
		if(val &lt; sp-&gt;value)
			top = mid;
		else if(mid &lt; nglob-1 &amp;&amp; val &gt;= globals[mid+1]-&gt;value)
			bot = mid;
		else
			return mid;
	}
	return -1;
}

/*
 * Find symbol containing val in specified search space
 * There is a special case when a value falls beyond the end
 * of the text segment; if the search space is CTEXT, that value
 * (usually etext) is returned.  If the search space is CANY, symbols in the
 * data space are searched for a match.
 */
int
findsym(uvlong val, int type, Symbol *s)
{
	int i;

	if(buildtbls() == 0)
		return 0;

	if(type == CTEXT || type == CANY) {
		i = srchtext(val);
		if(i &gt;= 0) {
			if(type == CTEXT || i != ntxt-1) {
				fillsym(txt[i].sym, s);
				s-&gt;handle = (void *) &amp;txt[i];
				s-&gt;index = i;
				return 1;
			}
		}
	}
	if(type == CDATA || type == CANY) {
		i = srchdata(val);
		if(i &gt;= 0) {
			fillsym(globals[i], s);
			s-&gt;index = i;
			return 1;
		}
	}
	return 0;
}

/*
 *	Find the start and end address of the function containing addr
 */
int
fnbound(uvlong addr, uvlong *bounds)
{
	int i;

	if(buildtbls() == 0)
		return 0;

	i = srchtext(addr);
	if(0 &lt;= i &amp;&amp; i &lt; ntxt-1) {
		bounds[0] = txt[i].sym-&gt;value;
		bounds[1] = txt[i+1].sym-&gt;value;
		return 1;
	}
	return 0;
}

/*
 * get the ith local symbol for a function
 * the input symbol table is reverse ordered, so we reverse
 * accesses here to maintain approx. parameter ordering in a stack trace.
 */
int
localsym(Symbol *s, int index)
{
	Txtsym *tp;

	if(s == 0 || index &lt; 0)
		return 0;
	if(buildtbls() == 0)
		return 0;

	tp = (Txtsym *)s-&gt;handle;
	if(tp &amp;&amp; tp-&gt;locals &amp;&amp; index &lt; tp-&gt;n) {
		fillsym(tp-&gt;locals[tp-&gt;n-index-1], s);	/* reverse */
		s-&gt;handle = (void *)tp;
		s-&gt;index = index;
		return 1;
	}
	return 0;
}

/*
 * get the ith global symbol
 */
int
globalsym(Symbol *s, int index)
{
	if(s == 0)
		return 0;
	if(buildtbls() == 0)
		return 0;

	if(index &gt;=0 &amp;&amp; index &lt; nglob) {
		fillsym(globals[index], s);
		s-&gt;index = index;
		return 1;
	}
	return 0;
}

/*
 *	find the pc given a file name and line offset into it.
 */
uvlong
file2pc(char *file, int32 line)
{
	File *fp;
	int32 i;
	uvlong pc, start, end;
	short *name;

	if(buildtbls() == 0 || files == 0)
		return ~0;
	name = encfname(file);
	if(name == 0) {			/* encode the file name */
		werrstr(&#34;file %s not found&#34;, file);
		return ~0;
	}
		/* find this history stack */
	for(i = 0, fp = files; i &lt; nfiles; i++, fp++)
		if (hline(fp, name, &amp;line))
			break;
	free(name);
	if(i &gt;= nfiles) {
		werrstr(&#34;line %ld in file %s not found&#34;, line, file);
		return ~0;
	}
	start = fp-&gt;addr;		/* first text addr this file */
	if(i &lt; nfiles-1)
		end = (fp+1)-&gt;addr;	/* first text addr next file */
	else
		end = 0;		/* last file in load module */
	/*
	 * At this point, line contains the offset into the file.
	 * run the state machine to locate the pc closest to that value.
	 */
	if(debug)
		print(&#34;find pc for %ld - between: %llux and %llux\n&#34;, line, start, end);
	pc = line2addr(line, start, end);
	if(pc == ~0) {
		werrstr(&#34;line %ld not in file %s&#34;, line, file);
		return ~0;
	}
	return pc;
}

/*
 *	search for a path component index
 */
static int
pathcomp(char *s, int n)
{
	int i;

	for(i = 0; i &lt;= fmaxi; i++)
		if(fnames[i] &amp;&amp; strncmp(s, fnames[i]-&gt;name, n) == 0)
			return i;
	return -1;
}

/*
 *	Encode a char file name as a sequence of short indices
 *	into the file name dictionary.
 */
static short*
encfname(char *file)
{
	int i, j;
	char *cp, *cp2;
	short *dest;

	if(*file == &#39;/&#39;)	/* always check first &#39;/&#39; */
		cp2 = file+1;
	else {
		cp2 = strchr(file, &#39;/&#39;);
		if(!cp2)
			cp2 = strchr(file, 0);
	}
	cp = file;
	dest = 0;
	for(i = 0; *cp; i++) {
		j = pathcomp(cp, cp2-cp);
		if(j &lt; 0)
			return 0;	/* not found */
		dest = realloc(dest, (i+1)*sizeof(short));
		dest[i] = j;
		cp = cp2;
		while(*cp == &#39;/&#39;)	/* skip embedded &#39;/&#39;s */
			cp++;
		cp2 = strchr(cp, &#39;/&#39;);
		if(!cp2)
			cp2 = strchr(cp, 0);
	}
	dest = realloc(dest, (i+1)*sizeof(short));
	dest[i] = 0;
	return dest;
}

/*
 *	Search a history stack for a matching file name accumulating
 *	the size of intervening files in the stack.
 */
static int
hline(File *fp, short *name, int32 *line)
{
	Hist *hp;
	int offset, depth;
	int32 ln;

	for(hp = fp-&gt;hist; hp-&gt;name; hp++)		/* find name in stack */
		if(hp-&gt;name[1] || hp-&gt;name[2]) {
			if(hcomp(hp, name))
				break;
		}
	if(!hp-&gt;name)		/* match not found */
		return 0;
	if(debug)
		printhist(&#34;hline found ... &#34;, hp, 1);
	/*
	 * unwind the stack until empty or we hit an entry beyond our line
	 */
	ln = *line;
	offset = hp-&gt;line-1;
	depth = 1;
	for(hp++; depth &amp;&amp; hp-&gt;name; hp++) {
		if(debug)
			printhist(&#34;hline inspect ... &#34;, hp, 1);
		if(hp-&gt;name[1] || hp-&gt;name[2]) {
			if(hp-&gt;offset){			/* Z record */
				offset = 0;
				if(hcomp(hp, name)) {
					if(*line &lt;= hp-&gt;offset)
						break;
					ln = *line+hp-&gt;line-hp-&gt;offset;
					depth = 1;	/* implicit pop */
				} else
					depth = 2;	/* implicit push */
			} else if(depth == 1 &amp;&amp; ln &lt; hp-&gt;line-offset)
					break;		/* Beyond our line */
			else if(depth++ == 1)		/* push	*/
				offset -= hp-&gt;line;
		} else if(--depth == 1)		/* pop */
			offset += hp-&gt;line;
	}
	*line = ln+offset;
	return 1;
}

/*
 *	compare two encoded file names
 */
static int
hcomp(Hist *hp, short *sp)
{
	uchar *cp;
	int i, j;
	short *s;

	cp = (uchar *)hp-&gt;name;
	s = sp;
	if (*s == 0)
		return 0;
	for (i = 1; j = (cp[i]&lt;&lt;8)|cp[i+1]; i += 2) {
		if(j == 0)
			break;
		if(*s == j)
			s++;
		else
			s = sp;
	}
	return *s == 0;
}

/*
 *	Convert a pc to a &#34;file:line {file:line}&#34; string.
 */
int32
fileline(char *str, int n, uvlong dot)
{
	int32 line, top, bot, mid;
	File *f;

	*str = 0;
	if(buildtbls() == 0)
		return 0;
		/* binary search assumes file list is sorted by addr */
	bot = 0;
	top = nfiles;
	for (mid = (bot+top)/2; mid &lt; top; mid = (bot+top)/2) {
		f = &amp;files[mid];
		if(dot &lt; f-&gt;addr)
			top = mid;
		else if(mid &lt; nfiles-1 &amp;&amp; dot &gt;= (f+1)-&gt;addr)
			bot = mid;
		else {
			line = pc2line(dot);
			if(line &gt; 0 &amp;&amp; fline(str, n, line, f-&gt;hist, 0) &gt;= 0)
				return 1;
			break;
		}
	}
	return 0;
}

/*
 *	Convert a line number within a composite file to relative line
 *	number in a source file.  A composite file is the source
 *	file with included files inserted in line.
 */
static int
fline(char *str, int n, int32 line, Hist *base, Hist **ret)
{
	Hist *start;			/* start of current level */
	Hist *h;			/* current entry */
	int32 delta;			/* sum of size of files this level */
	int k;

	start = base;
	h = base;
	delta = h-&gt;line;
	while(h &amp;&amp; h-&gt;name &amp;&amp; line &gt; h-&gt;line) {
		if(h-&gt;name[1] || h-&gt;name[2]) {
			if(h-&gt;offset != 0) {	/* #line Directive */
				delta = h-&gt;line-h-&gt;offset+1;
				start = h;
				base = h++;
			} else {		/* beginning of File */
				if(start == base)
					start = h++;
				else {
					k = fline(str, n, line, start, &amp;h);
					if(k &lt;= 0)
						return k;
				}
			}
		} else {
			if(start == base &amp;&amp; ret) {	/* end of recursion level */
				*ret = h;
				return 1;
			} else {			/* end of included file */
				delta += h-&gt;line-start-&gt;line;
				h++;
				start = base;
			}
		}
	}
	if(!h)
		return -1;
	if(start != base)
		line = line-start-&gt;line+1;
	else
		line = line-delta+1;
	if(!h-&gt;name)
		strncpy(str, &#34;&lt;eof&gt;&#34;, n);
	else {
		k = fileelem(fnames, (uchar*)start-&gt;name, str, n);
		if(k+8 &lt; n)
			sprint(str+k, &#34;:%ld&#34;, line);
	}
/**********Remove comments for complete back-trace of include sequence
 *	if(start != base) {
 *		k = strlen(str);
 *		if(k+2 &lt; n) {
 *			str[k++] = &#39; &#39;;
 *			str[k++] = &#39;{&#39;;
 *		}
 *		k += fileelem(fnames, (uchar*) base-&gt;name, str+k, n-k);
 *		if(k+10 &lt; n)
 *			sprint(str+k, &#34;:%ld}&#34;, start-&gt;line-delta);
 *	}
 ********************/
	return 0;
}

/*
 *	convert an encoded file name to a string.
 */
int
fileelem(Sym **fp, uchar *cp, char *buf, int n)
{
	int i, j;
	char *c, *bp, *end;

	bp = buf;
	end = buf+n-1;
	for(i = 1; j = (cp[i]&lt;&lt;8)|cp[i+1]; i+=2){
		c = fp[j]-&gt;name;
		if(bp != buf &amp;&amp; bp[-1] != &#39;/&#39; &amp;&amp; bp &lt; end)
			*bp++ = &#39;/&#39;;
		while(bp &lt; end &amp;&amp; *c)
			*bp++ = *c++;
	}
	*bp = 0;
	i =  bp-buf;
	if(i &gt; 1) {
		cleanname(buf);
		i = strlen(buf);
	}
	return i;
}

/*
 *	compare the values of two symbol table entries.
 */
static int
symcomp(const void *a, const void *b)
{
	int i;

	i = (*(Sym**)a)-&gt;value - (*(Sym**)b)-&gt;value;
	if (i)
		return i;
	return strcmp((*(Sym**)a)-&gt;name, (*(Sym**)b)-&gt;name);
}

/*
 *	compare the values of the symbols referenced by two text table entries
 */
static int
txtcomp(const void *a, const void *b)
{
	return ((Txtsym*)a)-&gt;sym-&gt;value - ((Txtsym*)b)-&gt;sym-&gt;value;
}

/*
 *	compare the values of the symbols referenced by two file table entries
 */
static int
filecomp(const void *a, const void *b)
{
	return ((File*)a)-&gt;addr - ((File*)b)-&gt;addr;
}

/*
 *	fill an interface Symbol structure from a symbol table entry
 */
static void
fillsym(Sym *sp, Symbol *s)
{
	s-&gt;type = sp-&gt;type;
	s-&gt;value = sp-&gt;value;
	s-&gt;name = sp-&gt;name;
	s-&gt;index = 0;
	switch(sp-&gt;type) {
	case &#39;b&#39;:
	case &#39;B&#39;:
	case &#39;D&#39;:
	case &#39;d&#39;:
		s-&gt;class = CDATA;
		break;
	case &#39;t&#39;:
	case &#39;T&#39;:
	case &#39;l&#39;:
	case &#39;L&#39;:
		s-&gt;class = CTEXT;
		break;
	case &#39;a&#39;:
		s-&gt;class = CAUTO;
		break;
	case &#39;p&#39;:
		s-&gt;class = CPARAM;
		break;
	case &#39;m&#39;:
		s-&gt;class = CSTAB;
		break;
	default:
		s-&gt;class = CNONE;
		break;
	}
	s-&gt;handle = 0;
}

/*
 *	find the stack frame, given the pc
 */
uvlong
pc2sp(uvlong pc)
{
	uchar *c, u;
	uvlong currpc, currsp;

	if(spoff == 0)
		return ~0;
	currsp = 0;
	currpc = txtstart - mach-&gt;pcquant;

	if(pc&lt;currpc || pc&gt;txtend)
		return ~0;
	for(c = spoff; c &lt; spoffend; c++) {
		if (currpc &gt;= pc)
			return currsp;
		u = *c;
		if (u == 0) {
			currsp += (c[1]&lt;&lt;24)|(c[2]&lt;&lt;16)|(c[3]&lt;&lt;8)|c[4];
			c += 4;
		}
		else if (u &lt; 65)
			currsp += 4*u;
		else if (u &lt; 129)
			currsp -= 4*(u-64);
		else
			currpc += mach-&gt;pcquant*(u-129);
		currpc += mach-&gt;pcquant;
	}
	return ~0;
}

/*
 *	find the source file line number for a given value of the pc
 */
int32
pc2line(uvlong pc)
{
	uchar *c, u;
	uvlong currpc;
	int32 currline;

	if(pcline == 0)
		return -1;
	currline = 0;
	if (firstinstr != 0)
		currpc = firstinstr-mach-&gt;pcquant;
	else
		currpc = txtstart-mach-&gt;pcquant;
	if(pc&lt;currpc || pc&gt;txtend)
		return ~0;

	for(c = pcline; c &lt; pclineend &amp;&amp; currpc &lt; pc; c++) {
		u = *c;
		if(u == 0) {
			currline += (c[1]&lt;&lt;24)|(c[2]&lt;&lt;16)|(c[3]&lt;&lt;8)|c[4];
			c += 4;
		}
		else if(u &lt; 65)
			currline += u;
		else if(u &lt; 129)
			currline -= (u-64);
		else
			currpc += mach-&gt;pcquant*(u-129);
		currpc += mach-&gt;pcquant;
	}
	return currline;
}

/*
 *	find the pc associated with a line number
 *	basepc and endpc are text addresses bounding the search.
 *	if endpc == 0, the end of the table is used (i.e., no upper bound).
 *	usually, basepc and endpc contain the first text address in
 *	a file and the first text address in the following file, respectively.
 */
uvlong
line2addr(int32 line, uvlong basepc, uvlong endpc)
{
	uchar *c,  u;
	uvlong currpc, pc;
	int32 currline;
	int32 delta, d;
	int found;

	if(pcline == 0 || line == 0)
		return ~0;

	currline = 0;
	currpc = txtstart-mach-&gt;pcquant;
	pc = ~0;
	found = 0;
	delta = HUGEINT;

	for(c = pcline; c &lt; pclineend; c++) {
		if(endpc &amp;&amp; currpc &gt;= endpc)	/* end of file of interest */
			break;
		if(currpc &gt;= basepc) {		/* proper file */
			if(currline &gt;= line) {
				d = currline-line;
				found = 1;
			} else
				d = line-currline;
			if(d &lt; delta) {
				delta = d;
				pc = currpc;
			}
		}
		u = *c;
		if(u == 0) {
			currline += (c[1]&lt;&lt;24)|(c[2]&lt;&lt;16)|(c[3]&lt;&lt;8)|c[4];
			c += 4;
		}
		else if(u &lt; 65)
			currline += u;
		else if(u &lt; 129)
			currline -= (u-64);
		else
			currpc += mach-&gt;pcquant*(u-129);
		currpc += mach-&gt;pcquant;
	}
	if(found)
		return pc;
	return ~0;
}

/*
 *	Print a history stack (debug). if count is 0, prints the whole stack
 */
static void
printhist(char *msg, Hist *hp, int count)
{
	int i;
	uchar *cp;
	char buf[128];

	i = 0;
	while(hp-&gt;name) {
		if(count &amp;&amp; ++i &gt; count)
			break;
		print(&#34;%s Line: %lx (%ld)  Offset: %lx (%ld)  Name: &#34;, msg,
			hp-&gt;line, hp-&gt;line, hp-&gt;offset, hp-&gt;offset);
		for(cp = (uchar *)hp-&gt;name+1; (*cp&lt;&lt;8)|cp[1]; cp += 2) {
			if (cp != (uchar *)hp-&gt;name+1)
				print(&#34;/&#34;);
			print(&#34;%x&#34;, (*cp&lt;&lt;8)|cp[1]);
		}
		fileelem(fnames, (uchar *) hp-&gt;name, buf, sizeof(buf));
		print(&#34; (%s)\n&#34;, buf);
		hp++;
	}
}

#ifdef DEBUG
/*
 *	print the history stack for a file. (debug only)
 *	if (name == 0) =&gt; print all history stacks.
 */
void
dumphist(char *name)
{
	int i;
	File *f;
	short *fname;

	if(buildtbls() == 0)
		return;
	if(name)
		fname = encfname(name);
	for(i = 0, f = files; i &lt; nfiles; i++, f++)
		if(fname == 0 || hcomp(f-&gt;hist, fname))
			printhist(&#34;&gt; &#34;, f-&gt;hist, f-&gt;n);

	if(fname)
		free(fname);
}
#endif
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
