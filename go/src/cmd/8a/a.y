<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8a/a.y</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8a/a.y</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8a/a.y
// http://code.google.com/p/inferno-os/source/browse/utils/8a/a.y
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
%union	{
	Sym	*sym;
	int32	lval;
	struct {
		int32 v1;
		int32 v2;
	} con2;
	double	dval;
	char	sval[8];
	Gen	gen;
	Gen2	gen2;
}
%left	&#39;|&#39;
%left	&#39;^&#39;
%left	&#39;&amp;&#39;
%left	&#39;&lt;&#39; &#39;&gt;&#39;
%left	&#39;+&#39; &#39;-&#39;
%left	&#39;*&#39; &#39;/&#39; &#39;%&#39;
%token	&lt;lval&gt;	LTYPE0 LTYPE1 LTYPE2 LTYPE3 LTYPE4
%token	&lt;lval&gt;	LTYPEC LTYPED LTYPEN LTYPER LTYPET LTYPES LTYPEM LTYPEI LTYPEG
%token	&lt;lval&gt;	LCONST LFP LPC LSB
%token	&lt;lval&gt;	LBREG LLREG LSREG LFREG
%token	&lt;dval&gt;	LFCONST
%token	&lt;sval&gt;	LSCONST LSP
%token	&lt;sym&gt;	LNAME LLAB LVAR
%type	&lt;lval&gt;	con expr pointer offset
%type	&lt;con2&gt;	con2
%type	&lt;gen&gt;	mem imm imm2 reg nam rel rem rim rom omem nmem
%type	&lt;gen2&gt;	nonnon nonrel nonrem rimnon rimrem remrim
%type	&lt;gen2&gt;	spec1 spec2 spec3 spec4 spec5 spec6 spec7 spec8
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
|	&#39;;&#39;
|	inst &#39;;&#39;
|	error &#39;;&#39;

inst:
	LNAME &#39;=&#39; expr
	{
		$1-&gt;type = LVAR;
		$1-&gt;value = $3;
	}
