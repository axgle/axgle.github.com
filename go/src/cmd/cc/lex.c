<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/lex.c</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/lex.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/lex.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/lex.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
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

#include	&#34;cc.h&#34;
#include	&#34;y.tab.h&#34;

#ifndef	CPP
#define	CPP	&#34;/bin/cpp&#34;
#endif

int
systemtype(int sys)
{

	return sys&amp;Plan9;
}

int
pathchar(void)
{
	return &#39;/&#39;;
}

/*
 * known debug flags
 *	-a		acid declaration output
 *	-A		!B
 *	-B		non ANSI
 *	-d		print declarations
 *	-D name		define
 *	-F		format specification check
 *	-i		print initialization
 *	-I path		include
 *	-l		generate little-endian code
 *	-L		print every NAME symbol
 *	-M		constant multiplication
 *	-m		print add/sub/mul trees
 *	-n		print acid to file (%.c=%.acid) (with -a or -aa)
 *	-o file		output file
 *	-p		use standard cpp ANSI preprocessor (not on windows)
 *	-r		print registerization
 *	-s		print structure offsets (with -a or -aa)
 *	-S		print assembly
 *	-t		print type trees
 *	-V		enable void* conversion warnings
 *	-v		verbose printing
 *	-w		print warnings
 *	-X		abort on error
 *	-.		Inhibit search for includes in source directory
 */

