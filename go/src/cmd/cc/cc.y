<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/cc.y</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/cc.y</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/cc.y
// http://code.google.com/p/inferno-os/source/browse/utils/cc/cc.y
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
#include &#34;cc.h&#34;
%}
%union	{
	Node*	node;
	Sym*	sym;
	Type*	type;
	struct
	{
		Type*	t;
		uchar	c;
	} tycl;
	struct
	{
		Type*	t1;
		Type*	t2;
		Type*	t3;
		uchar	c;
	} tyty;
	struct
	{
		char*	s;
		int32	l;
	} sval;
	int32	lval;
	double	dval;
	vlong	vval;
}
%type	&lt;sym&gt;	ltag
%type	&lt;lval&gt;	gctname gcname cname gname tname
%type	&lt;lval&gt;	gctnlist gcnlist zgnlist
%type	&lt;type&gt;	tlist sbody complex
%type	&lt;tycl&gt;	types
%type	&lt;node&gt;	zarglist arglist zcexpr
%type	&lt;node&gt;	name block stmnt cexpr expr xuexpr pexpr
%type	&lt;node&gt;	zelist elist adecl slist uexpr string lstring
%type	&lt;node&gt;	xdecor xdecor2 labels label ulstmnt
%type	&lt;node&gt;	adlist edecor tag qual qlist
%type	&lt;node&gt;	abdecor abdecor1 abdecor2 abdecor3
%type	&lt;node&gt;	zexpr lexpr init ilist forexpr

%left	&#39;;&#39;
%left	&#39;,&#39;
%right	&#39;=&#39; LPE LME LMLE LDVE LMDE LRSHE LLSHE LANDE LXORE LORE
%right	&#39;?&#39; &#39;:&#39;
%left	LOROR
%left	LANDAND
%left	&#39;|&#39;
%left	&#39;^&#39;
%left	&#39;&amp;&#39;
%left	LEQ LNE
%left	&#39;&lt;&#39; &#39;&gt;&#39; LLE LGE
%left	LLSH LRSH
%left	&#39;+&#39; &#39;-&#39;
%left	&#39;*&#39; &#39;/&#39; &#39;%&#39;
%right	LMM LPP LMG &#39;.&#39; &#39;[&#39; &#39;(&#39;

%token	&lt;sym&gt;	LNAME LTYPE
%token	&lt;dval&gt;	LFCONST LDCONST
%token	&lt;vval&gt;	LCONST LLCONST LUCONST LULCONST LVLCONST LUVLCONST
%token	&lt;sval&gt;	LSTRING LLSTRING
%token		LAUTO LBREAK LCASE LCHAR LCONTINUE LDEFAULT LDO
%token		LDOUBLE LELSE LEXTERN LFLOAT LFOR LGOTO
%token	LIF LINT LLONG LREGISTER LRETURN LSHORT LSIZEOF LUSED
%token	LSTATIC LSTRUCT LSWITCH LTYPEDEF LTYPESTR LUNION LUNSIGNED
%token	LWHILE LVOID LENUM LSIGNED LCONSTNT LVOLATILE LSET LSIGNOF
%token	LRESTRICT LINLINE
%%
prog:
|	prog xdecl

/*
 * external declarator
 */
xdecl:
	zctlist &#39;;&#39;
	{
		dodecl(xdecl, lastclass, lasttype, Z);
	}