|	LVAR &#39;=&#39; expr
	{
		if($1-&gt;value != $3)
			yyerror(&#34;redeclaration of %s&#34;, $1-&gt;name);
		$1-&gt;value = $3;
	}
|	LTYPE0 nonnon	{ outcode($1, &amp;$2); }
|	LTYPE1 nonrem	{ outcode($1, &amp;$2); }
|	LTYPE2 rimnon	{ outcode($1, &amp;$2); }
|	LTYPE3 rimrem	{ outcode($1, &amp;$2); }
|	LTYPE4 remrim	{ outcode($1, &amp;$2); }
|	LTYPER nonrel	{ outcode($1, &amp;$2); }
|	LTYPED spec1	{ outcode($1, &amp;$2); }
|	LTYPET spec2	{ outcode($1, &amp;$2); }
|	LTYPEC spec3	{ outcode($1, &amp;$2); }
|	LTYPEN spec4	{ outcode($1, &amp;$2); }
|	LTYPES spec5	{ outcode($1, &amp;$2); }
|	LTYPEM spec6	{ outcode($1, &amp;$2); }
|	LTYPEI spec7	{ outcode($1, &amp;$2); }
|	LTYPEG spec8	{ outcode($1, &amp;$2); }

nonnon:
	{
		$$.from = nullgen;
		$$.to = nullgen;
	}
|	&#39;,&#39;
	{
		$$.from = nullgen;
		$$.to = nullgen;
	}

rimrem:
	rim &#39;,&#39; rem
	{
		$$.from = $1;
		$$.to = $3;
	}

remrim:
	rem &#39;,&#39; rim
	{
		$$.from = $1;
		$$.to = $3;
	}

rimnon:
	rim &#39;,&#39;
	{
		$$.from = $1;
		$$.to = nullgen;
	}
|	rim
	{
		$$.from = $1;
		$$.to = nullgen;
	}

nonrem:
	&#39;,&#39; rem
	{
		$$.from = nullgen;
		$$.to = $2;
	}
|	rem
	{
		$$.from = nullgen;
		$$.to = $1;
	}

nonrel:
	&#39;,&#39; rel
	{
		$$.from = nullgen;
		$$.to = $2;
	}
|	rel
	{
		$$.from = nullgen;
		$$.to = $1;
	}

spec1:	/* DATA */
	nam &#39;/&#39; con &#39;,&#39; imm
	{
		$$.from = $1;
		$$.from.scale = $3;
		$$.to = $5;
	}

spec2:	/* TEXT */
	mem &#39;,&#39; imm2
	{
		$$.from = $1;
		$$.to = $3;
	}
|	mem &#39;,&#39; con &#39;,&#39; imm2
	{
		$$.from = $1;
		$$.from.scale = $3;
		$$.to = $5;
	}

spec3:	/* JMP/CALL */
	&#39;,&#39; rom
	{
		$$.from = nullgen;
		$$.to = $2;
	}
|	rom
	{
		$$.from = nullgen;
		$$.to = $1;
	}

spec4:	/* NOP */
	nonnon
|	nonrem

spec5:	/* SHL/SHR */
	rim &#39;,&#39; rem
	{
		$$.from = $1;
		$$.to = $3;
	}
|	rim &#39;,&#39; rem &#39;:&#39; LLREG
	{
		$$.from = $1;
		$$.to = $3;
		if($$.from.index != D_NONE)
			yyerror(&#34;dp shift with lhs index&#34;);
		$$.from.index = $5;
	}

spec6:	/* MOVW/MOVL */
	rim &#39;,&#39; rem
	{
		$$.from = $1;
		$$.to = $3;
	}
|	rim &#39;,&#39; rem &#39;:&#39; LSREG
	{
		$$.from = $1;
		$$.to = $3;
		if($$.to.index != D_NONE)
			yyerror(&#34;dp move with lhs index&#34;);
		$$.to.index = $5;
	}

spec7:
	rim &#39;,&#39;
	{
		$$.from = $1;
		$$.to = nullgen;
	}
|	rim
	{
		$$.from = $1;
		$$.to = nullgen;
	}
|	rim &#39;,&#39; rem
	{
		$$.from = $1;
		$$.to = $3;
	}

spec8:	/* GLOBL */
	mem &#39;,&#39; imm
	{
		$$.from = $1;
		$$.to = $3;
	}
|	mem &#39;,&#39; con &#39;,&#39; imm
	{
		$$.from = $1;
		$$.from.scale = $3;
		$$.to = $5;
	}

rem:
	reg
|	mem

rom:
	rel
|	nmem
|	&#39;*&#39; reg
	{
		$$ = $2;
	}
|	&#39;*&#39; omem
	{
		$$ = $2;
	}
|	reg
|	omem
|	imm

rim:
	rem
|	imm

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

reg:
	LBREG
	{
		$$ = nullgen;
		$$.type = $1;
	}
|	LFREG
	{
		$$ = nullgen;
		$$.type = $1;
	}
|	LLREG
	{
		$$ = nullgen;
		$$.type = $1;
	}
|	LSP
	{
		$$ = nullgen;
		$$.type = D_SP;
	}
|	LSREG
	{
		$$ = nullgen;
		$$.type = $1;
	}

imm:
	&#39;$&#39; con
	{
		$$ = nullgen;
		$$.type = D_CONST;
		$$.offset = $2;
	}
|	&#39;$&#39; nam
	{
		$$ = $2;
		$$.index = $2.type;
		$$.type = D_ADDR;
		/*
		if($2.type == D_AUTO || $2.type == D_PARAM)
			yyerror(&#34;constant cannot be automatic: %s&#34;,
				$2.sym-&gt;name);
		 */
	}
|	&#39;$&#39; LSCONST
	{
		$$ = nullgen;
		$$.type = D_SCONST;
		memcpy($$.sval, $2, sizeof($$.sval));
	}
|	&#39;$&#39; LFCONST
	{
		$$ = nullgen;
		$$.type = D_FCONST;
		$$.dval = $2;
	}
|	&#39;$&#39; &#39;(&#39; LFCONST &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_FCONST;
		$$.dval = $3;
	}
|	&#39;$&#39; &#39;-&#39; LFCONST
	{
		$$ = nullgen;
		$$.type = D_FCONST;
		$$.dval = -$3;
	}

imm2:
	&#39;$&#39; con2
	{
		$$ = nullgen;
		$$.type = D_CONST2;
		$$.offset = $2.v1;
		$$.offset2 = $2.v2;
	}

con2:
	LCONST
	{
		$$.v1 = $1;
		$$.v2 = 0;
	}
|	&#39;-&#39; LCONST
	{
		$$.v1 = -$2;
		$$.v2 = 0;
	}
|	LCONST &#39;-&#39; LCONST
	{
		$$.v1 = $1;
		$$.v2 = $3;
	}
|	&#39;-&#39; LCONST &#39;-&#39; LCONST
	{
		$$.v1 = -$2;
		$$.v2 = $4;
	}

mem:
	omem
|	nmem

omem:
	con
	{
		$$ = nullgen;
		$$.type = D_INDIR+D_NONE;
		$$.offset = $1;
	}
|	con &#39;(&#39; LLREG &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+$3;
		$$.offset = $1;
	}
|	con &#39;(&#39; LSP &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+D_SP;
		$$.offset = $1;
	}
|	con &#39;(&#39; LLREG &#39;*&#39; con &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+D_NONE;
		$$.offset = $1;
		$$.index = $3;
		$$.scale = $5;
		checkscale($$.scale);
	}
|	con &#39;(&#39; LLREG &#39;)&#39; &#39;(&#39; LLREG &#39;*&#39; con &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+$3;
		$$.offset = $1;
		$$.index = $6;
		$$.scale = $8;
		checkscale($$.scale);
	}
|	&#39;(&#39; LLREG &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+$2;
	}
|	&#39;(&#39; LSP &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+D_SP;
	}
|	con &#39;(&#39; LSREG &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+$3;
		$$.offset = $1;
	}
|	&#39;(&#39; LLREG &#39;*&#39; con &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+D_NONE;
		$$.index = $2;
		$$.scale = $4;
		checkscale($$.scale);
	}
|	&#39;(&#39; LLREG &#39;)&#39; &#39;(&#39; LLREG &#39;*&#39; con &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_INDIR+$2;
		$$.index = $5;
		$$.scale = $7;
		checkscale($$.scale);
	}

nmem:
	nam
	{
		$$ = $1;
	}
|	nam &#39;(&#39; LLREG &#39;*&#39; con &#39;)&#39;
	{
		$$ = $1;
		$$.index = $3;
		$$.scale = $5;
		checkscale($$.scale);
	}

nam:
	LNAME offset &#39;(&#39; pointer &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = $4;
		$$.sym = $1;
		$$.offset = $2;
	}
|	LNAME &#39;&lt;&#39; &#39;&gt;&#39; offset &#39;(&#39; LSB &#39;)&#39;
	{
		$$ = nullgen;
		$$.type = D_STATIC;
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
	{
		$$ = D_AUTO;
	}
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