void
main(int argc, char *argv[])
{
	char *defs[50], *p;
	int nproc, nout, i, c, ndef;

	memset(debug, 0, sizeof(debug));
	tinit();
	cinit();
	ginit();
	arginit();

	tufield = simplet((1L&lt;&lt;tfield-&gt;etype) | BUNSIGNED);
	ndef = 0;
	outfile = 0;
	include[ninclude++] = &#34;.&#34;;
	ARGBEGIN {
	default:
		c = ARGC();
		if(c &gt;= 0 &amp;&amp; c &lt; sizeof(debug))
			debug[c]++;
		break;

	case &#39;l&#39;:			/* for little-endian mips */
		if(thechar != &#39;v&#39;){
			print(&#34;can only use -l with vc&#34;);
			errorexit();
		}
		thechar = &#39;0&#39;;
		thestring = &#34;spim&#34;;
		break;

	case &#39;o&#39;:
		outfile = ARGF();
		break;

	case &#39;D&#39;:
		p = ARGF();
		if(p) {
			defs[ndef++] = p;
			dodefine(p);
		}
		break;

	case &#39;I&#39;:
		p = ARGF();
		setinclude(p);
		break;
	} ARGEND
	if(argc &lt; 1 &amp;&amp; outfile == 0) {
		print(&#34;usage: %cc [-options] files\n&#34;, thechar);
		errorexit();
	}
	if(argc &gt; 1 &amp;&amp; systemtype(Windows)){
		print(&#34;can&#39;t compile multiple files on windows\n&#34;);
		errorexit();
	}
	if(argc &gt; 1 &amp;&amp; !systemtype(Windows)) {
		nproc = 1;
		/*
		 * if we&#39;re writing acid to standard output, don&#39;t compile
		 * concurrently, to avoid interleaving output.
		 */
		if(((!debug[&#39;a&#39;] &amp;&amp; !debug[&#39;Z&#39;]) || debug[&#39;n&#39;]) &amp;&amp;
		    (p = getenv(&#34;NPROC&#34;)) != nil)
			nproc = atol(p);	/* */
		c = 0;
		nout = 0;
		for(;;) {
			Waitmsg *w;

			while(nout &lt; nproc &amp;&amp; argc &gt; 0) {
				i = fork();
				if(i &lt; 0) {
					print(&#34;cannot create a process\n&#34;);
					errorexit();
				}
				if(i == 0) {
					fprint(2, &#34;%s:\n&#34;, *argv);
					if (compile(*argv, defs, ndef))
						errorexit();
					exits(0);
				}
				nout++;
				argc--;
				argv++;
			}
			w = wait();
			if(w == nil) {
				if(c)
					errorexit();
				exits(0);
			}
			if(w-&gt;msg[0])
				c++;
			nout--;
		}
	}

	if(argc == 0)
		c = compile(&#34;stdin&#34;, defs, ndef);
	else
		c = compile(argv[0], defs, ndef);

	if(c)
		errorexit();
	exits(0);
}

int
compile(char *file, char **defs, int ndef)
{
	char ofile[400], incfile[20];
	char *p, *av[100], opt[256];
	int i, c, fd[2];
	static int first = 1;

	strcpy(ofile, file);
	p = utfrrune(ofile, pathchar());
	if(p) {
		*p++ = 0;
		if(!debug[&#39;.&#39;])
			include[0] = strdup(ofile);
	} else
		p = ofile;

	if(outfile == 0) {
		outfile = p;
		if(outfile) {
			if(p = utfrrune(outfile, &#39;.&#39;))
				if(p[1] == &#39;c&#39; &amp;&amp; p[2] == 0)
					p[0] = 0;
			p = utfrune(outfile, 0);
			if(debug[&#39;a&#39;] &amp;&amp; debug[&#39;n&#39;])
				strcat(p, &#34;.acid&#34;);
			else if(debug[&#39;Z&#39;] &amp;&amp; debug[&#39;n&#39;])
				strcat(p, &#34;_pickle.c&#34;);
			else {
				p[0] = &#39;.&#39;;
				p[1] = thechar;
				p[2] = 0;
			}
		} else
			outfile = &#34;/dev/null&#34;;
	}

	if(p = getenv(&#34;INCLUDE&#34;)) {
		setinclude(p);
	} else {
		if(systemtype(Plan9)) {
			sprint(incfile, &#34;/%s/include&#34;, thestring);
			setinclude(strdup(incfile));
			setinclude(&#34;/sys/include&#34;);
		}
	}
	if (first)
		Binit(&amp;diagbuf, 1, OWRITE);
	/*
	 * if we&#39;re writing acid to standard output, don&#39;t keep scratching
	 * outbuf.
	 */
	if((debug[&#39;a&#39;] || debug[&#39;Z&#39;]) &amp;&amp; !debug[&#39;n&#39;]) {
		if (first) {
			outfile = 0;
			Binit(&amp;outbuf, dup(1, -1), OWRITE);
			dup(2, 1);
		}
	} else {
		c = create(outfile, OWRITE, 0664);
		if(c &lt; 0) {
			diag(Z, &#34;cannot open %s - %r&#34;, outfile);
			outfile = 0;
			errorexit();
		}
		Binit(&amp;outbuf, c, OWRITE);
		outfile = strdup(outfile);
	}
	newio();
	first = 0;

	/* Use an ANSI preprocessor */
	if(debug[&#39;p&#39;]) {
		if(systemtype(Windows)) {
			diag(Z, &#34;-p option not supported on windows&#34;);
			errorexit();
		}
		if(access(file, AREAD) &lt; 0) {
			diag(Z, &#34;%s does not exist&#34;, file);
			errorexit();
		}
		if(pipe(fd) &lt; 0) {
			diag(Z, &#34;pipe failed&#34;);
			errorexit();
		}
		switch(fork()) {
		case -1:
			diag(Z, &#34;fork failed&#34;);
			errorexit();
		case 0:
			close(fd[0]);
			dup(fd[1], 1);
			close(fd[1]);
			av[0] = CPP;
			i = 1;
			if(debug[&#39;.&#39;]){
				sprint(opt, &#34;-.&#34;);
				av[i++] = strdup(opt);
			}
			if(debug[&#39;+&#39;]) {
				sprint(opt, &#34;-+&#34;);
				av[i++] = strdup(opt);
			}
			for(c = 0; c &lt; ndef; c++) {
				sprint(opt, &#34;-D%s&#34;, defs[c]);
				av[i++] = strdup(opt);
			}
			for(c = 0; c &lt; ninclude; c++) {
				sprint(opt, &#34;-I%s&#34;, include[c]);
				av[i++] = strdup(opt);
			}
			if(strcmp(file, &#34;stdin&#34;) != 0)
				av[i++] = file;
			av[i] = 0;
			if(debug[&#39;p&#39;] &gt; 1) {
				for(c = 0; c &lt; i; c++)
					fprint(2, &#34;%s &#34;, av[c]);
				fprint(2, &#34;\n&#34;);
			}
			exec(av[0], av);
			fprint(2, &#34;can&#39;t exec C preprocessor %s: %r\n&#34;, CPP);
			errorexit();
		default:
			close(fd[1]);
			newfile(file, fd[0]);
			break;
		}
	} else {
		if(strcmp(file, &#34;stdin&#34;) == 0)
			newfile(file, 0);
		else
			newfile(file, -1);
	}
	yyparse();
	if(!debug[&#39;a&#39;] &amp;&amp; !debug[&#39;Z&#39;])
		gclean();
	return nerrors;
}

void
errorexit(void)
{
	if(outfile)
		remove(outfile);
	exits(&#34;error&#34;);
}

void
pushio(void)
{
	Io *i;

	i = iostack;
	if(i == I) {
		yyerror(&#34;botch in pushio&#34;);
		errorexit();
	}
	i-&gt;p = fi.p;
	i-&gt;c = fi.c;
}

void
newio(void)
{
	Io *i;
	static int pushdepth = 0;

	i = iofree;
	if(i == I) {
		pushdepth++;
		if(pushdepth &gt; 1000) {
			yyerror(&#34;macro/io expansion too deep&#34;);
			errorexit();
		}
		i = alloc(sizeof(*i));
	} else
		iofree = i-&gt;link;
	i-&gt;c = 0;
	i-&gt;f = -1;
	ionext = i;
}

void
newfile(char *s, int f)
{
	Io *i;

	if(debug[&#39;e&#39;])
		print(&#34;%L: %s\n&#34;, lineno, s);

	i = ionext;
	i-&gt;link = iostack;
	iostack = i;
	i-&gt;f = f;
	if(f &lt; 0)
		i-&gt;f = open(s, 0);
	if(i-&gt;f &lt; 0) {
		yyerror(&#34;%cc: %r: %s&#34;, thechar, s);
		errorexit();
	}
	fi.c = 0;
	linehist(s, 0);
}

Sym*
slookup(char *s)
{

	strcpy(symb, s);
	return lookup();
}

Sym*
lookup(void)
{
	Sym *s;
	uint32 h;
	char *p;
	int c, n;

	h = 0;
	for(p=symb; *p;) {
		h = h * 3;
		h += *p++;
	}
	n = (p - symb) + 1;
	h &amp;= 0xffffff;
	h %= NHASH;
	c = symb[0];
	for(s = hash[h]; s != S; s = s-&gt;link) {
		if(s-&gt;name[0] != c)
			continue;
		if(strcmp(s-&gt;name, symb) == 0)
			return s;
	}
	s = alloc(sizeof(*s));
	s-&gt;name = alloc(n);
	memmove(s-&gt;name, symb, n);

	strcpy(s-&gt;name, symb);
	s-&gt;link = hash[h];
	hash[h] = s;
	syminit(s);

	return s;
}

void
syminit(Sym *s)
{
	s-&gt;lexical = LNAME;
	s-&gt;block = 0;
	s-&gt;offset = 0;
	s-&gt;type = T;
	s-&gt;suetag = T;
	s-&gt;class = CXXX;
	s-&gt;aused = 0;
	s-&gt;sig = SIGNONE;
}

#define	EOF	(-1)
#define	IGN	(-2)
#define	ESC	(1&lt;&lt;20)
#define	GETC()	((--fi.c &lt; 0)? filbuf(): (*fi.p++ &amp; 0xff))

enum
{
	Numdec		= 1&lt;&lt;0,
	Numlong		= 1&lt;&lt;1,
	Numuns		= 1&lt;&lt;2,
	Numvlong	= 1&lt;&lt;3,
	Numflt		= 1&lt;&lt;4,
};

int32
yylex(void)
{
	vlong vv;
	int32 c, c1, t;
	char *cp;
	Rune rune;
	Sym *s;

	if(peekc != IGN) {
		c = peekc;
		peekc = IGN;
		goto l1;
	}
l0:
	c = GETC();

l1:
	if(c &gt;= Runeself) {
		/*
		 * extension --
		 *	all multibyte runes are alpha
		 */
		cp = symb;
		goto talph;
	}
	if(isspace(c)) {
		if(c == &#39;\n&#39;)
			lineno++;
		goto l0;
	}
	if(isalpha(c)) {
		cp = symb;
		if(c != &#39;L&#39;)
			goto talph;
		*cp++ = c;
		c = GETC();
		if(c == &#39;\&#39;&#39;) {
			/* L&#39;x&#39; */
			c = escchar(&#39;\&#39;&#39;, 1, 0);
			if(c == EOF)
				c = &#39;\&#39;&#39;;
			c1 = escchar(&#39;\&#39;&#39;, 1, 0);
			if(c1 != EOF) {
				yyerror(&#34;missing &#39;&#34;);
				peekc = c1;
			}
			yylval.vval = convvtox(c, TUSHORT);
			return LUCONST;
		}
		if(c == &#39;&#34;&#39;) {
			goto caselq;
		}
		goto talph;
	}
	if(isdigit(c))
		goto tnum;
	switch(c)
	{

	case EOF:
		peekc = EOF;
		return -1;

	case &#39;_&#39;:
		cp = symb;
		goto talph;

	case &#39;#&#39;:
		domacro();
		goto l0;

	case &#39;.&#39;:
		c1 = GETC();
		if(isdigit(c1)) {
			cp = symb;
			*cp++ = c;
			c = c1;
			c1 = 0;
			goto casedot;
		}
		break;

	case &#39;&#34;&#39;:
		strcpy(symb, &#34;\&#34;&lt;string&gt;\&#34;&#34;);
		cp = alloc(0);
		c1 = 0;

		/* &#34;...&#34; */
		for(;;) {
			c = escchar(&#39;&#34;&#39;, 0, 1);
			if(c == EOF)
				break;
			if(c &amp; ESC) {
				cp = allocn(cp, c1, 1);
				cp[c1++] = c;
			} else {
				rune = c;
				c = runelen(rune);
				cp = allocn(cp, c1, c);
				runetochar(cp+c1, &amp;rune);
				c1 += c;
			}
		}
		yylval.sval.l = c1;
		do {
			cp = allocn(cp, c1, 1);
			cp[c1++] = 0;
		} while(c1 &amp; MAXALIGN);
		yylval.sval.s = cp;
		return LSTRING;

	caselq:
		/* L&#34;...&#34; */
		strcpy(symb, &#34;\&#34;L&lt;string&gt;\&#34;&#34;);
		cp = alloc(0);
		c1 = 0;
		for(;;) {
			c = escchar(&#39;&#34;&#39;, 1, 0);
			if(c == EOF)
				break;
			cp = allocn(cp, c1, sizeof(ushort));
			*(ushort*)(cp + c1) = c;
			c1 += sizeof(ushort);
		}
		yylval.sval.l = c1;
		do {
			cp = allocn(cp, c1, sizeof(ushort));
			*(ushort*)(cp + c1) = 0;
			c1 += sizeof(ushort);
		} while(c1 &amp; MAXALIGN);
		yylval.sval.s = cp;
		return LLSTRING;

	case &#39;\&#39;&#39;:
		/* &#39;.&#39; */
		c = escchar(&#39;\&#39;&#39;, 0, 0);
		if(c == EOF)
			c = &#39;\&#39;&#39;;
		c1 = escchar(&#39;\&#39;&#39;, 0, 0);
		if(c1 != EOF) {
			yyerror(&#34;missing &#39;&#34;);
			peekc = c1;
		}
		vv = c;
		yylval.vval = convvtox(vv, TUCHAR);
		if(yylval.vval != vv)
			yyerror(&#34;overflow in character constant: 0x%lx&#34;, c);
		else
		if(c &amp; 0x80){
			nearln = lineno;
			warn(Z, &#34;sign-extended character constant&#34;);
		}
		yylval.vval = convvtox(vv, TCHAR);
		return LCONST;

	case &#39;/&#39;:
		c1 = GETC();
		if(c1 == &#39;*&#39;) {
			for(;;) {
				c = getr();
				while(c == &#39;*&#39;) {
					c = getr();
					if(c == &#39;/&#39;)
						goto l0;
				}
				if(c == EOF) {
					yyerror(&#34;eof in comment&#34;);
					errorexit();
				}
			}
		}
		if(c1 == &#39;/&#39;) {
			for(;;) {
				c = getr();
				if(c == &#39;\n&#39;)
					goto l0;
				if(c == EOF) {
					yyerror(&#34;eof in comment&#34;);
					errorexit();
				}
			}
		}
		if(c1 == &#39;=&#39;)
			return LDVE;
		break;

	case &#39;*&#39;:
		c1 = GETC();
		if(c1 == &#39;=&#39;)
			return LMLE;
		break;

	case &#39;%&#39;:
		c1 = GETC();
		if(c1 == &#39;=&#39;)
			return LMDE;
		break;

	case &#39;+&#39;:
		c1 = GETC();
		if(c1 == &#39;+&#39;)
			return LPP;
		if(c1 == &#39;=&#39;)
			return LPE;
		break;

	case &#39;-&#39;:
		c1 = GETC();
		if(c1 == &#39;-&#39;)
			return LMM;
		if(c1 == &#39;=&#39;)
			return LME;
		if(c1 == &#39;&gt;&#39;)
			return LMG;
		break;

	case &#39;&gt;&#39;:
		c1 = GETC();
		if(c1 == &#39;&gt;&#39;) {
			c = LRSH;
			c1 = GETC();
			if(c1 == &#39;=&#39;)
				return LRSHE;
			break;
		}
		if(c1 == &#39;=&#39;)
			return LGE;
		break;

	case &#39;&lt;&#39;:
		c1 = GETC();
		if(c1 == &#39;&lt;&#39;) {
			c = LLSH;
			c1 = GETC();
			if(c1 == &#39;=&#39;)
				return LLSHE;
			break;
		}
		if(c1 == &#39;=&#39;)
			return LLE;
		break;

	case &#39;=&#39;:
		c1 = GETC();
		if(c1 == &#39;=&#39;)
			return LEQ;
		break;

	case &#39;!&#39;:
		c1 = GETC();
		if(c1 == &#39;=&#39;)
			return LNE;
		break;

	case &#39;&amp;&#39;:
		c1 = GETC();
		if(c1 == &#39;&amp;&#39;)
			return LANDAND;
		if(c1 == &#39;=&#39;)
			return LANDE;
		break;

	case &#39;|&#39;:
		c1 = GETC();
		if(c1 == &#39;|&#39;)
			return LOROR;
		if(c1 == &#39;=&#39;)
			return LORE;
		break;

	case &#39;^&#39;:
		c1 = GETC();
		if(c1 == &#39;=&#39;)
			return LXORE;
		break;

	default:
		return c;
	}
	peekc = c1;
	return c;

talph:
	/*
	 * cp is set to symb and some
	 * prefix has been stored
	 */
	for(;;) {
		if(c &gt;= Runeself) {
			for(c1=0;;) {
				cp[c1++] = c;
				if(fullrune(cp, c1))
					break;
				c = GETC();
			}
			cp += c1;
			c = GETC();
			continue;
		}
		if(!isalnum(c) &amp;&amp; c != &#39;_&#39;)
			break;
		*cp++ = c;
		c = GETC();
	}
	*cp = 0;
	if(debug[&#39;L&#39;])
		print(&#34;%L: %s\n&#34;, lineno, symb);
	peekc = c;
	s = lookup();
	if(s-&gt;macro) {
		newio();
		cp = ionext-&gt;b;
		macexpand(s, cp);
		pushio();
		ionext-&gt;link = iostack;
		iostack = ionext;
		fi.p = cp;
		fi.c = strlen(cp);
		if(peekc != IGN) {
			cp[fi.c++] = peekc;
			cp[fi.c] = 0;
			peekc = IGN;
		}
		goto l0;
	}
	yylval.sym = s;
	if(s-&gt;class == CTYPEDEF || s-&gt;class == CTYPESTR)
		return LTYPE;
	return s-&gt;lexical;

tnum:
	c1 = 0;
	cp = symb;
	if(c != &#39;0&#39;) {
		c1 |= Numdec;
		for(;;) {
			*cp++ = c;
			c = GETC();
			if(isdigit(c))
				continue;
			goto dc;
		}
	}
	*cp++ = c;
	c = GETC();
	if(c == &#39;x&#39; || c == &#39;X&#39;)
		for(;;) {
			*cp++ = c;
			c = GETC();
			if(isdigit(c))
				continue;
			if(c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;f&#39;)
				continue;
			if(c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;F&#39;)
				continue;
			if(cp == symb+2)
				yyerror(&#34;malformed hex constant&#34;);
			goto ncu;
		}
	if(c &lt; &#39;0&#39; || c &gt; &#39;7&#39;)
		goto dc;
	for(;;) {
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;7&#39;) {
			*cp++ = c;
			c = GETC();
			continue;
		}
		goto ncu;
	}

dc:
	if(c == &#39;.&#39;)
		goto casedot;
	if(c == &#39;e&#39; || c == &#39;E&#39;)
		goto casee;

ncu:
	if((c == &#39;U&#39; || c == &#39;u&#39;) &amp;&amp; !(c1 &amp; Numuns)) {
		c = GETC();
		c1 |= Numuns;
		goto ncu;
	}
	if((c == &#39;L&#39; || c == &#39;l&#39;) &amp;&amp; !(c1 &amp; Numvlong)) {
		c = GETC();
		if(c1 &amp; Numlong)
			c1 |= Numvlong;
		c1 |= Numlong;
		goto ncu;
	}
	*cp = 0;
	peekc = c;
	if(mpatov(symb, &amp;yylval.vval))
		yyerror(&#34;overflow in constant&#34;);

	vv = yylval.vval;
	if(c1 &amp; Numvlong) {
		if((c1 &amp; Numuns) || convvtox(vv, TVLONG) &lt; 0) {
			c = LUVLCONST;
			t = TUVLONG;
			goto nret;
		}
		c = LVLCONST;
		t = TVLONG;
		goto nret;
	}
	if(c1 &amp; Numlong) {
		if((c1 &amp; Numuns) || convvtox(vv, TLONG) &lt; 0) {
			c = LULCONST;
			t = TULONG;
			goto nret;
		}
		c = LLCONST;
		t = TLONG;
		goto nret;
	}
	if((c1 &amp; Numuns) || convvtox(vv, TINT) &lt; 0) {
		c = LUCONST;
		t = TUINT;
		goto nret;
	}
	c = LCONST;
	t = TINT;
	goto nret;

nret:
	yylval.vval = convvtox(vv, t);
	if(yylval.vval != vv){
		nearln = lineno;
		warn(Z, &#34;truncated constant: %T %s&#34;, types[t], symb);
	}
	return c;

casedot:
	for(;;) {
		*cp++ = c;
		c = GETC();
		if(!isdigit(c))
			break;
	}
	if(c != &#39;e&#39; &amp;&amp; c != &#39;E&#39;)
		goto caseout;

casee:
	*cp++ = &#39;e&#39;;
	c = GETC();
	if(c == &#39;+&#39; || c == &#39;-&#39;) {
		*cp++ = c;
		c = GETC();
	}
	if(!isdigit(c))
		yyerror(&#34;malformed fp constant exponent&#34;);
	while(isdigit(c)) {
		*cp++ = c;
		c = GETC();
	}

caseout:
	if(c == &#39;L&#39; || c == &#39;l&#39;) {
		c = GETC();
		c1 |= Numlong;
	} else
	if(c == &#39;F&#39; || c == &#39;f&#39;) {
		c = GETC();
		c1 |= Numflt;
	}
	*cp = 0;
	peekc = c;
	yylval.dval = strtod(symb, nil);
	if(isInf(yylval.dval, 1) || isInf(yylval.dval, -1)) {
		yyerror(&#34;overflow in float constant&#34;);
		yylval.dval = 0;
	}
	if(c1 &amp; Numflt)
		return LFCONST;
	return LDCONST;
}

/*
 * convert a string, s, to vlong in *v
 * return conversion overflow.
 * required syntax is [0[x]]d*
 */
int
mpatov(char *s, vlong *v)
{
	vlong n, nn;
	int c;

	n = 0;
	c = *s;
	if(c == &#39;0&#39;)
		goto oct;
	while(c = *s++) {
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;)
			nn = n*10 + c-&#39;0&#39;;
		else
			goto bad;
		if(n &lt; 0 &amp;&amp; nn &gt;= 0)
			goto bad;
		n = nn;
	}
	goto out;

oct:
	s++;
	c = *s;
	if(c == &#39;x&#39; || c == &#39;X&#39;)
		goto hex;
	while(c = *s++) {
		if(c &gt;= &#39;0&#39; || c &lt;= &#39;7&#39;)
			nn = n*8 + c-&#39;0&#39;;
		else
			goto bad;
		if(n &lt; 0 &amp;&amp; nn &gt;= 0)
			goto bad;
		n = nn;
	}
	goto out;

hex:
	s++;
	while(c = *s++) {
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;)
			c += 0-&#39;0&#39;;
		else
		if(c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;f&#39;)
			c += 10-&#39;a&#39;;
		else
		if(c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;F&#39;)
			c += 10-&#39;A&#39;;
		else
			goto bad;
		nn = n*16 + c;
		if(n &lt; 0 &amp;&amp; nn &gt;= 0)
			goto bad;
		n = nn;
	}
out:
	*v = n;
	return 0;

bad:
	*v = ~0;
	return 1;
}

int
getc(void)
{
	int c;

	if(peekc != IGN) {
		c = peekc;
		peekc = IGN;
	} else
		c = GETC();
	if(c == &#39;\n&#39;)
		lineno++;
	if(c == EOF) {
		yyerror(&#34;End of file&#34;);
		errorexit();
	}
	return c;
}

int32
getr(void)
{
	int c, i;
	char str[UTFmax+1];
	Rune rune;


	c = getc();
	if(c &lt; Runeself)
		return c;
	i = 0;
	str[i++] = c;

loop:
	c = getc();
	str[i++] = c;
	if(!fullrune(str, i))
		goto loop;
	c = chartorune(&amp;rune, str);
	if(rune == Runeerror &amp;&amp; c == 1) {
		nearln = lineno;
		diag(Z, &#34;illegal rune in string&#34;);
		for(c=0; c&lt;i; c++)
			print(&#34; %.2x&#34;, *(uchar*)(str+c));
		print(&#34;\n&#34;);
	}
	return rune;
}

int
getnsc(void)
{
	int c;

	if(peekc != IGN) {
		c = peekc;
		peekc = IGN;
	} else
		c = GETC();
	for(;;) {
		if(!isspace(c))
			return c;
		if(c == &#39;\n&#39;) {
			lineno++;
			return c;
		}
		c = GETC();
	}
}

void
unget(int c)
{

	peekc = c;
	if(c == &#39;\n&#39;)
		lineno--;
}

int32
escchar(int32 e, int longflg, int escflg)
{
	int32 c, l;
	int i;

loop:
	c = getr();
	if(c == &#39;\n&#39;) {
		yyerror(&#34;newline in string&#34;);
		return EOF;
	}
	if(c != &#39;\\&#39;) {
		if(c == e)
			c = EOF;
		return c;
	}
	c = getr();
	if(c == &#39;x&#39;) {
		/*
		 * note this is not ansi,
		 * supposed to only accept 2 hex
		 */
		i = 2;
		if(longflg)
			i = 4;
		l = 0;
		for(; i&gt;0; i--) {
			c = getc();
			if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) {
				l = l*16 + c-&#39;0&#39;;
				continue;
			}
			if(c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;f&#39;) {
				l = l*16 + c-&#39;a&#39; + 10;
				continue;
			}
			if(c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;F&#39;) {
				l = l*16 + c-&#39;A&#39; + 10;
				continue;
			}
			unget(c);
			break;
		}
		if(escflg)
			l |= ESC;
		return l;
	}
	if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;7&#39;) {
		/*
		 * note this is not ansi,
		 * supposed to only accept 3 oct
		 */
		i = 2;
		if(longflg)
			i = 5;
		l = c - &#39;0&#39;;
		for(; i&gt;0; i--) {
			c = getc();
			if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;7&#39;) {
				l = l*8 + c-&#39;0&#39;;
				continue;
			}
			unget(c);
		}
		if(escflg)
			l |= ESC;
		return l;
	}
	switch(c)
	{
	case &#39;\n&#39;:	goto loop;
	case &#39;n&#39;:	return &#39;\n&#39;;
	case &#39;t&#39;:	return &#39;\t&#39;;
	case &#39;b&#39;:	return &#39;\b&#39;;
	case &#39;r&#39;:	return &#39;\r&#39;;
	case &#39;f&#39;:	return &#39;\f&#39;;
	case &#39;a&#39;:	return &#39;\a&#39;;
	case &#39;v&#39;:	return &#39;\v&#39;;
	}
	return c;
}

struct
{
	char	*name;
	ushort	lexical;
	ushort	type;
} itab[] =
{
	&#34;auto&#34;,		LAUTO,		0,
	&#34;break&#34;,	LBREAK,		0,
	&#34;case&#34;,		LCASE,		0,
	&#34;char&#34;,		LCHAR,		TCHAR,
	&#34;const&#34;,	LCONSTNT,	0,
	&#34;continue&#34;,	LCONTINUE,	0,
	&#34;default&#34;,	LDEFAULT,	0,
	&#34;do&#34;,		LDO,		0,
	&#34;double&#34;,	LDOUBLE,	TDOUBLE,
	&#34;else&#34;,		LELSE,		0,
	&#34;enum&#34;,		LENUM,		0,
	&#34;extern&#34;,	LEXTERN,	0,
	&#34;float&#34;,	LFLOAT,		TFLOAT,
	&#34;for&#34;,		LFOR,		0,
	&#34;goto&#34;,		LGOTO,		0,
	&#34;if&#34;,		LIF,		0,
	&#34;inline&#34;,	LINLINE,	0,
	&#34;int&#34;,		LINT,		TINT,
	&#34;long&#34;,		LLONG,		TLONG,
	&#34;register&#34;,	LREGISTER,	0,
	&#34;restrict&#34;,	LRESTRICT,	0,
	&#34;return&#34;,	LRETURN,	0,
	&#34;SET&#34;,		LSET,		0,
	&#34;short&#34;,	LSHORT,		TSHORT,
	&#34;signed&#34;,	LSIGNED,	0,
	&#34;signof&#34;,	LSIGNOF,	0,
	&#34;sizeof&#34;,	LSIZEOF,	0,
	&#34;static&#34;,	LSTATIC,	0,
	&#34;struct&#34;,	LSTRUCT,	0,
	&#34;switch&#34;,	LSWITCH,	0,
	&#34;typedef&#34;,	LTYPEDEF,	0,
	&#34;typestr&#34;,	LTYPESTR,	0,
	&#34;union&#34;,	LUNION,		0,
	&#34;unsigned&#34;,	LUNSIGNED,	0,
	&#34;USED&#34;,		LUSED,		0,
	&#34;void&#34;,		LVOID,		TVOID,
	&#34;volatile&#34;,	LVOLATILE,	0,
	&#34;while&#34;,	LWHILE,		0,
	0
};

void
cinit(void)
{
	Sym *s;
	int i;
	Type *t;

	nerrors = 0;
	lineno = 1;
	iostack = I;
	iofree = I;
	peekc = IGN;
	nhunk = 0;

	types[TXXX] = T;
	types[TCHAR] = typ(TCHAR, T);
	types[TUCHAR] = typ(TUCHAR, T);
	types[TSHORT] = typ(TSHORT, T);
	types[TUSHORT] = typ(TUSHORT, T);
	types[TINT] = typ(TINT, T);
	types[TUINT] = typ(TUINT, T);
	types[TLONG] = typ(TLONG, T);
	types[TULONG] = typ(TULONG, T);
	types[TVLONG] = typ(TVLONG, T);
	types[TUVLONG] = typ(TUVLONG, T);
	types[TFLOAT] = typ(TFLOAT, T);
	types[TDOUBLE] = typ(TDOUBLE, T);
	types[TVOID] = typ(TVOID, T);
	types[TENUM] = typ(TENUM, T);
	types[TFUNC] = typ(TFUNC, types[TINT]);
	types[TIND] = typ(TIND, types[TVOID]);

	for(i=0; i&lt;NHASH; i++)
		hash[i] = S;
	for(i=0; itab[i].name; i++) {
		s = slookup(itab[i].name);
		s-&gt;lexical = itab[i].lexical;
		if(itab[i].type != 0)
			s-&gt;type = types[itab[i].type];
	}
	blockno = 0;
	autobn = 0;
	autoffset = 0;

	t = typ(TARRAY, types[TCHAR]);
	t-&gt;width = 0;
	symstring = slookup(&#34;.string&#34;);
	symstring-&gt;class = CSTATIC;
	symstring-&gt;type = t;

	t = typ(TARRAY, types[TCHAR]);
	t-&gt;width = 0;

	nodproto = new(OPROTO, Z, Z);
	dclstack = D;

	pathname = allocn(pathname, 0, 100);
	if(getwd(pathname, 99) == 0) {
		pathname = allocn(pathname, 100, 900);
		if(getwd(pathname, 999) == 0)
			strcpy(pathname, &#34;/???&#34;);
	}

	fmtinstall(&#39;O&#39;, Oconv);
	fmtinstall(&#39;T&#39;, Tconv);
	fmtinstall(&#39;F&#39;, FNconv);
	fmtinstall(&#39;L&#39;, Lconv);
	fmtinstall(&#39;Q&#39;, Qconv);
	fmtinstall(&#39;|&#39;, VBconv);
}

int
filbuf(void)
{
	Io *i;

loop:
	i = iostack;
	if(i == I)
		return EOF;
	if(i-&gt;f &lt; 0)
		goto pop;
	fi.c = read(i-&gt;f, i-&gt;b, BUFSIZ) - 1;
	if(fi.c &lt; 0) {
		close(i-&gt;f);
		linehist(0, 0);
		goto pop;
	}
	fi.p = i-&gt;b + 1;
	return i-&gt;b[0] &amp; 0xff;

pop:
	iostack = i-&gt;link;
	i-&gt;link = iofree;
	iofree = i;
	i = iostack;
	if(i == I)
		return EOF;
	fi.p = i-&gt;p;
	fi.c = i-&gt;c;
	if(--fi.c &lt; 0)
		goto loop;
	return *fi.p++ &amp; 0xff;
}

int
Oconv(Fmt *fp)
{
	int a;

	a = va_arg(fp-&gt;args, int);
	if(a &lt; OXXX || a &gt; OEND)
		return fmtprint(fp, &#34;***badO %d***&#34;, a);

	return fmtstrcpy(fp, onames[a]);
}

int
Lconv(Fmt *fp)
{
	char str[STRINGSZ], s[STRINGSZ];
	Hist *h;
	struct
	{
		Hist*	incl;	/* start of this include file */
		int32	idel;	/* delta line number to apply to include */
		Hist*	line;	/* start of this #line directive */
		int32	ldel;	/* delta line number to apply to #line */
	} a[HISTSZ];
	int32 l, d;
	int i, n;

	l = va_arg(fp-&gt;args, int32);
	n = 0;
	for(h = hist; h != H; h = h-&gt;link) {
		if(l &lt; h-&gt;line)
			break;
		if(h-&gt;name) {
			if(h-&gt;offset != 0) {		/* #line directive, not #pragma */
				if(n &gt; 0 &amp;&amp; n &lt; HISTSZ &amp;&amp; h-&gt;offset &gt;= 0) {
					a[n-1].line = h;
					a[n-1].ldel = h-&gt;line - h-&gt;offset + 1;
				}
			} else {
				if(n &lt; HISTSZ) {	/* beginning of file */
					a[n].incl = h;
					a[n].idel = h-&gt;line;
					a[n].line = 0;
				}
				n++;
			}
			continue;
		}
		n--;
		if(n &gt; 0 &amp;&amp; n &lt; HISTSZ) {
			d = h-&gt;line - a[n].incl-&gt;line;
			a[n-1].ldel += d;
			a[n-1].idel += d;
		}
	}
	if(n &gt; HISTSZ)
		n = HISTSZ;
	str[0] = 0;
	for(i=n-1; i&gt;=0; i--) {
		if(i != n-1) {
			if(fp-&gt;flags &amp; ~(FmtWidth|FmtPrec))	/* BUG ROB - was f3 */
				break;
			strcat(str, &#34; &#34;);
		}
		if(a[i].line)
			snprint(s, STRINGSZ, &#34;%s:%ld[%s:%ld]&#34;,
				a[i].line-&gt;name, l-a[i].ldel+1,
				a[i].incl-&gt;name, l-a[i].idel+1);
		else
			snprint(s, STRINGSZ, &#34;%s:%ld&#34;,
				a[i].incl-&gt;name, l-a[i].idel+1);
		if(strlen(s)+strlen(str) &gt;= STRINGSZ-10)
			break;
		strcat(str, s);
		l = a[i].incl-&gt;line - 1;	/* now print out start of this file */
	}
	if(n == 0)
		strcat(str, &#34;&lt;eof&gt;&#34;);
	return fmtstrcpy(fp, str);
}

int
Tconv(Fmt *fp)
{
	char str[STRINGSZ+20], s[STRINGSZ+20];
	Type *t, *t1;
	int et;
	int32 n;

	str[0] = 0;
	for(t = va_arg(fp-&gt;args, Type*); t != T; t = t-&gt;link) {
		et = t-&gt;etype;
		if(str[0])
			strcat(str, &#34; &#34;);
		if(t-&gt;garb&amp;~GINCOMPLETE) {
			sprint(s, &#34;%s &#34;, gnames[t-&gt;garb&amp;~GINCOMPLETE]);
			if(strlen(str) + strlen(s) &lt; STRINGSZ)
				strcat(str, s);
		}
		sprint(s, &#34;%s&#34;, tnames[et]);
		if(strlen(str) + strlen(s) &lt; STRINGSZ)
			strcat(str, s);
		if(et == TFUNC &amp;&amp; (t1 = t-&gt;down)) {
			sprint(s, &#34;(%T&#34;, t1);
			if(strlen(str) + strlen(s) &lt; STRINGSZ)
				strcat(str, s);
			while(t1 = t1-&gt;down) {
				sprint(s, &#34;, %T&#34;, t1);
				if(strlen(str) + strlen(s) &lt; STRINGSZ)
					strcat(str, s);
			}
			if(strlen(str) + strlen(s) &lt; STRINGSZ)
				strcat(str, &#34;)&#34;);
		}
		if(et == TARRAY) {
			n = t-&gt;width;
			if(t-&gt;link &amp;&amp; t-&gt;link-&gt;width)
				n /= t-&gt;link-&gt;width;
			sprint(s, &#34;[%ld]&#34;, n);
			if(strlen(str) + strlen(s) &lt; STRINGSZ)
				strcat(str, s);
		}
		if(t-&gt;nbits) {
			sprint(s, &#34; %d:%d&#34;, t-&gt;shift, t-&gt;nbits);
			if(strlen(str) + strlen(s) &lt; STRINGSZ)
				strcat(str, s);
		}
		if(typesu[et]) {
			if(t-&gt;tag) {
				strcat(str, &#34; &#34;);
				if(strlen(str) + strlen(t-&gt;tag-&gt;name) &lt; STRINGSZ)
					strcat(str, t-&gt;tag-&gt;name);
			} else
				strcat(str, &#34; {}&#34;);
			break;
		}
	}
	return fmtstrcpy(fp, str);
}

int
FNconv(Fmt *fp)
{
	char *str;
	Node *n;

	n = va_arg(fp-&gt;args, Node*);
	str = &#34;&lt;indirect&gt;&#34;;
	if(n != Z &amp;&amp; (n-&gt;op == ONAME || n-&gt;op == ODOT || n-&gt;op == OELEM))
		str = n-&gt;sym-&gt;name;
	return fmtstrcpy(fp, str);
}

int
Qconv(Fmt *fp)
{
	char str[STRINGSZ+20], *s;
	int32 b;
	int i;

	str[0] = 0;
	for(b = va_arg(fp-&gt;args, int32); b;) {
		i = bitno(b);
		if(str[0])
			strcat(str, &#34; &#34;);
		s = qnames[i];
		if(strlen(str) + strlen(s) &gt;= STRINGSZ)
			break;
		strcat(str, s);
		b &amp;= ~(1L &lt;&lt; i);
	}
	return fmtstrcpy(fp, str);
}

int
VBconv(Fmt *fp)
{
	char str[STRINGSZ];
	int i, n, t, pc;

	n = va_arg(fp-&gt;args, int);
	pc = 0;	/* BUG: was printcol */
	i = 0;
	while(pc &lt; n) {
		t = (pc+4) &amp; ~3;
		if(t &lt;= n) {
			str[i++] = &#39;\t&#39;;
			pc = t;
			continue;
		}
		str[i++] = &#39; &#39;;
		pc++;
	}
	str[i] = 0;

	return fmtstrcpy(fp, str);
}

void
setinclude(char *p)
{
	int i;
	char *e;

	while(*p != 0) {
		e = strchr(p, &#39; &#39;);
		if(e != 0)
			*e = &#39;\0&#39;;

		for(i=1; i &lt; ninclude; i++)
			if(strcmp(p, include[i]) == 0)
				break;

		if(i &gt;= ninclude)
			include[ninclude++] = p;

		if(ninclude &gt; nelem(include)) {
			diag(Z, &#34;ninclude too small %d&#34;, nelem(include));
			exits(&#34;ninclude&#34;);
		}

		if(e == 0)
			break;
		p = e+1;
	}
}

void*
alloc(int32 n)
{
	void *p;

	p = malloc(n);
	if(p == nil) {
		print(&#34;alloc out of mem\n&#34;);
		exit(1);
	}
	memset(p, 0, n);
	return p;
}

void*
allocn(void *p, int32 n, int32 d)
{
	if(p == nil)
		return alloc(n+d);
	p = realloc(p, n+d);
	if(p == nil) {
		print(&#34;allocn out of mem\n&#34;);
		exit(1);
	}
	if(d &gt; 0)
		memset((char*)p+n, 0, d);
	return p;
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
