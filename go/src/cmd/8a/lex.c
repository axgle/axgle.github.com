<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8a/lex.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8a/lex.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8a/lex.c
// http://code.google.com/p/inferno-os/source/browse/utils/8a/lex.c
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

int
pathchar(void)
{
	return &#39;/&#39;;
}

void
main(int argc, char *argv[])
{
	char *p;
	int nout, nproc, i, c;

	thechar = &#39;8&#39;;
	thestring = &#34;386&#34;;
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
	p = utfrrune(ofile, pathchar());
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

	&#34;AL&#34;,		LBREG,	D_AL,
	&#34;CL&#34;,		LBREG,	D_CL,
	&#34;DL&#34;,		LBREG,	D_DL,
	&#34;BL&#34;,		LBREG,	D_BL,
	&#34;AH&#34;,		LBREG,	D_AH,
	&#34;CH&#34;,		LBREG,	D_CH,
	&#34;DH&#34;,		LBREG,	D_DH,
	&#34;BH&#34;,		LBREG,	D_BH,

	&#34;AX&#34;,		LLREG,	D_AX,
	&#34;CX&#34;,		LLREG,	D_CX,
	&#34;DX&#34;,		LLREG,	D_DX,
	&#34;BX&#34;,		LLREG,	D_BX,
/*	&#34;SP&#34;,		LLREG,	D_SP,	*/
	&#34;BP&#34;,		LLREG,	D_BP,
	&#34;SI&#34;,		LLREG,	D_SI,
	&#34;DI&#34;,		LLREG,	D_DI,

	&#34;F0&#34;,		LFREG,	D_F0+0,
	&#34;F1&#34;,		LFREG,	D_F0+1,
	&#34;F2&#34;,		LFREG,	D_F0+2,
	&#34;F3&#34;,		LFREG,	D_F0+3,
	&#34;F4&#34;,		LFREG,	D_F0+4,
	&#34;F5&#34;,		LFREG,	D_F0+5,
	&#34;F6&#34;,		LFREG,	D_F0+6,
	&#34;F7&#34;,		LFREG,	D_F0+7,

	&#34;CS&#34;,		LSREG,	D_CS,
	&#34;SS&#34;,		LSREG,	D_SS,
	&#34;DS&#34;,		LSREG,	D_DS,
	&#34;ES&#34;,		LSREG,	D_ES,
	&#34;FS&#34;,		LSREG,	D_FS,
	&#34;GS&#34;,		LSREG,	D_GS,

	&#34;GDTR&#34;,		LBREG,	D_GDTR,
	&#34;IDTR&#34;,		LBREG,	D_IDTR,
	&#34;LDTR&#34;,		LBREG,	D_LDTR,
	&#34;MSW&#34;,		LBREG,	D_MSW,
	&#34;TASK&#34;,		LBREG,	D_TASK,

	&#34;CR0&#34;,		LBREG,	D_CR+0,
	&#34;CR1&#34;,		LBREG,	D_CR+1,
	&#34;CR2&#34;,		LBREG,	D_CR+2,
	&#34;CR3&#34;,		LBREG,	D_CR+3,
	&#34;CR4&#34;,		LBREG,	D_CR+4,
	&#34;CR5&#34;,		LBREG,	D_CR+5,
	&#34;CR6&#34;,		LBREG,	D_CR+6,
	&#34;CR7&#34;,		LBREG,	D_CR+7,

	&#34;DR0&#34;,		LBREG,	D_DR+0,
	&#34;DR1&#34;,		LBREG,	D_DR+1,
	&#34;DR2&#34;,		LBREG,	D_DR+2,
	&#34;DR3&#34;,		LBREG,	D_DR+3,
	&#34;DR4&#34;,		LBREG,	D_DR+4,
	&#34;DR5&#34;,		LBREG,	D_DR+5,
	&#34;DR6&#34;,		LBREG,	D_DR+6,
	&#34;DR7&#34;,		LBREG,	D_DR+7,

	&#34;TR0&#34;,		LBREG,	D_TR+0,
	&#34;TR1&#34;,		LBREG,	D_TR+1,
	&#34;TR2&#34;,		LBREG,	D_TR+2,
	&#34;TR3&#34;,		LBREG,	D_TR+3,
	&#34;TR4&#34;,		LBREG,	D_TR+4,
	&#34;TR5&#34;,		LBREG,	D_TR+5,
	&#34;TR6&#34;,		LBREG,	D_TR+6,
	&#34;TR7&#34;,		LBREG,	D_TR+7,

	&#34;AAA&#34;,		LTYPE0,	AAAA,
	&#34;AAD&#34;,		LTYPE0,	AAAD,
	&#34;AAM&#34;,		LTYPE0,	AAAM,
	&#34;AAS&#34;,		LTYPE0,	AAAS,
	&#34;ADCB&#34;,		LTYPE3,	AADCB,
	&#34;ADCL&#34;,		LTYPE3,	AADCL,
	&#34;ADCW&#34;,		LTYPE3,	AADCW,
	&#34;ADDB&#34;,		LTYPE3,	AADDB,
	&#34;ADDL&#34;,		LTYPE3,	AADDL,
	&#34;ADDW&#34;,		LTYPE3,	AADDW,
	&#34;ADJSP&#34;,	LTYPE2,	AADJSP,
	&#34;ANDB&#34;,		LTYPE3,	AANDB,
	&#34;ANDL&#34;,		LTYPE3,	AANDL,
	&#34;ANDW&#34;,		LTYPE3,	AANDW,
	&#34;ARPL&#34;,		LTYPE3,	AARPL,
	&#34;BOUNDL&#34;,	LTYPE3,	ABOUNDL,
	&#34;BOUNDW&#34;,	LTYPE3,	ABOUNDW,
	&#34;BSFL&#34;,		LTYPE3,	ABSFL,
	&#34;BSFW&#34;,		LTYPE3,	ABSFW,
	&#34;BSRL&#34;,		LTYPE3,	ABSRL,
	&#34;BSRW&#34;,		LTYPE3,	ABSRW,
	&#34;BTCL&#34;,		LTYPE3,	ABTCL,
	&#34;BTCW&#34;,		LTYPE3,	ABTCW,
	&#34;BTL&#34;,		LTYPE3,	ABTL,
	&#34;BTRL&#34;,		LTYPE3,	ABTRL,
	&#34;BTRW&#34;,		LTYPE3,	ABTRW,
	&#34;BTSL&#34;,		LTYPE3,	ABTSL,
	&#34;BTSW&#34;,		LTYPE3,	ABTSW,
	&#34;BTW&#34;,		LTYPE3,	ABTW,
	&#34;BYTE&#34;,		LTYPE2,	ABYTE,
	&#34;CALL&#34;,		LTYPEC,	ACALL,
	&#34;CLC&#34;,		LTYPE0,	ACLC,
	&#34;CLD&#34;,		LTYPE0,	ACLD,
	&#34;CLI&#34;,		LTYPE0,	ACLI,
	&#34;CLTS&#34;,		LTYPE0,	ACLTS,
	&#34;CMC&#34;,		LTYPE0,	ACMC,
	&#34;CMPB&#34;,		LTYPE4,	ACMPB,
	&#34;CMPL&#34;,		LTYPE4,	ACMPL,
	&#34;CMPW&#34;,		LTYPE4,	ACMPW,
	&#34;CMPSB&#34;,	LTYPE0,	ACMPSB,
	&#34;CMPSL&#34;,	LTYPE0,	ACMPSL,
	&#34;CMPSW&#34;,	LTYPE0,	ACMPSW,
	&#34;CMPXCHGB&#34;,	LTYPE3,	ACMPXCHGB,
	&#34;CMPXCHGL&#34;,	LTYPE3,	ACMPXCHGL,
	&#34;CMPXCHGW&#34;,	LTYPE3,	ACMPXCHGW,
	&#34;DAA&#34;,		LTYPE0,	ADAA,
	&#34;DAS&#34;,		LTYPE0,	ADAS,
	&#34;DATA&#34;,		LTYPED,	ADATA,
	&#34;DECB&#34;,		LTYPE1,	ADECB,
	&#34;DECL&#34;,		LTYPE1,	ADECL,
	&#34;DECW&#34;,		LTYPE1,	ADECW,
	&#34;DIVB&#34;,		LTYPE2,	ADIVB,
	&#34;DIVL&#34;,		LTYPE2,	ADIVL,
	&#34;DIVW&#34;,		LTYPE2,	ADIVW,
	&#34;END&#34;,		LTYPE0,	AEND,
	&#34;ENTER&#34;,	LTYPE2,	AENTER,
	&#34;GLOBL&#34;,	LTYPEG,	AGLOBL,
	&#34;HLT&#34;,		LTYPE0,	AHLT,
	&#34;IDIVB&#34;,	LTYPE2,	AIDIVB,
	&#34;IDIVL&#34;,	LTYPE2,	AIDIVL,
	&#34;IDIVW&#34;,	LTYPE2,	AIDIVW,
	&#34;IMULB&#34;,	LTYPE2,	AIMULB,
	&#34;IMULL&#34;,	LTYPE2,	AIMULL,
	&#34;IMULW&#34;,	LTYPE2,	AIMULW,
	&#34;INB&#34;,		LTYPE0,	AINB,
	&#34;INL&#34;,		LTYPE0,	AINL,
	&#34;INW&#34;,		LTYPE0,	AINW,
	&#34;INCB&#34;,		LTYPE1,	AINCB,
	&#34;INCL&#34;,		LTYPE1,	AINCL,
	&#34;INCW&#34;,		LTYPE1,	AINCW,
	&#34;INSB&#34;,		LTYPE0,	AINSB,
	&#34;INSL&#34;,		LTYPE0,	AINSL,
	&#34;INSW&#34;,		LTYPE0,	AINSW,
	&#34;INT&#34;,		LTYPE2,	AINT,
	&#34;INTO&#34;,		LTYPE0,	AINTO,
	&#34;IRETL&#34;,	LTYPE0,	AIRETL,
	&#34;IRETW&#34;,	LTYPE0,	AIRETW,

	&#34;JOS&#34;,		LTYPER,	AJOS,
	&#34;JO&#34;,		LTYPER,	AJOS,	/* alternate */
	&#34;JOC&#34;,		LTYPER,	AJOC,
	&#34;JNO&#34;,		LTYPER,	AJOC,	/* alternate */
	&#34;JCS&#34;,		LTYPER,	AJCS,
	&#34;JB&#34;,		LTYPER,	AJCS,	/* alternate */
	&#34;JC&#34;,		LTYPER,	AJCS,	/* alternate */
	&#34;JNAE&#34;,		LTYPER,	AJCS,	/* alternate */
	&#34;JLO&#34;,		LTYPER,	AJCS,	/* alternate */
	&#34;JCC&#34;,		LTYPER,	AJCC,
	&#34;JAE&#34;,		LTYPER,	AJCC,	/* alternate */
	&#34;JNB&#34;,		LTYPER,	AJCC,	/* alternate */
	&#34;JNC&#34;,		LTYPER,	AJCC,	/* alternate */
	&#34;JHS&#34;,		LTYPER,	AJCC,	/* alternate */
	&#34;JEQ&#34;,		LTYPER,	AJEQ,
	&#34;JE&#34;,		LTYPER,	AJEQ,	/* alternate */
	&#34;JZ&#34;,		LTYPER,	AJEQ,	/* alternate */
	&#34;JNE&#34;,		LTYPER,	AJNE,
	&#34;JNZ&#34;,		LTYPER,	AJNE,	/* alternate */
	&#34;JLS&#34;,		LTYPER,	AJLS,
	&#34;JBE&#34;,		LTYPER,	AJLS,	/* alternate */
	&#34;JNA&#34;,		LTYPER,	AJLS,	/* alternate */
	&#34;JHI&#34;,		LTYPER,	AJHI,
	&#34;JA&#34;,		LTYPER,	AJHI,	/* alternate */
	&#34;JNBE&#34;,		LTYPER,	AJHI,	/* alternate */
	&#34;JMI&#34;,		LTYPER,	AJMI,
	&#34;JS&#34;,		LTYPER,	AJMI,	/* alternate */
	&#34;JPL&#34;,		LTYPER,	AJPL,
	&#34;JNS&#34;,		LTYPER,	AJPL,	/* alternate */
	&#34;JPS&#34;,		LTYPER,	AJPS,
	&#34;JP&#34;,		LTYPER,	AJPS,	/* alternate */
	&#34;JPE&#34;,		LTYPER,	AJPS,	/* alternate */
	&#34;JPC&#34;,		LTYPER,	AJPC,
	&#34;JNP&#34;,		LTYPER,	AJPC,	/* alternate */
	&#34;JPO&#34;,		LTYPER,	AJPC,	/* alternate */
	&#34;JLT&#34;,		LTYPER,	AJLT,
	&#34;JL&#34;,		LTYPER,	AJLT,	/* alternate */
	&#34;JNGE&#34;,		LTYPER,	AJLT,	/* alternate */
	&#34;JGE&#34;,		LTYPER,	AJGE,
	&#34;JNL&#34;,		LTYPER,	AJGE,	/* alternate */
	&#34;JLE&#34;,		LTYPER,	AJLE,
	&#34;JNG&#34;,		LTYPER,	AJLE,	/* alternate */
	&#34;JGT&#34;,		LTYPER,	AJGT,
	&#34;JG&#34;,		LTYPER,	AJGT,	/* alternate */
	&#34;JNLE&#34;,		LTYPER,	AJGT,	/* alternate */

	&#34;JCXZ&#34;,		LTYPER,	AJCXZ,
	&#34;JMP&#34;,		LTYPEC,	AJMP,
	&#34;LAHF&#34;,		LTYPE0,	ALAHF,
	&#34;LARL&#34;,		LTYPE3,	ALARL,
	&#34;LARW&#34;,		LTYPE3,	ALARW,
	&#34;LEAL&#34;,		LTYPE3,	ALEAL,
	&#34;LEAW&#34;,		LTYPE3,	ALEAW,
	&#34;LEAVEL&#34;,	LTYPE0,	ALEAVEL,
	&#34;LEAVEW&#34;,	LTYPE0,	ALEAVEW,
	&#34;LOCK&#34;,		LTYPE0,	ALOCK,
	&#34;LODSB&#34;,	LTYPE0,	ALODSB,
	&#34;LODSL&#34;,	LTYPE0,	ALODSL,
	&#34;LODSW&#34;,	LTYPE0,	ALODSW,
	&#34;LONG&#34;,		LTYPE2,	ALONG,
	&#34;LOOP&#34;,		LTYPER,	ALOOP,
	&#34;LOOPEQ&#34;,	LTYPER,	ALOOPEQ,
	&#34;LOOPNE&#34;,	LTYPER,	ALOOPNE,
	&#34;LSLL&#34;,		LTYPE3,	ALSLL,
	&#34;LSLW&#34;,		LTYPE3,	ALSLW,
	&#34;MOVB&#34;,		LTYPE3,	AMOVB,
	&#34;MOVL&#34;,		LTYPEM,	AMOVL,
	&#34;MOVW&#34;,		LTYPEM,	AMOVW,
	&#34;MOVBLSX&#34;,	LTYPE3, AMOVBLSX,
	&#34;MOVBLZX&#34;,	LTYPE3, AMOVBLZX,
	&#34;MOVBWSX&#34;,	LTYPE3, AMOVBWSX,
	&#34;MOVBWZX&#34;,	LTYPE3, AMOVBWZX,
	&#34;MOVWLSX&#34;,	LTYPE3, AMOVWLSX,
	&#34;MOVWLZX&#34;,	LTYPE3, AMOVWLZX,
	&#34;MOVSB&#34;,	LTYPE0,	AMOVSB,
	&#34;MOVSL&#34;,	LTYPE0,	AMOVSL,
	&#34;MOVSW&#34;,	LTYPE0,	AMOVSW,
	&#34;MULB&#34;,		LTYPE2,	AMULB,
	&#34;MULL&#34;,		LTYPE2,	AMULL,
	&#34;MULW&#34;,		LTYPE2,	AMULW,
	&#34;NEGB&#34;,		LTYPE1,	ANEGB,
	&#34;NEGL&#34;,		LTYPE1,	ANEGL,
	&#34;NEGW&#34;,		LTYPE1,	ANEGW,
	&#34;NOP&#34;,		LTYPEN,	ANOP,
	&#34;NOTB&#34;,		LTYPE1,	ANOTB,
	&#34;NOTL&#34;,		LTYPE1,	ANOTL,
	&#34;NOTW&#34;,		LTYPE1,	ANOTW,
	&#34;ORB&#34;,		LTYPE3,	AORB,
	&#34;ORL&#34;,		LTYPE3,	AORL,
	&#34;ORW&#34;,		LTYPE3,	AORW,
	&#34;OUTB&#34;,		LTYPE0,	AOUTB,
	&#34;OUTL&#34;,		LTYPE0,	AOUTL,
	&#34;OUTW&#34;,		LTYPE0,	AOUTW,
	&#34;OUTSB&#34;,	LTYPE0,	AOUTSB,
	&#34;OUTSL&#34;,	LTYPE0,	AOUTSL,
	&#34;OUTSW&#34;,	LTYPE0,	AOUTSW,
	&#34;POPAL&#34;,	LTYPE0,	APOPAL,
	&#34;POPAW&#34;,	LTYPE0,	APOPAW,
	&#34;POPFL&#34;,	LTYPE0,	APOPFL,
	&#34;POPFW&#34;,	LTYPE0,	APOPFW,
	&#34;POPL&#34;,		LTYPE1,	APOPL,
	&#34;POPW&#34;,		LTYPE1,	APOPW,
	&#34;PUSHAL&#34;,	LTYPE0,	APUSHAL,
	&#34;PUSHAW&#34;,	LTYPE0,	APUSHAW,
	&#34;PUSHFL&#34;,	LTYPE0,	APUSHFL,
	&#34;PUSHFW&#34;,	LTYPE0,	APUSHFW,
	&#34;PUSHL&#34;,	LTYPE2,	APUSHL,
	&#34;PUSHW&#34;,	LTYPE2,	APUSHW,
	&#34;RCLB&#34;,		LTYPE3,	ARCLB,
	&#34;RCLL&#34;,		LTYPE3,	ARCLL,
	&#34;RCLW&#34;,		LTYPE3,	ARCLW,
	&#34;RCRB&#34;,		LTYPE3,	ARCRB,
	&#34;RCRL&#34;,		LTYPE3,	ARCRL,
	&#34;RCRW&#34;,		LTYPE3,	ARCRW,
	&#34;REP&#34;,		LTYPE0,	AREP,
	&#34;REPN&#34;,		LTYPE0,	AREPN,
	&#34;RET&#34;,		LTYPE0,	ARET,
	&#34;ROLB&#34;,		LTYPE3,	AROLB,
	&#34;ROLL&#34;,		LTYPE3,	AROLL,
	&#34;ROLW&#34;,		LTYPE3,	AROLW,
	&#34;RORB&#34;,		LTYPE3,	ARORB,
	&#34;RORL&#34;,		LTYPE3,	ARORL,
	&#34;RORW&#34;,		LTYPE3,	ARORW,
	&#34;SAHF&#34;,		LTYPE0,	ASAHF,
	&#34;SALB&#34;,		LTYPE3,	ASALB,
	&#34;SALL&#34;,		LTYPE3,	ASALL,
	&#34;SALW&#34;,		LTYPE3,	ASALW,
	&#34;SARB&#34;,		LTYPE3,	ASARB,
	&#34;SARL&#34;,		LTYPE3,	ASARL,
	&#34;SARW&#34;,		LTYPE3,	ASARW,
	&#34;SBBB&#34;,		LTYPE3,	ASBBB,
	&#34;SBBL&#34;,		LTYPE3,	ASBBL,
	&#34;SBBW&#34;,		LTYPE3,	ASBBW,
	&#34;SCASB&#34;,	LTYPE0,	ASCASB,
	&#34;SCASL&#34;,	LTYPE0,	ASCASL,
	&#34;SCASW&#34;,	LTYPE0,	ASCASW,
	&#34;SETCC&#34;,	LTYPE1,	ASETCC,
	&#34;SETCS&#34;,	LTYPE1,	ASETCS,
	&#34;SETEQ&#34;,	LTYPE1,	ASETEQ,
	&#34;SETGE&#34;,	LTYPE1,	ASETGE,
	&#34;SETGT&#34;,	LTYPE1,	ASETGT,
	&#34;SETHI&#34;,	LTYPE1,	ASETHI,
	&#34;SETLE&#34;,	LTYPE1,	ASETLE,
	&#34;SETLS&#34;,	LTYPE1,	ASETLS,
	&#34;SETLT&#34;,	LTYPE1,	ASETLT,
	&#34;SETMI&#34;,	LTYPE1,	ASETMI,
	&#34;SETNE&#34;,	LTYPE1,	ASETNE,
	&#34;SETOC&#34;,	LTYPE1,	ASETOC,
	&#34;SETOS&#34;,	LTYPE1,	ASETOS,
	&#34;SETPC&#34;,	LTYPE1,	ASETPC,
	&#34;SETPL&#34;,	LTYPE1,	ASETPL,
	&#34;SETPS&#34;,	LTYPE1,	ASETPS,
	&#34;CDQ&#34;,		LTYPE0,	ACDQ,
	&#34;CWD&#34;,		LTYPE0,	ACWD,
	&#34;SHLB&#34;,		LTYPE3,	ASHLB,
	&#34;SHLL&#34;,		LTYPES,	ASHLL,
	&#34;SHLW&#34;,		LTYPES,	ASHLW,
	&#34;SHRB&#34;,		LTYPE3,	ASHRB,
	&#34;SHRL&#34;,		LTYPES,	ASHRL,
	&#34;SHRW&#34;,		LTYPES,	ASHRW,
	&#34;STC&#34;,		LTYPE0,	ASTC,
	&#34;STD&#34;,		LTYPE0,	ASTD,
	&#34;STI&#34;,		LTYPE0,	ASTI,
	&#34;STOSB&#34;,	LTYPE0,	ASTOSB,
	&#34;STOSL&#34;,	LTYPE0,	ASTOSL,
	&#34;STOSW&#34;,	LTYPE0,	ASTOSW,
	&#34;SUBB&#34;,		LTYPE3,	ASUBB,
	&#34;SUBL&#34;,		LTYPE3,	ASUBL,
	&#34;SUBW&#34;,		LTYPE3,	ASUBW,
	&#34;SYSCALL&#34;,	LTYPE0,	ASYSCALL,
	&#34;TESTB&#34;,	LTYPE3,	ATESTB,
	&#34;TESTL&#34;,	LTYPE3,	ATESTL,
	&#34;TESTW&#34;,	LTYPE3,	ATESTW,
	&#34;TEXT&#34;,		LTYPET,	ATEXT,
	&#34;VERR&#34;,		LTYPE2,	AVERR,
	&#34;VERW&#34;,		LTYPE2,	AVERW,
	&#34;WAIT&#34;,		LTYPE0,	AWAIT,
	&#34;WORD&#34;,		LTYPE2,	AWORD,
	&#34;XCHGB&#34;,	LTYPE3,	AXCHGB,
	&#34;XCHGL&#34;,	LTYPE3,	AXCHGL,
	&#34;XCHGW&#34;,	LTYPE3,	AXCHGW,
	&#34;XLAT&#34;,		LTYPE2,	AXLAT,
	&#34;XORB&#34;,		LTYPE3,	AXORB,
	&#34;XORL&#34;,		LTYPE3,	AXORL,
	&#34;XORW&#34;,		LTYPE3,	AXORW,

	&#34;FMOVB&#34;,	LTYPE3, AFMOVB,
	&#34;FMOVBP&#34;,	LTYPE3, AFMOVBP,
	&#34;FMOVD&#34;,	LTYPE3, AFMOVD,
	&#34;FMOVDP&#34;,	LTYPE3, AFMOVDP,
	&#34;FMOVF&#34;,	LTYPE3, AFMOVF,
	&#34;FMOVFP&#34;,	LTYPE3, AFMOVFP,
	&#34;FMOVL&#34;,	LTYPE3, AFMOVL,
	&#34;FMOVLP&#34;,	LTYPE3, AFMOVLP,
	&#34;FMOVV&#34;,	LTYPE3, AFMOVV,
	&#34;FMOVVP&#34;,	LTYPE3, AFMOVVP,
	&#34;FMOVW&#34;,	LTYPE3, AFMOVW,
	&#34;FMOVWP&#34;,	LTYPE3, AFMOVWP,
	&#34;FMOVX&#34;,	LTYPE3, AFMOVX,
	&#34;FMOVXP&#34;,	LTYPE3, AFMOVXP,
	&#34;FCOMB&#34;,	LTYPE3, AFCOMB,
	&#34;FCOMBP&#34;,	LTYPE3, AFCOMBP,
	&#34;FCOMD&#34;,	LTYPE3, AFCOMD,
	&#34;FCOMDP&#34;,	LTYPE3, AFCOMDP,
	&#34;FCOMDPP&#34;,	LTYPE3, AFCOMDPP,
	&#34;FCOMF&#34;,	LTYPE3, AFCOMF,
	&#34;FCOMFP&#34;,	LTYPE3, AFCOMFP,
	&#34;FCOML&#34;,	LTYPE3, AFCOML,
	&#34;FCOMLP&#34;,	LTYPE3, AFCOMLP,
	&#34;FCOMW&#34;,	LTYPE3, AFCOMW,
	&#34;FCOMWP&#34;,	LTYPE3, AFCOMWP,
	&#34;FUCOM&#34;,	LTYPE3, AFUCOM,
	&#34;FUCOMP&#34;,	LTYPE3, AFUCOMP,
	&#34;FUCOMPP&#34;,	LTYPE3, AFUCOMPP,
	&#34;FADDW&#34;,	LTYPE3, AFADDW,
	&#34;FADDL&#34;,	LTYPE3, AFADDL,
	&#34;FADDF&#34;,	LTYPE3, AFADDF,
	&#34;FADDD&#34;,	LTYPE3, AFADDD,
	&#34;FADDDP&#34;,	LTYPE3, AFADDDP,
	&#34;FSUBDP&#34;,	LTYPE3, AFSUBDP,
	&#34;FSUBW&#34;,	LTYPE3, AFSUBW,
	&#34;FSUBL&#34;,	LTYPE3, AFSUBL,
	&#34;FSUBF&#34;,	LTYPE3, AFSUBF,
	&#34;FSUBD&#34;,	LTYPE3, AFSUBD,
	&#34;FSUBRDP&#34;,	LTYPE3, AFSUBRDP,
	&#34;FSUBRW&#34;,	LTYPE3, AFSUBRW,
	&#34;FSUBRL&#34;,	LTYPE3, AFSUBRL,
	&#34;FSUBRF&#34;,	LTYPE3, AFSUBRF,
	&#34;FSUBRD&#34;,	LTYPE3, AFSUBRD,
	&#34;FMULDP&#34;,	LTYPE3, AFMULDP,
	&#34;FMULW&#34;,	LTYPE3, AFMULW,
	&#34;FMULL&#34;,	LTYPE3, AFMULL,
	&#34;FMULF&#34;,	LTYPE3, AFMULF,
	&#34;FMULD&#34;,	LTYPE3, AFMULD,
	&#34;FDIVDP&#34;,	LTYPE3, AFDIVDP,
	&#34;FDIVW&#34;,	LTYPE3, AFDIVW,
	&#34;FDIVL&#34;,	LTYPE3, AFDIVL,
	&#34;FDIVF&#34;,	LTYPE3, AFDIVF,
	&#34;FDIVD&#34;,	LTYPE3, AFDIVD,
	&#34;FDIVRDP&#34;,	LTYPE3, AFDIVRDP,
	&#34;FDIVRW&#34;,	LTYPE3, AFDIVRW,
	&#34;FDIVRL&#34;,	LTYPE3, AFDIVRL,
	&#34;FDIVRF&#34;,	LTYPE3, AFDIVRF,
	&#34;FDIVRD&#34;,	LTYPE3, AFDIVRD,
	&#34;FXCHD&#34;,	LTYPE3, AFXCHD,
	&#34;FFREE&#34;,	LTYPE1, AFFREE,
	&#34;FLDCW&#34;,	LTYPE2, AFLDCW,
	&#34;FLDENV&#34;,	LTYPE1, AFLDENV,
	&#34;FRSTOR&#34;,	LTYPE2, AFRSTOR,
	&#34;FSAVE&#34;,	LTYPE1, AFSAVE,
	&#34;FSTCW&#34;,	LTYPE1, AFSTCW,
	&#34;FSTENV&#34;,	LTYPE1, AFSTENV,
	&#34;FSTSW&#34;,	LTYPE1, AFSTSW,
	&#34;F2XM1&#34;,	LTYPE0, AF2XM1,
	&#34;FABS&#34;,		LTYPE0, AFABS,
	&#34;FCHS&#34;,		LTYPE0, AFCHS,
	&#34;FCLEX&#34;,	LTYPE0, AFCLEX,
	&#34;FCOS&#34;,		LTYPE0, AFCOS,
	&#34;FDECSTP&#34;,	LTYPE0, AFDECSTP,
	&#34;FINCSTP&#34;,	LTYPE0, AFINCSTP,
	&#34;FINIT&#34;,	LTYPE0, AFINIT,
	&#34;FLD1&#34;,		LTYPE0, AFLD1,
	&#34;FLDL2E&#34;,	LTYPE0, AFLDL2E,
	&#34;FLDL2T&#34;,	LTYPE0, AFLDL2T,
	&#34;FLDLG2&#34;,	LTYPE0, AFLDLG2,
	&#34;FLDLN2&#34;,	LTYPE0, AFLDLN2,
	&#34;FLDPI&#34;,	LTYPE0, AFLDPI,
	&#34;FLDZ&#34;,		LTYPE0, AFLDZ,
	&#34;FNOP&#34;,		LTYPE0, AFNOP,
	&#34;FPATAN&#34;,	LTYPE0, AFPATAN,
	&#34;FPREM&#34;,	LTYPE0, AFPREM,
	&#34;FPREM1&#34;,	LTYPE0, AFPREM1,
	&#34;FPTAN&#34;,	LTYPE0, AFPTAN,
	&#34;FRNDINT&#34;,	LTYPE0, AFRNDINT,
	&#34;FSCALE&#34;,	LTYPE0, AFSCALE,
	&#34;FSIN&#34;,		LTYPE0, AFSIN,
	&#34;FSINCOS&#34;,	LTYPE0, AFSINCOS,
	&#34;FSQRT&#34;,	LTYPE0, AFSQRT,
	&#34;FTST&#34;,		LTYPE0, AFTST,
	&#34;FXAM&#34;,		LTYPE0, AFXAM,
	&#34;FXTRACT&#34;,	LTYPE0, AFXTRACT,
	&#34;FYL2X&#34;,	LTYPE0, AFYL2X,
	&#34;FYL2XP1&#34;,	LTYPE0, AFYL2XP1,

	0
};

