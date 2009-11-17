<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/lex.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/lex.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#define		EXTERN
#include	&#34;go.h&#34;
#include	&#34;y.tab.h&#34;
#include &lt;ar.h&gt;

extern int yychar;
char nopackage[] = &#34;____&#34;;
void lexfini(void);

#define	DBG	if(!debug[&#39;x&#39;]);else print
enum
{
	EOF		= -1,
};

int
main(int argc, char *argv[])
{
	int i, c;
	NodeList *l;

	outfile = nil;
	package = nopackage;
	ARGBEGIN {
	default:
		c = ARGC();
		if(c &gt;= 0 &amp;&amp; c &lt; sizeof(debug))
			debug[c]++;
		break;

	case &#39;o&#39;:
		outfile = ARGF();
		break;

	case &#39;k&#39;:
		package = ARGF();
		break;

	case &#39;I&#39;:
		addidir(ARGF());
		break;
	} ARGEND

	if(argc &lt; 1)
		goto usage;

	pathname = mal(1000);
	if(getwd(pathname, 999) == 0)
		strcpy(pathname, &#34;/???&#34;);

	fmtinstall(&#39;O&#39;, Oconv);		// node opcodes
	fmtinstall(&#39;E&#39;, Econv);		// etype opcodes
	fmtinstall(&#39;J&#39;, Jconv);		// all the node flags
	fmtinstall(&#39;S&#39;, Sconv);		// sym pointer
	fmtinstall(&#39;T&#39;, Tconv);		// type pointer
	fmtinstall(&#39;N&#39;, Nconv);		// node pointer
	fmtinstall(&#39;Z&#39;, Zconv);		// escaped string
	fmtinstall(&#39;L&#39;, Lconv);		// line number
	fmtinstall(&#39;B&#39;, Bconv);		// big numbers
	fmtinstall(&#39;F&#39;, Fconv);		// big float numbers

	betypeinit();
	if(maxround == 0 || widthptr == 0)
		fatal(&#34;betypeinit failed&#34;);

	lexinit();
	typeinit();

	blockgen = 1;
	dclcontext = PEXTERN;
	nerrors = 0;
	lexlineno = 1;

	for(i=0; i&lt;argc; i++) {
		infile = argv[i];
		linehist(infile, 0, 0);

		curio.infile = infile;
		curio.bin = Bopen(infile, OREAD);
		if(curio.bin == nil)
			fatal(&#34;open %s: %r&#34;, infile);
		curio.peekc = 0;
		curio.peekc1 = 0;

		block = 1;

		yyparse();
		if(nsyntaxerrors != 0)
			errorexit();

		linehist(nil, 0, 0);
		if(curio.bin != nil)
			Bterm(curio.bin);
	}
	testdclstack();
	mkpackage(package);	// final import not used checks
	lexfini();

	typecheckok = 1;
	if(debug[&#39;f&#39;])
		frame(1);
	defercheckwidth();
	typechecklist(xtop, Etop);
	resumecheckwidth();
	for(l=xtop; l; l=l-&gt;next)
		if(l-&gt;n-&gt;op == ODCLFUNC)
			funccompile(l-&gt;n);
	if(nerrors == 0)
		fninit(xtop);
	while(closures) {
		l = closures;
		closures = nil;
		for(; l; l=l-&gt;next)
			funccompile(l-&gt;n);
	}
	dclchecks();

	runifacechecks();
	if(nerrors)
		errorexit();

	dumpobj();

	if(nerrors)
		errorexit();

	exit(0);
	return 0;

usage:
	print(&#34;flags:\n&#34;);
	// -A is allow use of &#34;any&#34; type, for bootstrapping
	print(&#34;  -I DIR search for packages in DIR\n&#34;);
	print(&#34;  -d print declarations\n&#34;);
	print(&#34;  -e no limit on number of errors printed\n&#34;);
	print(&#34;  -f print stack frame structure\n&#34;);
	print(&#34;  -h panic on an error\n&#34;);
	print(&#34;  -k name specify package name\n&#34;);
	print(&#34;  -o file specify output file\n&#34;);
	print(&#34;  -S print the assembly language\n&#34;);
	print(&#34;  -w print the parse tree after typing\n&#34;);
	print(&#34;  -x print lex tokens\n&#34;);
	exit(0);
	return 0;
}

int
arsize(Biobuf *b, char *name)
{
	struct ar_hdr *a;

	if((a = Brdline(b, &#39;\n&#39;)) == nil)
		return -1;
	if(Blinelen(b) != sizeof(struct ar_hdr))
		return -1;
	if(strncmp(a-&gt;name, name, strlen(name)) != 0)
		return -1;
	return atoi(a-&gt;size);
}

int
skiptopkgdef(Biobuf *b)
{
	char *p;
	int sz;

	/* archive header */
	if((p = Brdline(b, &#39;\n&#39;)) == nil)
		return 0;
	if(Blinelen(b) != 8)
		return 0;
	if(memcmp(p, &#34;!&lt;arch&gt;\n&#34;, 8) != 0)
		return 0;
	/* symbol table is first; skip it */
	sz = arsize(b, &#34;__.SYMDEF&#34;);
	if(sz &lt; 0)
		return 0;
	Bseek(b, sz, 1);
	/* package export block is second */
	sz = arsize(b, &#34;__.PKGDEF&#34;);
	if(sz &lt;= 0)
		return 0;
	return 1;
}

void
addidir(char* dir)
{
	Idir** pp;

	if(dir == nil)
		return;

	for(pp = &amp;idirs; *pp != nil; pp = &amp;(*pp)-&gt;link)
		;
	*pp = mal(sizeof(Idir));
	(*pp)-&gt;link = nil;
	(*pp)-&gt;dir = dir;
}

// is this path a local name?  begins with ./ or ../ or /
int
islocalname(Strlit *name)
{
	if(name-&gt;len &gt;= 1 &amp;&amp; name-&gt;s[0] == &#39;/&#39;)
		return 1;
	if(name-&gt;len &gt;= 2 &amp;&amp; strncmp(name-&gt;s, &#34;./&#34;, 2) == 0)
		return 1;
	if(name-&gt;len &gt;= 3 &amp;&amp; strncmp(name-&gt;s, &#34;../&#34;, 3) == 0)
		return 1;
	return 0;
}

int
findpkg(Strlit *name)
{
	static char *goroot, *goos, *goarch;
	Idir *p;

	if(goroot == nil) {
		goroot = getenv(&#34;GOROOT&#34;);
		goos = getenv(&#34;GOOS&#34;);
		goarch = getenv(&#34;GOARCH&#34;);
	}

	if(islocalname(name)) {
		// try .a before .6.  important for building libraries:
		// if there is an array.6 in the array.a library,
		// want to find all of array.a, not just array.6.
		snprint(namebuf, sizeof(namebuf), &#34;%Z.a&#34;, name);
		if(access(namebuf, 0) &gt;= 0)
			return 1;
		snprint(namebuf, sizeof(namebuf), &#34;%Z.%c&#34;, name, thechar);
		if(access(namebuf, 0) &gt;= 0)
			return 1;
		return 0;
	}

	for(p = idirs; p != nil; p = p-&gt;link) {
		snprint(namebuf, sizeof(namebuf), &#34;%s/%Z.a&#34;, p-&gt;dir, name);
		if(access(namebuf, 0) &gt;= 0)
			return 1;
		snprint(namebuf, sizeof(namebuf), &#34;%s/%Z.%c&#34;, p-&gt;dir, name, thechar);
		if(access(namebuf, 0) &gt;= 0)
			return 1;
	}
	if(goroot != nil) {
		snprint(namebuf, sizeof(namebuf), &#34;%s/pkg/%s_%s/%Z.a&#34;, goroot, goos, goarch, name);
		if(access(namebuf, 0) &gt;= 0)
			return 1;
		snprint(namebuf, sizeof(namebuf), &#34;%s/pkg/%s_%s/%Z.%c&#34;, goroot, goos, goarch, name, thechar);
		if(access(namebuf, 0) &gt;= 0)
			return 1;
	}
	return 0;
}

void
importfile(Val *f, int line)
{
	Biobuf *imp;
	char *file, *p;
	int32 c;
	int len;

	// TODO(rsc): don&#39;t bother reloading imports more than once

	if(f-&gt;ctype != CTSTR) {
		yyerror(&#34;import statement not a string&#34;);
		return;
	}

	if(strcmp(f-&gt;u.sval-&gt;s, &#34;unsafe&#34;) == 0) {
		cannedimports(&#34;unsafe.6&#34;, unsafeimport);
		return;
	}

	if(!findpkg(f-&gt;u.sval))
		fatal(&#34;can&#39;t find import: %Z&#34;, f-&gt;u.sval);
	imp = Bopen(namebuf, OREAD);
	if(imp == nil)
		fatal(&#34;can&#39;t open import: %Z&#34;, f-&gt;u.sval);
	file = strdup(namebuf);

	len = strlen(namebuf);
	if(len &gt; 2 &amp;&amp; namebuf[len-2] == &#39;.&#39; &amp;&amp; namebuf[len-1] == &#39;a&#39;) {
		if(!skiptopkgdef(imp))
			fatal(&#34;import not package file: %s&#34;, namebuf);

		// assume .a files move (get installed)
		// so don&#39;t record the full path.
		p = file + len - f-&gt;u.sval-&gt;len - 2;
		linehist(p, -1, 1);	// acts as #pragma lib
	} else {
		// assume .6 files don&#39;t move around
		// so do record the full path
		linehist(file, -1, 0);
	}

	/*
	 * position the input right
	 * after $$ and return
	 */
	pushedio = curio;
	curio.bin = imp;
	curio.peekc = 0;
	curio.peekc1 = 0;
	curio.infile = file;
	typecheckok = 1;
	for(;;) {
		c = getc();
		if(c == EOF)
			break;
		if(c != &#39;$&#39;)
			continue;
		c = getc();
		if(c == EOF)
			break;
		if(c != &#39;$&#39;)
			continue;
		return;
	}
	yyerror(&#34;no import in: %Z&#34;, f-&gt;u.sval);
	unimportfile();
}

void
unimportfile(void)
{
	if(curio.bin != nil) {
		Bterm(curio.bin);
		curio.bin = nil;
	} else
		lexlineno--;	// re correct sys.6 line number

	curio = pushedio;
	pushedio.bin = nil;
	incannedimport = 0;
	typecheckok = 0;
}

void
cannedimports(char *file, char *cp)
{
	lexlineno++;		// if sys.6 is included on line 1,

	pushedio = curio;
	curio.bin = nil;
	curio.peekc = 0;
	curio.peekc1 = 0;
	curio.infile = file;
	curio.cp = cp;

	pkgmyname = S;
	typecheckok = 1;
	incannedimport = 1;
}

int
isfrog(int c)
{
	// complain about possibly invisible control characters
	if(c &lt; 0)
		return 1;
	if(c &lt; &#39; &#39;) {
		if(c == &#39;\n&#39; || c== &#39;\r&#39; || c == &#39;\t&#39;)	// good white space
			return 0;
		return 1;
	}
	if(0x80 &lt;= c &amp;&amp; c &lt;= 0xa0)	// unicode block including unbreakable space.
		return 1;
	return 0;
}

static int32
_yylex(void)
{
	int c, c1, clen, escflag;
	vlong v;
	char *cp;
	Rune rune;
	Sym *s;

	prevlineno = lineno;

l0:
	c = getc();
	if(isspace(c))
		goto l0;

	lineno = lexlineno;	/* start of token */

	if(c &gt;= Runeself) {
		/* all multibyte runes are alpha */
		cp = lexbuf;
		goto talph;
	}

	if(isalpha(c)) {
		cp = lexbuf;
		goto talph;
	}

	if(isdigit(c))
		goto tnum;

	switch(c) {
	case EOF:
		lineno = prevlineno;
		ungetc(EOF);
		return -1;

	case &#39;_&#39;:
		cp = lexbuf;
		goto talph;

	case &#39;.&#39;:
		c1 = getc();
		if(isdigit(c1)) {
			cp = lexbuf;
			*cp++ = c;
			c = c1;
			c1 = 0;
			goto casedot;
		}
		if(c1 == &#39;.&#39;) {
			c1 = getc();
			if(c1 == &#39;.&#39;) {
				c = LDDD;
				goto lx;
			}
			ungetc(c1);
			c1 = &#39;.&#39;;
		}
		break;

	case &#39;&#34;&#39;:
		/* &#34;...&#34; */
		strcpy(lexbuf, &#34;\&#34;&lt;string&gt;\&#34;&#34;);
		cp = mal(sizeof(int32));
		clen = sizeof(int32);

	caseq:
		for(;;) {
			if(escchar(&#39;&#34;&#39;, &amp;escflag, &amp;v))
				break;
			if(v &lt; Runeself || escflag) {
				cp = remal(cp, clen, 1);
				cp[clen++] = v;
			} else {
				// botch - this limits size of runes
				rune = v;
				c = runelen(rune);
				cp = remal(cp, clen, c);
				runetochar(cp+clen, &amp;rune);
				clen += c;
			}
		}
		goto catem;

	case &#39;`&#39;:
		/* `...` */
		strcpy(lexbuf, &#34;`&lt;string&gt;`&#34;);
		cp = mal(sizeof(int32));
		clen = sizeof(int32);

	casebq:
		for(;;) {
			c = getc();
			if(c == EOF) {
				yyerror(&#34;eof in string&#34;);
				break;
			}
			if(c == &#39;`&#39;)
				break;
			cp = remal(cp, clen, 1);
			cp[clen++] = c;
		}
		goto catem;

	catem:
		c = getc();
		if(isspace(c))
			goto catem;

		// skip comments
		if(c == &#39;/&#39;) {
			c1 = getc();
			if(c1 == &#39;*&#39;) {
				for(;;) {
					c = getr();
					while(c == &#39;*&#39;) {
						c = getr();
						if(c == &#39;/&#39;)
							goto catem;
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
						goto catem;
					if(c == EOF) {
						yyerror(&#34;eof in comment&#34;);
						errorexit();
					}
				}
			}
			ungetc(c1);
		}

		// cat adjacent strings
		if(c == &#39;&#34;&#39;)
			goto caseq;
		if(c == &#39;`&#39;)
			goto casebq;
		ungetc(c);

		*(int32*)cp = clen-sizeof(int32);	// length
		do {
			cp = remal(cp, clen, 1);
			cp[clen++] = 0;
		} while(clen &amp; MAXALIGN);
		yylval.val.u.sval = (Strlit*)cp;
		yylval.val.ctype = CTSTR;
		DBG(&#34;lex: string literal\n&#34;);
		return LLITERAL;

	case &#39;\&#39;&#39;:
		/* &#39;.&#39; */
		if(escchar(&#39;\&#39;&#39;, &amp;escflag, &amp;v)) {
			yyerror(&#34;empty character literal or unescaped &#39; in character literal&#34;);
			v = &#39;\&#39;&#39;;
		}
		if(!escchar(&#39;\&#39;&#39;, &amp;escflag, &amp;v)) {
			yyerror(&#34;missing &#39;&#34;);
			ungetc(v);
		}
		yylval.val.u.xval = mal(sizeof(*yylval.val.u.xval));
		mpmovecfix(yylval.val.u.xval, v);
		yylval.val.ctype = CTINT;
		DBG(&#34;lex: codepoint literal\n&#34;);
		return LLITERAL;

	case &#39;/&#39;:
		c1 = getc();
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
		if(c1 == &#39;=&#39;) {
			c = ODIV;
			goto asop;
		}
		break;

	case &#39;:&#39;:
		c1 = getc();
		if(c1 == &#39;=&#39;) {
			c = LCOLAS;
			goto lx;
		}
		break;

	case &#39;*&#39;:
		c1 = getc();
		if(c1 == &#39;=&#39;) {
			c = OMUL;
			goto asop;
		}
		break;

	case &#39;%&#39;:
		c1 = getc();
		if(c1 == &#39;=&#39;) {
			c = OMOD;
			goto asop;
		}
		break;

	case &#39;+&#39;:
		c1 = getc();
		if(c1 == &#39;+&#39;) {
			c = LINC;
			goto lx;
		}
		if(c1 == &#39;=&#39;) {
			c = OADD;
			goto asop;
		}
		break;

	case &#39;-&#39;:
		c1 = getc();
		if(c1 == &#39;-&#39;) {
			c = LDEC;
			goto lx;
		}
		if(c1 == &#39;=&#39;) {
			c = OSUB;
			goto asop;
		}
		break;

	case &#39;&gt;&#39;:
		c1 = getc();
		if(c1 == &#39;&gt;&#39;) {
			c = LRSH;
			c1 = getc();
			if(c1 == &#39;=&#39;) {
				c = ORSH;
				goto asop;
			}
			break;
		}
		if(c1 == &#39;=&#39;) {
			c = LGE;
			goto lx;
		}
		c = LGT;
		break;

	case &#39;&lt;&#39;:
		c1 = getc();
		if(c1 == &#39;&lt;&#39;) {
			c = LLSH;
			c1 = getc();
			if(c1 == &#39;=&#39;) {
				c = OLSH;
				goto asop;
			}
			break;
		}
		if(c1 == &#39;=&#39;) {
			c = LLE;
			goto lx;
		}
		if(c1 == &#39;-&#39;) {
			c = LCOMM;
			goto lx;
		}
		c = LLT;
		break;

	case &#39;=&#39;:
		c1 = getc();
		if(c1 == &#39;=&#39;) {
			c = LEQ;
			goto lx;
		}
		break;

	case &#39;!&#39;:
		c1 = getc();
		if(c1 == &#39;=&#39;) {
			c = LNE;
			goto lx;
		}
		break;

	case &#39;&amp;&#39;:
		c1 = getc();
		if(c1 == &#39;&amp;&#39;) {
			c = LANDAND;
			goto lx;
		}
		if(c1 == &#39;^&#39;) {
			c = LANDNOT;
			c1 = getc();
			if(c1 == &#39;=&#39;) {
				c = OANDNOT;
				goto asop;
			}
			break;
		}
		if(c1 == &#39;=&#39;) {
			c = OAND;
			goto asop;
		}
		break;

	case &#39;|&#39;:
		c1 = getc();
		if(c1 == &#39;|&#39;) {
			c = LOROR;
			goto lx;
		}
		if(c1 == &#39;=&#39;) {
			c = OOR;
			goto asop;
		}
		break;

	case &#39;^&#39;:
		c1 = getc();
		if(c1 == &#39;=&#39;) {
			c = OXOR;
			goto asop;
		}
		break;

	/*
	 * clumsy dance:
	 * to implement rule that disallows
	 *	if T{1}[0] { ... }
	 * but allows
	 * 	if (T{1}[0]) { ... }
	 * the block bodies for if/for/switch/select
	 * begin with an LBODY token, not &#39;{&#39;.
	 *
	 * when we see the keyword, the next
	 * non-parenthesized &#39;{&#39; becomes an LBODY.
	 * loophack is normally 0.
	 * a keyword makes it go up to 1.
	 * parens increment and decrement when loophack &gt; 0.
	 * a &#39;{&#39; with loophack == 1 becomes LBODY and disables loophack.
	 *
	 * i said it was clumsy.
	 */
	case &#39;(&#39;:
		if(loophack &gt; 0)
			loophack++;
		goto lx;
	case &#39;)&#39;:
		if(loophack &gt; 0)
			loophack--;
		goto lx;
	case &#39;{&#39;:
		if(loophack == 1) {
			DBG(&#34;%L lex: LBODY\n&#34;, lexlineno);
			loophack = 0;
			return LBODY;
		}
		goto lx;

	default:
		goto lx;
	}
	ungetc(c1);

lx:
	if(c &gt; 0xff)
		DBG(&#34;%L lex: TOKEN %s\n&#34;, lexlineno, lexname(c));
	else
		DBG(&#34;%L lex: TOKEN &#39;%c&#39;\n&#34;, lexlineno, c);
	if(isfrog(c)) {
		yyerror(&#34;illegal character 0x%ux&#34;, c);
		goto l0;
	}
	return c;

asop:
	yylval.lint = c;	// rathole to hold which asop
	DBG(&#34;lex: TOKEN ASOP %c\n&#34;, c);
	return LASOP;

talph:
	/*
	 * cp is set to lexbuf and some
	 * prefix has been stored
	 */
	for(;;) {
		if(c &gt;= Runeself) {
			for(c1=0;;) {
				cp[c1++] = c;
				if(fullrune(cp, c1)) {
					chartorune(&amp;rune, cp);
					if(isfrog(rune)) {
						yyerror(&#34;illegal character 0x%ux&#34;, rune);
						goto l0;
					}
					// 0xb7 · is used for internal names
					if(!isalpharune(rune) &amp;&amp; !isdigitrune(rune) &amp;&amp; rune != 0xb7)
						yyerror(&#34;invalid identifier character 0x%ux&#34;, rune);
					break;
				}
				c = getc();
			}
			cp += c1;
			c = getc();
			continue;
		}
		if(!isalnum(c) &amp;&amp; c != &#39;_&#39;)
			break;
		*cp++ = c;
		c = getc();
	}
	*cp = 0;
	ungetc(c);

	s = lookup(lexbuf);
	switch(s-&gt;lexical) {
	case LIGNORE:
		goto l0;

	case LFOR:
	case LIF:
	case LSWITCH:
	case LSELECT:
		loophack = 1;	// see comment about loophack above
		break;
	}

	DBG(&#34;lex: %S %s\n&#34;, s, lexname(s-&gt;lexical));
	yylval.sym = s;
	return s-&gt;lexical;

tnum:
	c1 = 0;
	cp = lexbuf;
	if(c != &#39;0&#39;) {
		for(;;) {
			*cp++ = c;
			c = getc();
			if(isdigit(c))
				continue;
			goto dc;
		}
	}
	*cp++ = c;
	c = getc();
	if(c == &#39;x&#39; || c == &#39;X&#39;) {
		for(;;) {
			*cp++ = c;
			c = getc();
			if(isdigit(c))
				continue;
			if(c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;f&#39;)
				continue;
			if(c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;F&#39;)
				continue;
			if(cp == lexbuf+2)
				yyerror(&#34;malformed hex constant&#34;);
			goto ncu;
		}
	}

	if(c == &#39;p&#39;)	// 0p begins floating point zero
		goto casep;

	c1 = 0;
	for(;;) {
		if(!isdigit(c))
			break;
		if(c &lt; &#39;0&#39; || c &gt; &#39;7&#39;)
			c1 = 1;		// not octal
		*cp++ = c;
		c = getc();
	}
	if(c == &#39;.&#39;)
		goto casedot;
	if(c == &#39;e&#39; || c == &#39;E&#39;)
		goto casee;
	if(c1)
		yyerror(&#34;malformed octal constant&#34;);
	goto ncu;

dc:
	if(c == &#39;.&#39;)
		goto casedot;
	if(c == &#39;e&#39; || c == &#39;E&#39;)
		goto casee;
	if(c == &#39;p&#39; || c == &#39;P&#39;)
		goto casep;

ncu:
	*cp = 0;
	ungetc(c);

	yylval.val.u.xval = mal(sizeof(*yylval.val.u.xval));
	mpatofix(yylval.val.u.xval, lexbuf);
	if(yylval.val.u.xval-&gt;ovf) {
		yyerror(&#34;overflow in constant&#34;);
		mpmovecfix(yylval.val.u.xval, 0);
	}
	yylval.val.ctype = CTINT;
	DBG(&#34;lex: integer literal\n&#34;);
	return LLITERAL;

casedot:
	for(;;) {
		*cp++ = c;
		c = getc();
		if(!isdigit(c))
			break;
	}
	if(c != &#39;e&#39; &amp;&amp; c != &#39;E&#39;)
		goto caseout;

casee:
	*cp++ = &#39;e&#39;;
	c = getc();
	if(c == &#39;+&#39; || c == &#39;-&#39;) {
		*cp++ = c;
		c = getc();
	}
	if(!isdigit(c))
		yyerror(&#34;malformed fp constant exponent&#34;);
	while(isdigit(c)) {
		*cp++ = c;
		c = getc();
	}
	goto caseout;

casep:
	*cp++ = &#39;p&#39;;
	c = getc();
	if(c == &#39;+&#39; || c == &#39;-&#39;) {
		*cp++ = c;
		c = getc();
	}
	if(!isdigit(c))
		yyerror(&#34;malformed fp constant exponent&#34;);
	while(isdigit(c)) {
		*cp++ = c;
		c = getc();
	}
	goto caseout;

caseout:
	*cp = 0;
	ungetc(c);

	yylval.val.u.fval = mal(sizeof(*yylval.val.u.fval));
	mpatoflt(yylval.val.u.fval, lexbuf);
	if(yylval.val.u.fval-&gt;val.ovf) {
		yyerror(&#34;overflow in float constant&#34;);
		mpmovecflt(yylval.val.u.fval, 0.0);
	}
	yylval.val.ctype = CTFLT;
	DBG(&#34;lex: floating literal\n&#34;);
	return LLITERAL;
}

/*
 * help the parser.  if the next token is not c and not &#39;;&#39;,
 * insert a &#39;;&#39; before it.
 */
void
yyoptsemi(int c)
{
	if(c == 0)
		c = -1;
	if(yychar &lt;= 0)
		yysemi = c;
}

int32
yylex(void)
{
	// if we delayed a token, return that one.
	if(yynext) {
		yylast = yynext;
		yynext = 0;
		return yylast;
	}

	yylast = _yylex();

	// if there&#39;s an optional semicolon needed,
	// delay the token we just read.
	if(yysemi) {
		if(yylast != &#39;;&#39; &amp;&amp; yylast != yysemi) {
			yynext = yylast;
			yylast = &#39;;&#39;;
		}
		yysemi = 0;
	}

	return yylast;
}

int
getc(void)
{
	int c;

	c = curio.peekc;
	if(c != 0) {
		curio.peekc = curio.peekc1;
		curio.peekc1 = 0;
		if(c == &#39;\n&#39; &amp;&amp; pushedio.bin == nil)
			lexlineno++;
		return c;
	}

	if(curio.bin == nil) {
		c = *curio.cp &amp; 0xff;
		if(c != 0)
			curio.cp++;
	} else
		c = Bgetc(curio.bin);

	switch(c) {
	case 0:
		if(curio.bin != nil)
			break;
	case EOF:
		return EOF;

	case &#39;\n&#39;:
		if(pushedio.bin == nil)
			lexlineno++;
		break;
	}
	return c;
}

void
ungetc(int c)
{
	curio.peekc1 = curio.peekc;
	curio.peekc = c;
	if(c == &#39;\n&#39; &amp;&amp; pushedio.bin == nil)
		lexlineno--;
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
		yyerror(&#34;illegal rune in string&#34;);
		for(c=0; c&lt;i; c++)
			print(&#34; %.2x&#34;, *(uchar*)(str+c));
		print(&#34;\n&#34;);
	}
	return rune;
}


int
escchar(int e, int *escflg, vlong *val)
{
	int i, c;
	vlong l;

	*escflg = 0;

loop:
	c = getr();
	switch(c) {
	case EOF:
		yyerror(&#34;eof in string&#34;);
		return 1;
	case &#39;\n&#39;:
		yyerror(&#34;newline in string&#34;);
		return 1;
	case &#39;\\&#39;:
		break;
	default:
		if(c == e)
			return 1;
		*val = c;
		return 0;
	}

	c = getr();
	switch(c) {
	case &#39;\n&#39;:
		goto loop;

	case &#39;x&#39;:
		*escflg = 1;	// it&#39;s a byte
		i = 2;
		goto hex;

	case &#39;u&#39;:
		i = 4;
		goto hex;

	case &#39;U&#39;:
		i = 8;
		goto hex;

	case &#39;0&#39;:
	case &#39;1&#39;:
	case &#39;2&#39;:
	case &#39;3&#39;:
	case &#39;4&#39;:
	case &#39;5&#39;:
	case &#39;6&#39;:
	case &#39;7&#39;:
		*escflg = 1;	// it&#39;s a byte
		goto oct;

	case &#39;a&#39;: c = &#39;\a&#39;; break;
	case &#39;b&#39;: c = &#39;\b&#39;; break;
	case &#39;f&#39;: c = &#39;\f&#39;; break;
	case &#39;n&#39;: c = &#39;\n&#39;; break;
	case &#39;r&#39;: c = &#39;\r&#39;; break;
	case &#39;t&#39;: c = &#39;\t&#39;; break;
	case &#39;v&#39;: c = &#39;\v&#39;; break;
	case &#39;\\&#39;: c = &#39;\\&#39;; break;

	default:
		if(c != e)
			yyerror(&#34;unknown escape sequence: %c&#34;, c);
	}
	*val = c;
	return 0;

hex:
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
		yyerror(&#34;non-hex character in escape sequence: %c&#34;, c);
		ungetc(c);
		break;
	}
	*val = l;
	return 0;

oct:
	l = c - &#39;0&#39;;
	for(i=2; i&gt;0; i--) {
		c = getc();
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;7&#39;) {
			l = l*8 + c-&#39;0&#39;;
			continue;
		}
		yyerror(&#34;non-oct character in escape sequence: %c&#34;, c);
		ungetc(c);
	}
	if(l &gt; 255)
		yyerror(&#34;oct escape value &gt; 255: %d&#34;, l);

	*val = l;
	return 0;
}

static	struct
{
	char*	name;
	int	lexical;
	int	etype;
	int	op;
} syms[] =
{
/*	name		lexical		etype		op
 */
/* basic types */
	&#34;int8&#34;,		LNAME,		TINT8,		OXXX,
	&#34;int16&#34;,	LNAME,		TINT16,		OXXX,
	&#34;int32&#34;,	LNAME,		TINT32,		OXXX,
	&#34;int64&#34;,	LNAME,		TINT64,		OXXX,

	&#34;uint8&#34;,	LNAME,		TUINT8,		OXXX,
	&#34;uint16&#34;,	LNAME,		TUINT16,	OXXX,
	&#34;uint32&#34;,	LNAME,		TUINT32,	OXXX,
	&#34;uint64&#34;,	LNAME,		TUINT64,	OXXX,

	&#34;float32&#34;,	LNAME,		TFLOAT32,	OXXX,
	&#34;float64&#34;,	LNAME,		TFLOAT64,	OXXX,

	&#34;bool&#34;,		LNAME,		TBOOL,		OXXX,
	&#34;byte&#34;,		LNAME,		TUINT8,		OXXX,
	&#34;string&#34;,	LNAME,		TSTRING,	OXXX,

	&#34;any&#34;,		LNAME,		TANY,		OXXX,

	&#34;break&#34;,	LBREAK,		Txxx,		OXXX,
	&#34;case&#34;,		LCASE,		Txxx,		OXXX,
	&#34;chan&#34;,		LCHAN,		Txxx,		OXXX,
	&#34;const&#34;,	LCONST,		Txxx,		OXXX,
	&#34;continue&#34;,	LCONTINUE,	Txxx,		OXXX,
	&#34;default&#34;,	LDEFAULT,	Txxx,		OXXX,
	&#34;else&#34;,		LELSE,		Txxx,		OXXX,
	&#34;defer&#34;,	LDEFER,		Txxx,		OXXX,
	&#34;fallthrough&#34;,	LFALL,		Txxx,		OXXX,
	&#34;for&#34;,		LFOR,		Txxx,		OXXX,
	&#34;func&#34;,		LFUNC,		Txxx,		OXXX,
	&#34;go&#34;,		LGO,		Txxx,		OXXX,
	&#34;goto&#34;,		LGOTO,		Txxx,		OXXX,
	&#34;if&#34;,		LIF,		Txxx,		OXXX,
	&#34;import&#34;,	LIMPORT,	Txxx,		OXXX,
	&#34;interface&#34;,	LINTERFACE,	Txxx,		OXXX,
	&#34;map&#34;,		LMAP,		Txxx,		OXXX,
	&#34;package&#34;,	LPACKAGE,	Txxx,		OXXX,
	&#34;range&#34;,	LRANGE,		Txxx,		OXXX,
	&#34;return&#34;,	LRETURN,	Txxx,		OXXX,
	&#34;select&#34;,	LSELECT,	Txxx,		OXXX,
	&#34;struct&#34;,	LSTRUCT,	Txxx,		OXXX,
	&#34;switch&#34;,	LSWITCH,	Txxx,		OXXX,
	&#34;type&#34;,		LTYPE,		Txxx,		OXXX,
	&#34;var&#34;,		LVAR,		Txxx,		OXXX,

	&#34;cap&#34;,		LNAME,		Txxx,		OCAP,
	&#34;close&#34;,	LNAME,		Txxx,		OCLOSE,
	&#34;closed&#34;,	LNAME,		Txxx,		OCLOSED,
	&#34;len&#34;,		LNAME,		Txxx,		OLEN,
	&#34;make&#34;,		LNAME,		Txxx,		OMAKE,
	&#34;new&#34;,		LNAME,		Txxx,		ONEW,
	&#34;panic&#34;,	LNAME,		Txxx,		OPANIC,
	&#34;panicln&#34;,	LNAME,		Txxx,		OPANICN,
	&#34;print&#34;,	LNAME,		Txxx,		OPRINT,
	&#34;println&#34;,	LNAME,		Txxx,		OPRINTN,

	&#34;notwithstanding&#34;,		LIGNORE,	Txxx,		OXXX,
	&#34;thetruthofthematter&#34;,		LIGNORE,	Txxx,		OXXX,
	&#34;despiteallobjections&#34;,		LIGNORE,	Txxx,		OXXX,
	&#34;whereas&#34;,			LIGNORE,	Txxx,		OXXX,
	&#34;insofaras&#34;,			LIGNORE,	Txxx,		OXXX,
};

void
lexinit(void)
{
	int i, lex;
	Sym *s, *s1;
	Type *t;
	int etype;

	/*
	 * initialize basic types array
	 * initialize known symbols
	 */
	for(i=0; i&lt;nelem(syms); i++) {
		lex = syms[i].lexical;
		s = lookup(syms[i].name);
		s-&gt;lexical = lex;

		etype = syms[i].etype;
		if(etype != Txxx) {
			if(etype &lt; 0 || etype &gt;= nelem(types))
				fatal(&#34;lexinit: %s bad etype&#34;, s-&gt;name);
			t = types[etype];
			if(t == T) {
				t = typ(etype);
				t-&gt;sym = s;

				if(etype != TANY &amp;&amp; etype != TSTRING)
					dowidth(t);
				types[etype] = t;
			}
			s1 = pkglookup(syms[i].name, &#34;/builtin/&#34;);	// impossible pkg name for builtins
			s1-&gt;lexical = LNAME;
			s1-&gt;def = typenod(t);
			continue;
		}
	}

	s = lookup(&#34;iota&#34;);
	s-&gt;def = nod(ONONAME, N, N);
	s-&gt;def-&gt;iota = 1;
	s-&gt;def-&gt;sym = s;

	// logically, the type of a string literal.
	// types[TSTRING] is the named type string
	// (the type of x in var x string or var x = &#34;hello&#34;).
	// this is the ideal form
	// (the type of x in const x = &#34;hello&#34;).
	// TODO(rsc): this may need some more thought.
	idealstring = typ(TSTRING);
	idealbool = typ(TBOOL);

	s = pkglookup(&#34;true&#34;, &#34;/builtin/&#34;);
	s-&gt;def = nodbool(1);
	s-&gt;def-&gt;sym = lookup(&#34;true&#34;);
	s-&gt;def-&gt;type = idealbool;

	s = pkglookup(&#34;false&#34;, &#34;/builtin/&#34;);
	s-&gt;def = nodbool(0);
	s-&gt;def-&gt;sym = lookup(&#34;false&#34;);
	s-&gt;def-&gt;type = idealbool;

	s = lookup(&#34;_&#34;);
	s-&gt;block = -100;
	s-&gt;def = nod(ONAME, N, N);
	s-&gt;def-&gt;sym = s;
	types[TBLANK] = typ(TBLANK);
	s-&gt;def-&gt;type = types[TBLANK];
	nblank = s-&gt;def;
}

void
lexfini(void)
{
	Sym *s;
	int lex, etype, i;
	Val v;

	for(i=0; i&lt;nelem(syms); i++) {
		lex = syms[i].lexical;
		if(lex != LNAME)
			continue;
		s = lookup(syms[i].name);
		s-&gt;lexical = lex;

		etype = syms[i].etype;
		if(etype != Txxx &amp;&amp; (etype != TANY || debug[&#39;A&#39;]))
		if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONONAME)
			*s-&gt;def = *typenod(types[etype]);

		etype = syms[i].op;
		if(etype != OXXX &amp;&amp; s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONONAME) {
			s-&gt;def-&gt;op = ONAME;
			s-&gt;def-&gt;sym = s;
			s-&gt;def-&gt;etype = etype;
			s-&gt;def-&gt;builtin = 1;
		}
	}

	for(i=0; typedefs[i].name; i++) {
		s = lookup(typedefs[i].name);
		if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONONAME)
			*s-&gt;def = *typenod(types[typedefs[i].etype]);
	}

	// there&#39;s only so much table-driven we can handle.
	// these are special cases.
	types[TNIL] = typ(TNIL);
	s = lookup(&#34;nil&#34;);
	if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONONAME) {
		v.ctype = CTNIL;
		*s-&gt;def = *nodlit(v);
		s-&gt;def-&gt;sym = s;
	}

	s = lookup(&#34;true&#34;);
	if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONONAME) {
		*s-&gt;def = *nodbool(1);
		s-&gt;def-&gt;sym = s;
	}

	s = lookup(&#34;false&#34;);
	if(s-&gt;def != N &amp;&amp; s-&gt;def-&gt;op == ONONAME) {
		*s-&gt;def = *nodbool(0);
		s-&gt;def-&gt;sym = s;
	}
}

struct
{
	int	lex;
	char*	name;
} lexn[] =
{
	LANDAND,	&#34;ANDAND&#34;,
	LASOP,		&#34;ASOP&#34;,
	LBREAK,		&#34;BREAK&#34;,
	LCASE,		&#34;CASE&#34;,
	LCHAN,		&#34;CHAN&#34;,
	LCOLAS,		&#34;COLAS&#34;,
	LCONST,		&#34;CONST&#34;,
	LCONTINUE,	&#34;CONTINUE&#34;,
	LDEC,		&#34;DEC&#34;,
	LDEFER,		&#34;DEFER&#34;,
	LELSE,		&#34;ELSE&#34;,
	LEQ,		&#34;EQ&#34;,
	LFALL,		&#34;FALL&#34;,
	LFOR,		&#34;FOR&#34;,
	LFUNC,		&#34;FUNC&#34;,
	LGE,		&#34;GE&#34;,
	LGO,		&#34;GO&#34;,
	LGOTO,		&#34;GOTO&#34;,
	LGT,		&#34;GT&#34;,
	LIF,		&#34;IF&#34;,
	LIMPORT,	&#34;IMPORT&#34;,
	LINC,		&#34;INC&#34;,
	LINTERFACE,	&#34;INTERFACE&#34;,
	LLE,		&#34;LE&#34;,
	LLITERAL,	&#34;LITERAL&#34;,
	LLSH,		&#34;LSH&#34;,
	LLT,		&#34;LT&#34;,
	LMAP,		&#34;MAP&#34;,
	LNAME,		&#34;NAME&#34;,
	LNE,		&#34;NE&#34;,
	LOROR,		&#34;OROR&#34;,
	LPACKAGE,	&#34;PACKAGE&#34;,
	LRANGE,		&#34;RANGE&#34;,
	LRETURN,	&#34;RETURN&#34;,
	LRSH,		&#34;RSH&#34;,
	LSTRUCT,	&#34;STRUCT&#34;,
	LSWITCH,	&#34;SWITCH&#34;,
	LTYPE,		&#34;TYPE&#34;,
	LVAR,		&#34;VAR&#34;,
};

char*
lexname(int lex)
{
	int i;
	static char buf[100];

	for(i=0; i&lt;nelem(lexn); i++)
		if(lexn[i].lex == lex)
			return lexn[i].name;
	snprint(buf, sizeof(buf), &#34;LEX-%d&#34;, lex);
	return buf;
}

void
mkpackage(char* pkg)
{
	Sym *s;
	int32 h;
	char *p;

	if(package == nopackage) {
		if(strcmp(pkg, &#34;_&#34;) == 0)
			yyerror(&#34;invalid package name _&#34;);

		// redefine all names to be this package.
		for(h=0; h&lt;NHASH; h++)
			for(s = hash[h]; s != S; s = s-&gt;link)
				if(s-&gt;package == nopackage)
					s-&gt;package = pkg;
		package = pkg;
	} else {
		if(strcmp(pkg, package) != 0)
			yyerror(&#34;package %s; expected %s&#34;, pkg, package);
		for(h=0; h&lt;NHASH; h++) {
			for(s = hash[h]; s != S; s = s-&gt;link) {
				if(s-&gt;def == N || s-&gt;package != package)
					continue;
				if(s-&gt;def-&gt;op == OPACK) {
					// throw away top-level package name leftover
					// from previous file.
					// TODO(rsc): remember that there was a package
					// name, so that the name cannot be redeclared
					// as a non-package in other files.
					if(!s-&gt;def-&gt;used &amp;&amp; !nsyntaxerrors)
						yyerrorl(s-&gt;def-&gt;lineno, &#34;imported and not used: %s&#34;, s-&gt;def-&gt;sym-&gt;name);
					s-&gt;def = N;
					continue;
				}
				if(s-&gt;def-&gt;sym != s) {
					// throw away top-level name left over
					// from previous import . &#34;x&#34;
					if(s-&gt;def-&gt;pack != N &amp;&amp; !s-&gt;def-&gt;pack-&gt;used &amp;&amp; !nsyntaxerrors) {
						yyerrorl(s-&gt;def-&gt;pack-&gt;lineno, &#34;imported and not used: %s&#34;, s-&gt;def-&gt;pack-&gt;sym-&gt;name);
						s-&gt;def-&gt;pack-&gt;used = 1;
					}
					s-&gt;def = N;
					continue;
				}
			}
		}
	}

	if(outfile == nil) {
		p = strrchr(infile, &#39;/&#39;);
		if(p == nil)
			p = infile;
		else
			p = p+1;
		snprint(namebuf, sizeof(namebuf), &#34;%s&#34;, p);
		p = strrchr(namebuf, &#39;.&#39;);
		if(p != nil)
			*p = 0;
		outfile = smprint(&#34;%s.%c&#34;, namebuf, thechar);
	}
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