|	zctlist xdlist &#39;;&#39;
|	zctlist xdecor
	{
		lastdcl = T;
		firstarg = S;
		dodecl(xdecl, lastclass, lasttype, $2);
		if(lastdcl == T || lastdcl-&gt;etype != TFUNC) {
			diag($2, &#34;not a function&#34;);
			lastdcl = types[TFUNC];
		}
		thisfn = lastdcl;
		markdcl();
		firstdcl = dclstack;
		argmark($2, 0);
	}
	pdecl
	{
		argmark($2, 1);
	}
	block
	{
		Node *n;

		n = revertdcl();
		if(n)
			$6 = new(OLIST, n, $6);
		if(!debug[&#39;a&#39;] &amp;&amp; !debug[&#39;Z&#39;])
			codgen($6, $2);
	}

xdlist:
	xdecor
	{
		dodecl(xdecl, lastclass, lasttype, $1);
	}
|	xdecor
	{
		$1 = dodecl(xdecl, lastclass, lasttype, $1);
	}
	&#39;=&#39; init
	{
		doinit($1-&gt;sym, $1-&gt;type, 0L, $4);
	}
|	xdlist &#39;,&#39; xdlist

xdecor:
	xdecor2
|	&#39;*&#39; zgnlist xdecor
	{
		$$ = new(OIND, $3, Z);
		$$-&gt;garb = simpleg($2);
	}

xdecor2:
	tag
|	&#39;(&#39; xdecor &#39;)&#39;
	{
		$$ = $2;
	}
|	xdecor2 &#39;(&#39; zarglist &#39;)&#39;
	{
		$$ = new(OFUNC, $1, $3);
	}
|	xdecor2 &#39;[&#39; zexpr &#39;]&#39;
	{
		$$ = new(OARRAY, $1, $3);
	}

/*
 * automatic declarator
 */
adecl:
	ctlist &#39;;&#39;
	{
		$$ = dodecl(adecl, lastclass, lasttype, Z);
	}
|	ctlist adlist &#39;;&#39;
	{
		$$ = $2;
	}

adlist:
	xdecor
	{
		dodecl(adecl, lastclass, lasttype, $1);
		$$ = Z;
	}
|	xdecor
	{
		$1 = dodecl(adecl, lastclass, lasttype, $1);
	}
	&#39;=&#39; init
	{
		int32 w;

		w = $1-&gt;sym-&gt;type-&gt;width;
		$$ = doinit($1-&gt;sym, $1-&gt;type, 0L, $4);
		$$ = contig($1-&gt;sym, $$, w);
	}
|	adlist &#39;,&#39; adlist
	{
		$$ = $1;
		if($3 != Z) {
			$$ = $3;
			if($1 != Z)
				$$ = new(OLIST, $1, $3);
		}
	}

/*
 * parameter declarator
 */
pdecl:
|	pdecl ctlist pdlist &#39;;&#39;

pdlist:
	xdecor
	{
		dodecl(pdecl, lastclass, lasttype, $1);
	}
|	pdlist &#39;,&#39; pdlist

/*
 * structure element declarator
 */
edecl:
	tlist
	{
		lasttype = $1;
	}
	zedlist &#39;;&#39;
|	edecl tlist
	{
		lasttype = $2;
	}
	zedlist &#39;;&#39;

zedlist:					/* extension */
	{
		lastfield = 0;
		edecl(CXXX, lasttype, S);
	}
|	edlist

edlist:
	edecor
	{
		dodecl(edecl, CXXX, lasttype, $1);
	}
|	edlist &#39;,&#39; edlist

edecor:
	xdecor
	{
		lastbit = 0;
		firstbit = 1;
	}
|	tag &#39;:&#39; lexpr
	{
		$$ = new(OBIT, $1, $3);
	}
|	&#39;:&#39; lexpr
	{
		$$ = new(OBIT, Z, $2);
	}

/*
 * abstract declarator
 */
abdecor:
	{
		$$ = (Z);
	}
|	abdecor1

abdecor1:
	&#39;*&#39; zgnlist
	{
		$$ = new(OIND, (Z), Z);
		$$-&gt;garb = simpleg($2);
	}
|	&#39;*&#39; zgnlist abdecor1
	{
		$$ = new(OIND, $3, Z);
		$$-&gt;garb = simpleg($2);
	}
|	abdecor2

abdecor2:
	abdecor3
|	abdecor2 &#39;(&#39; zarglist &#39;)&#39;
	{
		$$ = new(OFUNC, $1, $3);
	}
|	abdecor2 &#39;[&#39; zexpr &#39;]&#39;
	{
		$$ = new(OARRAY, $1, $3);
	}

abdecor3:
	&#39;(&#39; &#39;)&#39;
	{
		$$ = new(OFUNC, (Z), Z);
	}
|	&#39;[&#39; zexpr &#39;]&#39;
	{
		$$ = new(OARRAY, (Z), $2);
	}
|	&#39;(&#39; abdecor1 &#39;)&#39;
	{
		$$ = $2;
	}

init:
	expr
|	&#39;{&#39; ilist &#39;}&#39;
	{
		$$ = new(OINIT, invert($2), Z);
	}

qual:
	&#39;[&#39; lexpr &#39;]&#39;
	{
		$$ = new(OARRAY, $2, Z);
	}
|	&#39;.&#39; ltag
	{
		$$ = new(OELEM, Z, Z);
		$$-&gt;sym = $2;
	}
|	qual &#39;=&#39;

qlist:
	init &#39;,&#39;
|	qlist init &#39;,&#39;
	{
		$$ = new(OLIST, $1, $2);
	}
|	qual
|	qlist qual
	{
		$$ = new(OLIST, $1, $2);
	}

ilist:
	qlist
|	init
|	qlist init
	{
		$$ = new(OLIST, $1, $2);
	}

zarglist:
	{
		$$ = Z;
	}
|	arglist
	{
		$$ = invert($1);
	}


arglist:
	name
|	tlist abdecor
	{
		$$ = new(OPROTO, $2, Z);
		$$-&gt;type = $1;
	}
|	tlist xdecor
	{
		$$ = new(OPROTO, $2, Z);
		$$-&gt;type = $1;
	}
|	&#39;.&#39; &#39;.&#39; &#39;.&#39;
	{
		$$ = new(ODOTDOT, Z, Z);
	}
|	arglist &#39;,&#39; arglist
	{
		$$ = new(OLIST, $1, $3);
	}

block:
	&#39;{&#39; slist &#39;}&#39;
	{
		$$ = invert($2);
	//	if($2 != Z)
	//		$$ = new(OLIST, $2, $$);
		if($$ == Z)
			$$ = new(OLIST, Z, Z);
	}

slist:
	{
		$$ = Z;
	}
|	slist adecl
	{
		$$ = new(OLIST, $1, $2);
	}
|	slist stmnt
	{
		$$ = new(OLIST, $1, $2);
	}

labels:
	label
|	labels label
	{
		$$ = new(OLIST, $1, $2);
	}

label:
	LCASE expr &#39;:&#39;
	{
		$$ = new(OCASE, $2, Z);
	}
|	LDEFAULT &#39;:&#39;
	{
		$$ = new(OCASE, Z, Z);
	}
|	LNAME &#39;:&#39;
	{
		$$ = new(OLABEL, dcllabel($1, 1), Z);
	}

stmnt:
	error &#39;;&#39;
	{
		$$ = Z;
	}
|	ulstmnt
|	labels ulstmnt
	{
		$$ = new(OLIST, $1, $2);
	}

forexpr:
	zcexpr
|	ctlist adlist
	{
		$$ = $2;
	}

ulstmnt:
	zcexpr &#39;;&#39;
|	{
		markdcl();
	}
	block
	{
		$$ = revertdcl();
		if($$)
			$$ = new(OLIST, $$, $2);
		else
			$$ = $2;
	}
|	LIF &#39;(&#39; cexpr &#39;)&#39; stmnt
	{
		$$ = new(OIF, $3, new(OLIST, $5, Z));
		if($5 == Z)
			warn($3, &#34;empty if body&#34;);
	}
|	LIF &#39;(&#39; cexpr &#39;)&#39; stmnt LELSE stmnt
	{
		$$ = new(OIF, $3, new(OLIST, $5, $7));
		if($5 == Z)
			warn($3, &#34;empty if body&#34;);
		if($7 == Z)
			warn($3, &#34;empty else body&#34;);
	}
|	{ markdcl(); } LFOR &#39;(&#39; forexpr &#39;;&#39; zcexpr &#39;;&#39; zcexpr &#39;)&#39; stmnt
	{
		$$ = revertdcl();
		if($$){
			if($4)
				$4 = new(OLIST, $$, $4);
			else
				$4 = $$;
		}
		$$ = new(OFOR, new(OLIST, $6, new(OLIST, $4, $8)), $10);
	}
|	LWHILE &#39;(&#39; cexpr &#39;)&#39; stmnt
	{
		$$ = new(OWHILE, $3, $5);
	}
|	LDO stmnt LWHILE &#39;(&#39; cexpr &#39;)&#39; &#39;;&#39;
	{
		$$ = new(ODWHILE, $5, $2);
	}
|	LRETURN zcexpr &#39;;&#39;
	{
		$$ = new(ORETURN, $2, Z);
		$$-&gt;type = thisfn-&gt;link;
	}
|	LSWITCH &#39;(&#39; cexpr &#39;)&#39; stmnt
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;vconst = 0;
		$$-&gt;type = types[TINT];
		$3 = new(OSUB, $$, $3);

		$$ = new(OCONST, Z, Z);
		$$-&gt;vconst = 0;
		$$-&gt;type = types[TINT];
		$3 = new(OSUB, $$, $3);

		$$ = new(OSWITCH, $3, $5);
	}
|	LBREAK &#39;;&#39;
	{
		$$ = new(OBREAK, Z, Z);
	}
|	LCONTINUE &#39;;&#39;
	{
		$$ = new(OCONTINUE, Z, Z);
	}
|	LGOTO ltag &#39;;&#39;
	{
		$$ = new(OGOTO, dcllabel($2, 0), Z);
	}
|	LUSED &#39;(&#39; zelist &#39;)&#39; &#39;;&#39;
	{
		$$ = new(OUSED, $3, Z);
	}
|	LSET &#39;(&#39; zelist &#39;)&#39; &#39;;&#39;
	{
		$$ = new(OSET, $3, Z);
	}

zcexpr:
	{
		$$ = Z;
	}
|	cexpr

zexpr:
	{
		$$ = Z;
	}
|	lexpr

lexpr:
	expr
	{
		$$ = new(OCAST, $1, Z);
		$$-&gt;type = types[TLONG];
	}

cexpr:
	expr
|	cexpr &#39;,&#39; cexpr
	{
		$$ = new(OCOMMA, $1, $3);
	}

expr:
	xuexpr
|	expr &#39;*&#39; expr
	{
		$$ = new(OMUL, $1, $3);
	}
|	expr &#39;/&#39; expr
	{
		$$ = new(ODIV, $1, $3);
	}
|	expr &#39;%&#39; expr
	{
		$$ = new(OMOD, $1, $3);
	}
|	expr &#39;+&#39; expr
	{
		$$ = new(OADD, $1, $3);
	}
|	expr &#39;-&#39; expr
	{
		$$ = new(OSUB, $1, $3);
	}
|	expr LRSH expr
	{
		$$ = new(OASHR, $1, $3);
	}
|	expr LLSH expr
	{
		$$ = new(OASHL, $1, $3);
	}
|	expr &#39;&lt;&#39; expr
	{
		$$ = new(OLT, $1, $3);
	}
|	expr &#39;&gt;&#39; expr
	{
		$$ = new(OGT, $1, $3);
	}
|	expr LLE expr
	{
		$$ = new(OLE, $1, $3);
	}
|	expr LGE expr
	{
		$$ = new(OGE, $1, $3);
	}
|	expr LEQ expr
	{
		$$ = new(OEQ, $1, $3);
	}
|	expr LNE expr
	{
		$$ = new(ONE, $1, $3);
	}
|	expr &#39;&amp;&#39; expr
	{
		$$ = new(OAND, $1, $3);
	}
|	expr &#39;^&#39; expr
	{
		$$ = new(OXOR, $1, $3);
	}
|	expr &#39;|&#39; expr
	{
		$$ = new(OOR, $1, $3);
	}
|	expr LANDAND expr
	{
		$$ = new(OANDAND, $1, $3);
	}
|	expr LOROR expr
	{
		$$ = new(OOROR, $1, $3);
	}
|	expr &#39;?&#39; cexpr &#39;:&#39; expr
	{
		$$ = new(OCOND, $1, new(OLIST, $3, $5));
	}
|	expr &#39;=&#39; expr
	{
		$$ = new(OAS, $1, $3);
	}
|	expr LPE expr
	{
		$$ = new(OASADD, $1, $3);
	}
|	expr LME expr
	{
		$$ = new(OASSUB, $1, $3);
	}
|	expr LMLE expr
	{
		$$ = new(OASMUL, $1, $3);
	}
|	expr LDVE expr
	{
		$$ = new(OASDIV, $1, $3);
	}
|	expr LMDE expr
	{
		$$ = new(OASMOD, $1, $3);
	}
|	expr LLSHE expr
	{
		$$ = new(OASASHL, $1, $3);
	}
|	expr LRSHE expr
	{
		$$ = new(OASASHR, $1, $3);
	}
|	expr LANDE expr
	{
		$$ = new(OASAND, $1, $3);
	}
|	expr LXORE expr
	{
		$$ = new(OASXOR, $1, $3);
	}
|	expr LORE expr
	{
		$$ = new(OASOR, $1, $3);
	}

xuexpr:
	uexpr
|	&#39;(&#39; tlist abdecor &#39;)&#39; xuexpr
	{
		$$ = new(OCAST, $5, Z);
		dodecl(NODECL, CXXX, $2, $3);
		$$-&gt;type = lastdcl;
		$$-&gt;xcast = 1;
	}
|	&#39;(&#39; tlist abdecor &#39;)&#39; &#39;{&#39; ilist &#39;}&#39;	/* extension */
	{
		$$ = new(OSTRUCT, $6, Z);
		dodecl(NODECL, CXXX, $2, $3);
		$$-&gt;type = lastdcl;
	}

uexpr:
	pexpr
|	&#39;*&#39; xuexpr
	{
		$$ = new(OIND, $2, Z);
	}
|	&#39;&amp;&#39; xuexpr
	{
		$$ = new(OADDR, $2, Z);
	}
|	&#39;+&#39; xuexpr
	{
		$$ = new(OPOS, $2, Z);
	}
|	&#39;-&#39; xuexpr
	{
		$$ = new(ONEG, $2, Z);
	}
|	&#39;!&#39; xuexpr
	{
		$$ = new(ONOT, $2, Z);
	}
|	&#39;~&#39; xuexpr
	{
		$$ = new(OCOM, $2, Z);
	}
|	LPP xuexpr
	{
		$$ = new(OPREINC, $2, Z);
	}
|	LMM xuexpr
	{
		$$ = new(OPREDEC, $2, Z);
	}
|	LSIZEOF uexpr
	{
		$$ = new(OSIZE, $2, Z);
	}
|	LSIGNOF uexpr
	{
		$$ = new(OSIGN, $2, Z);
	}

pexpr:
	&#39;(&#39; cexpr &#39;)&#39;
	{
		$$ = $2;
	}
|	LSIZEOF &#39;(&#39; tlist abdecor &#39;)&#39;
	{
		$$ = new(OSIZE, Z, Z);
		dodecl(NODECL, CXXX, $3, $4);
		$$-&gt;type = lastdcl;
	}
|	LSIGNOF &#39;(&#39; tlist abdecor &#39;)&#39;
	{
		$$ = new(OSIGN, Z, Z);
		dodecl(NODECL, CXXX, $3, $4);
		$$-&gt;type = lastdcl;
	}
|	pexpr &#39;(&#39; zelist &#39;)&#39;
	{
		$$ = new(OFUNC, $1, Z);
		if($1-&gt;op == ONAME)
		if($1-&gt;type == T)
			dodecl(xdecl, CXXX, types[TINT], $$);
		$$-&gt;right = invert($3);
	}
|	pexpr &#39;[&#39; cexpr &#39;]&#39;
	{
		$$ = new(OIND, new(OADD, $1, $3), Z);
	}
|	pexpr LMG ltag
	{
		$$ = new(ODOT, new(OIND, $1, Z), Z);
		$$-&gt;sym = $3;
	}
|	pexpr &#39;.&#39; ltag
	{
		$$ = new(ODOT, $1, Z);
		$$-&gt;sym = $3;
	}
|	pexpr LPP
	{
		$$ = new(OPOSTINC, $1, Z);
	}
|	pexpr LMM
	{
		$$ = new(OPOSTDEC, $1, Z);
	}
|	name
|	LCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TINT];
		$$-&gt;vconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LLCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TLONG];
		$$-&gt;vconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LUCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TUINT];
		$$-&gt;vconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LULCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TULONG];
		$$-&gt;vconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LDCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TDOUBLE];
		$$-&gt;fconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LFCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TFLOAT];
		$$-&gt;fconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LVLCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TVLONG];
		$$-&gt;vconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	LUVLCONST
	{
		$$ = new(OCONST, Z, Z);
		$$-&gt;type = types[TUVLONG];
		$$-&gt;vconst = $1;
		$$-&gt;cstring = strdup(symb);
	}
