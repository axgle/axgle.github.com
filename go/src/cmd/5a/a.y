<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5a/a.y</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5a/a.y</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5a/a.y
// http://code.google.com/p/inferno-os/source/browse/utils/5a/a.y
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

%{
#include &#34;a.h&#34;
%}
%union
{
	Sym	*sym;
	int32	lval;
	double	dval;
	char	sval[8];
	Gen	gen;
}
%left	&#39;|&#39;
%left	&#39;^&#39;
%left	&#39;&amp;&#39;
%left	&#39;&lt;&#39; &#39;&gt;&#39;
%left	&#39;+&#39; &#39;-&#39;
%left	&#39;*&#39; &#39;/&#39; &#39;%&#39;
%token	&lt;lval&gt;	LTYPE1 LTYPE2 LTYPE3 LTYPE4 LTYPE5
%token	&lt;lval&gt;	LTYPE6 LTYPE7 LTYPE8 LTYPE9 LTYPEA
%token	&lt;lval&gt;	LTYPEB LTYPEC LTYPED LTYPEE LTYPEF
%token	&lt;lval&gt;	LTYPEG LTYPEH LTYPEI LTYPEJ LTYPEK
%token	&lt;lval&gt;	LTYPEL LTYPEM LTYPEN LTYPEBX
%token	&lt;lval&gt;	LCONST LSP LSB LFP LPC
%token	&lt;lval&gt;	LTYPEX LR LREG LF LFREG LC LCREG LPSR LFCR
%token	&lt;lval&gt;	LCOND LS LAT
%token	&lt;dval&gt;	LFCONST
%token	&lt;sval&gt;	LSCONST
%token	&lt;sym&gt;	LNAME LLAB LVAR
%type	&lt;lval&gt;	con expr oexpr pointer offset sreg spreg creg
%type	&lt;lval&gt;	rcon cond reglist
%type	&lt;gen&gt;	gen rel reg regreg freg shift fcon frcon
%type	&lt;gen&gt;	imm ximm name oreg ireg nireg ioreg imsr
%%
prog:
|	prog line

