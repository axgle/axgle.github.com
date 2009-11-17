<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/go.y</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/go.y</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * Go language grammar.
 *
 * The Go semicolon rules are:
 *
 *  1. all statements and declarations are terminated by semicolons
 *  2. semicolons can be omitted at top level.
 *  3. semicolons can be omitted before and after the closing ) or }
 *	on a list of statements or declarations.
 *
 * This is accomplished by calling yyoptsemi() to mark the places
 * where semicolons are optional.  That tells the lexer that if a
 * semicolon isn&#39;t the next token, it should insert one for us.
 */

%{
#include &#34;go.h&#34;
%}
%union	{
	Node*		node;
	NodeList*		list;
	Type*		type;
	Sym*		sym;
	struct	Val	val;
	int		lint;
}

// |sed &#39;s/.*	//&#39; |9 fmt -l1 |sort |9 fmt -l50 | sed &#39;s/^/%xxx		/&#39;

%token	&lt;val&gt;	LLITERAL
%token	&lt;lint&gt;	LASOP
%token	&lt;sym&gt;	LBREAK LCASE LCHAN LCOLAS LCONST LCONTINUE LDDD
%token	&lt;sym&gt;	LDEFAULT LDEFER LELSE LFALL LFOR LFUNC LGO LGOTO
%token	&lt;sym&gt;	LIF LIMPORT LINTERFACE LMAKE LMAP LNAME LNEW
%token	&lt;sym&gt;	LPACKAGE LRANGE LRETURN LSELECT LSTRUCT LSWITCH
%token	&lt;sym&gt;	LTYPE LVAR

%token		LANDAND LANDNOT LBODY LCOMM LDEC LEQ LGE LGT
%token		LIGNORE LINC LLE LLSH LLT LNE LOROR LRSH
%token		LSEMIBRACE

%type	&lt;lint&gt;	lbrace import_here
%type	&lt;sym&gt;	sym packname
%type	&lt;val&gt;	oliteral

%type	&lt;node&gt;	stmt ntype
%type	&lt;node&gt;	arg_type
%type	&lt;node&gt;	case caseblock
%type	&lt;node&gt;	compound_stmt dotname embed expr
%type	&lt;node&gt;	expr_or_type
%type	&lt;node&gt;	fndcl fnliteral
%type	&lt;node&gt;	for_body for_header for_stmt if_header if_stmt
%type	&lt;node&gt;	interfacedcl keyval labelname name
%type	&lt;node&gt;	name_or_type non_expr_type
%type	&lt;node&gt;	new_name dcl_name oexpr typedclname
%type	&lt;node&gt;	onew_name
%type	&lt;node&gt;	osimple_stmt pexpr
%type	&lt;node&gt;	pseudocall range_stmt select_stmt
%type	&lt;node&gt;	simple_stmt
%type	&lt;node&gt;	switch_stmt uexpr
%type	&lt;node&gt;	xfndcl typedcl

%type	&lt;list&gt;	xdcl fnbody fnres switch_body loop_body dcl_name_list
%type	&lt;list&gt;	new_name_list expr_list keyval_list braced_keyval_list expr_or_type_list xdcl_list
%type	&lt;list&gt;	oexpr_list oexpr_or_type_list caseblock_list stmt_list oarg_type_list arg_type_list
%type	&lt;list&gt;	interfacedcl_list vardcl vardcl_list structdcl structdcl_list
%type	&lt;list&gt;	common_dcl constdcl constdcl1 constdcl_list typedcl_list

%type	&lt;node&gt;	convtype dotdotdot
%type	&lt;node&gt;	indcl interfacetype structtype ptrtype
%type	&lt;node&gt;	chantype non_chan_type othertype non_fn_type fntype

%type	&lt;sym&gt;	hidden_importsym hidden_pkg_importsym

%type	&lt;node&gt;	hidden_constant hidden_dcl hidden_interfacedcl hidden_structdcl

%type	&lt;list&gt;	hidden_funres
%type	&lt;list&gt;	ohidden_funres
%type	&lt;list&gt;	hidden_funarg_list ohidden_funarg_list
%type	&lt;list&gt;	hidden_interfacedcl_list ohidden_interfacedcl_list
%type	&lt;list&gt;	hidden_structdcl_list ohidden_structdcl_list

%type	&lt;type&gt;	hidden_type hidden_type1 hidden_type2 hidden_pkgtype

%left		LOROR
%left		LANDAND
%left		LCOMM
%left		LEQ LNE LLE LGE LLT LGT
%left		&#39;+&#39; &#39;-&#39; &#39;|&#39; &#39;^&#39;
%left		&#39;*&#39; &#39;/&#39; &#39;%&#39; &#39;&amp;&#39; LLSH LRSH LANDNOT

/*
 * manual override of shift/reduce conflicts.
 * the general form is that we assign a precedence
 * to the token being shifted and then introduce
 * NotToken with lower precedence or PreferToToken with higher
 * and annotate the reducing rule accordingly.
 */
%left		NotPackage
%left		LPACKAGE

%left		NotParen
%left		&#39;(&#39;

%left		&#39;)&#39;
%left		PreferToRightParen

%left		&#39;.&#39;

%left		&#39;{&#39;

%left		NotSemi
%left		&#39;;&#39;

%%
file:
	loadsys
	package
	imports
	xdcl_list
	{
		xtop = concat(xtop, $4);
	}

