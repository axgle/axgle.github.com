<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8l/obj.c</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8l/obj.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8l/obj.c
// http://code.google.com/p/inferno-os/source/browse/utils/8l/obj.c
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

#define	EXTERN
#include	&#34;l.h&#34;
#include	&#34;../ld/lib.h&#34;
#include	&#34;../ld/elf.h&#34;
#include	&#34;../ld/macho.h&#34;
#include	&lt;ar.h&gt;

#ifndef	DEFAULT
#define	DEFAULT	&#39;9&#39;
#endif

char	*noname		= &#34;&lt;none&gt;&#34;;
char	thechar		= &#39;8&#39;;
char	*thestring 	= &#34;386&#34;;

/*
 *	-H0 -T0x40004C -D0x10000000	is garbage unix
 *	-H1 -T0xd0 -R4			is unix coff
 *	-H2 -T4128 -R4096		is plan9 format
 *	-H3 -Tx -Rx			is MS-DOS .COM
 *	-H4 -Tx -Rx			is fake MS-DOS .EXE
 */

static int
isobjfile(char *f)
{
	int n, v;
	Biobuf *b;
	char buf1[5], buf2[SARMAG];

	b = Bopen(f, OREAD);
	if(b == nil)
		return 0;
	n = Bread(b, buf1, 5);
	if(n == 5 &amp;&amp; (buf1[2] == 1 &amp;&amp; buf1[3] == &#39;&lt;&#39; || buf1[3] == 1 &amp;&amp; buf1[4] == &#39;&lt;&#39;))
		v = 1;	/* good enough for our purposes */
	else{
		Bseek(b, 0, 0);
		n = Bread(b, buf2, SARMAG);
		v = n == SARMAG &amp;&amp; strncmp(buf2, ARMAG, SARMAG) == 0;
	}
	Bterm(b);
	return v;
}

void
usage(void)
{
	fprint(2, &#34;usage: 8l [-options] [-E entry] [-H head] [-L dir] [-T text] [-R rnd] [-o out] files...\n&#34;);
	exits(&#34;usage&#34;);
}

void
main(int argc, char *argv[])
{
	int i, c;

	Binit(&amp;bso, 1, OWRITE);
	cout = -1;
	listinit();
	memset(debug, 0, sizeof(debug));
	nerrors = 0;
	outfile = &#34;8.out&#34;;
	HEADTYPE = -1;
	INITTEXT = -1;
	INITDAT = -1;
	INITRND = -1;
	INITENTRY = 0;

	ARGBEGIN {
	default:
		c = ARGC();
		if(c &gt;= 0 &amp;&amp; c &lt; sizeof(debug))
			debug[c]++;
		break;
	case &#39;o&#39;: /* output to (next arg) */
		outfile = EARGF(usage());
		break;
	case &#39;E&#39;:
		INITENTRY = EARGF(usage());
		break;
	case &#39;H&#39;:
		HEADTYPE = atolwhex(EARGF(usage()));
		break;
	case &#39;L&#39;:
		Lflag(EARGF(usage()));
		break;
	case &#39;T&#39;:
		INITTEXT = atolwhex(EARGF(usage()));
		break;
	case &#39;D&#39;:
		INITDAT = atolwhex(EARGF(usage()));
		break;
	case &#39;R&#39;:
		INITRND = atolwhex(EARGF(usage()));
		break;
		break;
	case &#39;x&#39;:	/* produce export table */
		doexp = 1;
		if(argv[1] != nil &amp;&amp; argv[1][0] != &#39;-&#39; &amp;&amp; !isobjfile(argv[1]))
			readundefs(ARGF(), SEXPORT);
		break;
	case &#39;u&#39;:	/* produce dynamically loadable module */
		dlm = 1;
		debug[&#39;l&#39;]++;
		if(argv[1] != nil &amp;&amp; argv[1][0] != &#39;-&#39; &amp;&amp; !isobjfile(argv[1]))
			readundefs(ARGF(), SIMPORT);
		break;
	} ARGEND
	USED(argc);
	if(*argv == 0)
		usage();

	libinit();

	if(HEADTYPE == -1) {
		HEADTYPE = 2;
		if(strcmp(goos, &#34;linux&#34;) == 0)
			HEADTYPE = 7;
		else
		if(strcmp(goos, &#34;darwin&#34;) == 0)
			HEADTYPE = 6;
		else
		if(strcmp(goos, &#34;nacl&#34;) == 0)
			HEADTYPE = 8;
		else
			print(&#34;goos is not known: %sn&#34;, goos);
	}

	switch(HEADTYPE) {
	default:
		diag(&#34;unknown -H option&#34;);
		errorexit();

	case 0:	/* this is garbage */
		HEADR = 20L+56L;
		if(INITTEXT == -1)
			INITTEXT = 0x40004CL;
		if(INITDAT == -1)
			INITDAT = 0x10000000L;
		if(INITRND == -1)
			INITRND = 0;
		break;
	case 1:	/* is unix coff */
		HEADR = 0xd0L;
		if(INITTEXT == -1)
			INITTEXT = 0xd0;
		if(INITDAT == -1)
			INITDAT = 0x400000;
		if(INITRND == -1)
			INITRND = 0;
		break;
	case 2:	/* plan 9 */
		HEADR = 32L;
		if(INITTEXT == -1)
			INITTEXT = 4096+32;
		if(INITDAT == -1)
			INITDAT = 0;
		if(INITRND == -1)
			INITRND = 4096;
		break;
	case 3:	/* MS-DOS .COM */
		HEADR = 0;
		if(INITTEXT == -1)
			INITTEXT = 0x0100;
		if(INITDAT == -1)
			INITDAT = 0;
		if(INITRND == -1)
			INITRND = 4;
		break;
	case 4:	/* fake MS-DOS .EXE */
		HEADR = 0x200;
		if(INITTEXT == -1)
			INITTEXT = 0x0100;
		if(INITDAT == -1)
			INITDAT = 0;
		if(INITRND == -1)
			INITRND = 4;
		HEADR += (INITTEXT &amp; 0xFFFF);
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;HEADR = 0x%ld\n&#34;, HEADR);
		break;
	case 6:	/* apple MACH */
		/*
		 * OS X system constant - offset from %gs to our TLS.
		 * Explained in ../../libcgo/darwin_386.c.
		 */
		tlsoffset = 0x468;
		machoinit();
		HEADR = MACHORESERVE;
		if(INITTEXT == -1)
			INITTEXT = 4096+HEADR;
		if(INITDAT == -1)
			INITDAT = 0;
		if(INITRND == -1)
			INITRND = 4096;
		break;
	case 7:	/* elf32 executable */
		/*
		 * Linux ELF uses TLS offsets negative from %gs.
		 * Translate 0(GS) and 4(GS) into -8(GS) and -4(GS).
		 * Also known to ../../pkg/runtime/linux/386/sys.s
		 * and ../../libcgo/linux_386.c.
		 */
		tlsoffset = -8;
		elfinit();
		HEADR = ELFRESERVE;
		if(INITTEXT == -1)
			INITTEXT = 0x08048000+HEADR;
		if(INITDAT == -1)
			INITDAT = 0;
		if(INITRND == -1)
			INITRND = 4096;
		break;
	case 8:	/* native client elf32 executable */
		elfinit();
		HEADR = 4096;
		if(INITTEXT == -1)
			INITTEXT = 0x20000;
		if(INITDAT == -1)
			INITDAT = 0;
		if(INITRND == -1)
			INITRND = 4096;
		break;
	}
	if(INITDAT != 0 &amp;&amp; INITRND != 0)
		print(&#34;warning: -D0x%lux is ignored because of -R0x%lux\n&#34;,
			INITDAT, INITRND);
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;HEADER = -H0x%ld -T0x%lux -D0x%lux -R0x%lux\n&#34;,
			HEADTYPE, INITTEXT, INITDAT, INITRND);
	Bflush(&amp;bso);
	for(i=1; optab[i].as; i++)
		if(i != optab[i].as) {
			diag(&#34;phase error in optab: %d&#34;, i);
			errorexit();
		}
	maxop = i;

	for(i=0; i&lt;Ymax; i++)
		ycover[i*Ymax + i] = 1;

	ycover[Yi0*Ymax + Yi8] = 1;
	ycover[Yi1*Ymax + Yi8] = 1;

	ycover[Yi0*Ymax + Yi32] = 1;
	ycover[Yi1*Ymax + Yi32] = 1;
	ycover[Yi8*Ymax + Yi32] = 1;

	ycover[Yal*Ymax + Yrb] = 1;
	ycover[Ycl*Ymax + Yrb] = 1;
	ycover[Yax*Ymax + Yrb] = 1;
	ycover[Ycx*Ymax + Yrb] = 1;
	ycover[Yrx*Ymax + Yrb] = 1;

	ycover[Yax*Ymax + Yrx] = 1;
	ycover[Ycx*Ymax + Yrx] = 1;

	ycover[Yax*Ymax + Yrl] = 1;
	ycover[Ycx*Ymax + Yrl] = 1;
	ycover[Yrx*Ymax + Yrl] = 1;

	ycover[Yf0*Ymax + Yrf] = 1;

	ycover[Yal*Ymax + Ymb] = 1;
	ycover[Ycl*Ymax + Ymb] = 1;
	ycover[Yax*Ymax + Ymb] = 1;
	ycover[Ycx*Ymax + Ymb] = 1;
	ycover[Yrx*Ymax + Ymb] = 1;
	ycover[Yrb*Ymax + Ymb] = 1;
	ycover[Ym*Ymax + Ymb] = 1;

	ycover[Yax*Ymax + Yml] = 1;
	ycover[Ycx*Ymax + Yml] = 1;
	ycover[Yrx*Ymax + Yml] = 1;
	ycover[Yrl*Ymax + Yml] = 1;
	ycover[Ym*Ymax + Yml] = 1;

	for(i=0; i&lt;D_NONE; i++) {
		reg[i] = -1;
		if(i &gt;= D_AL &amp;&amp; i &lt;= D_BH)
			reg[i] = (i-D_AL) &amp; 7;
		if(i &gt;= D_AX &amp;&amp; i &lt;= D_DI)
			reg[i] = (i-D_AX) &amp; 7;
		if(i &gt;= D_F0 &amp;&amp; i &lt;= D_F0+7)
			reg[i] = (i-D_F0) &amp; 7;
	}

	zprg.link = P;
	zprg.pcond = P;
	zprg.back = 2;
	zprg.as = AGOK;
	zprg.from.type = D_NONE;
	zprg.from.index = D_NONE;
	zprg.from.scale = 1;
	zprg.to = zprg.from;

	pcstr = &#34;%.6lux &#34;;
	nuxiinit();
	histgen = 0;
	textp = P;
	datap = P;
	edatap = P;
	pc = 0;
	dtype = 4;
	version = 0;
	cbp = buf.cbuf;
	cbc = sizeof(buf.cbuf);
	firstp = prg();
	lastp = firstp;

	while(*argv)
		objfile(*argv++);

	if(!debug[&#39;l&#39;])
		loadlib();

	deadcode();

	firstp = firstp-&gt;link;
	if(firstp == P)
		errorexit();
	if(doexp || dlm){
		EXPTAB = &#34;_exporttab&#34;;
		zerosig(EXPTAB);
		zerosig(&#34;etext&#34;);
		zerosig(&#34;edata&#34;);
		zerosig(&#34;end&#34;);
		if(dlm){
			import();
			HEADTYPE = 2;
			INITTEXT = INITDAT = 0;
			INITRND = 8;
			INITENTRY = EXPTAB;
		}
		export();
	}
	patch();
	follow();
	doelf();
	if(HEADTYPE == 6)
		domacho();
	dodata();
	dostkoff();
	if(debug[&#39;p&#39;])
		if(debug[&#39;1&#39;])
			doprof1();
		else
			doprof2();
	span();
	doinit();
	asmb();
	undef();
	if(debug[&#39;v&#39;]) {
		Bprint(&amp;bso, &#34;%5.2f cpu time\n&#34;, cputime());
		Bprint(&amp;bso, &#34;%ld symbols\n&#34;, nsymbol);
		Bprint(&amp;bso, &#34;%d sizeof adr\n&#34;, sizeof(Adr));
		Bprint(&amp;bso, &#34;%d sizeof prog\n&#34;, sizeof(Prog));
	}
	Bflush(&amp;bso);

	errorexit();
}

void
zaddr(Biobuf *f, Adr *a, Sym *h[])
{
	int t;
	int32 l;
	Sym *s;
	Auto *u;

	t = Bgetc(f);
	a-&gt;index = D_NONE;
	a-&gt;scale = 0;
	if(t &amp; T_INDEX) {
		a-&gt;index = Bgetc(f);
		a-&gt;scale = Bgetc(f);
	}
	a-&gt;type = D_NONE;
	a-&gt;offset = 0;
	if(t &amp; T_OFFSET)
		a-&gt;offset = Bget4(f);
	a-&gt;offset2 = 0;
	if(t &amp; T_OFFSET2) {
		a-&gt;offset2 = Bget4(f);
		a-&gt;type = D_CONST2;
	}
	a-&gt;sym = S;
	if(t &amp; T_SYM)
		a-&gt;sym = h[Bgetc(f)];
	if(t &amp; T_FCONST) {
		a-&gt;ieee.l = Bget4(f);
		a-&gt;ieee.h = Bget4(f);
		a-&gt;type = D_FCONST;
	} else
	if(t &amp; T_SCONST) {
		Bread(f, a-&gt;scon, NSNAME);
		a-&gt;type = D_SCONST;
	}
	if(t &amp; T_TYPE)
		a-&gt;type = Bgetc(f);
	adrgotype = S;
	if(t &amp; T_GOTYPE)
		adrgotype = h[Bgetc(f)];

	t = a-&gt;type;
	if(t == D_INDIR+D_GS)
		a-&gt;offset += tlsoffset;

	s = a-&gt;sym;
	if(s == S)
		return;
	if(t != D_AUTO &amp;&amp; t != D_PARAM) {
		if(adrgotype)
			s-&gt;gotype = adrgotype;
		return;
	}
	l = a-&gt;offset;
	for(u=curauto; u; u=u-&gt;link) {
		if(u-&gt;asym == s)
		if(u-&gt;type == t) {
			if(u-&gt;aoffset &gt; l)
				u-&gt;aoffset = l;
			if(adrgotype)
				u-&gt;gotype = adrgotype;
			return;
		}
	}

	u = mal(sizeof(*u));
	u-&gt;link = curauto;
	curauto = u;
	u-&gt;asym = s;
	u-&gt;aoffset = l;
	u-&gt;type = t;
	u-&gt;gotype = adrgotype;
}

void
nopout(Prog *p)
{
	p-&gt;as = ANOP;
	p-&gt;from.type = D_NONE;
	p-&gt;to.type = D_NONE;
}

void
ldobj1(Biobuf *f, int64 len, char *pn)
{
	int32 ipc;
	Prog *p, *t;
	int v, o, r, skip;
	Sym *h[NSYM], *s, *di;
	uint32 sig;
	int ntext;
	int32 eof;
	char *name;
	char src[1024];

	ntext = 0;
	eof = Boffset(f) + len;
	di = S;
	src[0] = 0;


newloop:
	memset(h, 0, sizeof(h));
	version++;
	histfrogp = 0;
	ipc = pc;
	skip = 0;

loop:
	if(f-&gt;state == Bracteof || Boffset(f) &gt;= eof)
		goto eof;
	o = Bgetc(f);
	if(o == Beof)
		goto eof;
	o |= Bgetc(f) &lt;&lt; 8;
	if(o &lt;= AXXX || o &gt;= ALAST) {
		if(o &lt; 0)
			goto eof;
		diag(&#34;%s:#%lld: opcode out of range: %#ux&#34;, pn, Boffset(f), o);
		print(&#34;	probably not a .%c file\n&#34;, thechar);
		errorexit();
	}

	if(o == ANAME || o == ASIGNAME) {
		sig = 0;
		if(o == ASIGNAME)
			sig = Bget4(f);
		v = Bgetc(f);	/* type */
		o = Bgetc(f);	/* sym */
		r = 0;
		if(v == D_STATIC)
			r = version;
		name = Brdline(f, &#39;\0&#39;);
		if(name == nil) {
			if(Blinelen(f) &gt; 0) {
				fprint(2, &#34;%s: name too long\n&#34;, pn);
				errorexit();
			}
			goto eof;
		}
		s = lookup(name, r);

		if(debug[&#39;S&#39;] &amp;&amp; r == 0)
			sig = 1729;
		if(sig != 0){
			if(s-&gt;sig != 0 &amp;&amp; s-&gt;sig != sig)
				diag(&#34;incompatible type signatures&#34;
					&#34;%lux(%s) and %lux(%s) for %s&#34;,
					s-&gt;sig, s-&gt;file, sig, pn, s-&gt;name);
			s-&gt;sig = sig;
			s-&gt;file = pn;
		}

		if(debug[&#39;W&#39;])
			print(&#34;	ANAME	%s\n&#34;, s-&gt;name);
		h[o] = s;
		if((v == D_EXTERN || v == D_STATIC) &amp;&amp; s-&gt;type == 0)
			s-&gt;type = SXREF;
		if(v == D_FILE) {
			if(s-&gt;type != SFILE) {
				histgen++;
				s-&gt;type = SFILE;
				s-&gt;value = histgen;
			}
			if(histfrogp &lt; MAXHIST) {
				histfrog[histfrogp] = s;
				histfrogp++;
			} else
				collapsefrog(s);
		}
		goto loop;
	}

	p = mal(sizeof(*p));
	p-&gt;as = o;
	p-&gt;line = Bget4(f);
	p-&gt;back = 2;
	p-&gt;ft = 0;
	p-&gt;tt = 0;
	zaddr(f, &amp;p-&gt;from, h);
	fromgotype = adrgotype;
	zaddr(f, &amp;p-&gt;to, h);

	if(debug[&#39;W&#39;])
		print(&#34;%P\n&#34;, p);

	switch(p-&gt;as) {
	case AHISTORY:
		if(p-&gt;to.offset == -1) {
			addlib(src, pn);
			histfrogp = 0;
			goto loop;
		}
		if(src[0] == &#39;\0&#39;)
			copyhistfrog(src, sizeof src);
		addhist(p-&gt;line, D_FILE);		/* &#39;z&#39; */
		if(p-&gt;to.offset)
			addhist(p-&gt;to.offset, D_FILE1);	/* &#39;Z&#39; */
		histfrogp = 0;
		goto loop;

	case AEND:
		histtoauto();
		if(curtext != P)
			curtext-&gt;to.autom = curauto;
		curauto = 0;
		curtext = P;
		if(Boffset(f) == eof)
			return;
		goto newloop;

	case AGLOBL:
		s = p-&gt;from.sym;
		if(s-&gt;type == 0 || s-&gt;type == SXREF) {
			s-&gt;type = SBSS;
			s-&gt;value = 0;
		}
		if(s-&gt;type != SBSS) {
			diag(&#34;%s: redefinition: %s in %s&#34;,
				pn, s-&gt;name, TNAME);
			s-&gt;type = SBSS;
			s-&gt;value = 0;
		}
		if(p-&gt;to.offset &gt; s-&gt;value)
			s-&gt;value = p-&gt;to.offset;
		if(p-&gt;from.scale &amp; DUPOK)
			s-&gt;dupok = 1;
		goto loop;

	case ADYNT:
		if(p-&gt;to.sym == S) {
			diag(&#34;DYNT without a sym\n%P&#34;, p);
			break;
		}
		di = p-&gt;to.sym;
		p-&gt;from.scale = 4;
		if(di-&gt;type == SXREF) {
			if(debug[&#39;z&#39;])
				Bprint(&amp;bso, &#34;%P set to %d\n&#34;, p, dtype);
			di-&gt;type = SCONST;
			di-&gt;value = dtype;
			dtype += 4;
		}
		if(p-&gt;from.sym == S)
			break;

		p-&gt;from.offset = di-&gt;value;
		p-&gt;from.sym-&gt;type = SDATA;
		if(curtext == P) {
			diag(&#34;DYNT not in text: %P&#34;, p);
			break;
		}
		p-&gt;to.sym = curtext-&gt;from.sym;
		p-&gt;to.type = D_ADDR;
		p-&gt;to.index = D_EXTERN;
		goto data;

	case AINIT:
		if(p-&gt;from.sym == S) {
			diag(&#34;INIT without a sym\n%P&#34;, p);
			break;
		}
		if(di == S) {
			diag(&#34;INIT without previous DYNT\n%P&#34;, p);
			break;
		}
		p-&gt;from.offset = di-&gt;value;
		p-&gt;from.sym-&gt;type = SDATA;
		goto data;

	case ADATA:
	data:
		// Assume that AGLOBL comes after ADATA.
		// If we&#39;ve seen an AGLOBL that said this sym was DUPOK,
		// ignore any more ADATA we see, which must be
		// redefinitions.
		s = p-&gt;from.sym;
		if(s != S &amp;&amp; s-&gt;dupok) {
//			if(debug[&#39;v&#39;])
//				Bprint(&amp;bso, &#34;skipping %s in %s: dupok\n&#34;, s-&gt;name, pn);
			goto loop;
		}
		if(s != S) {
			p-&gt;dlink = s-&gt;data;
			s-&gt;data = p;
		}
		if(edatap == P)
			datap = p;
		else
			edatap-&gt;link = p;
		edatap = p;
		p-&gt;link = P;
		goto loop;

	case AGOK:
		diag(&#34;%s: GOK opcode in %s&#34;, pn, TNAME);
		pc++;
		goto loop;

	case ATEXT:
		s = p-&gt;from.sym;
		if(s == S) {
			diag(&#34;%s: no TEXT symbol: %P&#34;, pn, p);
			errorexit();
		}
		if(ntext++ == 0 &amp;&amp; s-&gt;type != 0 &amp;&amp; s-&gt;type != SXREF) {
			/* redefinition, so file has probably been seen before */
			if(debug[&#39;v&#39;])
				diag(&#34;skipping: %s: redefinition: %s&#34;, pn, s-&gt;name);
			return;
		}
		if(curtext != P) {
			histtoauto();
			curtext-&gt;to.autom = curauto;
			curauto = 0;
		}
		skip = 0;
		curtext = p;
		if(s-&gt;type != 0 &amp;&amp; s-&gt;type != SXREF) {
			if(p-&gt;from.scale &amp; DUPOK) {
				skip = 1;
				goto casdef;
			}
			diag(&#34;%s: redefinition: %s\n%P&#34;, pn, s-&gt;name, p);
		}
		newtext(p, s);
		goto loop;

	case AFMOVF:
	case AFADDF:
	case AFSUBF:
	case AFSUBRF:
	case AFMULF:
	case AFDIVF:
	case AFDIVRF:
	case AFCOMF:
	case AFCOMFP:
		if(skip)
			goto casdef;
		if(p-&gt;from.type == D_FCONST) {
			/* size sb 9 max */
			sprint(literal, &#34;$%lux&#34;, ieeedtof(&amp;p-&gt;from.ieee));
			s = lookup(literal, 0);
			if(s-&gt;type == 0) {
				s-&gt;type = SBSS;
				s-&gt;value = 4;
				t = prg();
				t-&gt;as = ADATA;
				t-&gt;line = p-&gt;line;
				t-&gt;from.type = D_EXTERN;
				t-&gt;from.sym = s;
				t-&gt;from.scale = 4;
				t-&gt;to = p-&gt;from;
				if(edatap == P)
					datap = t;
				else
					edatap-&gt;link = t;
				edatap = t;
				t-&gt;link = P;
			}
			p-&gt;from.type = D_EXTERN;
			p-&gt;from.sym = s;
			p-&gt;from.offset = 0;
		}
		goto casdef;

	case AFMOVD:
	case AFADDD:
	case AFSUBD:
	case AFSUBRD:
	case AFMULD:
	case AFDIVD:
	case AFDIVRD:
	case AFCOMD:
	case AFCOMDP:
		if(skip)
			goto casdef;
		if(p-&gt;from.type == D_FCONST) {
			/* size sb 18 max */
			sprint(literal, &#34;$%lux.%lux&#34;,
				p-&gt;from.ieee.l, p-&gt;from.ieee.h);
			s = lookup(literal, 0);
			if(s-&gt;type == 0) {
				s-&gt;type = SBSS;
				s-&gt;value = 8;
				t = prg();
				t-&gt;as = ADATA;
				t-&gt;line = p-&gt;line;
				t-&gt;from.type = D_EXTERN;
				t-&gt;from.sym = s;
				t-&gt;from.scale = 8;
				t-&gt;to = p-&gt;from;
				if(edatap == P)
					datap = t;
				else
					edatap-&gt;link = t;
				edatap = t;
				t-&gt;link = P;
			}
			p-&gt;from.type = D_EXTERN;
			p-&gt;from.sym = s;
			p-&gt;from.offset = 0;
		}
		goto casdef;

	casdef:
	default:
		if(skip)
			nopout(p);

		if(p-&gt;to.type == D_BRANCH)
			p-&gt;to.offset += ipc;
		lastp-&gt;link = p;
		lastp = p;
		p-&gt;pc = pc;
		pc++;
		goto loop;
	}
	goto loop;

eof:
	diag(&#34;truncated object file: %s&#34;, pn);
}

Prog*
prg(void)
{
	Prog *p;

	p = mal(sizeof(Prog));
	*p = zprg;
	return p;
}

Prog*
copyp(Prog *q)
{
	Prog *p;

	p = prg();
	*p = *q;
	return p;
}

Prog*
appendp(Prog *q)
{
	Prog *p;

	p = prg();
	p-&gt;link = q-&gt;link;
	q-&gt;link = p;
	p-&gt;line = q-&gt;line;
	return p;
}

void
doprof1(void)
{
	Sym *s;
	int32 n;
	Prog *p, *q;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f profile 1\n&#34;, cputime());
	Bflush(&amp;bso);
	s = lookup(&#34;__mcount&#34;, 0);
	n = 1;
	for(p = firstp-&gt;link; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT) {
			q = prg();
			q-&gt;line = p-&gt;line;
			q-&gt;link = datap;
			datap = q;
			q-&gt;as = ADATA;
			q-&gt;from.type = D_EXTERN;
			q-&gt;from.offset = n*4;
			q-&gt;from.sym = s;
			q-&gt;from.scale = 4;
			q-&gt;to = p-&gt;from;
			q-&gt;to.type = D_CONST;

			q = prg();
			q-&gt;line = p-&gt;line;
			q-&gt;pc = p-&gt;pc;
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;
			p-&gt;as = AADDL;
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = 1;
			p-&gt;to.type = D_EXTERN;
			p-&gt;to.sym = s;
			p-&gt;to.offset = n*4 + 4;

			n += 2;
			continue;
		}
	}
	q = prg();
	q-&gt;line = 0;
	q-&gt;link = datap;
	datap = q;

	q-&gt;as = ADATA;
	q-&gt;from.type = D_EXTERN;
	q-&gt;from.sym = s;
	q-&gt;from.scale = 4;
	q-&gt;to.type = D_CONST;
	q-&gt;to.offset = n;

	s-&gt;type = SBSS;
	s-&gt;value = n*4;
}

void
doprof2(void)
{
	Sym *s2, *s4;
	Prog *p, *q, *ps2, *ps4;

	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f profile 2\n&#34;, cputime());
	Bflush(&amp;bso);

	s2 = lookup(&#34;_profin&#34;, 0);
	s4 = lookup(&#34;_profout&#34;, 0);
	if(s2-&gt;type != STEXT || s4-&gt;type != STEXT) {
		diag(&#34;_profin/_profout not defined&#34;);
		return;
	}

	ps2 = P;
	ps4 = P;
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT) {
			if(p-&gt;from.sym == s2) {
				p-&gt;from.scale = 1;
				ps2 = p;
			}
			if(p-&gt;from.sym == s4) {
				p-&gt;from.scale = 1;
				ps4 = p;
			}
		}
	}
	for(p = firstp; p != P; p = p-&gt;link) {
		if(p-&gt;as == ATEXT) {
			curtext = p;

			if(p-&gt;from.scale &amp; NOPROF) {	/* dont profile */
				for(;;) {
					q = p-&gt;link;
					if(q == P)
						break;
					if(q-&gt;as == ATEXT)
						break;
					p = q;
				}
				continue;
			}

			/*
			 * JMPL	profin
			 */
			q = prg();
			q-&gt;line = p-&gt;line;
			q-&gt;pc = p-&gt;pc;
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;
			p = q;
			p-&gt;as = ACALL;
			p-&gt;to.type = D_BRANCH;
			p-&gt;pcond = ps2;
			p-&gt;to.sym = s2;

			continue;
		}
		if(p-&gt;as == ARET) {
			/*
			 * RET
			 */
			q = prg();
			q-&gt;as = ARET;
			q-&gt;from = p-&gt;from;
			q-&gt;to = p-&gt;to;
			q-&gt;link = p-&gt;link;
			p-&gt;link = q;

			/*
			 * JAL	profout
			 */
			p-&gt;as = ACALL;
			p-&gt;from = zprg.from;
			p-&gt;to = zprg.to;
			p-&gt;to.type = D_BRANCH;
			p-&gt;pcond = ps4;
			p-&gt;to.sym = s4;

			p = q;

			continue;
		}
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