void
cinit(void)
{
	Sym *s;
	int i;

	nullgen.sym = S;
	nullgen.offset = 0;
	if(FPCHIP)
		nullgen.dval = 0;
	for(i=0; i&lt;sizeof(nullgen.sval); i++)
		nullgen.sval[i] = 0;
	nullgen.type = D_NONE;
	nullgen.index = D_NONE;
	nullgen.scale = 0;

	nerrors = 0;
	iostack = I;
	iofree = I;
	peekc = IGN;
	nhunk = 0;
	for(i=0; i&lt;NHASH; i++)
		hash[i] = S;
	for(i=0; itab[i].name; i++) {
		s = slookup(itab[i].name);
		if(s-&gt;type != LNAME)
			yyerror(&#34;double initialization %s&#34;, itab[i].name);
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
checkscale(int scale)
{

	switch(scale) {
	case 1:
	case 2:
	case 4:
	case 8:
		return;
	}
	yyerror(&#34;scale must be 1248: %d&#34;, scale);
}

void
syminit(Sym *s)
{

	s-&gt;type = LNAME;
	s-&gt;value = 0;
}

void
cclean(void)
{
	Gen2 g2;

	g2.from = nullgen;
	g2.to = nullgen;
	outcode(AEND, &amp;g2);
	Bflush(&amp;obuf);
}

void
zname(char *n, int t, int s)
{

	Bputc(&amp;obuf, ANAME);		/* as(2) */
	Bputc(&amp;obuf, ANAME&gt;&gt;8);
	Bputc(&amp;obuf, t);		/* type */
	Bputc(&amp;obuf, s);		/* sym */
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
	int i, t;
	char *n;
	Ieee e;

	t = 0;
	if(a-&gt;index != D_NONE || a-&gt;scale != 0)
		t |= T_INDEX;
	if(a-&gt;offset != 0)
		t |= T_OFFSET;
	if(s != 0)
		t |= T_SYM;

	switch(a-&gt;type) {
	default:
		t |= T_TYPE;
		break;
	case D_FCONST:
		t |= T_FCONST;
		break;
	case D_CONST2:
		t |= T_OFFSET|T_OFFSET2;
		break;
	case D_SCONST:
		t |= T_SCONST;
		break;
	case D_NONE:
		break;
	}
	Bputc(&amp;obuf, t);

	if(t &amp; T_INDEX) {	/* implies index, scale */
		Bputc(&amp;obuf, a-&gt;index);
		Bputc(&amp;obuf, a-&gt;scale);
	}
	if(t &amp; T_OFFSET) {	/* implies offset */
		l = a-&gt;offset;
		Bputc(&amp;obuf, l);
		Bputc(&amp;obuf, l&gt;&gt;8);
		Bputc(&amp;obuf, l&gt;&gt;16);
		Bputc(&amp;obuf, l&gt;&gt;24);
	}
	if(t &amp; T_OFFSET2) {
		l = a-&gt;offset2;
		Bputc(&amp;obuf, l);
		Bputc(&amp;obuf, l&gt;&gt;8);
		Bputc(&amp;obuf, l&gt;&gt;16);
		Bputc(&amp;obuf, l&gt;&gt;24);
	}
	if(t &amp; T_SYM)		/* implies sym */
		Bputc(&amp;obuf, s);
	if(t &amp; T_FCONST) {
		ieeedtod(&amp;e, a-&gt;dval);
		l = e.l;
		Bputc(&amp;obuf, l);
		Bputc(&amp;obuf, l&gt;&gt;8);
		Bputc(&amp;obuf, l&gt;&gt;16);
		Bputc(&amp;obuf, l&gt;&gt;24);
		l = e.h;
		Bputc(&amp;obuf, l);
		Bputc(&amp;obuf, l&gt;&gt;8);
		Bputc(&amp;obuf, l&gt;&gt;16);
		Bputc(&amp;obuf, l&gt;&gt;24);
		return;
	}
	if(t &amp; T_SCONST) {
		n = a-&gt;sval;
		for(i=0; i&lt;NSNAME; i++) {
			Bputc(&amp;obuf, *n);
			n++;
		}
		return;
	}
	if(t &amp; T_TYPE)
		Bputc(&amp;obuf, a-&gt;type);
}

void
outcode(int a, Gen2 *g2)
{
	int sf, st, t;
	Sym *s;

	if(pass == 1)
		goto out;

jackpot:
	sf = 0;
	s = g2-&gt;from.sym;
	while(s != S) {
		sf = s-&gt;sym;
		if(sf &lt; 0 || sf &gt;= NSYM)
			sf = 0;
		t = g2-&gt;from.type;
		if(t == D_ADDR)
			t = g2-&gt;from.index;
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
	s = g2-&gt;to.sym;
	while(s != S) {
		st = s-&gt;sym;
		if(st &lt; 0 || st &gt;= NSYM)
			st = 0;
		t = g2-&gt;to.type;
		if(t == D_ADDR)
			t = g2-&gt;to.index;
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
	Bputc(&amp;obuf, a&gt;&gt;8);
	Bputc(&amp;obuf, lineno);
	Bputc(&amp;obuf, lineno&gt;&gt;8);
	Bputc(&amp;obuf, lineno&gt;&gt;16);
	Bputc(&amp;obuf, lineno&gt;&gt;24);
	zaddr(&amp;g2-&gt;from, sf);
	zaddr(&amp;g2-&gt;to, st);

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
	c = pathchar();
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
				Bputc(&amp;obuf, ANAME&gt;&gt;8);
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
		Bputc(&amp;obuf, AHISTORY&gt;&gt;8);
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
