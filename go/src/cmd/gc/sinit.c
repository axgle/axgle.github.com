<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/sinit.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/sinit.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * static initialization
 */

#include	&#34;go.h&#34;

static void
init1(Node *n, NodeList **out)
{
	NodeList *l;

	if(n == N)
		return;
	init1(n-&gt;left, out);
	init1(n-&gt;right, out);
	for(l=n-&gt;list; l; l=l-&gt;next)
		init1(l-&gt;n, out);

	if(n-&gt;op != ONAME)
		return;
	switch(n-&gt;class) {
	case PEXTERN:
	case PFUNC:
		break;
	default:
		if(isblank(n))
			*out = list(*out, n-&gt;defn);
		return;
	}

	if(n-&gt;initorder == 1)
		return;
	if(n-&gt;initorder == 2)
		fatal(&#34;init loop&#34;);

	// make sure that everything n depends on is initialized.
	// n-&gt;defn is an assignment to n
	n-&gt;initorder = 2;
	if(n-&gt;defn != N) {
		switch(n-&gt;defn-&gt;op) {
		default:
			goto bad;

		case ODCLFUNC:
			for(l=n-&gt;defn-&gt;nbody; l; l=l-&gt;next)
				init1(l-&gt;n, out);
			break;

		case OAS:
			if(n-&gt;defn-&gt;left != n)
				goto bad;
			n-&gt;defn-&gt;dodata = 1;
			init1(n-&gt;defn-&gt;right, out);
			if(debug[&#39;j&#39;])
				print(&#34;%S\n&#34;, n-&gt;sym);
			*out = list(*out, n-&gt;defn);
			break;
		}
	}
	n-&gt;initorder = 1;
	return;

bad:
	dump(&#34;defn&#34;, n-&gt;defn);
	fatal(&#34;init1: bad defn&#34;);
}

static void
initreorder(NodeList *l, NodeList **out)
{
	Node *n;

	for(; l; l=l-&gt;next) {
		n = l-&gt;n;
		switch(n-&gt;op) {
		case ODCLFUNC:
		case ODCLCONST:
		case ODCLTYPE:
			continue;
		}
		initreorder(n-&gt;ninit, out);
		n-&gt;ninit = nil;
		init1(n, out);
	}
}

NodeList*
initfix(NodeList *l)
{
	NodeList *lout;

	lout = nil;
	initreorder(l, &amp;lout);
	return lout;
}

/*
 * from here down is the walk analysis
 * of composit literals.
 * most of the work is to generate
 * data statements for the constant
 * part of the composit literal.
 */

static int
isliteral(Node *n)
{
	if(n-&gt;op == OLITERAL)
		if(n-&gt;val.ctype != CTNIL)
			return 1;
	return 0;
}

static int
simplename(Node *n)
{
	if(n-&gt;op != ONAME)
		goto no;
	if(!n-&gt;addable)
		goto no;
	if(n-&gt;class &amp; PHEAP)
		goto no;
	if(n-&gt;class == PPARAMREF)
		goto no;
	return 1;

no:
	return 0;
}

static void	arraylit(Node *n, Node *var, int pass, NodeList **init);

static void
structlit(Node *n, Node *var, int pass, NodeList **init)
{
	Node *r, *a;
	NodeList *nl;
	Node *index, *value;

	for(nl=n-&gt;list; nl; nl=nl-&gt;next) {
		r = nl-&gt;n;
		if(r-&gt;op != OKEY)
			fatal(&#34;structlit: rhs not OKEY: %N&#34;, r);
		index = r-&gt;left;
		value = r-&gt;right;

		switch(value-&gt;op) {
		case OARRAYLIT:
			if(value-&gt;type-&gt;bound &lt; 0)
				break;
			a = nod(ODOT, var, newname(index-&gt;sym));
			arraylit(value, a, pass, init);
			continue;

		case OSTRUCTLIT:
			a = nod(ODOT, var, newname(index-&gt;sym));
			structlit(value, a, pass, init);
			continue;
		}

		if(isliteral(value)) {
			if(pass == 2)
				continue;
		} else
			if(pass == 1)
				continue;

		// build list of var.field = expr
		a = nod(ODOT, var, newname(index-&gt;sym));
		a = nod(OAS, a, value);
		typecheck(&amp;a, Etop);
		walkexpr(&amp;a, init);
		if(pass == 1) {
			if(a-&gt;op != OAS)
				fatal(&#34;structlit: not as&#34;);
			a-&gt;dodata = 2;
		}
		*init = list(*init, a);
	}
}

static void
arraylit(Node *n, Node *var, int pass, NodeList **init)
{
	Node *r, *a;
	NodeList *l;
	Node *index, *value;

	for(l=n-&gt;list; l; l=l-&gt;next) {
		r = l-&gt;n;
		if(r-&gt;op != OKEY)
			fatal(&#34;arraylit: rhs not OKEY: %N&#34;, r);
		index = r-&gt;left;
		value = r-&gt;right;

		switch(value-&gt;op) {
		case OARRAYLIT:
			if(value-&gt;type-&gt;bound &lt; 0)
				break;
			a = nod(OINDEX, var, index);
			arraylit(value, a, pass, init);
			continue;

		case OSTRUCTLIT:
			a = nod(OINDEX, var, index);
			structlit(value, a, pass, init);
			continue;
		}

		if(isliteral(index) &amp;&amp; isliteral(value)) {
			if(pass == 2)
				continue;
		} else
			if(pass == 1)
				continue;

		// build list of var[index] = value
		a = nod(OINDEX, var, index);
		a = nod(OAS, a, value);
		typecheck(&amp;a, Etop);
		walkexpr(&amp;a, init);	// add any assignments in r to top
		if(pass == 1) {
			if(a-&gt;op != OAS)
				fatal(&#34;structlit: not as&#34;);
			a-&gt;dodata = 2;
		}
		*init = list(*init, a);
	}
}

static void
slicelit(Node *n, Node *var, NodeList **init)
{
	Node *r, *a;
	NodeList *l;
	Type *t;
	Node *vstat, *vheap;
	Node *index, *value;

	// make an array type
	t = shallow(n-&gt;type);
	t-&gt;bound = mpgetfix(n-&gt;right-&gt;val.u.xval);
	t-&gt;width = 0;
	dowidth(t);

	// make static initialized array
	vstat = staticname(t);
	arraylit(n, vstat, 1, init);

	// make new *array heap
	vheap = nod(OXXX, N, N);
	tempname(vheap, ptrto(t));

	a = nod(ONEW, N, N);
	a-&gt;list = list1(typenod(t));
	a = nod(OAS, vheap, a);
	typecheck(&amp;a, Etop);
	walkexpr(&amp;a, init);
	*init = list(*init, a);

	// copy static to heap
	a = nod(OIND, vheap, N);
	a = nod(OAS, a, vstat);
	typecheck(&amp;a, Etop);
	walkexpr(&amp;a, init);
	*init = list(*init, a);

	// make slice out of heap
	a = nod(OAS, var, vheap);
	typecheck(&amp;a, Etop);
	walkexpr(&amp;a, init);
	*init = list(*init, a);

	// put dynamics into slice
	for(l=n-&gt;list; l; l=l-&gt;next) {
		r = l-&gt;n;
		if(r-&gt;op != OKEY)
			fatal(&#34;slicelit: rhs not OKEY: %N&#34;, r);
		index = r-&gt;left;
		value = r-&gt;right;

		switch(value-&gt;op) {
		case OARRAYLIT:
			if(value-&gt;type-&gt;bound &lt; 0)
				break;
			a = nod(OINDEX, var, index);
			arraylit(value, a, 2, init);
			continue;

		case OSTRUCTLIT:
			a = nod(OINDEX, var, index);
			structlit(value, a, 2, init);
			continue;
		}

		if(isliteral(index) &amp;&amp; isliteral(value))
			continue;

		// build list of var[c] = expr
		a = nod(OINDEX, var, index);
		a = nod(OAS, a, value);
		typecheck(&amp;a, Etop);
		walkexpr(&amp;a, init);	// add any assignments in r to top
		*init = list(*init, a);
	}
}

static void
maplit(Node *n, Node *var, NodeList **init)
{
	Node *r, *a;
	NodeList *l;
	int nerr, b;
	Type *t, *tk, *tv, *t1;
	Node *vstat, *index, *value;
	Sym *syma, *symb;

	// make the map var
	nerr = nerrors;

	a = nod(OMAKE, N, N);
	a-&gt;list = list1(typenod(n-&gt;type));
	a = nod(OAS, var, a);
	typecheck(&amp;a, Etop);
	walkexpr(&amp;a, init);
	*init = list(*init, a);

	// count the initializers
	b = 0;
	for(l=n-&gt;list; l; l=l-&gt;next) {
		r = l-&gt;n;

		if(r-&gt;op != OKEY)
			fatal(&#34;slicelit: rhs not OKEY: %N&#34;, r);
		index = r-&gt;left;
		value = r-&gt;right;

		if(isliteral(index) &amp;&amp; isliteral(value))
			b++;
	}

	t = T;
	if(b != 0) {
		// build type [count]struct { a Tindex, b Tvalue }
		t = n-&gt;type;
		tk = t-&gt;down;
		tv = t-&gt;type;

		symb = lookup(&#34;b&#34;);
		t = typ(TFIELD);
		t-&gt;type = tv;
		t-&gt;sym = symb;

		syma = lookup(&#34;a&#34;);
		t1 = t;
		t = typ(TFIELD);
		t-&gt;type = tk;
		t-&gt;sym = syma;
		t-&gt;down = t1;

		t1 = t;
		t = typ(TSTRUCT);
		t-&gt;type = t1;

		t1 = t;
		t = typ(TARRAY);
		t-&gt;bound = b;
		t-&gt;type = t1;

		dowidth(t);

		// make and initialize static array
		vstat = staticname(t);
		b = 0;
		for(l=n-&gt;list; l; l=l-&gt;next) {
			r = l-&gt;n;

			if(r-&gt;op != OKEY)
				fatal(&#34;slicelit: rhs not OKEY: %N&#34;, r);
			index = r-&gt;left;
			value = r-&gt;right;

			if(isliteral(index) &amp;&amp; isliteral(value)) {
				// build vstat[b].a = key;
				a = nodintconst(b);
				a = nod(OINDEX, vstat, a);
				a = nod(ODOT, a, newname(syma));
				a = nod(OAS, a, index);
				typecheck(&amp;a, Etop);
				walkexpr(&amp;a, init);
				a-&gt;dodata = 2;
				*init = list(*init, a);

				// build vstat[b].b = value;
				a = nodintconst(b);
				a = nod(OINDEX, vstat, a);
				a = nod(ODOT, a, newname(symb));
				a = nod(OAS, a, value);
				typecheck(&amp;a, Etop);
				walkexpr(&amp;a, init);
				a-&gt;dodata = 2;
				*init = list(*init, a);

				b++;
			}
		}

		// loop adding structure elements to map
		// for i = 0; i &lt; len(vstat); i++ {
		//	map[vstat[i].a] = vstat[i].b
		// }
		index = nod(OXXX, N, N);
		tempname(index, types[TINT]);

		a = nod(OINDEX, vstat, index);
		a = nod(ODOT, a, newname(symb));

		r = nod(OINDEX, vstat, index);
		r = nod(ODOT, r, newname(syma));
		r = nod(OINDEX, var, r);

		r = nod(OAS, r, a);

		a = nod(OFOR, N, N);
		a-&gt;nbody = list1(r);

		a-&gt;ninit = list1(nod(OAS, index, nodintconst(0)));
		a-&gt;ntest = nod(OLT, index, nodintconst(t-&gt;bound));
		a-&gt;nincr = nod(OASOP, index, nodintconst(1));
		a-&gt;nincr-&gt;etype = OADD;

		typecheck(&amp;a, Etop);
		walkstmt(&amp;a);
		*init = list(*init, a);
	}

	// put in dynamic entries one-at-a-time
	for(l=n-&gt;list; l; l=l-&gt;next) {
		r = l-&gt;n;

		if(r-&gt;op != OKEY)
			fatal(&#34;slicelit: rhs not OKEY: %N&#34;, r);
		index = r-&gt;left;
		value = r-&gt;right;

		if(isliteral(index) &amp;&amp; isliteral(value))
			continue;

		// build list of var[c] = expr
		a = nod(OINDEX, var, r-&gt;left);
		a = nod(OAS, a, r-&gt;right);
		typecheck(&amp;a, Etop);
		walkexpr(&amp;a, init);
		if(nerr != nerrors)
			break;

		*init = list(*init, a);
	}
}

void
anylit(Node *n, Node *var, NodeList **init)
{
	Type *t;
	Node *a, *vstat;

	t = n-&gt;type;
	switch(n-&gt;op) {
	default:
		fatal(&#34;anylit: not lit&#34;);

	case OSTRUCTLIT:
		if(t-&gt;etype != TSTRUCT)
			fatal(&#34;anylit: not struct&#34;);

		if(simplename(var)) {

			// lay out static data
			vstat = staticname(t);
			structlit(n, vstat, 1, init);

			// copy static to automatic
			a = nod(OAS, var, vstat);
			typecheck(&amp;a, Etop);
			walkexpr(&amp;a, init);
			*init = list(*init, a);

			// add expressions to automatic
			structlit(n, var, 2, init);
			break;
		}

		// initialize of not completely specified
		if(count(n-&gt;list) &lt; structcount(t)) {
			a = nod(OAS, var, N);
			typecheck(&amp;a, Etop);
			walkexpr(&amp;a, init);
			*init = list(*init, a);
		}
		structlit(n, var, 3, init);
		break;

	case OARRAYLIT:
		if(t-&gt;etype != TARRAY)
			fatal(&#34;anylit: not array&#34;);
		if(t-&gt;bound &lt; 0) {
			slicelit(n, var, init);
			break;
		}

		if(simplename(var)) {

			// lay out static data
			vstat = staticname(t);
			arraylit(n, vstat, 1, init);

			// copy static to automatic
			a = nod(OAS, var, vstat);
			typecheck(&amp;a, Etop);
			walkexpr(&amp;a, init);
			*init = list(*init, a);

			// add expressions to automatic
			arraylit(n, var, 2, init);
			break;
		}

		// initialize of not completely specified
		if(count(n-&gt;list) &lt; t-&gt;bound) {
			a = nod(OAS, var, N);
			typecheck(&amp;a, Etop);
			walkexpr(&amp;a, init);
			*init = list(*init, a);
		}
		arraylit(n, var, 3, init);
		break;

	case OMAPLIT:
		if(t-&gt;etype != TMAP)
			fatal(&#34;anylit: not map&#34;);
		maplit(n, var, init);
		break;
	}
}

int
oaslit(Node *n, NodeList **init)
{
	Type *t;
	Node *vstat, *a;

	if(n-&gt;left == N || n-&gt;right == N)
		goto no;
	if(n-&gt;left-&gt;type == T || n-&gt;right-&gt;type == T)
		goto no;
	if(!simplename(n-&gt;left))
		goto no;
	if(!eqtype(n-&gt;left-&gt;type, n-&gt;right-&gt;type))
		goto no;
	if(n-&gt;dodata == 1)
		goto initctxt;

	switch(n-&gt;right-&gt;op) {
	default:
		goto no;

	case OSTRUCTLIT:
	case OARRAYLIT:
	case OMAPLIT:
		if(vmatch1(n-&gt;left, n-&gt;right))
			goto no;
		anylit(n-&gt;right, n-&gt;left, init);
		break;
	}
	n-&gt;op = OEMPTY;
	return 1;

no:
	// not a special composit literal assignment
	return 0;

initctxt:
	// in the initialization context
	// we are trying to put data statements
	// right into the initialized variables
	switch(n-&gt;right-&gt;op) {
	default:
		goto no;

	case OSTRUCTLIT:
		structlit(n-&gt;right, n-&gt;left, 1, init);
		structlit(n-&gt;right, n-&gt;left, 2, init);
		break;

	case OARRAYLIT:
		t = n-&gt;right-&gt;type;
		if(t == T)
			goto no;
		if(t-&gt;bound &gt;= 0) {
			arraylit(n-&gt;right, n-&gt;left, 1, init);
			arraylit(n-&gt;right, n-&gt;left, 2, init);
			break;
		}

		// make a static slice
		// make an array type
		t = shallow(t);
		t-&gt;bound = mpgetfix(n-&gt;right-&gt;right-&gt;val.u.xval);
		t-&gt;width = 0;
		dowidth(t);

		// make static initialized array
		vstat = staticname(t);
		arraylit(n-&gt;right, vstat, 1, init);
		arraylit(n-&gt;right, vstat, 2, init);

		// copy static to slice
		a = nod(OADDR, vstat, N);
		a = nod(OAS, n-&gt;left, a);
		typecheck(&amp;a, Etop);
// turns into a function that is hard to parse
// in ggen where it is turned into DATA statements
//		walkexpr(&amp;a, init);
		a-&gt;dodata = 2;
		*init = list(*init, a);
		break;

	case OMAPLIT:
		maplit(n-&gt;right, n-&gt;left, init);
		break;
	}
	n-&gt;op = OEMPTY;
	return 1;
}

int
getlit(Node *lit)
{
	if(smallintconst(lit))
		return mpgetfix(lit-&gt;val.u.xval);
	return -1;
}

int
stataddr(Node *nam, Node *n)
{
	int l;

	if(n == N)
		goto no;

	switch(n-&gt;op) {

	case ONAME:
		*nam = *n;
		return n-&gt;addable;

	case ODOT:
		if(!stataddr(nam, n-&gt;left))
			break;
		nam-&gt;xoffset += n-&gt;xoffset;
		nam-&gt;type = n-&gt;type;
		return 1;

	case OINDEX:
		if(n-&gt;left-&gt;type-&gt;bound &lt; 0)
			break;
		if(!stataddr(nam, n-&gt;left))
			break;
		l = getlit(n-&gt;right);
		if(l &lt; 0)
			break;
		nam-&gt;xoffset += l*n-&gt;type-&gt;width;
		nam-&gt;type = n-&gt;type;
		return 1;
	}

no:
	return 0;
}

int
gen_as_init(Node *n)
{
	Node *nr, *nl;
	Node nam, nod1;

	if(n-&gt;dodata == 0)
		goto no;

	nr = n-&gt;right;
	nl = n-&gt;left;
	if(nr == N) {
		if(!stataddr(&amp;nam, nl))
			goto no;
		if(nam.class != PEXTERN)
			goto no;
		goto yes;
	}

	if(nr-&gt;type == T || !eqtype(nl-&gt;type, nr-&gt;type))
		goto no;

	if(!stataddr(&amp;nam, nl))
		goto no;

	if(nam.class != PEXTERN)
		goto no;

	switch(nr-&gt;op) {
	default:
		goto no;

	case OCONVSLICE:
		goto slice;

	case OLITERAL:
		break;
	}

	switch(nr-&gt;type-&gt;etype) {
	default:
		goto no;

	case TBOOL:
	case TINT8:
	case TUINT8:
	case TINT16:
	case TUINT16:
	case TINT32:
	case TUINT32:
	case TINT64:
	case TUINT64:
	case TINT:
	case TUINT:
	case TUINTPTR:
	case TPTR32:
	case TPTR64:
	case TFLOAT32:
	case TFLOAT64:
	case TFLOAT:
		gused(N); // in case the data is the dest of a goto
		gdata(&amp;nam, nr, nr-&gt;type-&gt;width);
		break;

	case TSTRING:
		gused(N); // in case the data is the dest of a goto
		gdatastring(&amp;nam, nr-&gt;val.u.sval);
		break;
	}

yes:
	return 1;

slice:
	gused(N); // in case the data is the dest of a goto
	nr = n-&gt;right-&gt;left;
	if(nr == N || nr-&gt;op != OADDR)
		goto no;
	nr = nr-&gt;left;
	if(nr == N || nr-&gt;op != ONAME)
		goto no;

	// nr is the array being converted to a slice
	if(nr-&gt;type == T || nr-&gt;type-&gt;etype != TARRAY || nr-&gt;type-&gt;bound &lt; 0)
		goto no;

	nam.xoffset += Array_array;
	gdata(&amp;nam, n-&gt;right-&gt;left, types[tptr]-&gt;width);

	nam.xoffset += Array_nel-Array_array;
	nodconst(&amp;nod1, types[TINT32], nr-&gt;type-&gt;bound);
	gdata(&amp;nam, &amp;nod1, types[TINT32]-&gt;width);

	nam.xoffset += Array_cap-Array_nel;
	gdata(&amp;nam, &amp;nod1, types[TINT32]-&gt;width);

	goto yes;

no:
	if(n-&gt;dodata == 2) {
		dump(&#34;\ngen_as_init&#34;, n);
		fatal(&#34;gen_as_init couldnt make data statement&#34;);
	}
	return 0;
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