|	string
|	lstring

string:
	LSTRING
	{
		$$ = new(OSTRING, Z, Z);
		$$-&gt;type = typ(TARRAY, types[TCHAR]);
		$$-&gt;type-&gt;width = $1.l + 1;
		$$-&gt;cstring = $1.s;
		$$-&gt;sym = symstring;
		$$-&gt;etype = TARRAY;
		$$-&gt;class = CSTATIC;
	}
|	string LSTRING
	{
		char *s;
		int n;

		n = $1-&gt;type-&gt;width - 1;
		s = alloc(n+$2.l+MAXALIGN);

		memcpy(s, $1-&gt;cstring, n);
		memcpy(s+n, $2.s, $2.l);
		s[n+$2.l] = 0;

		$$ = $1;
		$$-&gt;type-&gt;width += $2.l;
		$$-&gt;cstring = s;
	}

lstring:
	LLSTRING
	{
		$$ = new(OLSTRING, Z, Z);
		$$-&gt;type = typ(TARRAY, types[TUSHORT]);
		$$-&gt;type-&gt;width = $1.l + sizeof(ushort);
		$$-&gt;rstring = (ushort*)$1.s;
		$$-&gt;sym = symstring;
		$$-&gt;etype = TARRAY;
		$$-&gt;class = CSTATIC;
	}
