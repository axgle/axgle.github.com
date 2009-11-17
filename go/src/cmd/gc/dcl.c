<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/dcl.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/dcl.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;
#include	&#34;y.tab.h&#34;

static	void	funcargs(Node*);

int
dflag(void)
{
	if(!debug[&#39;d&#39;])
		return 0;
	if(debug[&#39;y&#39;])
		return 1;
	if(incannedimport)
		return 0;
	return 1;
}

/*
 * declaration stack &amp; operations
 */
static	Sym*	dclstack;

void
dcopy(Sym *a, Sym *b)
{
	a-&gt;package = b-&gt;package;
	a-&gt;name = b-&gt;name;
	a-&gt;def = b-&gt;def;
	a-&gt;block = b-&gt;block;
	a-&gt;lastlineno = b-&gt;lastlineno;
}

Sym*
push(void)
{
	Sym *d;

	d = mal(sizeof(*d));
	d-&gt;link = dclstack;
	dclstack = d;
	return d;
}

Sym*
pushdcl(Sym *s)
{
	Sym *d;

	d = push();
	dcopy(d, s);
	if(dflag())
		print(&#34;\t%L push %S %p\n&#34;, lineno, s, s-&gt;def);
	return d;
}

void
popdcl(void)
{
	Sym *d, *s;

//	if(dflag())
//		print(&#34;revert\n&#34;);

	for(d=dclstack; d!=S; d=d-&gt;link) {
		if(d-&gt;name == nil)
			break;
		s = pkglookup(d-&gt;name, d-&gt;package);
		dcopy(s, d);
		if(dflag())
			print(&#34;\t%L pop %S %p\n&#34;, lineno, s, s-&gt;def);
	}
	if(d == S)
		fatal(&#34;popdcl: no mark&#34;);
	dclstack = d-&gt;link;
	block = d-&gt;block;
}

void
poptodcl(void)
{
	Sym *d, *s;

	for(d=dclstack; d!=S; d=d-&gt;link) {
		if(d-&gt;name == nil)
			break;
		s = pkglookup(d-&gt;name, d-&gt;package);
		dcopy(s, d);
		if(dflag())
			print(&#34;\t%L pop %S\n&#34;, lineno, s);
	}
	if(d == S)
		fatal(&#34;poptodcl: no mark&#34;);
	dclstack = d;
}

void
markdcl(void)
{
	Sym *d;

	d = push();
	d-&gt;name = nil;		// used as a mark in fifo
	d-&gt;block = block;

	blockgen++;
	block = blockgen;

//	if(dflag())
//		print(&#34;markdcl\n&#34;);
}

void
dumpdcl(char *st)
{
	Sym *s, *d;
	int i;

	i = 0;
	for(d=dclstack; d!=S; d=d-&gt;link) {
		i++;
		print(&#34;    %.2d %p&#34;, i, d);
		if(d-&gt;name == nil) {
			print(&#34;\n&#34;);
			continue;
		}
		print(&#34; &#39;%s&#39;&#34;, d-&gt;name);
		s = pkglookup(d-&gt;name, d-&gt;package);
		print(&#34; %lS\n&#34;, s);
	}
}

void
testdclstack(void)
{
	Sym *d;

	for(d=dclstack; d!=S; d=d-&gt;link) {
		if(d-&gt;name == nil) {
			yyerror(&#34;mark left on the stack&#34;);
			continue;
		}
	}
}

void
redeclare(Sym *s, char *where)
{
	yyerror(&#34;%S redeclared %s\n&#34;
		&#34;\tprevious declaration at %L&#34;,
		s, where, s-&gt;lastlineno);
}

/*
 * declare individual names - var, typ, const
 */
void
declare(Node *n, int ctxt)
{
	Sym *s;
	int gen;
	static int typegen, vargen;

	if(isblank(n))
		return;

	s = n-&gt;sym;
	gen = 0;
	if(ctxt == PEXTERN) {
		externdcl = list(externdcl, n);
		if(dflag())
			print(&#34;\t%L global decl %S %p\n&#34;, lineno, s, n);
	} else {
		if(curfn == nil &amp;&amp; ctxt == PAUTO)
			fatal(&#34;automatic outside function&#34;);
		if(curfn != nil)
			curfn-&gt;dcl = list(curfn-&gt;dcl, n);
		if(n-&gt;op == OTYPE)
			gen = ++typegen;
		else if(n-&gt;op == ONAME)
			gen = ++vargen;
		pushdcl(s);
	}
	if(ctxt == PAUTO)
		n-&gt;xoffset = BADWIDTH;

	if(s-&gt;block == block)
		redeclare(s, &#34;in this block&#34;);

	s-&gt;block = block;
	s-&gt;lastlineno = parserline();
	s-&gt;def = n;
	n-&gt;vargen = gen;
	n-&gt;funcdepth = funcdepth;
	n-&gt;class = ctxt;

	autoexport(n, ctxt);
}

void
addvar(Node *n, Type *t, int ctxt)
{
	if(n==N || n-&gt;sym == S || (n-&gt;op != ONAME &amp;&amp; n-&gt;op != ONONAME) || t == T)
		fatal(&#34;addvar: n=%N t=%T nil&#34;, n, t);

	n-&gt;op = ONAME;
	declare(n, ctxt);
	n-&gt;type = t;
}

// TODO: cut use of below in sigtype and then delete
void
addtyp(Type *n, int ctxt)
{
	Node *def;

	if(n==T || n-&gt;sym == S)
		fatal(&#34;addtyp: n=%T t=%T nil&#34;, n);

	def = typenod(n);
	declare(def, ctxt);
	n-&gt;vargen = def-&gt;vargen;

	typelist = list(typelist, def);
}

/*
 * introduce a type named n
 * but it is an unknown type for now
 */
// TODO(rsc): cut use of this in sigtype and then delete
Type*
dodcltype(Type *n)
{
	addtyp(n, dclcontext);
	n-&gt;local = 1;
	autoexport(typenod(n), dclcontext);
	return n;
}

/*
 * now we know what n is: it&#39;s t
 */
// TODO(rsc): cut use of this in sigtype and then delete
void
updatetype(Type *n, Type *t)
{
	Sym *s;
	int local, vargen;
	int maplineno, lno, etype;

	if(t == T)
		return;
	s = n-&gt;sym;
	if(s == S || s-&gt;def == N || s-&gt;def-&gt;op != OTYPE || s-&gt;def-&gt;type != n)
		fatal(&#34;updatetype %T = %T&#34;, n, t);

	etype = n-&gt;etype;
	switch(n-&gt;etype) {
	case TFORW:
		break;

	default:
		fatal(&#34;updatetype %T / %T&#34;, n, t);
	}

	// decl was
	//	type n t;
	// copy t, but then zero out state associated with t
	// that is no longer associated with n.
	maplineno = n-&gt;maplineno;
	local = n-&gt;local;
	vargen = n-&gt;vargen;
	*n = *t;
	n-&gt;sym = s;
	n-&gt;local = local;
	n-&gt;siggen = 0;
	n-&gt;printed = 0;
	n-&gt;method = nil;
	n-&gt;vargen = vargen;
	n-&gt;nod = N;

	checkwidth(n);

	// double-check use of type as map key
	if(maplineno) {
		lno = lineno;
		lineno = maplineno;
		maptype(n, types[TBOOL]);
		lineno = lno;
	}
}

/*
 * declare variables from grammar
 * new_name_list (type | [type] = expr_list)
 */
NodeList*
variter(NodeList *vl, Node *t, NodeList *el)
{
	int doexpr;
	Node *v, *e;
	NodeList *init;

	init = nil;
	doexpr = el != nil;
	for(; vl; vl=vl-&gt;next) {
		if(doexpr) {
			if(el == nil) {
				yyerror(&#34;missing expr in var dcl&#34;);
				break;
			}
			e = el-&gt;n;
			el = el-&gt;next;
		} else
			e = N;

		v = vl-&gt;n;
		v-&gt;op = ONAME;
		declare(v, dclcontext);
		v-&gt;ntype = t;

		if(e != N || funcdepth &gt; 0 || isblank(v)) {
			if(funcdepth &gt; 0)
				init = list(init, nod(ODCL, v, N));
			e = nod(OAS, v, e);
			init = list(init, e);
			if(e-&gt;right != N)
				v-&gt;defn = e;
		}
	}
	if(el != nil)
		yyerror(&#34;extra expr in var dcl&#34;);
	return init;
}

/*
 * declare constants from grammar
 * new_name_list [[type] = expr_list]
 */
NodeList*
constiter(NodeList *vl, Node *t, NodeList *cl)
{
	Node *v, *c;
	NodeList *vv;

	vv = nil;
	if(cl == nil) {
		if(t != N)
			yyerror(&#34;constdcl cannot have type without expr&#34;);
		cl = lastconst;
		t = lasttype;
	} else {
		lastconst = cl;
		lasttype = t;
	}
	cl = listtreecopy(cl);

	for(; vl; vl=vl-&gt;next) {
		if(cl == nil) {
			yyerror(&#34;missing expr in const dcl&#34;);
			break;
		}
		c = cl-&gt;n;
		cl = cl-&gt;next;

		v = vl-&gt;n;
		v-&gt;op = OLITERAL;
		declare(v, dclcontext);

		v-&gt;ntype = t;
		v-&gt;defn = c;

		vv = list(vv, nod(ODCLCONST, v, N));
	}
	if(cl != nil)
		yyerror(&#34;extra expr in const dcl&#34;);
	iota += 1;
	return vv;
}

/*
 * this generates a new name node,
 * typically for labels or other one-off names.
 */
Node*
newname(Sym *s)
{
	Node *n;

	if(s == S)
		fatal(&#34;newname nil&#34;);

	n = nod(ONAME, N, N);
	n-&gt;sym = s;
	n-&gt;type = T;
	n-&gt;addable = 1;
	n-&gt;ullman = 1;
	n-&gt;xoffset = 0;
	return n;
}

/*
 * this generates a new name node for a name
 * being declared.  if at the top level, it might return
 * an ONONAME node created by an earlier reference.
 */
Node*
dclname(Sym *s)
{
	Node *n;

	// top-level name: might already have been
	// referred to, in which case s-&gt;def is already
	// set to an ONONAME.
	if(dclcontext == PEXTERN &amp;&amp; s-&gt;block &lt;= 1) {
		if(s-&gt;def == N)
			oldname(s);
		if(s-&gt;def-&gt;op == ONONAME)
			return s-&gt;def;
	}

	n = newname(s);
	n-&gt;op = ONONAME;	// caller will correct it
	return n;
}

Node*
typenod(Type *t)
{
	// if we copied another type with *t = *u
	// then t-&gt;nod might be out of date, so
	// check t-&gt;nod-&gt;type too
	if(t-&gt;nod == N || t-&gt;nod-&gt;type != t) {
		t-&gt;nod = nod(OTYPE, N, N);
		t-&gt;nod-&gt;type = t;
		t-&gt;nod-&gt;sym = t-&gt;sym;
	}
	return t-&gt;nod;
}


/*
 * this will return an old name
 * that has already been pushed on the
 * declaration list. a diagnostic is
 * generated if no name has been defined.
 */
Node*
oldname(Sym *s)
{
	Node *n;
	Node *c;

	n = s-&gt;def;
	if(n == N) {
		// maybe a top-level name will come along
		// to give this a definition later.
		n = newname(s);
		n-&gt;op = ONONAME;
		s-&gt;def = n;
	}
	if(n-&gt;oldref &lt; 100)
		n-&gt;oldref++;
	if(curfn != nil &amp;&amp; n-&gt;funcdepth &gt; 0 &amp;&amp; n-&gt;funcdepth != funcdepth &amp;&amp; n-&gt;op == ONAME) {
		// inner func is referring to var in outer func.
		//
		// TODO(rsc): If there is an outer variable x and we
		// are parsing x := 5 inside the closure, until we get to
		// the := it looks like a reference to the outer x so we&#39;ll
		// make x a closure variable unnecessarily.
		if(n-&gt;closure == N || n-&gt;closure-&gt;funcdepth != funcdepth) {
			// create new closure var.
			c = nod(ONAME, N, N);
			c-&gt;sym = s;
			c-&gt;class = PPARAMREF;
			c-&gt;defn = n;
			c-&gt;addable = 0;
			c-&gt;ullman = 2;
			c-&gt;funcdepth = funcdepth;
			c-&gt;outer = n-&gt;closure;
			n-&gt;closure = c;
			c-&gt;closure = n;
			c-&gt;xoffset = 0;
			curfn-&gt;cvars = list(curfn-&gt;cvars, c);
		}
		// return ref to closure var, not original
		return n-&gt;closure;
	}
	return n;
}

/*
 * same for types
 */
Type*
newtype(Sym *s)
{
	Type *t;

	t = typ(TFORW);
	t-&gt;sym = s;
	t-&gt;type = T;
	return t;
}

/*
 * type check top level declarations
 */
void
dclchecks(void)
{
	NodeList *l;

	for(l=externdcl; l; l=l-&gt;next) {
		if(l-&gt;n-&gt;op != ONAME)
			continue;
		typecheck(&amp;l-&gt;n, Erv);
	}
}

/*
 * := declarations
 */

static int
colasname(Node *n)
{
	// TODO(rsc): can probably simplify
	// once late-binding of names goes in
	switch(n-&gt;op) {
	case ONAME:
	case ONONAME:
	case OPACK:
	case OTYPE:
	case OLITERAL:
		return n-&gt;sym != S;
	}
	return 0;
}

void
colasdefn(NodeList *left, Node *defn)
{
	int nnew;
	NodeList *l;
	Node *n;

	nnew = 0;
	for(l=left; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(isblank(n))
			continue;
		if(!colasname(n)) {
			yyerror(&#34;non-name %#N on left side of :=&#34;, n);
			continue;
		}
		if(n-&gt;sym-&gt;block == block)
			continue;

		// If we created an ONONAME just for this :=,
		// delete it, to avoid confusion with top-level imports.
		if(n-&gt;op == ONONAME &amp;&amp; n-&gt;oldref &lt; 100 &amp;&amp; --n-&gt;oldref == 0)
			n-&gt;sym-&gt;def = N;

		nnew++;
		n = newname(n-&gt;sym);
		declare(n, dclcontext);
		n-&gt;defn = defn;
		defn-&gt;ninit = list(defn-&gt;ninit, nod(ODCL, n, N));
		l-&gt;n = n;
	}
	if(nnew == 0)
		yyerror(&#34;no new variables on left side of :=&#34;);
}

Node*
colas(NodeList *left, NodeList *right)
{
	Node *as;

	as = nod(OAS2, N, N);
	as-&gt;list = left;
	as-&gt;rlist = right;
	as-&gt;colas = 1;
	colasdefn(left, as);

	// make the tree prettier; not necessary
	if(count(left) == 1 &amp;&amp; count(right) == 1) {
		as-&gt;left = as-&gt;list-&gt;n;
		as-&gt;right = as-&gt;rlist-&gt;n;
		as-&gt;list = nil;
		as-&gt;rlist = nil;
		as-&gt;op = OAS;
	}

	return as;
}

/*
 * declare the function proper
 * and declare the arguments.
 * called in extern-declaration context
 * returns in auto-declaration context.
 */
void
funchdr(Node *n)
{

	if(n-&gt;nname != N) {
		n-&gt;nname-&gt;op = ONAME;
		declare(n-&gt;nname, PFUNC);
		n-&gt;nname-&gt;defn = n;
	}

	// change the declaration context from extern to auto
	if(funcdepth == 0 &amp;&amp; dclcontext != PEXTERN)
		fatal(&#34;funchdr: dclcontext&#34;);

	dclcontext = PAUTO;
	markdcl();
	funcdepth++;

	n-&gt;outer = curfn;
	curfn = n;
	if(n-&gt;nname)
		funcargs(n-&gt;nname-&gt;ntype);
	else
		funcargs(n-&gt;ntype);
}

static void
funcargs(Node *nt)
{
	Node *n;
	NodeList *l;

	if(nt-&gt;op != OTFUNC)
		fatal(&#34;funcargs %O&#34;, nt-&gt;op);

	// declare the receiver and in arguments.
	// no n-&gt;defn because type checking of func header
	// will fill in the types before we can demand them.
	if(nt-&gt;left != N) {
		n = nt-&gt;left;
		if(n-&gt;op != ODCLFIELD)
			fatal(&#34;funcargs1 %O&#34;, n-&gt;op);
		if(n-&gt;left != N) {
			n-&gt;left-&gt;op = ONAME;
			n-&gt;left-&gt;ntype = n-&gt;right;
			declare(n-&gt;left, PPARAM);
		}
	}
	for(l=nt-&gt;list; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op != ODCLFIELD)
			fatal(&#34;funcargs2 %O&#34;, n-&gt;op);
		if(n-&gt;left != N) {
			n-&gt;left-&gt;op = ONAME;
			n-&gt;left-&gt;ntype = n-&gt;right;
			declare(n-&gt;left, PPARAM);
		}
	}

	// declare the out arguments.
	for(l=nt-&gt;rlist; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op != ODCLFIELD)
			fatal(&#34;funcargs3 %O&#34;, n-&gt;op);
		if(n-&gt;left != N) {
			n-&gt;left-&gt;op = ONAME;
			n-&gt;left-&gt;ntype = n-&gt;right;
			declare(n-&gt;left, PPARAMOUT);
		}
	}
}

/*
 * finish the body.
 * called in auto-declaration context.
 * returns in extern-declaration context.
 */
void
funcbody(Node *n)
{
	// change the declaration context from auto to extern
	if(dclcontext != PAUTO)
		fatal(&#34;funcbody: dclcontext&#34;);
	popdcl();
	funcdepth--;
	curfn = n-&gt;outer;
	n-&gt;outer = N;
	if(funcdepth == 0)
		dclcontext = PEXTERN;
}

/*
 * new type being defined with name s.
 */
Node*
typedcl0(Sym *s)
{
	Node *n;

	n = dclname(s);
	n-&gt;op = OTYPE;
	declare(n, dclcontext);
	return n;
}

/*
 * node n, which was returned by typedcl0
 * is being declared to have uncompiled type t.
 * return the ODCLTYPE node to use.
 */
Node*
typedcl1(Node *n, Node *t, int local)
{
	n-&gt;ntype = t;
	n-&gt;local = local;
	return nod(ODCLTYPE, n, N);
}

/*
 * typedcl1 but during imports
 */
void
typedcl2(Type *pt, Type *t)
{
	Node *n;

	if(pt-&gt;etype == TFORW)
		goto ok;
	if(!cvttype(pt, t))
		yyerror(&#34;inconsistent definition for type %S during import\n\t%lT\n\t%lT&#34;, pt-&gt;sym, pt, t);
	return;

ok:
	n = pt-&gt;nod;
	*pt = *t;
	pt-&gt;method = nil;
	pt-&gt;nod = n;
	pt-&gt;sym = n-&gt;sym;
	pt-&gt;sym-&gt;lastlineno = parserline();
	declare(n, PEXTERN);

	checkwidth(pt);
}

/*
 * structs, functions, and methods.
 * they don&#39;t belong here, but where do they belong?
 */


/*
 * turn a parsed struct into a type
 */
Type**
stotype(NodeList *l, int et, Type **t)
{
	Type *f, *t1, *t2, **t0;
	Strlit *note;
	int lno;
	NodeList *init;
	Node *n;
	char *what;

	t0 = t;
	init = nil;
	lno = lineno;
	what = &#34;field&#34;;
	if(et == TINTER)
		what = &#34;method&#34;;

	for(; l; l=l-&gt;next) {
		n = l-&gt;n;
		lineno = n-&gt;lineno;
		note = nil;

		if(n-&gt;op != ODCLFIELD)
			fatal(&#34;stotype: oops %N\n&#34;, n);
		if(n-&gt;right != N) {
			typecheck(&amp;n-&gt;right, Etype);
			n-&gt;type = n-&gt;right-&gt;type;
			if(n-&gt;type == T) {
				*t0 = T;
				return t0;
			}
			if(n-&gt;left != N)
				n-&gt;left-&gt;type = n-&gt;type;
			n-&gt;right = N;
			if(n-&gt;embedded &amp;&amp; n-&gt;type != T) {
				t1 = n-&gt;type;
				if(t1-&gt;sym == S &amp;&amp; isptr[t1-&gt;etype])
					t1 = t1-&gt;type;
				if(isptr[t1-&gt;etype])
					yyerror(&#34;embedded type cannot be a pointer&#34;);
				else if(t1-&gt;etype == TFORW &amp;&amp; t1-&gt;embedlineno == 0)
					t1-&gt;embedlineno = lineno;
			}
		}

		if(n-&gt;type == T) {
			// assume error already printed
			continue;
		}

		switch(n-&gt;val.ctype) {
		case CTSTR:
			if(et != TSTRUCT)
				yyerror(&#34;interface method cannot have annotation&#34;);
			note = n-&gt;val.u.sval;
			break;
		default:
			if(et != TSTRUCT)
				yyerror(&#34;interface method cannot have annotation&#34;);
			else
				yyerror(&#34;field annotation must be string&#34;);
		case CTxxx:
			note = nil;
			break;
		}

		if(et == TINTER &amp;&amp; n-&gt;left == N) {
			// embedded interface - inline the methods
			if(n-&gt;type-&gt;etype != TINTER) {
				yyerror(&#34;interface contains embedded non-interface %T&#34;, n-&gt;type);
				continue;
			}
			for(t1=n-&gt;type-&gt;type; t1!=T; t1=t1-&gt;down) {
				f = typ(TFIELD);
				f-&gt;type = t1-&gt;type;
				f-&gt;width = BADWIDTH;
				f-&gt;nname = newname(t1-&gt;sym);
				f-&gt;sym = t1-&gt;sym;
				for(t2=*t0; t2!=T; t2=t2-&gt;down) {
					if(t2-&gt;sym == f-&gt;sym) {
						yyerror(&#34;duplicate method %s&#34;, t2-&gt;sym-&gt;name);
						break;
					}
				}
				*t = f;
				t = &amp;f-&gt;down;
			}
			continue;
		}

		f = typ(TFIELD);
		f-&gt;type = n-&gt;type;
		f-&gt;note = note;
		f-&gt;width = BADWIDTH;

		if(n-&gt;left != N &amp;&amp; n-&gt;left-&gt;op == ONAME) {
			f-&gt;nname = n-&gt;left;
			f-&gt;embedded = n-&gt;embedded;
			f-&gt;sym = f-&gt;nname-&gt;sym;
			if(pkgimportname != S &amp;&amp; !exportname(f-&gt;sym-&gt;name))
				f-&gt;sym = pkglookup(f-&gt;sym-&gt;name, structpkg);
			if(f-&gt;sym &amp;&amp; !isblank(f-&gt;nname)) {
				for(t1=*t0; t1!=T; t1=t1-&gt;down) {
					if(t1-&gt;sym == f-&gt;sym) {
						yyerror(&#34;duplicate %s %s&#34;, what, t1-&gt;sym-&gt;name);
						break;
					}
				}
			}
		}

		*t = f;
		t = &amp;f-&gt;down;
	}

	*t = T;
	lineno = lno;
	return t;
}

Type*
dostruct(NodeList *l, int et)
{
	Type *t;
	int funarg;

	/*
	 * convert a parsed id/type list into
	 * a type for struct/interface/arglist
	 */

	funarg = 0;
	if(et == TFUNC) {
		funarg = 1;
		et = TSTRUCT;
	}
	t = typ(et);
	t-&gt;funarg = funarg;
	stotype(l, et, &amp;t-&gt;type);
	if(t-&gt;type == T &amp;&amp; l != nil) {
		t-&gt;broke = 1;
		return t;
	}
	if(!funarg)
		checkwidth(t);
	return t;
}


Node*
embedded(Sym *s)
{
	Node *n;
	char *name;

	// Names sometimes have disambiguation junk
	// appended after a center dot.  Discard it when
	// making the name for the embedded struct field.
	enum { CenterDot = 0xB7 };
	name = s-&gt;name;
	if(utfrune(s-&gt;name, CenterDot)) {
		name = strdup(s-&gt;name);
		*utfrune(name, CenterDot) = 0;
	}

	n = newname(lookup(name));
	n = nod(ODCLFIELD, n, oldname(s));
	n-&gt;embedded = 1;
	return n;
}

/*
 * check that the list of declarations is either all anonymous or all named
 */

static Node*
findtype(NodeList *l)
{
	for(; l; l=l-&gt;next)
		if(l-&gt;n-&gt;op == OKEY)
			return l-&gt;n-&gt;right;
	return N;
}

NodeList*
checkarglist(NodeList *all)
{
	int named;
	Node *n, *t, *nextt;
	NodeList *l;

	named = 0;
	for(l=all; l; l=l-&gt;next) {
		if(l-&gt;n-&gt;op == OKEY) {
			named = 1;
			break;
		}
	}
	if(named) {
		n = N;
		for(l=all; l; l=l-&gt;next) {
			n = l-&gt;n;
			if(n-&gt;op != OKEY &amp;&amp; n-&gt;sym == S) {
				yyerror(&#34;mixed named and unnamed function parameters&#34;);
				break;
			}
		}
		if(l == nil &amp;&amp; n != N &amp;&amp; n-&gt;op != OKEY)
			yyerror(&#34;final function parameter must have type&#34;);
	}

	nextt = nil;
	for(l=all; l; l=l-&gt;next) {
		// can cache result from findtype to avoid
		// quadratic behavior here, but unlikely to matter.
		n = l-&gt;n;
		if(named) {
			if(n-&gt;op == OKEY) {
				t = n-&gt;right;
				n = n-&gt;left;
				nextt = nil;
			} else {
				if(nextt == nil)
					nextt = findtype(l);
				t = nextt;
			}
		} else {
			t = n;
			n = N;
		}
		if(isblank(n))
			n = N;
		if(n != N &amp;&amp; n-&gt;sym == S) {
			t = n;
			n = N;
		}
		if(n != N)
			n = newname(n-&gt;sym);
		n = nod(ODCLFIELD, n, t);
		if(l-&gt;next != nil &amp;&amp; n-&gt;right != N &amp;&amp; n-&gt;right-&gt;op == OTYPE &amp;&amp; isddd(n-&gt;right-&gt;type))
			yyerror(&#34;only last argument can have type ...&#34;);
		l-&gt;n = n;
	}
	return all;
}


Node*
fakethis(void)
{
	Node *n;

	n = nod(ODCLFIELD, N, typenod(ptrto(typ(TSTRUCT))));
	return n;
}

/*
 * Is this field a method on an interface?
 * Those methods have an anonymous
 * *struct{} as the receiver.
 * (See fakethis above.)
 */
int
isifacemethod(Type *f)
{
	Type *rcvr;
	Type *t;

	rcvr = getthisx(f-&gt;type)-&gt;type;
	if(rcvr-&gt;sym != S)
		return 0;
	t = rcvr-&gt;type;
	if(!isptr[t-&gt;etype])
		return 0;
	t = t-&gt;type;
	if(t-&gt;sym != S || t-&gt;etype != TSTRUCT || t-&gt;type != T)
		return 0;
	return 1;
}

/*
 * turn a parsed function declaration
 * into a type
 */
Type*
functype(Node *this, NodeList *in, NodeList *out)
{
	Type *t;
	NodeList *rcvr;

	t = typ(TFUNC);

	rcvr = nil;
	if(this)
		rcvr = list1(this);
	t-&gt;type = dostruct(rcvr, TFUNC);
	t-&gt;type-&gt;down = dostruct(out, TFUNC);
	t-&gt;type-&gt;down-&gt;down = dostruct(in, TFUNC);

	if(this)
		t-&gt;thistuple = 1;
	t-&gt;outtuple = count(out);
	t-&gt;intuple = count(in);
	t-&gt;outnamed = t-&gt;outtuple &gt; 0 &amp;&amp; out-&gt;n-&gt;left != N;

	return t;
}

int
methcmp(Type *t1, Type *t2)
{
	if(t1-&gt;etype != TFUNC)
		return 0;
	if(t2-&gt;etype != TFUNC)
		return 0;

	t1 = t1-&gt;type-&gt;down;	// skip this arg
	t2 = t2-&gt;type-&gt;down;	// skip this arg
	for(;;) {
		if(t1 == t2)
			break;
		if(t1 == T || t2 == T)
			return 0;
		if(t1-&gt;etype != TSTRUCT || t2-&gt;etype != TSTRUCT)
			return 0;

		if(!eqtype(t1-&gt;type, t2-&gt;type))
			return 0;

		t1 = t1-&gt;down;
		t2 = t2-&gt;down;
	}
	return 1;
}

Sym*
methodsym(Sym *nsym, Type *t0)
{
	Sym *s;
	char buf[NSYMB];
	Type *t;

	t = t0;
	if(t == T)
		goto bad;
	s = t-&gt;sym;
	if(s == S) {
		if(!isptr[t-&gt;etype])
			goto bad;
		t = t-&gt;type;
		if(t == T)
			goto bad;
		s = t-&gt;sym;
		if(s == S)
			goto bad;
	}

	// if t0 == *t and t0 has a sym,
	// we want to see *t, not t0, in the method name.
	if(t != t0 &amp;&amp; t0-&gt;sym)
		t0 = ptrto(t);

	snprint(buf, sizeof(buf), &#34;%#hT·%s&#34;, t0, nsym-&gt;name);
	return pkglookup(buf, s-&gt;package);

bad:
	yyerror(&#34;illegal &lt;this&gt; type: %T&#34;, t);
	return S;
}

Node*
methodname(Node *n, Type *t)
{
	Sym *s;

	s = methodsym(n-&gt;sym, t);
	if(s == S)
		return n;
	return newname(s);
}

Node*
methodname1(Node *n, Node *t)
{
	char *star;
	char buf[NSYMB];

	star = &#34;&#34;;
	if(t-&gt;op == OIND) {
		star = &#34;*&#34;;
		t = t-&gt;left;
	}
	if(t-&gt;sym == S || isblank(n))
		return newname(n-&gt;sym);
	snprint(buf, sizeof(buf), &#34;%s%S·%S&#34;, star, t-&gt;sym, n-&gt;sym);
	return newname(pkglookup(buf, t-&gt;sym-&gt;package));
}

/*
 * add a method, declared as a function,
 * n is fieldname, pa is base type, t is function type
 */
void
addmethod(Sym *sf, Type *t, int local)
{
	Type *f, *d, *pa;
	Node *n;

	pa = nil;

	// get field sym
	if(sf == S)
		fatal(&#34;no method symbol&#34;);

	// get parent type sym
	pa = getthisx(t)-&gt;type;	// ptr to this structure
	if(pa == T) {
		yyerror(&#34;missing receiver&#34;);
		return;
	}

	pa = pa-&gt;type;
	f = methtype(pa);
	if(f == T) {
		yyerror(&#34;invalid receiver type %T&#34;, pa);
		return;
	}

	pa = f;
	if(pkgimportname != S &amp;&amp; !exportname(sf-&gt;name))
		sf = pkglookup(sf-&gt;name, pkgimportname-&gt;name);

	n = nod(ODCLFIELD, newname(sf), N);
	n-&gt;type = t;

	d = T;	// last found
	for(f=pa-&gt;method; f!=T; f=f-&gt;down) {
		d = f;
		if(f-&gt;etype != TFIELD)
			fatal(&#34;addmethod: not TFIELD: %N&#34;, f);
		if(strcmp(sf-&gt;name, f-&gt;sym-&gt;name) != 0)
			continue;
		if(!eqtype(t, f-&gt;type))
			yyerror(&#34;method redeclared: %T.%S\n\t%T\n\t%T&#34;, pa, sf, f-&gt;type, t);
		return;
	}

	if(local &amp;&amp; !pa-&gt;local) {
		// defining method on non-local type.
		yyerror(&#34;cannot define new methods on non-local type %T&#34;, pa);
		return;
	}

	if(d == T)
		stotype(list1(n), 0, &amp;pa-&gt;method);
	else
		stotype(list1(n), 0, &amp;d-&gt;down);
	return;
}

void
funccompile(Node *n)
{
	stksize = BADWIDTH;
	maxarg = 0;

	if(n-&gt;type == T) {
		if(nerrors == 0)
			fatal(&#34;funccompile missing type&#34;);
		return;
	}

	// assign parameter offsets
	checkwidth(n-&gt;type);

	if(curfn)
		fatal(&#34;funccompile %S inside %S&#34;, n-&gt;nname-&gt;sym, curfn-&gt;nname-&gt;sym);
	curfn = n;
	typechecklist(n-&gt;nbody, Etop);
	curfn = nil;

	stksize = 0;
	dclcontext = PAUTO;
	funcdepth = n-&gt;funcdepth + 1;
	compile(n);
	curfn = nil;
	funcdepth = 0;
	dclcontext = PEXTERN;
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