package:
	%prec NotPackage
	{
		prevlineno = lineno;
		yyerror(&#34;package statement must be first&#34;);
		mkpackage(&#34;main&#34;);
	}
|	LPACKAGE sym
	{
		mkpackage($2-&gt;name);
	}

/*
 * this loads the definitions for the low-level runtime functions,
 * so that the compiler can generate calls to them,
 * but does not make the name &#34;runtime&#34; visible as a package.
 */
loadsys:
	{
		cannedimports(&#34;runtime.builtin&#34;, runtimeimport);
	}
	import_package
	import_there
	{
		pkgimportname = S;
	}

imports:
|	imports import

import:
	LIMPORT import_stmt osemi
|	LIMPORT &#39;(&#39; import_stmt_list osemi &#39;)&#39; osemi
|	LIMPORT &#39;(&#39; &#39;)&#39; osemi

import_stmt:
	import_here import_package import_there
	{
		Sym *import, *my;
		Node *pack;

		import = pkgimportname;
		my = pkgmyname;
		pkgmyname = S;
		pkgimportname = S;

		if(import == S)
			break;

		pack = nod(OPACK, N, N);
		pack-&gt;sym = import;
		pack-&gt;lineno = $1;

		if(my == S)
			my = import;
		if(my-&gt;name[0] == &#39;.&#39;) {
			importdot(import, pack);
			break;
		}
		if(my-&gt;name[0] == &#39;_&#39; &amp;&amp; my-&gt;name[1] == &#39;\0&#39;)
			break;

		// Can end up with my-&gt;def-&gt;op set to ONONAME
		// if one package refers to p without importing it.
		// Don&#39;t want to give an error on a good import
		// in another file.
		if(my-&gt;def &amp;&amp; my-&gt;def-&gt;op != ONONAME) {
			lineno = $1;
			redeclare(my, &#34;as imported package name&#34;);
		}
		my-&gt;def = pack;
		my-&gt;lastlineno = $1;
		import-&gt;block = 1;	// at top level
	}


import_stmt_list:
	import_stmt
|	import_stmt_list &#39;;&#39; import_stmt

import_here:
	LLITERAL
	{
		// import with original name
		$$ = parserline();
		pkgimportname = S;
		pkgmyname = S;
		importfile(&amp;$1, $$);
	}
|	sym LLITERAL
	{
		// import with given name
		$$ = parserline();
		pkgimportname = S;
		pkgmyname = $1;
		importfile(&amp;$2, $$);
	}
|	&#39;.&#39; LLITERAL
	{
		// import into my name space
		$$ = parserline();
		pkgmyname = lookup(&#34;.&#34;);
		importfile(&amp;$2, $$);
	}

import_package:
	LPACKAGE sym
	{
		pkgimportname = $2;
		if(strcmp($2-&gt;name, &#34;main&#34;) == 0)
			yyerror(&#34;cannot import package main&#34;);

		// TODO(rsc): This should go away once we get
		// rid of the global package name space.
		if(strcmp($2-&gt;name, package) == 0 &amp;&amp; strcmp(package, &#34;runtime&#34;) != 0)
			yyerror(&#34;package cannot import itself&#34;);
	}

import_there:
	{
		defercheckwidth();
	}
	hidden_import_list &#39;$&#39; &#39;$&#39;
	{
		resumecheckwidth();
		checkimports();
		unimportfile();
	}
|	LIMPORT &#39;$&#39; &#39;$&#39;
	{
		defercheckwidth();
	}
	hidden_import_list &#39;$&#39; &#39;$&#39;
	{
		resumecheckwidth();
		checkimports();
	}

/*
 * declarations
 */
xdcl:
	common_dcl osemi
|	xfndcl osemi
	{
		$$ = list1($1);
	}
|	error osemi
	{
		$$ = nil;
	}
|	&#39;;&#39;
	{
		yyerror(&#34;empty top-level declaration&#34;);
		$$ = nil;
	}

common_dcl:
	LVAR vardcl
	{
		$$ = $2;
		if(yylast == LSEMIBRACE)
			yyoptsemi(0);
	}
|	LVAR &#39;(&#39; vardcl_list osemi &#39;)&#39;
	{
		$$ = $3;
		yyoptsemi(0);
	}
|	LVAR &#39;(&#39; &#39;)&#39;
	{
		$$ = nil;
		yyoptsemi(0);
	}
|	LCONST constdcl
	{
		$$ = $2;
		iota = 0;
		lastconst = nil;
	}
|	LCONST &#39;(&#39; constdcl osemi &#39;)&#39;
	{
		$$ = $3;
		iota = 0;
		lastconst = nil;
		yyoptsemi(0);
	}
|	LCONST &#39;(&#39; constdcl &#39;;&#39; constdcl_list osemi &#39;)&#39;
	{
		$$ = concat($3, $5);
		iota = 0;
		lastconst = nil;
		yyoptsemi(0);
	}
|	LCONST &#39;(&#39; &#39;)&#39;
	{
		$$ = nil;
		yyoptsemi(0);
	}
|	LTYPE typedcl
	{
		$$ = list1($2);
		if(yylast == LSEMIBRACE)
			yyoptsemi(0);
	}
|	LTYPE &#39;(&#39; typedcl_list osemi &#39;)&#39;
	{
		$$ = $3;
		yyoptsemi(0);
	}
|	LTYPE &#39;(&#39; &#39;)&#39;
	{
		$$ = nil;
		yyoptsemi(0);
	}

varoptsemi:
	{
		if(yylast == LSEMIBRACE)
			yyoptsemi(&#39;=&#39;);
	}

vardcl:
	dcl_name_list ntype varoptsemi
	{
		$$ = variter($1, $2, nil);
	}
|	dcl_name_list ntype varoptsemi &#39;=&#39; expr_list
	{
		$$ = variter($1, $2, $5);
	}
|	dcl_name_list &#39;=&#39; expr_list
	{
		$$ = variter($1, nil, $3);
	}

constdcl:
	dcl_name_list ntype &#39;=&#39; expr_list
	{
		$$ = constiter($1, $2, $4);
	}
|	dcl_name_list &#39;=&#39; expr_list
	{
		$$ = constiter($1, N, $3);
	}

constdcl1:
	constdcl
|	dcl_name_list ntype
	{
		$$ = constiter($1, $2, nil);
	}
|	dcl_name_list
	{
		$$ = constiter($1, N, nil);
	}

typedclname:
	sym
	{
		// different from dclname because the name
		// becomes visible right here, not at the end
		// of the declaration.
		$$ = typedcl0($1);
	}

typedcl:
	typedclname ntype
	{
		$$ = typedcl1($1, $2, 1);
	}

simple_stmt:
	expr
	{
		$$ = $1;
	}
|	expr LASOP expr
	{
		$$ = nod(OASOP, $1, $3);
		$$-&gt;etype = $2;			// rathole to pass opcode
	}
|	expr_list &#39;=&#39; expr_list
	{
		if($1-&gt;next == nil &amp;&amp; $3-&gt;next == nil) {
			// simple
			$$ = nod(OAS, $1-&gt;n, $3-&gt;n);
			break;
		}
		// multiple
		$$ = nod(OAS2, N, N);
		$$-&gt;list = $1;
		$$-&gt;rlist = $3;
	}
|	expr_list LCOLAS expr_list
	{
		if($3-&gt;n-&gt;op == OTYPESW) {
			if($3-&gt;next != nil)
				yyerror(&#34;expr.(type) must be alone in list&#34;);
			else if($1-&gt;next != nil)
				yyerror(&#34;argument count mismatch: %d = %d&#34;, count($1), 1);
			$$ = nod(OTYPESW, $1-&gt;n, $3-&gt;n-&gt;right);
			break;
		}
		$$ = colas($1, $3);
	}
|	expr LINC
	{
		$$ = nod(OASOP, $1, nodintconst(1));
		$$-&gt;etype = OADD;
	}
|	expr LDEC
	{
		$$ = nod(OASOP, $1, nodintconst(1));
		$$-&gt;etype = OSUB;
	}

case:
	LCASE expr_or_type_list &#39;:&#39;
	{
		Node *n;

		// will be converted to OCASE
		// right will point to next case
		// done in casebody()
		poptodcl();
		$$ = nod(OXCASE, N, N);
		$$-&gt;list = $2;
		if(typesw != N &amp;&amp; typesw-&gt;right != N &amp;&amp; (n=typesw-&gt;right-&gt;left) != N) {
			// type switch - declare variable
			n = newname(n-&gt;sym);
			n-&gt;used = 1;	// TODO(rsc): better job here
			declare(n, dclcontext);
			$$-&gt;nname = n;
		}
		break;
	}
|	LCASE name &#39;=&#39; expr &#39;:&#39;
	{
		// will be converted to OCASE
		// right will point to next case
		// done in casebody()
		poptodcl();
		$$ = nod(OXCASE, N, N);
		$$-&gt;list = list1(nod(OAS, $2, $4));
	}
|	LCASE name LCOLAS expr &#39;:&#39;
	{
		// will be converted to OCASE
		// right will point to next case
		// done in casebody()
		poptodcl();
		$$ = nod(OXCASE, N, N);
		$$-&gt;list = list1(colas(list1($2), list1($4)));
	}
|	LDEFAULT &#39;:&#39;
	{
		Node *n;

		poptodcl();
		$$ = nod(OXCASE, N, N);
		if(typesw != N &amp;&amp; typesw-&gt;right != N &amp;&amp; (n=typesw-&gt;right-&gt;left) != N) {
			// type switch - declare variable
			n = newname(n-&gt;sym);
			n-&gt;used = 1;	// TODO(rsc): better job here
			declare(n, dclcontext);
			$$-&gt;nname = n;
		}
	}

compound_stmt:
	&#39;{&#39;
	{
		markdcl();
	}
	stmt_list &#39;}&#39;
	{
		$$ = liststmt($3);
		popdcl();
		yyoptsemi(0);
	}

switch_body:
	LBODY
	{
		markdcl();
	}
	caseblock_list &#39;}&#39;
	{
		$$ = $3;
		popdcl();
		yyoptsemi(0);
	}

caseblock:
	case stmt_list
	{
		$$ = $1;
		$$-&gt;nbody = $2;
	}

caseblock_list:
	{
		$$ = nil;
	}
|	caseblock_list caseblock
	{
		$$ = list($1, $2);
	}

loop_body:
	LBODY
	{
		markdcl();
	}
	stmt_list &#39;}&#39;
	{
		$$ = $3;
		popdcl();
	}

range_stmt:
	expr_list &#39;=&#39; LRANGE expr
	{
		$$ = nod(ORANGE, N, $4);
		$$-&gt;list = $1;
		$$-&gt;etype = 0;	// := flag
	}
|	expr_list LCOLAS LRANGE expr
	{
		$$ = nod(ORANGE, N, $4);
		$$-&gt;list = $1;
		$$-&gt;colas = 1;
		colasdefn($1, $$);
	}

for_header:
	osimple_stmt &#39;;&#39; osimple_stmt &#39;;&#39; osimple_stmt
	{
		// init ; test ; incr
		if($5 != N &amp;&amp; $5-&gt;colas != 0)
			yyerror(&#34;cannot declare in the for-increment&#34;);
		$$ = nod(OFOR, N, N);
		if($1 != N)
			$$-&gt;ninit = list1($1);
		$$-&gt;ntest = $3;
		$$-&gt;nincr = $5;
	}
|	osimple_stmt
	{
		// normal test
		$$ = nod(OFOR, N, N);
		$$-&gt;ntest = $1;
	}
|	range_stmt

for_body:
	for_header loop_body
	{
		$$ = $1;
		$$-&gt;nbody = concat($$-&gt;nbody, $2);
		yyoptsemi(0);
	}

for_stmt:
	LFOR
	{
		markdcl();
	}
	for_body
	{
		$$ = $3;
		popdcl();
	}

if_header:
	osimple_stmt
	{
		// test
		$$ = nod(OIF, N, N);
		$$-&gt;ntest = $1;
	}
|	osimple_stmt &#39;;&#39; osimple_stmt
	{
		// init ; test
		$$ = nod(OIF, N, N);
		if($1 != N)
			$$-&gt;ninit = list1($1);
		$$-&gt;ntest = $3;
	}

if_stmt:
	LIF
	{
		markdcl();
	}
	if_header loop_body
	{
		$$ = $3;
		$$-&gt;nbody = $4;
		// no popdcl; maybe there&#39;s an LELSE
		yyoptsemi(LELSE);
	}

switch_stmt:
	LSWITCH
	{
		markdcl();
	}
	if_header
	{
		Node *n;
		n = $3-&gt;ntest;
		if(n != N &amp;&amp; n-&gt;op != OTYPESW)
			n = N;
		typesw = nod(OXXX, typesw, n);
	}
	switch_body
	{
		$$ = $3;
		$$-&gt;op = OSWITCH;
		$$-&gt;list = $5;
		typesw = typesw-&gt;left;
		popdcl();
	}

select_stmt:
	LSELECT
	{
		markdcl();
	}
	switch_body
	{
		$$ = nod(OSELECT, N, N);
		$$-&gt;list = $3;
		popdcl();
	}

/*
 * expressions
 */
expr:
	uexpr
|	expr LOROR expr
	{
		$$ = nod(OOROR, $1, $3);
	}
|	expr LANDAND expr
	{
		$$ = nod(OANDAND, $1, $3);
	}
|	expr LEQ expr
	{
		$$ = nod(OEQ, $1, $3);
	}
|	expr LNE expr
	{
		$$ = nod(ONE, $1, $3);
	}
|	expr LLT expr
	{
		$$ = nod(OLT, $1, $3);
	}
|	expr LLE expr
	{
		$$ = nod(OLE, $1, $3);
	}
|	expr LGE expr
	{
		$$ = nod(OGE, $1, $3);
	}
|	expr LGT expr
	{
		$$ = nod(OGT, $1, $3);
	}
|	expr &#39;+&#39; expr
	{
		$$ = nod(OADD, $1, $3);
	}
|	expr &#39;-&#39; expr
	{
		$$ = nod(OSUB, $1, $3);
	}
|	expr &#39;|&#39; expr
	{
		$$ = nod(OOR, $1, $3);
	}
|	expr &#39;^&#39; expr
	{
		$$ = nod(OXOR, $1, $3);
	}
|	expr &#39;*&#39; expr
	{
		$$ = nod(OMUL, $1, $3);
	}
|	expr &#39;/&#39; expr
	{
		$$ = nod(ODIV, $1, $3);
	}
|	expr &#39;%&#39; expr
	{
		$$ = nod(OMOD, $1, $3);
	}
|	expr &#39;&amp;&#39; expr
	{
		$$ = nod(OAND, $1, $3);
	}
|	expr LANDNOT expr
	{
		$$ = nod(OANDNOT, $1, $3);
	}
|	expr LLSH expr
	{
		$$ = nod(OLSH, $1, $3);
	}
|	expr LRSH expr
	{
		$$ = nod(ORSH, $1, $3);
	}
|	expr LCOMM expr
	{
		$$ = nod(OSEND, $1, $3);
	}

uexpr:
	pexpr
|	&#39;*&#39; uexpr
	{
		$$ = nod(OIND, $2, N);
	}
|	&#39;&amp;&#39; uexpr
	{
		$$ = nod(OADDR, $2, N);
	}
|	&#39;+&#39; uexpr
	{
		$$ = nod(OPLUS, $2, N);
	}
|	&#39;-&#39; uexpr
	{
		$$ = nod(OMINUS, $2, N);
	}
|	&#39;!&#39; uexpr
	{
		$$ = nod(ONOT, $2, N);
	}
|	&#39;~&#39; uexpr
	{
		yyerror(&#34;the OCOM operator is ^&#34;);
		$$ = nod(OCOM, $2, N);
	}
|	&#39;^&#39; uexpr
	{
		$$ = nod(OCOM, $2, N);
	}
|	LCOMM uexpr
	{
		$$ = nod(ORECV, $2, N);
	}

/*
 * call-like statements that
 * can be preceded by &#39;defer&#39; and &#39;go&#39;
 */
pseudocall:
	pexpr &#39;(&#39; oexpr_or_type_list &#39;)&#39;
	{
		$$ = nod(OCALL, $1, N);
		$$-&gt;list = $3;
	}

pexpr:
	LLITERAL
	{
		$$ = nodlit($1);
	}
|	name
|	pexpr &#39;.&#39; sym
	{
		if($1-&gt;op == OPACK) {
			Sym *s;
			s = restrictlookup($3-&gt;name, $1-&gt;sym-&gt;name);
			$1-&gt;used = 1;
			$$ = oldname(s);
			break;
		}
		$$ = nod(OXDOT, $1, newname($3));
	}
|	&#39;(&#39; expr_or_type &#39;)&#39;
	{
		$$ = $2;
	}
|	pexpr &#39;.&#39; &#39;(&#39; expr_or_type &#39;)&#39;
	{
		$$ = nod(ODOTTYPE, $1, $4);
	}
|	pexpr &#39;.&#39; &#39;(&#39; LTYPE &#39;)&#39;
	{
		$$ = nod(OTYPESW, N, $1);
	}
|	pexpr &#39;[&#39; expr &#39;]&#39;
	{
		$$ = nod(OINDEX, $1, $3);
	}
|	pexpr &#39;[&#39; keyval &#39;]&#39;
	{
		$$ = nod(OSLICE, $1, $3);
	}
|	pseudocall
|	convtype &#39;(&#39; expr &#39;)&#39;
	{
		// conversion
		$$ = nod(OCALL, $1, N);
		$$-&gt;list = list1($3);
	}
|	convtype lbrace braced_keyval_list &#39;}&#39;
	{
		// composite expression
		$$ = nod(OCOMPLIT, N, $1);
		$$-&gt;list = $3;

		// If the opening brace was an LBODY,
		// set up for another one now that we&#39;re done.
		// See comment in lex.c about loophack.
		if($2 == LBODY)
			loophack = 1;
	}
|	pexpr &#39;{&#39; braced_keyval_list &#39;}&#39;
	{
		// composite expression
		$$ = nod(OCOMPLIT, N, $1);
		$$-&gt;list = $3;
	}
|	fnliteral

expr_or_type:
	expr
|	non_expr_type	%prec PreferToRightParen

name_or_type:
	ntype

lbrace:
	LBODY
	{
		$$ = LBODY;
	}
|	&#39;{&#39;
	{
		$$ = &#39;{&#39;;
	}

/*
 * names and types
 *	newname is used before declared
 *	oldname is used after declared
 */
new_name:
	sym
	{
		$$ = newname($1);
	}

dcl_name:
	sym
	{
		$$ = dclname($1);
	}

onew_name:
	{
		$$ = N;
	}
|	new_name

sym:
	LNAME

name:
	sym
	{
		$$ = oldname($1);
		if($$-&gt;pack != N)
			$$-&gt;pack-&gt;used = 1;
	}

labelname:
	new_name

convtype:
	&#39;[&#39; oexpr &#39;]&#39; ntype
	{
		// array literal
		$$ = nod(OTARRAY, $2, $4);
	}
|	&#39;[&#39; dotdotdot &#39;]&#39; ntype
	{
		// array literal of nelem
		$$ = nod(OTARRAY, $2, $4);
	}
|	LMAP &#39;[&#39; ntype &#39;]&#39; ntype
	{
		// map literal
		$$ = nod(OTMAP, $3, $5);
	}
|	structtype

/*
 * to avoid parsing conflicts, type is split into
 *	channel types
 *	function types
 *	parenthesized types
 *	any other type
 * the type system makes additional restrictions,
 * but those are not implemented in the grammar.
 */
dotdotdot:
	LDDD
	{
		$$ = typenod(typ(TDDD));
	}

ntype:
	chantype
|	fntype
|	othertype
|	ptrtype
|	dotname
|	&#39;(&#39; ntype &#39;)&#39;
	{
		$$ = $2;
	}

non_expr_type:
	chantype
|	fntype
|	othertype
|	&#39;*&#39; non_expr_type
	{
		$$ = nod(OIND, $2, N);
	}
|	&#39;(&#39; non_expr_type &#39;)&#39;
	{
		$$ = $2;
	}

non_chan_type:
	fntype
|	othertype
|	ptrtype
|	dotname
|	&#39;(&#39; ntype &#39;)&#39;
	{
		$$ = $2;
	}

non_fn_type:
	chantype
|	othertype
|	ptrtype
|	dotname

dotname:
	name
|	name &#39;.&#39; sym
	{
		if($1-&gt;op == OPACK) {
			Sym *s;
			s = restrictlookup($3-&gt;name, $1-&gt;sym-&gt;name);
			$1-&gt;used = 1;
			$$ = oldname(s);
			break;
		}
		$$ = nod(OXDOT, $1, newname($3));
	}

othertype:
	&#39;[&#39; oexpr &#39;]&#39; ntype
	{
		$$ = nod(OTARRAY, $2, $4);
	}
|	&#39;[&#39; dotdotdot &#39;]&#39; ntype
	{
		// array literal of nelem
		$$ = nod(OTARRAY, $2, $4);
	}
|	LCOMM LCHAN ntype
	{
		$$ = nod(OTCHAN, $3, N);
		$$-&gt;etype = Crecv;
	}
|	LCHAN LCOMM non_chan_type
	{
		$$ = nod(OTCHAN, $3, N);
		$$-&gt;etype = Csend;
	}
|	LMAP &#39;[&#39; ntype &#39;]&#39; ntype
	{
		$$ = nod(OTMAP, $3, $5);
	}
|	structtype
|	interfacetype

ptrtype:
	&#39;*&#39; ntype
	{
		$$ = nod(OIND, $2, N);
	}

chantype:
	LCHAN ntype
	{
		$$ = nod(OTCHAN, $2, N);
		$$-&gt;etype = Cboth;
	}

structtype:
	LSTRUCT &#39;{&#39; structdcl_list osemi &#39;}&#39;
	{
		$$ = nod(OTSTRUCT, N, N);
		$$-&gt;list = $3;
		// Distinguish closing brace in struct from
		// other closing braces by explicitly marking it.
		// Used above (yylast == LSEMIBRACE).
		yylast = LSEMIBRACE;
	}
|	LSTRUCT &#39;{&#39; &#39;}&#39;
	{
		$$ = nod(OTSTRUCT, N, N);
		yylast = LSEMIBRACE;
	}

interfacetype:
	LINTERFACE &#39;{&#39; interfacedcl_list osemi &#39;}&#39;
	{
		$$ = nod(OTINTER, N, N);
		$$-&gt;list = $3;
		yylast = LSEMIBRACE;
	}
|	LINTERFACE &#39;{&#39; &#39;}&#39;
	{
		$$ = nod(OTINTER, N, N);
		yylast = LSEMIBRACE;
	}

keyval:
	expr &#39;:&#39; expr
	{
		$$ = nod(OKEY, $1, $3);
	}


/*
 * function stuff
 * all in one place to show how crappy it all is
 */
xfndcl:
	LFUNC fndcl fnbody
	{
		$$ = $2;
		$$-&gt;nbody = $3;
		funcbody($$);
	}

fndcl:
	dcl_name &#39;(&#39; oarg_type_list &#39;)&#39; fnres
	{
		Node *n;

		$$ = nod(ODCLFUNC, N, N);
		$$-&gt;nname = $1;
		if($3 == nil &amp;&amp; $5 == nil)
			$$-&gt;nname = renameinit($1);
		n = nod(OTFUNC, N, N);
		n-&gt;list = $3;
		n-&gt;rlist = $5;
		// TODO: check if nname already has an ntype
		$$-&gt;nname-&gt;ntype = n;
		funchdr($$);
	}
|	&#39;(&#39; oarg_type_list &#39;)&#39; new_name &#39;(&#39; oarg_type_list &#39;)&#39; fnres
	{
		Node *rcvr, *t;

		rcvr = $2-&gt;n;
		if($2-&gt;next != nil || $2-&gt;n-&gt;op != ODCLFIELD) {
			yyerror(&#34;bad receiver in method&#34;);
			rcvr = N;
		}

		$$ = nod(ODCLFUNC, N, N);
		$$-&gt;nname = methodname1($4, rcvr-&gt;right);
		t = nod(OTFUNC, rcvr, N);
		t-&gt;list = $6;
		t-&gt;rlist = $8;
		$$-&gt;nname-&gt;ntype = t;
		$$-&gt;shortname = $4;
		funchdr($$);
	}

fntype:
	LFUNC &#39;(&#39; oarg_type_list &#39;)&#39; fnres
	{
		$$ = nod(OTFUNC, N, N);
		$$-&gt;list = $3;
		$$-&gt;rlist = $5;
	}

fnbody:
	{
		$$ = nil;
	}
|	&#39;{&#39; stmt_list &#39;}&#39;
	{
		$$ = $2;
		if($$ == nil)
			$$ = list1(nod(OEMPTY, N, N));
		yyoptsemi(0);
	}

fnres:
	%prec NotParen
	{
		$$ = nil;
	}
|	non_fn_type
	{
		$$ = list1(nod(ODCLFIELD, N, $1));
	}
|	&#39;(&#39; oarg_type_list &#39;)&#39;
	{
		$$ = $2;
	}

fnlitdcl:
	fntype
	{
		closurehdr($1);
	}

fnliteral:
	fnlitdcl &#39;{&#39; stmt_list &#39;}&#39;
	{
		$$ = closurebody($3);
	}


/*
 * lists of things
 * note that they are left recursive
 * to conserve yacc stack. they need to
 * be reversed to interpret correctly
 */
xdcl_list:
	{
		$$ = nil;
	}
|	xdcl_list xdcl
	{
		$$ = concat($1, $2);
		if(nsyntaxerrors == 0)
			testdclstack();
	}

vardcl_list:
	vardcl
|	vardcl_list &#39;;&#39; vardcl
	{
		$$ = concat($1, $3);
	}

constdcl_list:
	constdcl1
|	constdcl_list &#39;;&#39; constdcl1
	{
		$$ = concat($1, $3);
	}

typedcl_list:
	typedcl
	{
		$$ = list1($1);
	}
|	typedcl_list &#39;;&#39; typedcl
	{
		$$ = list($1, $3);
	}

structdcl_list:
	structdcl
|	structdcl_list &#39;;&#39; structdcl
	{
		$$ = concat($1, $3);
	}

interfacedcl_list:
	interfacedcl
	{
		$$ = list1($1);
	}
|	interfacedcl_list &#39;;&#39; interfacedcl
	{
		$$ = list($1, $3);
	}

structdcl:
	new_name_list ntype oliteral
	{
		NodeList *l;

		for(l=$1; l; l=l-&gt;next) {
			l-&gt;n = nod(ODCLFIELD, l-&gt;n, $2);
			l-&gt;n-&gt;val = $3;
		}
	}
|	embed oliteral
	{
		$1-&gt;val = $2;
		$$ = list1($1);
	}
|	&#39;*&#39; embed oliteral
	{
		$2-&gt;right = nod(OIND, $2-&gt;right, N);
		$2-&gt;val = $3;
		$$ = list1($2);
	}

packname:
	LNAME
	{
		Node *n;

		$$ = $1;
		n = oldname($1);
		if(n-&gt;pack != N)
			n-&gt;pack-&gt;used = 1;
	}
|	LNAME &#39;.&#39; sym
	{
		char *pkg;

		if($1-&gt;def == N || $1-&gt;def-&gt;op != OPACK) {
			yyerror(&#34;%S is not a package&#34;, $1);
			pkg = $1-&gt;name;
		} else {
			$1-&gt;def-&gt;used = 1;
			pkg = $1-&gt;def-&gt;sym-&gt;name;
		}
		$$ = restrictlookup($3-&gt;name, pkg);
	}

embed:
	packname
	{
		$$ = embedded($1);
	}

interfacedcl:
	new_name indcl
	{
		$$ = nod(ODCLFIELD, $1, $2);
	}
|	packname
	{
		$$ = nod(ODCLFIELD, N, oldname($1));
	}

indcl:
	&#39;(&#39; oarg_type_list &#39;)&#39; fnres
	{
		// without func keyword
		$$ = nod(OTFUNC, fakethis(), N);
		$$-&gt;list = $2;
		$$-&gt;rlist = $4;
	}

/*
 * function arguments.
 */
arg_type:
	name_or_type
|	sym name_or_type
	{
		$$ = nod(ONONAME, N, N);
		$$-&gt;sym = $1;
		$$ = nod(OKEY, $$, $2);
	}
|	sym dotdotdot
	{
		$$ = nod(ONONAME, N, N);
		$$-&gt;sym = $1;
		$$ = nod(OKEY, $$, $2);
	}
|	dotdotdot

arg_type_list:
	arg_type
	{
		$$ = list1($1);
	}
|	arg_type_list &#39;,&#39; arg_type
	{
		$$ = list($1, $3);
	}

oarg_type_list:
	{
		$$ = nil;
	}
|	arg_type_list
	{
		$$ = checkarglist($1);
	}

/*
 * statement
 */
stmt:
	{
		$$ = N;
	}
|	simple_stmt
|	compound_stmt
|	common_dcl
	{
		$$ = liststmt($1);
	}
|	for_stmt
|	switch_stmt
|	select_stmt
|	if_stmt
	{
		popdcl();
		$$ = $1;
	}
|	if_stmt LELSE stmt
	{
		popdcl();
		$$ = $1;
		$$-&gt;nelse = list1($3);
	}
|	error
	{
		$$ = N;
	}
|	labelname &#39;:&#39; stmt
	{
		NodeList *l;

		l = list1(nod(OLABEL, $1, $3));
		if($3)
			l = list(l, $3);
		$$ = liststmt(l);
	}
|	LFALL
	{
		// will be converted to OFALL
		$$ = nod(OXFALL, N, N);
	}
|	LBREAK onew_name
	{
		$$ = nod(OBREAK, $2, N);
	}
|	LCONTINUE onew_name
	{
		$$ = nod(OCONTINUE, $2, N);
	}
|	LGO pseudocall
	{
		$$ = nod(OPROC, $2, N);
	}
|	LDEFER pseudocall
	{
		$$ = nod(ODEFER, $2, N);
	}
|	LGOTO new_name
	{
		$$ = nod(OGOTO, $2, N);
	}
|	LRETURN oexpr_list
	{
		$$ = nod(ORETURN, N, N);
		$$-&gt;list = $2;
	}

stmt_list:
	stmt
	{
		$$ = nil;
		if($1 != N)
			$$ = list1($1);
	}
|	stmt_list &#39;;&#39; stmt
	{
		$$ = $1;
		if($3 != N)
			$$ = list($$, $3);
	}

new_name_list:
	new_name
	{
		$$ = list1($1);
	}
|	new_name_list &#39;,&#39; new_name
	{
		$$ = list($1, $3);
	}

dcl_name_list:
	dcl_name
	{
		$$ = list1($1);
	}
|	dcl_name_list &#39;,&#39; dcl_name
	{
		$$ = list($1, $3);
	}

expr_list:
	expr
	{
		$$ = list1($1);
	}
|	expr_list &#39;,&#39; expr
	{
		$$ = list($1, $3);
	}

expr_or_type_list:
	expr_or_type
	{
		$$ = list1($1);
	}
|	expr_or_type_list &#39;,&#39; expr_or_type
	{
		$$ = list($1, $3);
	}

/*
 * list of combo of keyval and val
 */
keyval_list:
	keyval
	{
		$$ = list1($1);
	}
|	expr
	{
		$$ = list1($1);
	}
|	keyval_list &#39;,&#39; keyval
	{
		$$ = list($1, $3);
	}
|	keyval_list &#39;,&#39; expr
	{
		$$ = list($1, $3);
	}

braced_keyval_list:
	{
		$$ = nil;
	}
|	keyval_list ocomma
	{
		$$ = $1;
	}

/*
 * optional things
 */
osemi:
	%prec NotSemi
|	&#39;;&#39;

ocomma:
|	&#39;,&#39;

oexpr:
	{
		$$ = N;
	}
|	expr

oexpr_list:
	{
		$$ = nil;
	}
|	expr_list

oexpr_or_type_list:
	{
		$$ = nil;
	}
|	expr_or_type_list

osimple_stmt:
	{
		$$ = N;
	}
|	simple_stmt

ohidden_funarg_list:
	{
		$$ = nil;
	}
|	hidden_funarg_list

ohidden_structdcl_list:
	{
		$$ = nil;
	}
|	hidden_structdcl_list

ohidden_interfacedcl_list:
	{
		$$ = nil;
	}
|	hidden_interfacedcl_list

oliteral:
	{
		$$.ctype = CTxxx;
	}
|	LLITERAL

/*
 * import syntax from header of
 * an output package
 */
hidden_import:
	LPACKAGE sym
	/* variables */
|	LVAR hidden_pkg_importsym hidden_type
	{
		importvar($2, $3, PEXTERN);
	}
|	LCONST hidden_pkg_importsym &#39;=&#39; hidden_constant
	{
		importconst($2, types[TIDEAL], $4);
	}
|	LCONST hidden_pkg_importsym hidden_type &#39;=&#39; hidden_constant
	{
		importconst($2, $3, $5);
	}
|	LTYPE hidden_pkgtype hidden_type
	{
		importtype($2, $3);
	}
|	LFUNC hidden_pkg_importsym &#39;(&#39; ohidden_funarg_list &#39;)&#39; ohidden_funres
	{
		importvar($2, functype(N, $4, $6), PFUNC);
	}
|	LFUNC &#39;(&#39; hidden_funarg_list &#39;)&#39; sym &#39;(&#39; ohidden_funarg_list &#39;)&#39; ohidden_funres
	{
		if($3-&gt;next != nil || $3-&gt;n-&gt;op != ODCLFIELD) {
			yyerror(&#34;bad receiver in method&#34;);
			YYERROR;
		}
		importmethod($5, functype($3-&gt;n, $7, $9));
	}

hidden_pkgtype:
	hidden_pkg_importsym
	{
		$$ = pkgtype($1);
		importsym($1, OTYPE);
	}

hidden_type:
	hidden_type1
|	hidden_type2

hidden_type1:
	hidden_importsym
	{
		$$ = pkgtype($1);
	}
|	LNAME
	{
		// predefined name like uint8
		$1 = pkglookup($1-&gt;name, &#34;/builtin/&#34;);
		if($1-&gt;def == N || $1-&gt;def-&gt;op != OTYPE) {
			yyerror(&#34;%s is not a type&#34;, $1-&gt;name);
			$$ = T;
		} else
			$$ = $1-&gt;def-&gt;type;
	}
|	&#39;[&#39; &#39;]&#39; hidden_type
	{
		$$ = aindex(N, $3);
	}
|	&#39;[&#39; LLITERAL &#39;]&#39; hidden_type
	{
		$$ = aindex(nodlit($2), $4);
	}
|	LMAP &#39;[&#39; hidden_type &#39;]&#39; hidden_type
	{
		$$ = maptype($3, $5);
	}
|	LSTRUCT &#39;{&#39; ohidden_structdcl_list &#39;}&#39;
	{
		$$ = dostruct($3, TSTRUCT);
	}
|	LINTERFACE &#39;{&#39; ohidden_interfacedcl_list &#39;}&#39;
	{
		$$ = dostruct($3, TINTER);
		$$ = sortinter($$);
	}
|	&#39;*&#39; hidden_type
	{
		$$ = ptrto($2);
	}
|	LCOMM LCHAN hidden_type
	{
		$$ = typ(TCHAN);
		$$-&gt;type = $3;
		$$-&gt;chan = Crecv;
	}
|	LCHAN LCOMM hidden_type1
	{
		$$ = typ(TCHAN);
		$$-&gt;type = $3;
		$$-&gt;chan = Csend;
	}
|	LDDD
	{
		$$ = typ(TDDD);
	}

hidden_type2:
	LCHAN hidden_type
	{
		$$ = typ(TCHAN);
		$$-&gt;type = $2;
		$$-&gt;chan = Cboth;
	}
|	LFUNC &#39;(&#39; ohidden_funarg_list &#39;)&#39; ohidden_funres
	{
		$$ = functype(nil, $3, $5);
	}

hidden_dcl:
	sym hidden_type
	{
		$$ = nod(ODCLFIELD, newname($1), typenod($2));
	}
|	&#39;?&#39; hidden_type
	{
		$$ = nod(ODCLFIELD, N, typenod($2));
	}

hidden_structdcl:
	sym hidden_type oliteral
	{
		$$ = nod(ODCLFIELD, newname($1), typenod($2));
		$$-&gt;val = $3;
	}
|	&#39;?&#39; hidden_type oliteral
	{
		Sym *s;

		s = $2-&gt;sym;
		if(s == S &amp;&amp; isptr[$2-&gt;etype])
			s = $2-&gt;type-&gt;sym;
		if(s &amp;&amp; strcmp(s-&gt;package, &#34;/builtin/&#34;) == 0)
			s = lookup(s-&gt;name);
		$$ = embedded(s);
		$$-&gt;right = typenod($2);
		$$-&gt;val = $3;
	}

hidden_interfacedcl:
	sym &#39;(&#39; ohidden_funarg_list &#39;)&#39; ohidden_funres
	{
		$$ = nod(ODCLFIELD, newname($1), typenod(functype(fakethis(), $3, $5)));
	}

ohidden_funres:
	{
		$$ = nil;
	}
|	hidden_funres

hidden_funres:
	&#39;(&#39; ohidden_funarg_list &#39;)&#39;
	{
		$$ = $2;
	}
|	hidden_type1
	{
		$$ = list1(nod(ODCLFIELD, N, typenod($1)));
	}

hidden_constant:
	LLITERAL
	{
		$$ = nodlit($1);
	}
|	&#39;-&#39; LLITERAL
	{
		$$ = nodlit($2);
		switch($$-&gt;val.ctype){
		case CTINT:
			mpnegfix($$-&gt;val.u.xval);
			break;
		case CTFLT:
			mpnegflt($$-&gt;val.u.fval);
			break;
		default:
			yyerror(&#34;bad negated constant&#34;);
		}
	}
|	sym
	{
		$$ = oldname(pkglookup($1-&gt;name, &#34;/builtin/&#34;));
		if($$-&gt;op != OLITERAL)
			yyerror(&#34;bad constant %S&#34;, $$-&gt;sym);
	}

hidden_importsym:
	sym &#39;.&#39; sym
	{
		$$ = pkglookup($3-&gt;name, $1-&gt;name);
	}

hidden_pkg_importsym:
	hidden_importsym
	{
		$$ = $1;
		structpkg = $$-&gt;package;
	}

hidden_import_list:
|	hidden_import_list hidden_import

hidden_funarg_list:
	hidden_dcl
	{
		$$ = list1($1);
	}
|	hidden_funarg_list &#39;,&#39; hidden_dcl
	{
		$$ = list($1, $3);
	}

hidden_structdcl_list:
	hidden_structdcl
	{
		$$ = list1($1);
	}
|	hidden_structdcl_list &#39;;&#39; hidden_structdcl
	{
		$$ = list($1, $3);
	}

hidden_interfacedcl_list:
	hidden_interfacedcl
	{
		$$ = list1($1);
	}
|	hidden_interfacedcl_list &#39;;&#39; hidden_interfacedcl
	{
		$$ = list($1, $3);
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