line:
	LLAB &#39;:&#39;
	{
		if($1-&gt;value != pc)
			yyerror(&#34;redeclaration of %s&#34;, $1-&gt;name);
		$1-&gt;value = pc;
	}
	line
|	LNAME &#39;:&#39;
	{
		$1-&gt;type = LLAB;
		$1-&gt;value = pc;
	}
	line
|	LNAME &#39;=&#39; expr &#39;;&#39;
	{
		$1-&gt;type = LVAR;
		$1-&gt;value = $3;
	}
|	LVAR &#39;=&#39; expr &#39;;&#39;
	{
		if($1-&gt;value != $3)
			yyerror(&#34;redeclaration of %s&#34;, $1-&gt;name);
		$1-&gt;value = $3;
	}
|	&#39;;&#39;
|	inst &#39;;&#39;
|	error &#39;;&#39;

inst:
/*
 * ADD
 */
	LTYPE1 cond imsr &#39;,&#39; spreg &#39;,&#39; reg
	{
		outcode($1, $2, &amp;$3, $5, &amp;$7);
	}
|	LTYPE1 cond imsr &#39;,&#39; spreg &#39;,&#39;
	{
		outcode($1, $2, &amp;$3, $5, &amp;nullgen);
	}
|	LTYPE1 cond imsr &#39;,&#39; reg
	{
		outcode($1, $2, &amp;$3, NREG, &amp;$5);
	}
/*
 * MVN
 */
|	LTYPE2 cond imsr &#39;,&#39; reg
	{
		outcode($1, $2, &amp;$3, NREG, &amp;$5);
	}
/*
 * MOVW
 */
|	LTYPE3 cond gen &#39;,&#39; gen
	{
		outcode($1, $2, &amp;$3, NREG, &amp;$5);
	}
/*
 * B/BL
 */
|	LTYPE4 cond comma rel
	{
		outcode($1, $2, &amp;nullgen, NREG, &amp;$4);
	}
|	LTYPE4 cond comma nireg
	{
		outcode($1, $2, &amp;nullgen, NREG, &amp;$4);
	}
/*
 * BX
 */
|	LTYPEBX comma ireg
	{
		outcode($1, Always, &amp;nullgen, NREG, &amp;$3);
	}
/*
 * BEQ
 */
|	LTYPE5 comma rel
	{
		outcode($1, Always, &amp;nullgen, NREG, &amp;$3);
	}
/*
 * SWI
 */
|	LTYPE6 cond comma gen
	{
		outcode($1, $2, &amp;nullgen, NREG, &amp;$4);
	}
/*
 * CMP
 */
|	LTYPE7 cond imsr &#39;,&#39; spreg comma
	{
		outcode($1, $2, &amp;$3, $5, &amp;nullgen);
	}
/*
 * MOVM
 */
|	LTYPE8 cond ioreg &#39;,&#39; &#39;[&#39; reglist &#39;]&#39;
	{
		Gen g;

		g = nullgen;
		g.type = D_CONST;
		g.offset = $6;
		outcode($1, $2, &amp;$3, NREG, &amp;g);
	}
|	LTYPE8 cond &#39;[&#39; reglist &#39;]&#39; &#39;,&#39; ioreg
	{
		Gen g;

		g = nullgen;
		g.type = D_CONST;
		g.offset = $4;
		outcode($1, $2, &amp;g, NREG, &amp;$7);
	}
/*
 * SWAP
 */
|	LTYPE9 cond reg &#39;,&#39; ireg &#39;,&#39; reg
	{
		outcode($1, $2, &amp;$5, $3.reg, &amp;$7);
	}
|	LTYPE9 cond reg &#39;,&#39; ireg comma
	{
		outcode($1, $2, &amp;$5, $3.reg, &amp;$3);
	}
|	LTYPE9 cond comma ireg &#39;,&#39; reg
	{
		outcode($1, $2, &amp;$4, $6.reg, &amp;$6);
	}
/*
 * RET
 */
|	LTYPEA cond comma
	{
		outcode($1, $2, &amp;nullgen, NREG, &amp;nullgen);
	}
/*
 * TEXT/GLOBL
 */
|	LTYPEB name &#39;,&#39; imm
	{
		outcode($1, Always, &amp;$2, NREG, &amp;$4);
	}
|	LTYPEB name &#39;,&#39; con &#39;,&#39; imm
	{
		outcode($1, Always, &amp;$2, $4, &amp;$6);
	}
/*
 * DATA
 */
|	LTYPEC name &#39;/&#39; con &#39;,&#39; ximm
	{
		outcode($1, Always, &amp;$2, $4, &amp;$6);
	}
/*
 * CASE
 */
|	LTYPED cond reg comma
	{
		outcode($1, $2, &amp;$3, NREG, &amp;nullgen);
	}
/*
 * word
 */
|	LTYPEH comma ximm
	{
		outcode($1, Always, &amp;nullgen, NREG, &amp;$3);
	}
/*
 * floating-point coprocessor
 */
|	LTYPEI cond freg &#39;,&#39; freg
	{
		outcode($1, $2, &amp;$3, NREG, &amp;$5);
	}
|	LTYPEK cond frcon &#39;,&#39; freg
	{
		outcode($1, $2, &amp;$3, NREG, &amp;$5);
	}
|	LTYPEK cond frcon &#39;,&#39; LFREG &#39;,&#39; freg
	{
		outcode($1, $2, &amp;$3, $5, &amp;$7);
	}
|	LTYPEL cond freg &#39;,&#39; freg comma
	{
		outcode($1, $2, &amp;$3, $5.reg, &amp;nullgen);
	}
/*
 * MCR MRC
 */
|	LTYPEJ cond con &#39;,&#39; expr &#39;,&#39; spreg &#39;,&#39; creg &#39;,&#39; creg oexpr
	{
		Gen g;

		g = nullgen;
		g.type = D_CONST;
		g.offset =
			(0xe &lt;&lt; 24) |		/* opcode */
			($1 &lt;&lt; 20) |		/* MCR/MRC */
			($2 &lt;&lt; 28) |		/* scond */
			(($3 &amp; 15) &lt;&lt; 8) |	/* coprocessor number */
			(($5 &amp; 7) &lt;&lt; 21) |	/* coprocessor operation */
			(($7 &amp; 15) &lt;&lt; 12) |	/* arm register */
			(($9 &amp; 15) &lt;&lt; 16) |	/* Crn */
			(($11 &amp; 15) &lt;&lt; 0) |	/* Crm */
			(($12 &amp; 7) &lt;&lt; 5) |	/* coprocessor information */
			(1&lt;&lt;4);			/* must be set */
		outcode(AWORD, Always, &amp;nullgen, NREG, &amp;g);
	}
/*
 * MULL hi,lo,r1,r2
 */
|	LTYPEM cond reg &#39;,&#39; reg &#39;,&#39; regreg
	{
		outcode($1, $2, &amp;$3, $5.reg, &amp;$7);
	}
/*
 * MULA hi,lo,r1,r2
 */
|	LTYPEN cond reg &#39;,&#39; reg &#39;,&#39; reg &#39;,&#39; spreg
	{
		$7.type = D_REGREG;
		$7.offset = $9;
		outcode($1, $2, &amp;$3, $5.reg, &amp;$7);
	}
/*
 * END
 */
|	LTYPEE comma
	{
		outcode($1, Always, &amp;nullgen, NREG, &amp;nullgen);
	}

cond:
	{
		$$ = Always;
	}
|	cond LCOND
	{
		$$ = ($1 &amp; ~C_SCOND) | $2;
	}
|	cond LS
	{
		$$ = $1 | $2;
	}

comma:
|	&#39;,&#39; comma

rel:
	con &#39;(&#39; LPC &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_BRANCH;
		$$.offset = $1 + pc;
	}
|	LNAME offset
	{
		$$ = nullgen;
		if(pass == 2)
			yyerror(&#34;undefined label: %s&#34;, $1-&gt;name);
		$$.type = D_BRANCH;
		$$.sym = $1;
		$$.offset = $2;
	}
|	LLAB offset
	{
		$$ = nullgen;
		$$.type = D_BRANCH;
		$$.sym = $1;
		$$.offset = $1-&gt;value + $2;
	}

ximm:	&#39;$&#39; con
	{
		$$ = nullgen;
		$$.type = D_CONST;
		$$.offset = $2;
	}
|	&#39;$&#39; oreg
	{
		$$ = $2;
		$$.type = D_CONST;
	}
|	&#39;$&#39; &#39;*&#39; &#39;$&#39; oreg
	{
		$$ = $4;
		$$.type = D_OCONST;
	}
|	&#39;$&#39; LSCONST
	{
		$$ = nullgen;
		$$.type = D_SCONST;
		memcpy($$.sval, $2, sizeof($$.sval));
	}
|	fcon

fcon:
	&#39;$&#39; LFCONST
	{
		$$ = nullgen;
		$$.type = D_FCONST;
		$$.dval = $2;
	}
|	&#39;$&#39; &#39;-&#39; LFCONST
	{
		$$ = nullgen;
		$$.type = D_FCONST;
		$$.dval = -$3;
	}

reglist:
	spreg
	{
		$$ = 1 &lt;&lt; $1;
	}
|	spreg &#39;-&#39; spreg
	{
		int i;
		$$=0;
		for(i=$1; i&lt;=$3; i++)
			$$ |= 1&lt;&lt;i;
		for(i=$3; i&lt;=$1; i++)
			$$ |= 1&lt;&lt;i;
	}
|	spreg comma reglist
	{
		$$ = (1&lt;&lt;$1) | $3;
	}

gen:
	reg
|	ximm
|	shift
|	shift &#39;(&#39; spreg &#39;)&#39;
	{
		$$ = $1;
		$$.reg = $3;
	}
|	LPSR
	{
		$$ = nullgen;
		$$.type = D_PSR;
		$$.reg = $1;
	}
|	LFCR
	{
		$$ = nullgen;
		$$.type = D_FPCR;
		$$.reg = $1;
	}
|	con
	{
		$$ = nullgen;
		$$.type = D_OREG;
		$$.offset = $1;
	}
|	oreg
|	freg

nireg:
	ireg
|	name
	{
		$$ = $1;
		if($1.name != D_EXTERN &amp;&amp; $1.name != D_STATIC) {
		}
	}

ireg:
	&#39;(&#39; spreg &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_OREG;
		$$.reg = $2;
		$$.offset = 0;
	}

ioreg:
	ireg
|	con &#39;(&#39; sreg &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_OREG;
		$$.reg = $3;
		$$.offset = $1;
	}

