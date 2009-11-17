<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5a/lex.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5a/lex.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5a/lex.c
// http://code.google.com/p/inferno-os/source/browse/utils/5a/lex.c
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
#include &#34;a.h&#34;
#include &#34;y.tab.h&#34;
#include &lt;ctype.h&gt;

enum
{
	Plan9	= 1&lt;&lt;0,
	Unix	= 1&lt;&lt;1,
	Windows	= 1&lt;&lt;2,
};

int
systemtype(int sys)
{
	return sys&amp;Plan9;
}

void*	alloc(int32);
void*	allocn(void*, int32, int32);

void
main(int argc, char *argv[])
{
	char *p;
	int nout, nproc, i, c;

	thechar = &#39;5&#39;;
	thestring = &#34;arm&#34;;
	memset(debug, 0, sizeof(debug));
	cinit();
	outfile = 0;
	include[ninclude++] = &#34;.&#34;;
	ARGBEGIN {
	default:
		c = ARGC();
		if(c &gt;= 0 || c &lt; sizeof(debug))
			debug[c] = 1;
		break;

	case &#39;o&#39;:
		outfile = ARGF();
		break;

	case &#39;D&#39;:
		p = ARGF();
		if(p)
			Dlist[nDlist++] = p;
		break;

	case &#39;I&#39;:
		p = ARGF();
		setinclude(p);
		break;
	case &#39;t&#39;:
		thechar = &#39;t&#39;;
		thestring = &#34;thumb&#34;;
		break;
	} ARGEND
	if(*argv == 0) {
		print(&#34;usage: %ca [-options] file.s\n&#34;, thechar);
		errorexit();
	}
	if(argc &gt; 1 &amp;&amp; systemtype(Windows)){
		print(&#34;can&#39;t assemble multiple files on windows\n&#34;);
		errorexit();
	}
	if(argc &gt; 1 &amp;&amp; !systemtype(Windows)) {
		nproc = 1;
		if(p = getenv(&#34;NPROC&#34;))
			nproc = atol(p);	/* */
		c = 0;
		nout = 0;
		for(;;) {
			Waitmsg *w;

			while(nout &lt; nproc &amp;&amp; argc &gt; 0) {
				i = fork();
				if(i &lt; 0) {
					fprint(2, &#34;fork: %r\n&#34;);
					errorexit();
				}
				if(i == 0) {
					print(&#34;%s:\n&#34;, *argv);
					if(assemble(*argv))
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
	if(assemble(argv[0]))
		errorexit();
	exits(0);
}

int
assemble(char *file)
{
	char ofile[100], incfile[20], *p;
	int i, of;

	strcpy(ofile, file);
	p = utfrrune(ofile, &#39;/&#39;);
	if(p) {
		include[0] = ofile;
		*p++ = 0;
	} else
		p = ofile;
	if(outfile == 0) {
		outfile = p;
		if(outfile){
			p = utfrrune(outfile, &#39;.&#39;);
			if(p)
				if(p[1] == &#39;s&#39; &amp;&amp; p[2] == 0)
					p[0] = 0;
			p = utfrune(outfile, 0);
			p[0] = &#39;.&#39;;
			p[1] = thechar;
			p[2] = 0;
		} else
			outfile = &#34;/dev/null&#34;;
	}
	p = getenv(&#34;INCLUDE&#34;);
	if(p) {
		setinclude(p);
	} else {
		if(systemtype(Plan9)) {
			sprint(incfile,&#34;/%s/include&#34;, thestring);
			setinclude(strdup(incfile));
		}
	}

	of = create(outfile, OWRITE, 0664);
	if(of &lt; 0) {
		yyerror(&#34;%ca: cannot create %s&#34;, thechar, outfile);
		errorexit();
	}
	Binit(&amp;obuf, of, OWRITE);

	pass = 1;
	pinit(file);

	Bprint(&amp;obuf, &#34;%s\n&#34;, thestring);

	for(i=0; i&lt;nDlist; i++)
		dodefine(Dlist[i]);
	yyparse();
	if(nerrors) {
		cclean();
		return nerrors;
	}

	Bprint(&amp;obuf, &#34;\n!\n&#34;);

	pass = 2;
	outhist();
	pinit(file);
	for(i=0; i&lt;nDlist; i++)
		dodefine(Dlist[i]);
	yyparse();
	cclean();
	return nerrors;
}

struct
{
	char	*name;
	ushort	type;
	ushort	value;
} itab[] =
{
	&#34;SP&#34;,		LSP,	D_AUTO,
	&#34;SB&#34;,		LSB,	D_EXTERN,
	&#34;FP&#34;,		LFP,	D_PARAM,
	&#34;PC&#34;,		LPC,	D_BRANCH,

	&#34;R&#34;,		LR,	0,
	&#34;R0&#34;,		LREG,	0,
	&#34;R1&#34;,		LREG,	1,
	&#34;R2&#34;,		LREG,	2,
	&#34;R3&#34;,		LREG,	3,
	&#34;R4&#34;,		LREG,	4,
	&#34;R5&#34;,		LREG,	5,
	&#34;R6&#34;,		LREG,	6,
	&#34;R7&#34;,		LREG,	7,
	&#34;R8&#34;,		LREG,	8,
	&#34;R9&#34;,		LREG,	9,
	&#34;R10&#34;,		LREG,	10,
	&#34;R11&#34;,		LREG,	11,
	&#34;R12&#34;,		LREG,	12,
	&#34;R13&#34;,		LREG,	13,
	&#34;R14&#34;,		LREG,	14,
	&#34;R15&#34;,		LREG,	15,

	&#34;F&#34;,		LF,	0,

	&#34;F0&#34;,		LFREG,	0,
	&#34;F1&#34;,		LFREG,	1,
	&#34;F2&#34;,		LFREG,	2,
	&#34;F3&#34;,		LFREG,	3,
	&#34;F4&#34;,		LFREG,	4,
	&#34;F5&#34;,		LFREG,	5,
	&#34;F6&#34;,		LFREG,	6,
	&#34;F7&#34;,		LFREG,	7,
	&#34;F8&#34;,		LFREG,	8,
	&#34;F9&#34;,		LFREG,	9,
	&#34;F10&#34;,		LFREG,	10,
	&#34;F11&#34;,		LFREG,	11,
	&#34;F12&#34;,		LFREG,	12,
	&#34;F13&#34;,		LFREG,	13,
	&#34;F14&#34;,		LFREG,	14,
	&#34;F15&#34;,		LFREG,	15,

	&#34;C&#34;,		LC,	0,

	&#34;C0&#34;,		LCREG,	0,
	&#34;C1&#34;,		LCREG,	1,
	&#34;C2&#34;,		LCREG,	2,
	&#34;C3&#34;,		LCREG,	3,
	&#34;C4&#34;,		LCREG,	4,
	&#34;C5&#34;,		LCREG,	5,
	&#34;C6&#34;,		LCREG,	6,
	&#34;C7&#34;,		LCREG,	7,
	&#34;C8&#34;,		LCREG,	8,
	&#34;C9&#34;,		LCREG,	9,
	&#34;C10&#34;,		LCREG,	10,
	&#34;C11&#34;,		LCREG,	11,
	&#34;C12&#34;,		LCREG,	12,
	&#34;C13&#34;,		LCREG,	13,
	&#34;C14&#34;,		LCREG,	14,
	&#34;C15&#34;,		LCREG,	15,

	&#34;CPSR&#34;,		LPSR,	0,
	&#34;SPSR&#34;,		LPSR,	1,

	&#34;FPSR&#34;,		LFCR,	0,
	&#34;FPCR&#34;,		LFCR,	1,

	&#34;.EQ&#34;,		LCOND,	0,
	&#34;.NE&#34;,		LCOND,	1,
	&#34;.CS&#34;,		LCOND,	2,
	&#34;.HS&#34;,		LCOND,	2,
	&#34;.CC&#34;,		LCOND,	3,
	&#34;.LO&#34;,		LCOND,	3,
	&#34;.MI&#34;,		LCOND,	4,
	&#34;.PL&#34;,		LCOND,	5,
	&#34;.VS&#34;,		LCOND,	6,
	&#34;.VC&#34;,		LCOND,	7,
	&#34;.HI&#34;,		LCOND,	8,
	&#34;.LS&#34;,		LCOND,	9,
	&#34;.GE&#34;,		LCOND,	10,
	&#34;.LT&#34;,		LCOND,	11,
	&#34;.GT&#34;,		LCOND,	12,
	&#34;.LE&#34;,		LCOND,	13,
	&#34;.AL&#34;,		LCOND,	Always,

	&#34;.U&#34;,		LS,	C_UBIT,
	&#34;.S&#34;,		LS,	C_SBIT,
	&#34;.W&#34;,		LS,	C_WBIT,
	&#34;.P&#34;,		LS,	C_PBIT,
	&#34;.PW&#34;,		LS,	C_WBIT|C_PBIT,
	&#34;.WP&#34;,		LS,	C_WBIT|C_PBIT,

	&#34;.F&#34;,		LS,	C_FBIT,

	&#34;.IBW&#34;,		LS,	C_WBIT|C_PBIT|C_UBIT,
	&#34;.IAW&#34;,		LS,	C_WBIT|C_UBIT,
	&#34;.DBW&#34;,		LS,	C_WBIT|C_PBIT,
	&#34;.DAW&#34;,		LS,	C_WBIT,
	&#34;.IB&#34;,		LS,	C_PBIT|C_UBIT,
	&#34;.IA&#34;,		LS,	C_UBIT,
	&#34;.DB&#34;,		LS,	C_PBIT,
	&#34;.DA&#34;,		LS,	0,

	&#34;@&#34;,		LAT,	0,

	&#34;AND&#34;,		LTYPE1,	AAND,
	&#34;EOR&#34;,		LTYPE1,	AEOR,
	&#34;SUB&#34;,		LTYPE1,	ASUB,
	&#34;RSB&#34;,		LTYPE1,	ARSB,
	&#34;ADD&#34;,		LTYPE1,	AADD,
	&#34;ADC&#34;,		LTYPE1,	AADC,
	&#34;SBC&#34;,		LTYPE1,	ASBC,
	&#34;RSC&#34;,		LTYPE1,	ARSC,
	&#34;ORR&#34;,		LTYPE1,	AORR,
	&#34;BIC&#34;,		LTYPE1,	ABIC,

	&#34;SLL&#34;,		LTYPE1,	ASLL,
	&#34;SRL&#34;,		LTYPE1,	ASRL,
	&#34;SRA&#34;,		LTYPE1,	ASRA,

	&#34;MUL&#34;,		LTYPE1, AMUL,
	&#34;MULA&#34;,		LTYPEN, AMULA,
	&#34;DIV&#34;,		LTYPE1,	ADIV,
	&#34;MOD&#34;,		LTYPE1,	AMOD,

	&#34;MULL&#34;,		LTYPEM, AMULL,
	&#34;MULAL&#34;,	LTYPEM, AMULAL,
	&#34;MULLU&#34;,	LTYPEM, AMULLU,
	&#34;MULALU&#34;,	LTYPEM, AMULALU,

	&#34;MVN&#34;,		LTYPE2, AMVN,	/* op2 ignored */

	&#34;MOVB&#34;,		LTYPE3, AMOVB,
	&#34;MOVBU&#34;,	LTYPE3, AMOVBU,
	&#34;MOVH&#34;,		LTYPE3, AMOVH,
	&#34;MOVHU&#34;,	LTYPE3, AMOVHU,
	&#34;MOVW&#34;,		LTYPE3, AMOVW,

	&#34;MOVD&#34;,		LTYPE3, AMOVD,
	&#34;MOVDF&#34;,		LTYPE3, AMOVDF,
	&#34;MOVDW&#34;,	LTYPE3, AMOVDW,
	&#34;MOVF&#34;,		LTYPE3, AMOVF,
	&#34;MOVFD&#34;,		LTYPE3, AMOVFD,
	&#34;MOVFW&#34;,		LTYPE3, AMOVFW,
	&#34;MOVWD&#34;,	LTYPE3, AMOVWD,
	&#34;MOVWF&#34;,		LTYPE3, AMOVWF,

	&#34;LDREX&#34;,		LTYPE3, ALDREX,
	&#34;STREX&#34;,		LTYPE9, ASTREX,

/*
	&#34;ABSF&#34;,		LTYPEI, AABSF,
	&#34;ABSD&#34;,		LTYPEI, AABSD,
	&#34;NEGF&#34;,		LTYPEI, ANEGF,
	&#34;NEGD&#34;,		LTYPEI, ANEGD,
	&#34;SQTF&#34;,		LTYPEI,	ASQTF,
	&#34;SQTD&#34;,		LTYPEI,	ASQTD,
	&#34;RNDF&#34;,		LTYPEI,	ARNDF,
	&#34;RNDD&#34;,		LTYPEI,	ARNDD,
	&#34;URDF&#34;,		LTYPEI,	AURDF,
	&#34;URDD&#34;,		LTYPEI,	AURDD,
	&#34;NRMF&#34;,		LTYPEI,	ANRMF,
	&#34;NRMD&#34;,		LTYPEI,	ANRMD,
*/

	&#34;CMPF&#34;,		LTYPEL, ACMPF,
	&#34;CMPD&#34;,		LTYPEL, ACMPD,
	&#34;ADDF&#34;,		LTYPEK,	AADDF,
	&#34;ADDD&#34;,		LTYPEK,	AADDD,
	&#34;SUBF&#34;,		LTYPEK,	ASUBF,
	&#34;SUBD&#34;,		LTYPEK,	ASUBD,
	&#34;MULF&#34;,		LTYPEK,	AMULF,
	&#34;MULD&#34;,		LTYPEK,	AMULD,
	&#34;DIVF&#34;,		LTYPEK,	ADIVF,
	&#34;DIVD&#34;,		LTYPEK,	ADIVD,

	&#34;B&#34;,		LTYPE4, AB,
	&#34;BL&#34;,		LTYPE4, ABL,
	&#34;BX&#34;,		LTYPEBX,	ABX,

	&#34;BEQ&#34;,		LTYPE5,	ABEQ,
	&#34;BNE&#34;,		LTYPE5,	ABNE,
	&#34;BCS&#34;,		LTYPE5,	ABCS,
	&#34;BHS&#34;,		LTYPE5,	ABHS,
	&#34;BCC&#34;,		LTYPE5,	ABCC,
	&#34;BLO&#34;,		LTYPE5,	ABLO,
	&#34;BMI&#34;,		LTYPE5,	ABMI,
	&#34;BPL&#34;,		LTYPE5,	ABPL,
	&#34;BVS&#34;,		LTYPE5,	ABVS,
	&#34;BVC&#34;,		LTYPE5,	ABVC,
	&#34;BHI&#34;,		LTYPE5,	ABHI,
	&#34;BLS&#34;,		LTYPE5,	ABLS,
	&#34;BGE&#34;,		LTYPE5,	ABGE,
	&#34;BLT&#34;,		LTYPE5,	ABLT,
	&#34;BGT&#34;,		LTYPE5,	ABGT,
	&#34;BLE&#34;,		LTYPE5,	ABLE,
	&#34;BCASE&#34;,	LTYPE5,	ABCASE,

	&#34;SWI&#34;,		LTYPE6, ASWI,

	&#34;CMP&#34;,		LTYPE7,	ACMP,
	&#34;TST&#34;,		LTYPE7,	ATST,
	&#34;TEQ&#34;,		LTYPE7,	ATEQ,
	&#34;CMN&#34;,		LTYPE7,	ACMN,

	&#34;MOVM&#34;,		LTYPE8, AMOVM,

	&#34;SWPBU&#34;,	LTYPE9, ASWPBU,
	&#34;SWPW&#34;,		LTYPE9, ASWPW,

	&#34;RET&#34;,		LTYPEA, ARET,
	&#34;RFE&#34;,		LTYPEA, ARFE,

	&#34;TEXT&#34;,		LTYPEB, ATEXT,
	&#34;GLOBL&#34;,	LTYPEB, AGLOBL,
	&#34;DATA&#34;,		LTYPEC, ADATA,
	&#34;CASE&#34;,		LTYPED, ACASE,
	&#34;END&#34;,		LTYPEE, AEND,
	&#34;WORD&#34;,		LTYPEH, AWORD,
	&#34;NOP&#34;,		LTYPEI, ANOP,

	&#34;MCR&#34;,		LTYPEJ, 0,
	&#34;MRC&#34;,		LTYPEJ, 1,
	0
};

void
cinit(void)
{
	Sym *s;
	int i;

	nullgen.sym = S;
	nullgen.offset = 0;
	nullgen.type = D_NONE;
	nullgen.name = D_NONE;
	nullgen.reg = NREG;
	if(FPCHIP)
		nullgen.dval = 0;
	for(i=0; i&lt;sizeof(nullgen.sval); i++)
		nullgen.sval[i] = 0;

	nerrors = 0;
	iostack = I;
	iofree = I;
	peekc = IGN;
	nhunk = 0;
	for(i=0; i&lt;NHASH; i++)
		hash[i] = S;
	for(i=0; itab[i].name; i++) {
		s = slookup(itab[i].name);
		s-&gt;type = itab[i].type;
		s-&gt;value = itab[i].value;
	}

	pathname = allocn(pathname, 0, 100);
	if(getwd(pathname, 99) == 0) {
		pathname = allocn(pathname, 100, 900);
		if(getwd(pathname, 999) == 0)
			strcpy(pathname, &#34;/???&#34;);
	}
}

void
syminit(Sym *s)
{

	s-&gt;type = LNAME;
	s-&gt;value = 0;
}

int
isreg(Gen *g)
{

	USED(g);
	return 1;
}

void
cclean(void)
{

	outcode(AEND, Always, &amp;nullgen, NREG, &amp;nullgen);
	Bflush(&amp;obuf);
}

void
zname(char *n, int t, int s)
{

	Bputc(&amp;obuf, ANAME);
	Bputc(&amp;obuf, t);	/* type */
	Bputc(&amp;obuf, s);	/* sym */
	while(*n) {
		Bputc(&amp;obuf, *n);
		n++;
	}
	Bputc(&amp;obuf, 0);
}

void
zaddr(Gen *a, int s)
{
	int32 l;
	int i;
	char *n;
	Ieee e;

	Bputc(&amp;obuf, a-&gt;type);
	Bputc(&amp;obuf, a-&gt;reg);
	Bputc(&amp;obuf, s);
	Bputc(&amp;obuf, a-&gt;name);
	switch(a-&gt;type) {
	default:
		print(&#34;unknown type %d\n&#34;, a-&gt;type);
		exits(&#34;arg&#34;);

	case D_NONE:
	case D_REG:
	case D_FREG:
	case D_PSR:
	case D_FPCR:
		break;

	case D_REGREG:
		Bputc(&amp;obuf, a-&gt;offset);
		break;

	case D_OREG:
	case D_CONST:
	case D_BRANCH:
	case D_SHIFT:
		l = a-&gt;offset;
		Bputc(&amp;obuf, l);
		Bputc(&amp;obuf, l&gt;&gt;8);
		Bputc(&amp;obuf, l&gt;&gt;16);
		Bputc(&amp;obuf, l&gt;&gt;24);
		break;

	case D_SCONST:
		n = a-&gt;sval;
		for(i=0; i&lt;NSNAME; i++) {
			Bputc(&amp;obuf, *n);
			n++;
		}
		break;

	case D_FCONST:
		ieeedtod(&amp;e, a-&gt;dval);
		Bputc(&amp;obuf, e.l);
		Bputc(&amp;obuf, e.l&gt;&gt;8);
		Bputc(&amp;obuf, e.l&gt;&gt;16);
		Bputc(&amp;obuf, e.l&gt;&gt;24);
		Bputc(&amp;obuf, e.h);
		Bputc(&amp;obuf, e.h&gt;&gt;8);
		Bputc(&amp;obuf, e.h&gt;&gt;16);
		Bputc(&amp;obuf, e.h&gt;&gt;24);
		break;
	}
}

static int bcode[] =
{
	ABEQ,
	ABNE,
	ABCS,
	ABCC,
	ABMI,
	ABPL,
	ABVS,
	ABVC,
	ABHI,
	ABLS,
	ABGE,
	ABLT,
	ABGT,
	ABLE,
	AB,
	ANOP,
};

void
outcode(int a, int scond, Gen *g1, int reg, Gen *g2)
{
	int sf, st, t;
	Sym *s;

	/* hack to make B.NE etc. work: turn it into the corresponding conditional */
	if(a == AB){
		a = bcode[scond&amp;0xf];
		scond = (scond &amp; ~0xf) | Always;
	}

	if(pass == 1)
		goto out;
jackpot:
	sf = 0;
	s = g1-&gt;sym;
	while(s != S) {
		sf = s-&gt;sym;
		if(sf &lt; 0 || sf &gt;= NSYM)
			sf = 0;
		t = g1-&gt;name;
		if(h[sf].type == t)
		if(h[sf].sym == s)
			break;
		zname(s-&gt;name, t, sym);
		s-&gt;sym = sym;
		h[sym].sym = s;
		h[sym].type = t;
		sf = sym;
		sym++;
		if(sym &gt;= NSYM)
			sym = 1;
		break;
	}
	st = 0;
	s = g2-&gt;sym;
	while(s != S) {
		st = s-&gt;sym;
		if(st &lt; 0 || st &gt;= NSYM)
			st = 0;
		t = g2-&gt;name;
		if(h[st].type == t)
		if(h[st].sym == s)
			break;
		zname(s-&gt;name, t, sym);
		s-&gt;sym = sym;
		h[sym].sym = s;
		h[sym].type = t;
		st = sym;
		sym++;
		if(sym &gt;= NSYM)
			sym = 1;
		if(st == sf)
			goto jackpot;
		break;
	}
	Bputc(&amp;obuf, a);
	Bputc(&amp;obuf, scond);
	Bputc(&amp;obuf, reg);
	Bputc(&amp;obuf, lineno);
	Bputc(&amp;obuf, lineno&gt;&gt;8);
	Bputc(&amp;obuf, lineno&gt;&gt;16);
	Bputc(&amp;obuf, lineno&gt;&gt;24);
	zaddr(g1, sf);
	zaddr(g2, st);

out:
	if(a != AGLOBL &amp;&amp; a != ADATA)
		pc++;
}

void
outhist(void)
{
	Gen g;
	Hist *h;
	char *p, *q, *op, c;
	int n;

	g = nullgen;
	c = &#39;/&#39;;
	for(h = hist; h != H; h = h-&gt;link) {
		p = h-&gt;name;
		op = 0;
		/* on windows skip drive specifier in pathname */
		if(systemtype(Windows) &amp;&amp; p &amp;&amp; p[1] == &#39;:&#39;){
			p += 2;
			c = *p;
		}
		if(p &amp;&amp; p[0] != c &amp;&amp; h-&gt;offset == 0 &amp;&amp; pathname){
			/* on windows skip drive specifier in pathname */
			if(systemtype(Windows) &amp;&amp; pathname[1] == &#39;:&#39;) {
				op = p;
				p = pathname+2;
				c = *p;
			} else if(pathname[0] == c){
				op = p;
				p = pathname;
			}
		}
		while(p) {
			q = strchr(p, c);
			if(q) {
				n = q-p;
				if(n == 0){
					n = 1;	/* leading &#34;/&#34; */
					*p = &#39;/&#39;;	/* don&#39;t emit &#34;\&#34; on windows */
				}
				q++;
			} else {
				n = strlen(p);
				q = 0;
			}
			if(n) {
				Bputc(&amp;obuf, ANAME);
				Bputc(&amp;obuf, D_FILE);	/* type */
				Bputc(&amp;obuf, 1);	/* sym */
				Bputc(&amp;obuf, &#39;&lt;&#39;);
				Bwrite(&amp;obuf, p, n);
				Bputc(&amp;obuf, 0);
			}
			p = q;
			if(p == 0 &amp;&amp; op) {
				p = op;
				op = 0;
			}
		}
		g.offset = h-&gt;offset;

		Bputc(&amp;obuf, AHISTORY);
		Bputc(&amp;obuf, Always);
		Bputc(&amp;obuf, 0);
		Bputc(&amp;obuf, h-&gt;line);
		Bputc(&amp;obuf, h-&gt;line&gt;&gt;8);
		Bputc(&amp;obuf, h-&gt;line&gt;&gt;16);
		Bputc(&amp;obuf, h-&gt;line&gt;&gt;24);
		zaddr(&amp;nullgen, 0);
		zaddr(&amp;g, 0);
	}
}

#include &#34;../cc/lexbody&#34;
#include &#34;../cc/macbody&#34;
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