|	lstring LLSTRING
	{
		char *s;
		int n;

		n = $1-&gt;type-&gt;width - sizeof(ushort);
		s = alloc(n+$2.l+MAXALIGN);

		memcpy(s, $1-&gt;rstring, n);
		memcpy(s+n, $2.s, $2.l);
		*(ushort*)(s+n+$2.l) = 0;

		$$ = $1;
		$$-&gt;type-&gt;width += $2.l;
		$$-&gt;rstring = (ushort*)s;
	}

zelist:
	{
		$$ = Z;
	}
|	elist

elist:
	expr
|	elist &#39;,&#39; elist
	{
		$$ = new(OLIST, $1, $3);
	}

sbody:
	&#39;{&#39;
	{
		$&lt;tyty&gt;$.t1 = strf;
		$&lt;tyty&gt;$.t2 = strl;
		$&lt;tyty&gt;$.t3 = lasttype;
		$&lt;tyty&gt;$.c = lastclass;
		strf = T;
		strl = T;
		lastbit = 0;
		firstbit = 1;
		lastclass = CXXX;
		lasttype = T;
	}
	edecl &#39;}&#39;
	{
		$$ = strf;
		strf = $&lt;tyty&gt;2.t1;
		strl = $&lt;tyty&gt;2.t2;
		lasttype = $&lt;tyty&gt;2.t3;
		lastclass = $&lt;tyty&gt;2.c;
	}