oreg:
	name
|	name &#39;(&#39; sreg &#39;)&#39;
	{
		$$ = $1;
		$$.type = D_OREG;
		$$.reg = $3;
	}
|	ioreg

imsr:
	reg
|	imm
|	shift

imm:	&#39;$&#39; con
	{
		$$ = nullgen;
		$$.type = D_CONST;
		$$.offset = $2;
	}

reg:
	spreg
	{
		$$ = nullgen;
		$$.type = D_REG;
		$$.reg = $1;
	}

regreg:
	&#39;(&#39; spreg &#39;,&#39; spreg &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_REGREG;
		$$.reg = $2;
		$$.offset = $4;
	}

shift:
	spreg &#39;&lt;&#39; &#39;&lt;&#39; rcon
	{
		$$ = nullgen;
		$$.type = D_SHIFT;
		$$.offset = $1 | $4 | (0 &lt;&lt; 5);
	}
|	spreg &#39;&gt;&#39; &#39;&gt;&#39; rcon
	{
		$$ = nullgen;
		$$.type = D_SHIFT;
		$$.offset = $1 | $4 | (1 &lt;&lt; 5);
	}
|	spreg &#39;-&#39; &#39;&gt;&#39; rcon
	{
		$$ = nullgen;
		$$.type = D_SHIFT;
		$$.offset = $1 | $4 | (2 &lt;&lt; 5);
	}
