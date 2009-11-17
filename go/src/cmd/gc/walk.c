<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/walk.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/walk.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;

static	Node*	walkprint(Node*, NodeList**);
static	Node*	conv(Node*, Type*);
static	Node*	mapfn(char*, Type*);
static	Node*	makenewvar(Type*, NodeList**, Node**);
enum
{
	Inone,
	I2T,
	I2T2,
	I2I,
	I2Ix,
	I2I2,
	T2I,
	I2Isame,
	E2T,
	E2T2,
	E2I,
	E2I2,
	I2E,
	I2E2,
	T2E,
	E2Esame,
};

// can this code branch reach the end
// without an undcontitional RETURN
// this is hard, so it is conservative
int
walkret(NodeList *l)
{
	Node *n;

loop:
	while(l &amp;&amp; l-&gt;next)
		l = l-&gt;next;
	if(l == nil)
		return 1;

	// at this point, we have the last
	// statement of the function
	n = l-&gt;n;
	switch(n-&gt;op) {
	case OBLOCK:
		l = n-&gt;list;
		goto loop;

	case OGOTO:
	case ORETURN:
	case OPANIC:
	case OPANICN:
		return 0;
		break;
	}

	// all other statements
	// will flow to the end
	return 1;
}

void
walk(Node *fn)
{
	char s[50];
	NodeList *l;
	Node *n;
	int lno;

	curfn = fn;
	if(debug[&#39;W&#39;]) {
		snprint(s, sizeof(s), &#34;\nbefore %S&#34;, curfn-&gt;nname-&gt;sym);
		dumplist(s, curfn-&gt;nbody);
	}
	if(curfn-&gt;type-&gt;outtuple)
		if(walkret(curfn-&gt;nbody))
			yyerror(&#34;function ends without a return statement&#34;);
	typechecklist(curfn-&gt;nbody, Etop);
	lno = lineno;
	for(l=fn-&gt;dcl; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op != ONAME || n-&gt;class != PAUTO)
			continue;
		lineno = n-&gt;lineno;
		typecheck(&amp;n, Erv | Easgn);	// only needed for unused variables
		if(!n-&gt;used &amp;&amp; n-&gt;sym-&gt;name[0] != &#39;&amp;&#39; &amp;&amp; !nsyntaxerrors)
			yyerror(&#34;%S declared and not used&#34;, n-&gt;sym);
	}
	lineno = lno;
	if(nerrors != 0)
		return;
	walkstmtlist(curfn-&gt;nbody);
	if(debug[&#39;W&#39;]) {
		snprint(s, sizeof(s), &#34;after walk %S&#34;, curfn-&gt;nname-&gt;sym);
		dumplist(s, curfn-&gt;nbody);
	}
	heapmoves();
	if(debug[&#39;W&#39;] &amp;&amp; curfn-&gt;enter != nil) {
		snprint(s, sizeof(s), &#34;enter %S&#34;, curfn-&gt;nname-&gt;sym);
		dumplist(s, curfn-&gt;enter);
	}
}

void
gettype(Node **np, NodeList **init)
{
	if(debug[&#39;W&#39;])
		dump(&#34;\nbefore gettype&#34;, *np);
	typecheck(np, Erv);
	if(debug[&#39;W&#39;])
		dump(&#34;after gettype&#34;, *np);
}

void
walkdeflist(NodeList *l)
{
	for(; l; l=l-&gt;next)
		walkdef(l-&gt;n);
}

void
walkdef(Node *n)
{
	int lno, maplineno, embedlineno;
	NodeList *init;
	Node *e;
	Type *t;

	lno = lineno;
	setlineno(n);

	if(n-&gt;op == ONONAME) {
		if(!n-&gt;diag) {
			n-&gt;diag = 1;
			yyerror(&#34;undefined: %S&#34;, n-&gt;sym);
		}
		return;
	}

	if(n-&gt;walkdef == 1)
		return;
	if(n-&gt;walkdef == 2) {
		// TODO(rsc): better loop message
		fatal(&#34;loop&#34;);
	}
	n-&gt;walkdef = 2;

	if(n-&gt;type != T || n-&gt;sym == S)	// builtin or no name
		goto ret;

	init = nil;
	switch(n-&gt;op) {
	default:
		fatal(&#34;walkdef %O&#34;, n-&gt;op);

	case OLITERAL:
		if(n-&gt;ntype != N) {
			typecheck(&amp;n-&gt;ntype, Etype);
			n-&gt;type = n-&gt;ntype-&gt;type;
			n-&gt;ntype = N;
			if(n-&gt;type == T) {
				n-&gt;diag = 1;
				goto ret;
			}
		}
		e = n-&gt;defn;
		n-&gt;defn = N;
		if(e == N) {
			lineno = n-&gt;lineno;
			dump(&#34;walkdef nil defn&#34;, n);
			yyerror(&#34;xxx&#34;);
		}
		typecheck(&amp;e, Erv | Eiota);
		if(e-&gt;type != T &amp;&amp; e-&gt;op != OLITERAL) {
			yyerror(&#34;const initializer must be constant&#34;);
			goto ret;
		}
		t = n-&gt;type;
		if(t != T)
			convlit(&amp;e, t);
		n-&gt;val = e-&gt;val;
		n-&gt;type = e-&gt;type;
		break;

	case ONAME:
		if(n-&gt;ntype != N) {
			typecheck(&amp;n-&gt;ntype, Etype);
			n-&gt;type = n-&gt;ntype-&gt;type;
			if(n-&gt;type == T) {
				n-&gt;diag = 1;
				goto ret;
			}
		}
		if(n-&gt;type != T)
			break;
		if(n-&gt;defn == N)
			fatal(&#34;var without type, init: %S&#34;, n-&gt;sym);
		if(n-&gt;defn-&gt;op == ONAME) {
			typecheck(&amp;n-&gt;defn, Erv);
			n-&gt;type = n-&gt;defn-&gt;type;
			break;
		}
		typecheck(&amp;n-&gt;defn, Etop);	// fills in n-&gt;type
		break;

	case OTYPE:
		n-&gt;walkdef = 1;
		n-&gt;type = typ(TFORW);
		n-&gt;type-&gt;sym = n-&gt;sym;
		n-&gt;typecheck = 1;
		typecheck(&amp;n-&gt;ntype, Etype);
		if((t = n-&gt;ntype-&gt;type) == T) {
			n-&gt;diag = 1;
			goto ret;
		}

		// copy new type and clear fields
		// that don&#39;t come along
		maplineno = n-&gt;type-&gt;maplineno;
		embedlineno = n-&gt;type-&gt;embedlineno;
		*n-&gt;type = *t;
		t = n-&gt;type;
		t-&gt;sym = n-&gt;sym;
		t-&gt;local = n-&gt;local;
		t-&gt;vargen = n-&gt;vargen;
		t-&gt;siggen = 0;
		t-&gt;printed = 0;
		t-&gt;method = nil;
		t-&gt;nod = N;

		// double-check use of type as map key
		// TODO(rsc): also use of type as receiver?
		if(maplineno) {
			lineno = maplineno;
			maptype(n-&gt;type, types[TBOOL]);
		}
		if(embedlineno) {
			lineno = embedlineno;
			if(isptr[t-&gt;etype])
				yyerror(&#34;embedded type cannot be a pointer&#34;);
		}
		break;

	case OPACK:
		// nothing to see here
		break;
	}

ret:
	lineno = lno;
	n-&gt;walkdef = 1;
}

void
walkstmtlist(NodeList *l)
{
	for(; l; l=l-&gt;next)
		walkstmt(&amp;l-&gt;n);
}

static int
samelist(NodeList *a, NodeList *b)
{
	for(; a &amp;&amp; b; a=a-&gt;next, b=b-&gt;next)
		if(a-&gt;n != b-&gt;n)
			return 0;
	return a == b;
}


void
walkstmt(Node **np)
{
	NodeList *init;
	NodeList *ll, *rl;
	int cl, lno;
	Node *n;

	n = *np;
	if(n == N)
		return;

	lno = lineno;
	setlineno(n);

	switch(n-&gt;op) {
	default:
		if(n-&gt;op == ONAME)
			yyerror(&#34;%S is not a top level statement&#34;, n-&gt;sym);
		else
			yyerror(&#34;%O is not a top level statement&#34;, n-&gt;op);
		dump(&#34;nottop&#34;, n);
		break;

	case OAPPENDSTR:
	case OASOP:
	case OAS:
	case OAS2:
	case OAS2DOTTYPE:
	case OAS2RECV:
	case OAS2FUNC:
	case OAS2MAPW:
	case OAS2MAPR:
	case OCLOSE:
	case OCLOSED:
	case OCALLMETH:
	case OCALLINTER:
	case OCALL:
	case OCALLFUNC:
	case OSEND:
	case ORECV:
	case OPRINT:
	case OPRINTN:
	case OPANIC:
	case OPANICN:
	case OEMPTY:
		if(n-&gt;typecheck == 0)
			fatal(&#34;missing typecheck&#34;);
		init = n-&gt;ninit;
		n-&gt;ninit = nil;
		walkexpr(&amp;n, &amp;init);
		n-&gt;ninit = concat(init, n-&gt;ninit);
		break;

	case OBREAK:
	case ODCL:
	case OCONTINUE:
	case OFALL:
	case OGOTO:
	case OLABEL:
	case ODCLCONST:
	case ODCLTYPE:
		break;

	case OBLOCK:
		walkstmtlist(n-&gt;list);
		break;

	case OXCASE:
		yyerror(&#34;case statement out of place&#34;);
		n-&gt;op = OCASE;
	case OCASE:
		walkstmt(&amp;n-&gt;right);
		break;

	case ODEFER:
		hasdefer = 1;
		walkexpr(&amp;n-&gt;left, &amp;n-&gt;ninit);
		break;

	case OFOR:
		walkstmtlist(n-&gt;ninit);
		if(n-&gt;ntest != N) {
			walkstmtlist(n-&gt;ntest-&gt;ninit);
			walkexpr(&amp;n-&gt;ntest, &amp;n-&gt;ninit);
		}
		walkstmt(&amp;n-&gt;nincr);
		walkstmtlist(n-&gt;nbody);
		break;

	case OIF:
		walkstmtlist(n-&gt;ninit);
		walkexpr(&amp;n-&gt;ntest, &amp;n-&gt;ninit);
		walkstmtlist(n-&gt;nbody);
		walkstmtlist(n-&gt;nelse);
		break;

	case OPROC:
		walkexpr(&amp;n-&gt;left, &amp;n-&gt;ninit);
		break;

	case ORETURN:
		walkexprlist(n-&gt;list, &amp;n-&gt;ninit);
		if(curfn-&gt;type-&gt;outnamed &amp;&amp; count(n-&gt;list) != 1) {
			if(n-&gt;list == nil) {
				// print(&#34;special return\n&#34;);
				break;
			}
			// assign to the function out parameters,
			// so that reorder3 can fix up conflicts
			rl = nil;
			for(ll=curfn-&gt;dcl; ll != nil; ll=ll-&gt;next) {
				cl = ll-&gt;n-&gt;class &amp; ~PHEAP;
				if(cl == PAUTO)
					break;
				if(cl == PPARAMOUT)
					rl = list(rl, ll-&gt;n);
			}
			if(samelist(rl, n-&gt;list)) {
				// special return in disguise
				n-&gt;list = nil;
				break;
			}
			ll = ascompatee(n-&gt;op, rl, n-&gt;list, &amp;n-&gt;ninit);
			n-&gt;list = reorder3(ll);
			break;
		}
		ll = ascompatte(n-&gt;op, getoutarg(curfn-&gt;type), n-&gt;list, 1, &amp;n-&gt;ninit);
		n-&gt;list = reorder4(ll);
		break;

	case OSELECT:
		walkselect(n);
		break;

	case OSWITCH:
		walkswitch(n);
		break;

	case ORANGE:
		walkrange(n);
		break;

	case OXFALL:
		yyerror(&#34;fallthrough statement out of place&#34;);
		n-&gt;op = OFALL;
		break;
	}

	*np = n;
}


/*
 * walk the whole tree of the body of an
 * expression or simple statement.
 * the types expressions are calculated.
 * compile-time constants are evaluated.
 * complex side effects like statements are appended to init
 */

void
walkexprlist(NodeList *l, NodeList **init)
{
	for(; l; l=l-&gt;next)
		walkexpr(&amp;l-&gt;n, init);
}

void
walkexpr(Node **np, NodeList **init)
{
	Node *r, *l;
	NodeList *ll, *lr;
	Type *t;
	int et;
	int32 lno;
	Node *n, *fn;

	n = *np;

	if(n == N)
		return;

	if(init == &amp;n-&gt;ninit) {
		// not okay to use n-&gt;ninit when walking n,
		// because we might replace n with some other node
		// and would lose the init list.
		fatal(&#34;walkexpr init == &amp;n-&gt;ninit&#34;);
	}

	// annoying case - not typechecked
	if(n-&gt;op == OKEY) {
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right, init);
		return;
	}

	lno = setlineno(n);

	if(debug[&#39;w&#39;] &gt; 1)
		dump(&#34;walk-before&#34;, n);

	if(n-&gt;typecheck != 1) {
		dump(&#34;missed typecheck&#34;, n);
		fatal(&#34;missed typecheck&#34;);
	}

	t = T;
	et = Txxx;

	switch(n-&gt;op) {
	default:
		dump(&#34;walk&#34;, n);
		fatal(&#34;walkexpr: switch 1 unknown op %N&#34;, n);
		goto ret;

	case OTYPE:
	case ONONAME:
	case OINDREG:
	case OEMPTY:
		goto ret;

	case ONOT:
	case OMINUS:
	case OPLUS:
	case OCOM:
	case OLEN:
	case OCAP:
	case ODOT:
	case ODOTPTR:
	case ODOTMETH:
	case ODOTINTER:
	case OIND:
		walkexpr(&amp;n-&gt;left, init);
		goto ret;

	case OLSH:
	case ORSH:
	case OAND:
	case OOR:
	case OXOR:
	case OANDAND:
	case OOROR:
	case OSUB:
	case OMUL:
	case OEQ:
	case ONE:
	case OLT:
	case OLE:
	case OGE:
	case OGT:
	case OADD:
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right, init);
		goto ret;

	case OPRINT:
	case OPRINTN:
	case OPANIC:
	case OPANICN:
		walkexprlist(n-&gt;list, init);
		n = walkprint(n, init);
		goto ret;

	case OLITERAL:
		n-&gt;addable = 1;
		goto ret;

	case ONAME:
		if(!(n-&gt;class &amp; PHEAP) &amp;&amp; n-&gt;class != PPARAMREF)
			n-&gt;addable = 1;
		goto ret;

	case OCALLINTER:
		t = n-&gt;left-&gt;type;
		if(n-&gt;list &amp;&amp; n-&gt;list-&gt;n-&gt;op == OAS)
			goto ret;
		walkexpr(&amp;n-&gt;left, init);
		walkexprlist(n-&gt;list, init);
		ll = ascompatte(n-&gt;op, getinarg(t), n-&gt;list, 0, init);
		n-&gt;list = reorder1(ll);
		goto ret;

	case OCALLFUNC:
		t = n-&gt;left-&gt;type;
		if(n-&gt;list &amp;&amp; n-&gt;list-&gt;n-&gt;op == OAS)
			goto ret;
		walkexpr(&amp;n-&gt;left, init);
		walkexprlist(n-&gt;list, init);
		ll = ascompatte(n-&gt;op, getinarg(t), n-&gt;list, 0, init);
		n-&gt;list = reorder1(ll);
		if(isselect(n)) {
			// special prob with selectsend and selectrecv:
			// if chan is nil, they don&#39;t know big the channel
			// element is and therefore don&#39;t know how to find
			// the output bool, so we clear it before the call.
			Node *b;
			b = nodbool(0);
			lr = ascompatte(n-&gt;op, getoutarg(t), list1(b), 0, init);
			n-&gt;list = concat(n-&gt;list, lr);
		}
		goto ret;

	case OCALLMETH:
		t = n-&gt;left-&gt;type;
		if(n-&gt;list &amp;&amp; n-&gt;list-&gt;n-&gt;op == OAS)
			goto ret;
		walkexpr(&amp;n-&gt;left, init);
		walkexprlist(n-&gt;list, init);
		ll = ascompatte(n-&gt;op, getinarg(t), n-&gt;list, 0, init);
		lr = ascompatte(n-&gt;op, getthis(t), list1(n-&gt;left-&gt;left), 0, init);
		ll = concat(ll, lr);
		n-&gt;left-&gt;left = N;
		ullmancalc(n-&gt;left);
		n-&gt;list = reorder1(ll);
		goto ret;

	case OAS:
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		walkexpr(&amp;n-&gt;left, init);
		if(oaslit(n, init))
			goto ret;
		walkexpr(&amp;n-&gt;right, init);
		l = n-&gt;left;
		r = n-&gt;right;
		if(l == N || r == N)
			goto ret;
		r = ascompatee1(n-&gt;op, l, r, init);
		if(r != N)
			n = r;
		goto ret;

	case OAS2:
	as2:
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		walkexprlist(n-&gt;list, init);
		walkexprlist(n-&gt;rlist, init);
		ll = ascompatee(OAS, n-&gt;list, n-&gt;rlist, init);
		ll = reorder3(ll);
		n = liststmt(ll);
		goto ret;

	case OAS2FUNC:
	as2func:
		// a,b,... = fn()
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		r = n-&gt;rlist-&gt;n;
		walkexprlist(n-&gt;list, init);
		walkexpr(&amp;r, init);
		ll = ascompatet(n-&gt;op, n-&gt;list, &amp;r-&gt;type, 0, init);
		n = liststmt(concat(list1(r), ll));
		goto ret;

	case OAS2RECV:
		// a,b = &lt;-c
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		r = n-&gt;rlist-&gt;n;
		walkexprlist(n-&gt;list, init);
		walkexpr(&amp;r-&gt;left, init);
		fn = chanfn(&#34;chanrecv2&#34;, 2, r-&gt;left-&gt;type);
		r = mkcall1(fn, getoutargx(fn-&gt;type), init, r-&gt;left);
		n-&gt;rlist-&gt;n = r;
		n-&gt;op = OAS2FUNC;
		goto as2func;

	case OAS2MAPR:
		// a,b = m[i];
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		r = n-&gt;rlist-&gt;n;
		walkexprlist(n-&gt;list, init);
		walkexpr(&amp;r-&gt;left, init);
		fn = mapfn(&#34;mapaccess2&#34;, r-&gt;left-&gt;type);
		r = mkcall1(fn, getoutargx(fn-&gt;type), init, r-&gt;left, r-&gt;right);
		n-&gt;rlist = list1(r);
		n-&gt;op = OAS2FUNC;
		goto as2func;

	case OAS2MAPW:
		// map[] = a,b - mapassign2
		// a,b = m[i];
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		walkexprlist(n-&gt;list, init);
		l = n-&gt;list-&gt;n;
		t = l-&gt;left-&gt;type;
		n = mkcall1(mapfn(&#34;mapassign2&#34;, t), T, init, l-&gt;left, l-&gt;right, n-&gt;rlist-&gt;n, n-&gt;rlist-&gt;next-&gt;n);
		goto ret;

	case OAS2DOTTYPE:
		// a,b = i.(T)
		*init = concat(*init, n-&gt;ninit);
		n-&gt;ninit = nil;
		r = n-&gt;rlist-&gt;n;
		walkexprlist(n-&gt;list, init);
		walkdottype(r, init);
		et = ifaceas1(r-&gt;type, r-&gt;left-&gt;type, 1);
		switch(et) {
		case I2Isame:
		case E2Esame:
			n-&gt;rlist = list(list1(r-&gt;left), nodbool(1));
			typechecklist(n-&gt;rlist, Erv);
			goto as2;
		case I2E:
			n-&gt;list = list(list1(n-&gt;right), nodbool(1));
			typechecklist(n-&gt;rlist, Erv);
			goto as2;
		case I2T:
			et = I2T2;
			break;
		case I2Ix:
			et = I2I2;
			break;
		case E2I:
			et = E2I2;
			break;
		case E2T:
			et = E2T2;
			break;
		default:
			et = Inone;
			break;
		}
		if(et == Inone)
			break;
		r = ifacecvt(r-&gt;type, r-&gt;left, et, init);
		ll = ascompatet(n-&gt;op, n-&gt;list, &amp;r-&gt;type, 0, init);
		n = liststmt(concat(list1(r), ll));
		goto ret;

	case ODOTTYPE:
		walkdottype(n, init);
		walkconv(&amp;n, init);
		goto ret;

	case OCONV:
	case OCONVNOP:
		if(thechar == &#39;5&#39;) {
			if(isfloat[n-&gt;left-&gt;type-&gt;etype] &amp;&amp; (n-&gt;type-&gt;etype == TINT64 || n-&gt;type-&gt;etype == TUINT64)) {
				n = mkcall(&#34;float64toint64&#34;, n-&gt;type, init, conv(n-&gt;left, types[TFLOAT64]));
				goto ret;
			}
			if((n-&gt;left-&gt;type-&gt;etype == TINT64 || n-&gt;left-&gt;type-&gt;etype == TUINT64) &amp;&amp; isfloat[n-&gt;type-&gt;etype]) {
				n = mkcall(&#34;int64tofloat64&#34;, n-&gt;type, init, conv(n-&gt;left, types[TINT64]));
				goto ret;
			}
		}
		walkexpr(&amp;n-&gt;left, init);
		goto ret;

	case OASOP:
		walkexpr(&amp;n-&gt;left, init);
		l = n-&gt;left;
		if(l-&gt;op == OINDEXMAP)
			n = mapop(n, init);
		walkexpr(&amp;n-&gt;right, init);
		if(n-&gt;etype == OANDNOT) {
			n-&gt;etype = OAND;
			n-&gt;right = nod(OCOM, n-&gt;right, N);
			typecheck(&amp;n-&gt;right, Erv);
			goto ret;
		}

		/*
		 * on 32-bit arch, rewrite 64-bit ops into l = l op r
		 */
		et = n-&gt;left-&gt;type-&gt;etype;
		if(widthptr == 4 &amp;&amp; (et == TUINT64 || et == TINT64)) {
			l = saferef(n-&gt;left, init);
			r = nod(OAS, l, nod(n-&gt;etype, l, n-&gt;right));
			typecheck(&amp;r, Etop);
			walkexpr(&amp;r, init);
			n = r;
		}
		goto ret;

	case OANDNOT:
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right, init);
		n-&gt;op = OAND;
		n-&gt;right = nod(OCOM, n-&gt;right, N);
		typecheck(&amp;n-&gt;right, Erv);
		goto ret;

	case ODIV:
	case OMOD:
		/*
		 * rewrite div and mod into function calls
		 * on 32-bit architectures.
		 */
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right, init);
		et = n-&gt;left-&gt;type-&gt;etype;
		if(widthptr &gt; 4 || (et != TUINT64 &amp;&amp; et != TINT64))
			goto ret;
		if(et == TINT64)
			strcpy(namebuf, &#34;int64&#34;);
		else
			strcpy(namebuf, &#34;uint64&#34;);
		if(n-&gt;op == ODIV)
			strcat(namebuf, &#34;div&#34;);
		else
			strcat(namebuf, &#34;mod&#34;);
		n = mkcall(namebuf, n-&gt;type, init,
			conv(n-&gt;left, types[et]), conv(n-&gt;right, types[et]));
		goto ret;

	case OINDEX:
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right, init);
		goto ret;

	case OINDEXMAP:
		if(n-&gt;etype == 1)
			goto ret;
		t = n-&gt;left-&gt;type;
		n = mkcall1(mapfn(&#34;mapaccess1&#34;, t), t-&gt;type, init, n-&gt;left, n-&gt;right);
		goto ret;

	case ORECV:
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right, init);
		n = mkcall1(chanfn(&#34;chanrecv1&#34;, 2, n-&gt;left-&gt;type), n-&gt;type, init, n-&gt;left);
		goto ret;

	case OSLICE:
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right-&gt;left, init);
		walkexpr(&amp;n-&gt;right-&gt;right, init);
		// dynamic slice
		// sliceslice(old []any, lb int, hb int, width int) (ary []any)
		t = n-&gt;type;
		fn = syslook(&#34;sliceslice&#34;, 1);
		argtype(fn, t-&gt;type);			// any-1
		argtype(fn, t-&gt;type);			// any-2
		n = mkcall1(fn, t, init,
			n-&gt;left,
			conv(n-&gt;right-&gt;left, types[TINT]),
			conv(n-&gt;right-&gt;right, types[TINT]),
			nodintconst(t-&gt;type-&gt;width));
		goto ret;

	case OSLICEARR:
		walkexpr(&amp;n-&gt;left, init);
		walkexpr(&amp;n-&gt;right-&gt;left, init);
		walkexpr(&amp;n-&gt;right-&gt;right, init);
		// static slice
		// slicearray(old *any, nel int, lb int, hb int, width int) (ary []any)
		t = n-&gt;type;
		fn = syslook(&#34;slicearray&#34;, 1);
		argtype(fn, n-&gt;left-&gt;type);	// any-1
		argtype(fn, t-&gt;type);			// any-2
		n = mkcall1(fn, t, init,
			nod(OADDR, n-&gt;left, N), nodintconst(n-&gt;left-&gt;type-&gt;bound),
			conv(n-&gt;right-&gt;left, types[TINT]),
			conv(n-&gt;right-&gt;right, types[TINT]),
			nodintconst(t-&gt;type-&gt;width));
		goto ret;

	case OADDR:;
		Node *nvar, *nstar;

		// turn &amp;Point(1, 2) or &amp;[]int(1, 2) or &amp;[...]int(1, 2) into allocation.
		// initialize with
		//	nvar := new(*Point);
		//	*nvar = Point(1, 2);
		// and replace expression with nvar
		switch(n-&gt;left-&gt;op) {
		case OARRAYLIT:
		case OMAPLIT:
		case OSTRUCTLIT:
			nvar = makenewvar(n-&gt;type, init, &amp;nstar);
			anylit(n-&gt;left, nstar, init);
			n = nvar;
			goto ret;
		}

		walkexpr(&amp;n-&gt;left, init);
		goto ret;

	case ONEW:
		n = callnew(n-&gt;type-&gt;type);
		goto ret;

	case OCMPSTR:
		// sys_cmpstring(s1, s2) :: 0
		r = mkcall(&#34;cmpstring&#34;, types[TINT], init,
			conv(n-&gt;left, types[TSTRING]),
			conv(n-&gt;right, types[TSTRING]));
		r = nod(n-&gt;etype, r, nodintconst(0));
		typecheck(&amp;r, Erv);
		n = r;
		goto ret;

	case OADDSTR:
		// sys_catstring(s1, s2)
		n = mkcall(&#34;catstring&#34;, n-&gt;type, init,
			conv(n-&gt;left, types[TSTRING]),
			conv(n-&gt;right, types[TSTRING]));
		goto ret;

	case OAPPENDSTR:
		// s1 = sys_catstring(s1, s2)
		if(n-&gt;etype != OADD)
			fatal(&#34;walkasopstring: not add&#34;);
		r = mkcall(&#34;catstring&#34;, n-&gt;left-&gt;type, init,
			conv(n-&gt;left, types[TSTRING]),
			conv(n-&gt;right, types[TSTRING]));
		r = nod(OAS, n-&gt;left, r);
		n = r;
		goto ret;

	case OSLICESTR:
		// sys_slicestring(s, lb, hb)
		n = mkcall(&#34;slicestring&#34;, n-&gt;type, init,
			conv(n-&gt;left, types[TSTRING]),
			conv(n-&gt;right-&gt;left, types[TINT]),
			conv(n-&gt;right-&gt;right, types[TINT]));
		goto ret;

	case OINDEXSTR:
		// TODO(rsc): should be done in back end
		// sys_indexstring(s, i)
		n = mkcall(&#34;indexstring&#34;, n-&gt;type, init,
			conv(n-&gt;left, types[TSTRING]),
			conv(n-&gt;right, types[TINT]));
		goto ret;

	case OCLOSE:
		// cannot use chanfn - closechan takes any, not chan any
		fn = syslook(&#34;closechan&#34;, 1);
		argtype(fn, n-&gt;left-&gt;type);
		n = mkcall1(fn, T, init, n-&gt;left);
		goto ret;

	case OCLOSED:
		// cannot use chanfn - closechan takes any, not chan any
		fn = syslook(&#34;closedchan&#34;, 1);
		argtype(fn, n-&gt;left-&gt;type);
		n = mkcall1(fn, n-&gt;type, init, n-&gt;left);
		goto ret;

	case OMAKECHAN:
		n = mkcall1(chanfn(&#34;makechan&#34;, 1, n-&gt;type), n-&gt;type, init,
			typename(n-&gt;type-&gt;type),
			conv(n-&gt;left, types[TINT]));
		goto ret;

	case OMAKEMAP:
		t = n-&gt;type;

		fn = syslook(&#34;makemap&#34;, 1);
		argtype(fn, t-&gt;down);	// any-1
		argtype(fn, t-&gt;type);	// any-2

		n = mkcall1(fn, n-&gt;type, init,
			typename(t-&gt;down),	// key type
			typename(t-&gt;type),		// value type
			conv(n-&gt;left, types[TINT]));
		goto ret;

	case OMAKESLICE:
		// makeslice(nel int, max int, width int) (ary []any)
		t = n-&gt;type;
		fn = syslook(&#34;makeslice&#34;, 1);
		argtype(fn, t-&gt;type);			// any-1
		n = mkcall1(fn, n-&gt;type, nil,
			conv(n-&gt;left, types[TINT]),
			conv(n-&gt;right, types[TINT]),
			nodintconst(t-&gt;type-&gt;width));
		goto ret;

	case ORUNESTR:
		// sys_intstring(v)
		n = mkcall(&#34;intstring&#34;, n-&gt;type, init, conv(n-&gt;left, types[TINT64]));	// TODO(rsc): int64?!
		goto ret;

	case OARRAYBYTESTR:
		// slicebytetostring([]byte) string;
		n = mkcall(&#34;slicebytetostring&#34;, n-&gt;type, init, n-&gt;left);
		goto ret;

	case OARRAYRUNESTR:
		// sliceinttostring([]byte) string;
		n = mkcall(&#34;sliceinttostring&#34;, n-&gt;type, init, n-&gt;left);
		goto ret;

	case OCMPIFACE:
		// ifaceeq(i1 any-1, i2 any-2) (ret bool);
		if(!eqtype(n-&gt;left-&gt;type, n-&gt;right-&gt;type))
			fatal(&#34;ifaceeq %O %T %T&#34;, n-&gt;op, n-&gt;left-&gt;type, n-&gt;right-&gt;type);
		if(isnilinter(n-&gt;left-&gt;type))
			fn = syslook(&#34;efaceeq&#34;, 1);
		else
			fn = syslook(&#34;ifaceeq&#34;, 1);
		argtype(fn, n-&gt;right-&gt;type);
		argtype(fn, n-&gt;left-&gt;type);
		r = mkcall1(fn, n-&gt;type, init, n-&gt;left, n-&gt;right);
		if(n-&gt;etype == ONE) {
			r = nod(ONOT, r, N);
			typecheck(&amp;r, Erv);
		}
		n = r;
		goto ret;

	case OARRAYLIT:
	case OMAPLIT:
	case OSTRUCTLIT:
		nvar = nod(OXXX, N, N);
		tempname(nvar, n-&gt;type);
		anylit(n, nvar, init);
		n = nvar;
		goto ret;

	case OSEND:
		n = mkcall1(chanfn(&#34;chansend1&#34;, 2, n-&gt;left-&gt;type), T, init, n-&gt;left, n-&gt;right);
		goto ret;

	case OSENDNB:
		n = mkcall1(chanfn(&#34;chansend2&#34;, 2, n-&gt;left-&gt;type), n-&gt;type, init, n-&gt;left, n-&gt;right);
		goto ret;

	case OCONVIFACE:
		walkexpr(&amp;n-&gt;left, init);
		n = ifacecvt(n-&gt;type, n-&gt;left, n-&gt;etype, init);
		goto ret;

	case OCONVSLICE:
		// arraytoslice(old *any, nel int) (ary []any)
		fn = syslook(&#34;arraytoslice&#34;, 1);
		argtype(fn, n-&gt;left-&gt;type-&gt;type);		// any-1
		argtype(fn, n-&gt;type-&gt;type);			// any-2
		n = mkcall1(fn, n-&gt;type, init, n-&gt;left, nodintconst(n-&gt;left-&gt;type-&gt;type-&gt;bound));
		goto ret;

	case OCLOSURE:
		n = walkclosure(n, init);
		goto ret;
	}
	fatal(&#34;missing switch %O&#34;, n-&gt;op);

ret:
	if(debug[&#39;w&#39;] &amp;&amp; n != N)
		dump(&#34;walk&#34;, n);

	ullmancalc(n);
	lineno = lno;
	*np = n;
}

Node*
makenewvar(Type *t, NodeList **init, Node **nstar)
{
	Node *nvar, *nas;

	nvar = nod(OXXX, N, N);
	tempname(nvar, t);
	nas = nod(OAS, nvar, callnew(t-&gt;type));
	typecheck(&amp;nas, Etop);
	walkexpr(&amp;nas, init);
	*init = list(*init, nas);

	*nstar = nod(OIND, nvar, N);
	typecheck(nstar, Erv);
	return nvar;
}

// TODO(rsc): cut
void
walkdottype(Node *n, NodeList **init)
{
	walkexpr(&amp;n-&gt;left, init);
	if(n-&gt;left == N)
		return;
	if(n-&gt;right != N) {
		walkexpr(&amp;n-&gt;right, init);
		n-&gt;type = n-&gt;right-&gt;type;
		n-&gt;right = N;
	}
}

// TODO(rsc): cut
void
walkconv(Node **np, NodeList **init)
{
	int et;
	char *what;
	Type *t;
	Node *l;
	Node *n;

	n = *np;
	t = n-&gt;type;
	if(t == T)
		return;
	walkexpr(&amp;n-&gt;left, init);
	l = n-&gt;left;
	if(l == N)
		return;
	if(l-&gt;type == T)
		return;

	// if using .(T), interface assertion.
	if(n-&gt;op == ODOTTYPE) {
		et = ifaceas1(t, l-&gt;type, 1);
		if(et == I2Isame || et == E2Esame) {
			n-&gt;op = OCONVNOP;
			return;
		}
		if(et != Inone) {
			n = ifacecvt(t, l, et, init);
			*np = n;
			return;
		}
		goto bad;
	}

	fatal(&#34;walkconv&#34;);

bad:
	if(n-&gt;diag)
		return;
	n-&gt;diag = 1;
	if(n-&gt;op == ODOTTYPE)
		what = &#34;type assertion&#34;;
	else
		what = &#34;conversion&#34;;
	if(l-&gt;type != T)
		yyerror(&#34;invalid %s: %T to %T&#34;, what, l-&gt;type, t);
}

Node*
ascompatee1(int op, Node *l, Node *r, NodeList **init)
{
	return convas(nod(OAS, l, r), init);
}

NodeList*
ascompatee(int op, NodeList *nl, NodeList *nr, NodeList **init)
{
	NodeList *ll, *lr, *nn;

	/*
	 * check assign expression list to
	 * a expression list. called in
	 *	expr-list = expr-list
	 */
	nn = nil;
	for(ll=nl, lr=nr; ll &amp;&amp; lr; ll=ll-&gt;next, lr=lr-&gt;next)
		nn = list(nn, ascompatee1(op, ll-&gt;n, lr-&gt;n, init));

	// cannot happen: caller checked that lists had same length
	if(ll || lr)
		yyerror(&#34;error in shape across %O&#34;, op);
	return nn;
}

/*
 * l is an lv and rt is the type of an rv
 * return 1 if this implies a function call
 * evaluating the lv or a function call
 * in the conversion of the types
 */
int
fncall(Node *l, Type *rt)
{
	if(l-&gt;ullman &gt;= UINF)
		return 1;
	if(eqtype(l-&gt;type, rt))
		return 0;
	return 1;
}

NodeList*
ascompatet(int op, NodeList *nl, Type **nr, int fp, NodeList **init)
{
	Node *l, *tmp, *a;
	NodeList *ll;
	Type *r;
	Iter saver;
	int ucount;
	NodeList *nn, *mm;

	/*
	 * check assign type list to
	 * a expression list. called in
	 *	expr-list = func()
	 */
	r = structfirst(&amp;saver, nr);
	nn = nil;
	mm = nil;
	ucount = 0;
	for(ll=nl; ll; ll=ll-&gt;next) {
		if(r == T)
			break;
		l = ll-&gt;n;
		if(isblank(l)) {
			r = structnext(&amp;saver);
			continue;
		}

		// any lv that causes a fn call must be
		// deferred until all the return arguments
		// have been pulled from the output arguments
		if(fncall(l, r-&gt;type)) {
			tmp = nod(OXXX, N, N);
			tempname(tmp, r-&gt;type);
			a = nod(OAS, l, tmp);
			a = convas(a, init);
			mm = list(mm, a);
			l = tmp;
		}

		a = nod(OAS, l, nodarg(r, fp));
		a = convas(a, init);
		ullmancalc(a);
		if(a-&gt;ullman &gt;= UINF)
			ucount++;
		nn = list(nn, a);
		r = structnext(&amp;saver);
	}

	if(ll != nil || r != T)
		yyerror(&#34;assignment count mismatch: %d = %d&#34;,
			count(nl), structcount(*nr));
	if(ucount)
		yyerror(&#34;reorder2: too many function calls evaluating parameters&#34;);
	return concat(nn, mm);
}

/*
 * make a tsig for the structure
 * carrying the ... arguments
 */
Type*
sigtype(Type *st)
{
	Sym *s;
	Type *t;
	static int sigdddgen;

	dowidth(st);

	sigdddgen++;
	snprint(namebuf, sizeof(namebuf), &#34;dsigddd_%d&#34;, sigdddgen);
	s = lookup(namebuf);
	t = newtype(s);
	t = dodcltype(t);
	updatetype(t, st);
	t-&gt;local = 1;
	return t;
}

/*
 * package all the arguments that
 * match a ... parameter into an
 * automatic structure.
 * then call the ... arg (interface)
 * with a pointer to the structure.
 */
NodeList*
mkdotargs(NodeList *lr0, NodeList *nn, Type *l, int fp, NodeList **init)
{
	Node *r;
	Type *t, *st, *ft;
	Node *a, *var;
	NodeList *lr, *n;

	n = nil;			// list of assignments

	st = typ(TSTRUCT);	// generated structure
	ft = T;			// last field
	for(lr=lr0; lr; lr=lr-&gt;next) {
		r = lr-&gt;n;
		if(r-&gt;op == OLITERAL &amp;&amp; r-&gt;val.ctype == CTNIL) {
			if(r-&gt;type == T || r-&gt;type-&gt;etype == TNIL) {
				yyerror(&#34;inappropriate use of nil in ... argument&#34;);
				return nil;
			}
		}
		defaultlit(&amp;r, T);
		lr-&gt;n = r;
		if(r-&gt;type == T)	// type check failed
			return nil;

		// generate the next structure field
		t = typ(TFIELD);
		t-&gt;type = r-&gt;type;
		if(ft == T)
			st-&gt;type = t;
		else
			ft-&gt;down = t;
		ft = t;

		a = nod(OAS, N, r);
		n = list(n, a);
	}

	// make a named type for the struct
	st = sigtype(st);
	dowidth(st);

	// now we have the size, make the struct
	var = nod(OXXX, N, N);
	tempname(var, st);
	var-&gt;sym = lookup(&#34;.ddd&#34;);

	// assign the fields to the struct.
	// use the init list so that reorder1 doesn&#39;t reorder
	// these assignments after the interface conversion
	// below.
	t = st-&gt;type;
	for(lr=n; lr; lr=lr-&gt;next) {
		r = lr-&gt;n;
		r-&gt;left = nod(OXXX, N, N);
		*r-&gt;left = *var;
		r-&gt;left-&gt;type = r-&gt;right-&gt;type;
		r-&gt;left-&gt;xoffset += t-&gt;width;
		typecheck(&amp;r, Etop);
		walkexpr(&amp;r, init);
		lr-&gt;n = r;
		t = t-&gt;down;
	}
	*init = concat(*init, n);

	// last thing is to put assignment
	// of the structure to the DDD parameter
	a = nod(OAS, nodarg(l, fp), var);
	nn = list(nn, convas(a, init));
	return nn;
}

/*
 * helpers for shape errors
 */
static void
dumptypes(Type **nl, char *what)
{
	int first;
	Type *l;
	Iter savel;

	l = structfirst(&amp;savel, nl);
	print(&#34;\t&#34;);
	first = 1;
	for(l = structfirst(&amp;savel, nl); l != T; l = structnext(&amp;savel)) {
		if(first)
			first = 0;
		else
			print(&#34;, &#34;);
		print(&#34;%T&#34;, l);
	}
	if(first)
		print(&#34;[no arguments %s]&#34;, what);
	print(&#34;\n&#34;);
}

static void
dumpnodetypes(NodeList *l, char *what)
{
	int first;
	Node *r;

	print(&#34;\t&#34;);
	first = 1;
	for(; l; l=l-&gt;next) {
		r = l-&gt;n;
		if(first)
			first = 0;
		else
			print(&#34;, &#34;);
		print(&#34;%T&#34;, r-&gt;type);
	}
	if(first)
		print(&#34;[no arguments %s]&#34;, what);
	print(&#34;\n&#34;);
}

/*
 * check assign expression list to
 * a type list. called in
 *	return expr-list
 *	func(expr-list)
 */
NodeList*
ascompatte(int op, Type **nl, NodeList *lr, int fp, NodeList **init)
{
	Type *l, *ll;
	Node *r, *a;
	NodeList *nn, *lr0, *alist;
	Iter savel, peekl;

	lr0 = lr;
	l = structfirst(&amp;savel, nl);
	r = N;
	if(lr)
		r = lr-&gt;n;
	nn = nil;

	// 1 to many
	peekl = savel;
	if(l != T &amp;&amp; r != N &amp;&amp; structnext(&amp;peekl) != T &amp;&amp; lr-&gt;next == nil
	&amp;&amp; r-&gt;type-&gt;etype == TSTRUCT &amp;&amp; r-&gt;type-&gt;funarg) {
		// optimization - can do block copy
		if(eqtypenoname(r-&gt;type, *nl)) {
			a = nodarg(*nl, fp);
			a-&gt;type = r-&gt;type;
			nn = list1(convas(nod(OAS, a, r), init));
			goto ret;
		}
		// conversions involved.
		// copy into temporaries.
		alist = nil;
		for(l=structfirst(&amp;savel, &amp;r-&gt;type); l; l=structnext(&amp;savel)) {
			a = nod(OXXX, N, N);
			tempname(a, l-&gt;type);
			alist = list(alist, a);
		}
		a = nod(OAS2, N, N);
		a-&gt;list = alist;
		a-&gt;rlist = lr;
		typecheck(&amp;a, Etop);
		walkstmt(&amp;a);
		*init = list(*init, a);
		lr = alist;
		r = lr-&gt;n;
		l = structfirst(&amp;savel, nl);
	}

loop:
	if(l != T &amp;&amp; isddd(l-&gt;type)) {
		// the ddd parameter must be last
		ll = structnext(&amp;savel);
		if(ll != T)
			yyerror(&#34;... must be last argument&#34;);

		// special case --
		// only if we are assigning a single ddd
		// argument to a ddd parameter then it is
		// passed thru unencapsulated
		if(r != N &amp;&amp; lr-&gt;next == nil &amp;&amp; isddd(r-&gt;type)) {
			a = nod(OAS, nodarg(l, fp), r);
			a = convas(a, init);
			nn = list(nn, a);
			goto ret;
		}

		// normal case -- make a structure of all
		// remaining arguments and pass a pointer to
		// it to the ddd parameter (empty interface)
		nn = mkdotargs(lr, nn, l, fp, init);
		goto ret;
	}

	if(l == T || r == N) {
		if(l != T || r != N) {
			if(l != T)
				yyerror(&#34;xxx not enough arguments to %O&#34;, op);
			else
				yyerror(&#34;xxx too many arguments to %O&#34;, op);
			dumptypes(nl, &#34;expected&#34;);
			dumpnodetypes(lr0, &#34;given&#34;);
		}
		goto ret;
	}

	a = nod(OAS, nodarg(l, fp), r);
	a = convas(a, init);
	nn = list(nn, a);

	l = structnext(&amp;savel);
	r = N;
	lr = lr-&gt;next;
	if(lr != nil)
		r = lr-&gt;n;
	goto loop;

ret:
	for(lr=nn; lr; lr=lr-&gt;next)
		lr-&gt;n-&gt;typecheck = 1;
	return nn;
}

// generate code for print
static Node*
walkprint(Node *nn, NodeList **init)
{
	Node *r;
	Node *n;
	NodeList *l, *all;
	Node *on;
	Type *t;
	int notfirst, et, op;
	NodeList *calls;

	op = nn-&gt;op;
	all = nn-&gt;list;
	calls = nil;
	notfirst = 0;

	for(l=all; l; l=l-&gt;next) {
		if(notfirst)
			calls = list(calls, mkcall(&#34;printsp&#34;, T, init));
		notfirst = op == OPRINTN || op == OPANICN;

		n = l-&gt;n;
		if(n-&gt;op == OLITERAL) {
			switch(n-&gt;val.ctype) {
			case CTINT:
				defaultlit(&amp;n, types[TINT64]);
				break;
			case CTFLT:
				defaultlit(&amp;n, types[TFLOAT64]);
				break;
			}
		}
		if(n-&gt;op != OLITERAL &amp;&amp; n-&gt;type &amp;&amp; n-&gt;type-&gt;etype == TIDEAL)
			defaultlit(&amp;n, types[TINT64]);
		defaultlit(&amp;n, nil);
		l-&gt;n = n;
		if(n-&gt;type == T || n-&gt;type-&gt;etype == TFORW)
			continue;

		et = n-&gt;type-&gt;etype;
		if(isinter(n-&gt;type)) {
			if(isnilinter(n-&gt;type))
				on = syslook(&#34;printeface&#34;, 1);
			else
				on = syslook(&#34;printiface&#34;, 1);
			argtype(on, n-&gt;type);		// any-1
		} else if(isptr[et] || et == TCHAN || et == TMAP || et == TFUNC) {
			on = syslook(&#34;printpointer&#34;, 1);
			argtype(on, n-&gt;type);	// any-1
		} else if(isslice(n-&gt;type)) {
			on = syslook(&#34;printslice&#34;, 1);
			argtype(on, n-&gt;type);	// any-1
		} else if(isint[et]) {
			if(et == TUINT64)
				on = syslook(&#34;printuint&#34;, 0);
			else
				on = syslook(&#34;printint&#34;, 0);
		} else if(isfloat[et]) {
			on = syslook(&#34;printfloat&#34;, 0);
		} else if(et == TBOOL) {
			on = syslook(&#34;printbool&#34;, 0);
		} else if(et == TSTRING) {
			on = syslook(&#34;printstring&#34;, 0);
		} else {
			badtype(OPRINT, n-&gt;type, T);
			continue;
		}

		t = *getinarg(on-&gt;type);
		if(t != nil)
			t = t-&gt;type;
		if(t != nil)
			t = t-&gt;type;

		if(!eqtype(t, n-&gt;type)) {
			n = nod(OCONV, n, N);
			n-&gt;type = t;
		}
		r = nod(OCALL, on, N);
		r-&gt;list = list1(n);
		calls = list(calls, r);
	}

	if(op == OPRINTN)
		calls = list(calls, mkcall(&#34;printnl&#34;, T, nil));
	typechecklist(calls, Etop);
	walkexprlist(calls, init);

	if(op == OPANIC || op == OPANICN)
		r = mkcall(&#34;panicl&#34;, T, nil);
	else
		r = nod(OEMPTY, N, N);
	typecheck(&amp;r, Etop);
	walkexpr(&amp;r, init);
	r-&gt;ninit = calls;
	return r;
}

Node*
callnew(Type *t)
{
	Node *fn;

	dowidth(t);
	fn = syslook(&#34;mal&#34;, 1);
	argtype(fn, t);
	return mkcall1(fn, ptrto(t), nil, nodintconst(t-&gt;width));
}

Type*
fixchan(Type *t)
{
	if(t == T)
		goto bad;
	if(t-&gt;etype != TCHAN)
		goto bad;
	if(t-&gt;type == T)
		goto bad;

	dowidth(t-&gt;type);

	return t;

bad:
	yyerror(&#34;not a channel: %lT&#34;, t);
	return T;
}

Node*
mapop(Node *n, NodeList **init)
{
	Node *r, *a;

	r = n;
	switch(n-&gt;op) {
	default:
		fatal(&#34;mapop: unknown op %O&#34;, n-&gt;op);
	case OASOP:
		// rewrite map[index] op= right
		// into tmpi := index; map[tmpi] = map[tmpi] op right

		// make it ok to double-evaluate map[tmpi]
		n-&gt;left-&gt;left = safeval(n-&gt;left-&gt;left, init);
		n-&gt;left-&gt;right = safeval(n-&gt;left-&gt;right, init);

		a = nod(OXXX, N, N);
		*a = *n-&gt;left;		// copy of map[tmpi]
		a-&gt;etype = 0;
		a = nod(n-&gt;etype, a, n-&gt;right);		// m[tmpi] op right
		r = nod(OAS, n-&gt;left, a);		// map[tmpi] = map[tmpi] op right
		typecheck(&amp;r, Etop);
		walkexpr(&amp;r, init);
		break;
	}
	return r;
}

/*
 * assigning src to dst involving interfaces?
 * return op to use.
 */
int
ifaceas1(Type *dst, Type *src, int explicit)
{
	if(src == T || dst == T)
		return Inone;

	if(explicit &amp;&amp; !isinter(src))
		yyerror(&#34;cannot use .(T) on non-interface type %T&#34;, src);

	if(isinter(dst)) {
		if(isinter(src)) {
			if(isnilinter(dst)) {
				if(isnilinter(src))
					return E2Esame;
				return I2E;
			}
			if(eqtype(dst, src))
				return I2Isame;
			ifacecheck(dst, src, lineno, explicit);
			if(isnilinter(src))
				return E2I;
			if(explicit)
				return I2Ix;
			return I2I;
		}
		if(isnilinter(dst))
			return T2E;
		ifacecheck(dst, src, lineno, explicit);
		return T2I;
	}
	if(isinter(src)) {
		ifacecheck(dst, src, lineno, explicit);
		if(isnilinter(src))
			return E2T;
		return I2T;
	}
	return Inone;
}

/*
 * treat convert T to T as noop
 */
int
ifaceas(Type *dst, Type *src, int explicit)
{
	int et;

	et = ifaceas1(dst, src, explicit);
	if(et == I2Isame || et == E2Esame)
		et = Inone;
	return et;
}

static	char*
ifacename[] =
{
	[I2T]		= &#34;ifaceI2T&#34;,
	[I2T2]		= &#34;ifaceI2T2&#34;,
	[I2I]		= &#34;ifaceI2I&#34;,
	[I2Ix]		= &#34;ifaceI2Ix&#34;,
	[I2I2]		= &#34;ifaceI2I2&#34;,
	[I2Isame]	= &#34;ifaceI2Isame&#34;,
	[E2T]		= &#34;ifaceE2T&#34;,
	[E2T2]		= &#34;ifaceE2T2&#34;,
	[E2I]		= &#34;ifaceE2I&#34;,
	[E2I2]		= &#34;ifaceE2I2&#34;,
	[I2E]		= &#34;ifaceI2E&#34;,
	[I2E2]		= &#34;ifaceI2E2&#34;,
	[T2I]		= &#34;ifaceT2I&#34;,
	[T2E]		= &#34;ifaceT2E&#34;,
	[E2Esame]	= &#34;ifaceE2Esame&#34;,
};

Node*
ifacecvt(Type *tl, Node *n, int et, NodeList **init)
{
	Type *tr;
	Node *r, *on;
	NodeList *args;

	tr = n-&gt;type;

	switch(et) {
	default:
		fatal(&#34;ifacecvt: unknown op %d\n&#34;, et);

	case I2Isame:
	case E2Esame:
		return n;

	case T2I:
		// ifaceT2I(sigi *byte, sigt *byte, elem any) (ret any);
		args = list1(typename(tl));	// sigi
		args = list(args, typename(tr));	// sigt
		args = list(args, n);	// elem

		on = syslook(&#34;ifaceT2I&#34;, 1);
		argtype(on, tr);
		argtype(on, tl);
		dowidth(on-&gt;type);
		break;

	case I2T:
	case I2T2:
	case I2I:
	case I2Ix:
	case I2I2:
	case E2T:
	case E2T2:
	case E2I:
	case E2I2:
		// iface[IT]2[IT][2](sigt *byte, iface any) (ret any[, ok bool]);
		args = list1(typename(tl));	// sigi or sigt
		args = list(args, n);		// iface

		on = syslook(ifacename[et], 1);
		argtype(on, tr);
		argtype(on, tl);
		break;

	case I2E:
		// TODO(rsc): Should do this in back end, without a call.
		// ifaceI2E(elem any) (ret any);
		args = list1(n);	// elem

		on = syslook(&#34;ifaceI2E&#34;, 1);
		argtype(on, tr);
		argtype(on, tl);
		break;

	case T2E:
		// TODO(rsc): Should do this in back end for pointer case, without a call.
		// ifaceT2E(sigt *byte, elem any) (ret any);
		args = list1(typename(tr));	// sigt
		args = list(args, n);		// elem

		on = syslook(&#34;ifaceT2E&#34;, 1);
		argtype(on, tr);
		argtype(on, tl);
		break;
	}

	dowidth(on-&gt;type);
	r = nod(OCALL, on, N);
	r-&gt;list = args;
	typecheck(&amp;r, Erv | Efnstruct);
	walkexpr(&amp;r, init);
	return r;
}

Node*
convas(Node *n, NodeList **init)
{
	Node *l, *r;
	Type *lt, *rt;
	int et;

	if(n-&gt;op != OAS)
		fatal(&#34;convas: not OAS %O&#34;, n-&gt;op);
	n-&gt;typecheck = 1;

	lt = T;
	rt = T;

	l = n-&gt;left;
	r = n-&gt;right;
	if(l == N || r == N)
		goto out;

	lt = l-&gt;type;
	rt = r-&gt;type;
	if(lt == T || rt == T)
		goto out;

	if(isblank(n-&gt;left))
		goto out;

	if(n-&gt;left-&gt;op == OINDEXMAP) {
		n = mkcall1(mapfn(&#34;mapassign1&#34;, n-&gt;left-&gt;left-&gt;type), T, init,
			n-&gt;left-&gt;left, n-&gt;left-&gt;right, n-&gt;right);
		goto out;
	}

	if(eqtype(lt, rt))
		goto out;

	et = ifaceas(lt, rt, 0);
	if(et != Inone) {
		n-&gt;right = ifacecvt(lt, r, et, init);
		goto out;
	}

out:
	ullmancalc(n);
	return n;
}

/*
 * from ascompat[te]
 * evaluating actual function arguments.
 *	f(a,b)
 * if there is exactly one function expr,
 * then it is done first. otherwise must
 * make temp variables
 */
NodeList*
reorder1(NodeList *all)
{
	Node *f, *a, *n;
	NodeList *l, *r, *g;
	int c, t;

	c = 0;	// function calls
	t = 0;	// total parameters

	for(l=all; l; l=l-&gt;next) {
		n = l-&gt;n;
		t++;
		ullmancalc(n);
		if(n-&gt;ullman &gt;= UINF)
			c++;
	}
	if(c == 0 || t == 1)
		return all;

	g = nil;	// fncalls assigned to tempnames
	f = N;	// one fncall assigned to stack
	r = nil;	// non fncalls and tempnames assigned to stack

	for(l=all; l; l=l-&gt;next) {
		n = l-&gt;n;
		ullmancalc(n);
		if(n-&gt;ullman &lt; UINF) {
			r = list(r, n);
			continue;
		}
		if(f == N) {
			f = n;
			continue;
		}

		// make assignment of fncall to tempname
		a = nod(OXXX, N, N);
		tempname(a, n-&gt;right-&gt;type);
		a = nod(OAS, a, n-&gt;right);
		g = list(g, a);

		// put normal arg assignment on list
		// with fncall replaced by tempname
		n-&gt;right = a-&gt;left;
		r = list(r, n);
	}

	if(f != N)
		g = list(g, f);
	return concat(g, r);
}

/*
 * from ascompat[ee]
 *	a,b = c,d
 * simultaneous assignment. there cannot
 * be later use of an earlier lvalue.
 */

int
vmatch2(Node *l, Node *r)
{
	NodeList *ll;

	/*
	 * isolate all right sides
	 */
	if(r == N)
		return 0;
	switch(r-&gt;op) {
	case ONAME:
		// match each right given left
		if(l == r)
			return 1;
	case OLITERAL:
		return 0;
	}
	if(vmatch2(l, r-&gt;left))
		return 1;
	if(vmatch2(l, r-&gt;right))
		return 1;
	for(ll=r-&gt;list; ll; ll=ll-&gt;next)
		if(vmatch2(l, ll-&gt;n))
			return 1;
	return 0;
}

int
vmatch1(Node *l, Node *r)
{
	NodeList *ll;

	/*
	 * isolate all left sides
	 */
	if(l == N)
		return 0;
	switch(l-&gt;op) {
	case ONAME:
		switch(l-&gt;class) {
		case PPARAM:
		case PPARAMREF:
		case PAUTO:
			break;
		default:
			// assignment to non-stack variable
			// must be delayed if right has function calls.
			if(r-&gt;ullman &gt;= UINF)
				return 1;
			break;
		}
		return vmatch2(l, r);
	case OLITERAL:
		return 0;
	}
	if(vmatch1(l-&gt;left, r))
		return 1;
	if(vmatch1(l-&gt;right, r))
		return 1;
	for(ll=l-&gt;list; ll; ll=ll-&gt;next)
		if(vmatch1(ll-&gt;n, r))
			return 1;
	return 0;
}

NodeList*
reorder3(NodeList *all)
{
	Node *n1, *n2, *q;
	int c1, c2;
	NodeList *l1, *l2, *r;

	r = nil;
	for(l1=all, c1=0; l1; l1=l1-&gt;next, c1++) {
		n1 = l1-&gt;n;
		for(l2=all, c2=0; l2; l2=l2-&gt;next, c2++) {
			n2 = l2-&gt;n;
			if(c2 &gt; c1) {
				if(vmatch1(n1-&gt;left, n2-&gt;right)) {
					// delay assignment to n1-&gt;left
					q = nod(OXXX, N, N);
					tempname(q, n1-&gt;right-&gt;type);
					q = nod(OAS, n1-&gt;left, q);
					n1-&gt;left = q-&gt;right;
					r = list(r, q);
					break;
				}
			}
		}
	}
	return concat(all, r);
}

NodeList*
reorder4(NodeList *ll)
{
	/*
	 * from ascompat[te]
	 *	return c,d
	 * return expression assigned to output
	 * parameters. there may be no problems.
	 *
	 * TODO(rsc): i don&#39;t believe that.
	 *	func f() (a, b int) {
	 *		a, b = 1, 2;
	 *		return b, a;
	 *	}
	 */
	return ll;
}

/*
 * walk through argin parameters.
 * generate and return code to allocate
 * copies of escaped parameters to the heap.
 */
NodeList*
paramstoheap(Type **argin)
{
	Type *t;
	Iter savet;
	Node *v;
	NodeList *nn;

	nn = nil;
	for(t = structfirst(&amp;savet, argin); t != T; t = structnext(&amp;savet)) {
		v = t-&gt;nname;
		if(v == N || !(v-&gt;class &amp; PHEAP))
			continue;

		// generate allocation &amp; copying code
		nn = list(nn, nod(OAS, v-&gt;heapaddr, v-&gt;alloc));
		nn = list(nn, nod(OAS, v, v-&gt;stackparam));
	}
	return nn;
}

/*
 * take care of migrating any function in/out args
 * between the stack and the heap.  adds code to
 * curfn&#39;s before and after lists.
 */
void
heapmoves(void)
{
	NodeList *nn;

	nn = paramstoheap(getthis(curfn-&gt;type));
	nn = concat(nn, paramstoheap(getinarg(curfn-&gt;type)));
	curfn-&gt;enter = concat(curfn-&gt;enter, nn);
}

static Node*
vmkcall(Node *fn, Type *t, NodeList **init, va_list va)
{
	int i, n;
	Node *r;
	NodeList *args;

	if(fn-&gt;type == T || fn-&gt;type-&gt;etype != TFUNC)
		fatal(&#34;mkcall %#N %T&#34;, fn, fn-&gt;type);

	args = nil;
	n = fn-&gt;type-&gt;intuple;
	for(i=0; i&lt;n; i++)
		args = list(args, va_arg(va, Node*));

	r = nod(OCALL, fn, N);
	r-&gt;list = args;
	if(fn-&gt;type-&gt;outtuple &gt; 0)
		typecheck(&amp;r, Erv | Efnstruct);
	else
		typecheck(&amp;r, Etop);
	walkexpr(&amp;r, init);
	r-&gt;type = t;
	return r;
}

Node*
mkcall(char *name, Type *t, NodeList **init, ...)
{
	Node *r;
	va_list va;

	va_start(va, init);
	r = vmkcall(syslook(name, 0), t, init, va);
	va_end(va);
	return r;
}

Node*
mkcall1(Node *fn, Type *t, NodeList **init, ...)
{
	Node *r;
	va_list va;

	va_start(va, init);
	r = vmkcall(fn, t, init, va);
	va_end(va);
	return r;
}

static Node*
conv(Node *n, Type *t)
{
	if(eqtype(n-&gt;type, t))
		return n;
	n = nod(OCONV, n, N);
	n-&gt;type = t;
	typecheck(&amp;n, Erv);
	return n;
}

Node*
chanfn(char *name, int n, Type *t)
{
	Node *fn;
	int i;

	if(t-&gt;etype != TCHAN)
		fatal(&#34;chanfn %T&#34;, t);
	fn = syslook(name, 1);
	for(i=0; i&lt;n; i++)
		argtype(fn, t-&gt;type);
	return fn;
}

static Node*
mapfn(char *name, Type *t)
{
	Node *fn;

	if(t-&gt;etype != TMAP)
		fatal(&#34;mapfn %T&#34;, t);
	fn = syslook(name, 1);
	argtype(fn, t-&gt;down);
	argtype(fn, t-&gt;type);
	argtype(fn, t-&gt;down);
	argtype(fn, t-&gt;type);
	return fn;
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