zctlist:
	{
		lastclass = CXXX;
		lasttype = types[TINT];
	}
|	ctlist

types:
	complex
	{
		$$.t = $1;
		$$.c = CXXX;
	}
|	tname
	{
		$$.t = simplet($1);
		$$.c = CXXX;
	}
|	gcnlist
	{
		$$.t = simplet($1);
		$$.c = simplec($1);
		$$.t = garbt($$.t, $1);
	}
|	complex gctnlist
	{
		$$.t = $1;
		$$.c = simplec($2);
		$$.t = garbt($$.t, $2);
		if($2 &amp; ~BCLASS &amp; ~BGARB)
			diag(Z, &#34;duplicate types given: %T and %Q&#34;, $1, $2);
	}
|	tname gctnlist
	{
		$$.t = simplet(typebitor($1, $2));
		$$.c = simplec($2);
		$$.t = garbt($$.t, $2);
	}
|	gcnlist complex zgnlist
	{
		$$.t = $2;
		$$.c = simplec($1);
		$$.t = garbt($$.t, $1|$3);
	}
|	gcnlist tname
	{
		$$.t = simplet($2);
		$$.c = simplec($1);
		$$.t = garbt($$.t, $1);
	}
|	gcnlist tname gctnlist
	{
		$$.t = simplet(typebitor($2, $3));
		$$.c = simplec($1|$3);
		$$.t = garbt($$.t, $1|$3);
	}

tlist:
	types
	{
		$$ = $1.t;
		if($1.c != CXXX)
			diag(Z, &#34;illegal combination of class 4: %s&#34;, cnames[$1.c]);
	}

ctlist:
	types
	{
		lasttype = $1.t;
		lastclass = $1.c;
	}

complex:
	LSTRUCT ltag
	{
		dotag($2, TSTRUCT, 0);
		$$ = $2-&gt;suetag;
	}
|	LSTRUCT ltag
	{
		dotag($2, TSTRUCT, autobn);
	}
	sbody
	{
		$$ = $2-&gt;suetag;
		if($$-&gt;link != T)
			diag(Z, &#34;redeclare tag: %s&#34;, $2-&gt;name);
		$$-&gt;link = $4;
		suallign($$);
	}
|	LSTRUCT sbody
	{
		taggen++;
		sprint(symb, &#34;_%d_&#34;, taggen);
		$$ = dotag(lookup(), TSTRUCT, autobn);
		$$-&gt;link = $2;
		suallign($$);
	}
|	LUNION ltag
	{
		dotag($2, TUNION, 0);
		$$ = $2-&gt;suetag;
	}
|	LUNION ltag
	{
		dotag($2, TUNION, autobn);
	}
	sbody
	{
		$$ = $2-&gt;suetag;
		if($$-&gt;link != T)
			diag(Z, &#34;redeclare tag: %s&#34;, $2-&gt;name);
		$$-&gt;link = $4;
		suallign($$);
	}
|	LUNION sbody
	{
		taggen++;
		sprint(symb, &#34;_%d_&#34;, taggen);
		$$ = dotag(lookup(), TUNION, autobn);
		$$-&gt;link = $2;
		suallign($$);
	}
|	LENUM ltag
	{
		dotag($2, TENUM, 0);
		$$ = $2-&gt;suetag;
		if($$-&gt;link == T)
			$$-&gt;link = types[TINT];
		$$ = $$-&gt;link;
	}
|	LENUM ltag
	{
		dotag($2, TENUM, autobn);
	}
	&#39;{&#39;
	{
		en.tenum = T;
		en.cenum = T;
	}
	enum &#39;}&#39;
	{
		$$ = $2-&gt;suetag;
		if($$-&gt;link != T)
			diag(Z, &#34;redeclare tag: %s&#34;, $2-&gt;name);
		if(en.tenum == T) {
			diag(Z, &#34;enum type ambiguous: %s&#34;, $2-&gt;name);
			en.tenum = types[TINT];
		}
		$$-&gt;link = en.tenum;
		$$ = en.tenum;
	}
|	LENUM &#39;{&#39;
	{
		en.tenum = T;
		en.cenum = T;
	}
	enum &#39;}&#39;
	{
		$$ = en.tenum;
	}
|	LTYPE
	{
		$$ = tcopy($1-&gt;type);
	}

gctnlist:
	gctname
|	gctnlist gctname
	{
		$$ = typebitor($1, $2);
	}

zgnlist:
	{
		$$ = 0;
	}
|	zgnlist gname
	{
		$$ = typebitor($1, $2);
	}

gctname:
	tname
|	gname
|	cname

gcnlist:
	gcname
|	gcnlist gcname
	{
		$$ = typebitor($1, $2);
	}

gcname:
	gname
|	cname

enum:
	LNAME
	{
		doenum($1, Z);
	}
|	LNAME &#39;=&#39; expr
	{
		doenum($1, $3);
	}
|	enum &#39;,&#39;
|	enum &#39;,&#39; enum

tname:	/* type words */
	LCHAR { $$ = BCHAR; }
|	LSHORT { $$ = BSHORT; }
|	LINT { $$ = BINT; }
|	LLONG { $$ = BLONG; }
|	LSIGNED { $$ = BSIGNED; }
|	LUNSIGNED { $$ = BUNSIGNED; }
|	LFLOAT { $$ = BFLOAT; }
|	LDOUBLE { $$ = BDOUBLE; }
|	LVOID { $$ = BVOID; }

cname:	/* class words */
	LAUTO { $$ = BAUTO; }
|	LSTATIC { $$ = BSTATIC; }
|	LEXTERN { $$ = BEXTERN; }
|	LTYPEDEF { $$ = BTYPEDEF; }
|	LTYPESTR { $$ = BTYPESTR; }
|	LREGISTER { $$ = BREGISTER; }
|	LINLINE { $$ = 0; }

gname:	/* garbage words */
	LCONSTNT { $$ = BCONSTNT; }
|	LVOLATILE { $$ = BVOLATILE; }
|	LRESTRICT { $$ = 0; }

name:
	LNAME
	{
		$$ = new(ONAME, Z, Z);
		if($1-&gt;class == CLOCAL)
			$1 = mkstatic($1);
		$$-&gt;sym = $1;
		$$-&gt;type = $1-&gt;type;
		$$-&gt;etype = TVOID;
		if($$-&gt;type != T)
			$$-&gt;etype = $$-&gt;type-&gt;etype;
		$$-&gt;xoffset = $1-&gt;offset;
		$$-&gt;class = $1-&gt;class;
		$1-&gt;aused = 1;
	}
tag:
	ltag
	{
		$$ = new(ONAME, Z, Z);
		$$-&gt;sym = $1;
		$$-&gt;type = $1-&gt;type;
		$$-&gt;etype = TVOID;
		if($$-&gt;type != T)
			$$-&gt;etype = $$-&gt;type-&gt;etype;
		$$-&gt;xoffset = $1-&gt;offset;
		$$-&gt;class = $1-&gt;class;
	}
ltag:
	LNAME
|	LTYPE
%%
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