|	spreg LAT &#39;&gt;&#39; rcon
	{
		$$ = nullgen;
		$$.type = D_SHIFT;
		$$.offset = $1 | $4 | (3 &lt;&lt; 5);
	}

rcon:
	spreg
	{
		if($$ &lt; 0 || $$ &gt;= 16)
			print(&#34;register value out of range\n&#34;);
		$$ = (($1&amp;15) &lt;&lt; 8) | (1 &lt;&lt; 4);
	}
|	con
	{
		if($$ &lt; 0 || $$ &gt;= 32)
			print(&#34;shift value out of range\n&#34;);
		$$ = ($1&amp;31) &lt;&lt; 7;
	}

sreg:
	LREG
|	LPC
	{
		$$ = REGPC;
	}
|	LR &#39;(&#39; expr &#39;)&#39;
	{
		if($3 &lt; 0 || $3 &gt;= NREG)
			print(&#34;register value out of range\n&#34;);
		$$ = $3;
	}

spreg:
	sreg
|	LSP
	{
		$$ = REGSP;
	}

creg:
	LCREG
|	LC &#39;(&#39; expr &#39;)&#39;
	{
		if($3 &lt; 0 || $3 &gt;= NREG)
			print(&#34;register value out of range\n&#34;);
		$$ = $3;
	}

frcon:
	freg
|	fcon

freg:
	LFREG
	{
		$$ = nullgen;
		$$.type = D_FREG;
		$$.reg = $1;
	}
|	LF &#39;(&#39; con &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_FREG;
		$$.reg = $3;
	}

name:
	con &#39;(&#39; pointer &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_OREG;
		$$.name = $3;
		$$.sym = S;
		$$.offset = $1;
	}
|	LNAME offset &#39;(&#39; pointer &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_OREG;
		$$.name = $4;
		$$.sym = $1;
		$$.offset = $2;
	}
|	LNAME &#39;&lt;&#39; &#39;&gt;&#39; offset &#39;(&#39; LSB &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_OREG;
		$$.name = D_STATIC;
		$$.sym = $1;
		$$.offset = $4;
	}

offset:
	{
		$$ = 0;
	}
|	&#39;+&#39; con
	{
		$$ = $2;
	}
|	&#39;-&#39; con
	{
		$$ = -$2;
	}

pointer:
	LSB
|	LSP
|	LFP

con:
	LCONST
|	LVAR
	{
		$$ = $1-&gt;value;
	}
|	&#39;-&#39; con
	{
		$$ = -$2;
	}
|	&#39;+&#39; con
	{
		$$ = $2;
	}
|	&#39;~&#39; con
	{
		$$ = ~$2;
	}
|	&#39;(&#39; expr &#39;)&#39;
	{
		$$ = $2;
	}

oexpr:
	{
		$$ = 0;
	}
|	&#39;,&#39; expr
	{
		$$ = $2;
	}

expr:
	con
|	expr &#39;+&#39; expr
	{
		$$ = $1 + $3;
	}
|	expr &#39;-&#39; expr
	{
		$$ = $1 - $3;
	}
|	expr &#39;*&#39; expr
	{
		$$ = $1 * $3;
	}
|	expr &#39;/&#39; expr
	{
		$$ = $1 / $3;
	}
|	expr &#39;%&#39; expr
	{
		$$ = $1 % $3;
	}
|	expr &#39;&lt;&#39; &#39;&lt;&#39; expr
	{
		$$ = $1 &lt;&lt; $4;
	}
|	expr &#39;&gt;&#39; &#39;&gt;&#39; expr
	{
		$$ = $1 &gt;&gt; $4;
	}
|	expr &#39;&amp;&#39; expr
	{
		$$ = $1 &amp; $3;
	}
|	expr &#39;^&#39; expr
	{
		$$ = $1 ^ $3;
	}
|	expr &#39;|&#39; expr
	{
		$$ = $1 | $3;
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
