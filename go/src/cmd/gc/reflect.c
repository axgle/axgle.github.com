<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/reflect.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/reflect.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;go.h&#34;

/*
 * runtime interface and reflection data structures
 */

static	NodeList*	signatlist;
static	Sym*	dtypesym(Type*);

static int
sigcmp(Sig *a, Sig *b)
{
	return strcmp(a-&gt;name, b-&gt;name);
}

static Sig*
lsort(Sig *l, int(*f)(Sig*, Sig*))
{
	Sig *l1, *l2, *le;

	if(l == 0 || l-&gt;link == 0)
		return l;

	l1 = l;
	l2 = l;
	for(;;) {
		l2 = l2-&gt;link;
		if(l2 == 0)
			break;
		l2 = l2-&gt;link;
		if(l2 == 0)
			break;
		l1 = l1-&gt;link;
	}

	l2 = l1-&gt;link;
	l1-&gt;link = 0;
	l1 = lsort(l, f);
	l2 = lsort(l2, f);

	/* set up lead element */
	if((*f)(l1, l2) &lt; 0) {
		l = l1;
		l1 = l1-&gt;link;
	} else {
		l = l2;
		l2 = l2-&gt;link;
	}
	le = l;

	for(;;) {
		if(l1 == 0) {
			while(l2) {
				le-&gt;link = l2;
				le = l2;
				l2 = l2-&gt;link;
			}
			le-&gt;link = 0;
			break;
		}
		if(l2 == 0) {
			while(l1) {
				le-&gt;link = l1;
				le = l1;
				l1 = l1-&gt;link;
			}
			break;
		}
		if((*f)(l1, l2) &lt; 0) {
			le-&gt;link = l1;
			le = l1;
			l1 = l1-&gt;link;
		} else {
			le-&gt;link = l2;
			le = l2;
			l2 = l2-&gt;link;
		}
	}
	le-&gt;link = 0;
	return l;
}

/*
 * f is method type, with receiver.
 * return function type, receiver as first argument.
 */
static Type*
methodfunc(Type *f)
{
	NodeList *in, *out;
	Node *d;
	Type *t;

	in = nil;
	if(!isifacemethod(f)) {
		d = nod(ODCLFIELD, N, N);
		d-&gt;type = getthisx(f-&gt;type)-&gt;type-&gt;type;
		in = list(in, d);
	}
	for(t=getinargx(f-&gt;type)-&gt;type; t; t=t-&gt;down) {
		d = nod(ODCLFIELD, N, N);
		d-&gt;type = t-&gt;type;
		in = list(in, d);
	}

	out = nil;
	for(t=getoutargx(f-&gt;type)-&gt;type; t; t=t-&gt;down) {
		d = nod(ODCLFIELD, N, N);
		d-&gt;type = t-&gt;type;
		out = list(out, d);
	}

	return functype(N, in, out);
}

/*
 * return methods of non-interface type t,
 * sorted by hash.
 * generates stub functions as needed.
 */
static Sig*
methods(Type *t)
{
	int o;
	Type *f, *mt, *it, *this;
	Sig *a, *b;
	Sym *method;
	Prog *oldlist;

	// named method type
	mt = methtype(t);
	if(mt == T)
		return nil;
	expandmeth(mt-&gt;sym, mt);

	// type stored in interface word
	it = t;
	if(it-&gt;width &gt; widthptr)
		it = ptrto(t);

	// make list of methods for t,
	// generating code if necessary.
	a = nil;
	o = 0;
	oldlist = nil;
	for(f=mt-&gt;xmethod; f; f=f-&gt;down) {
		if(f-&gt;type-&gt;etype != TFUNC)
			continue;
		if(f-&gt;etype != TFIELD)
			fatal(&#34;methods: not field&#34;);
		method = f-&gt;sym;
		if(method == nil)
			continue;

		// get receiver type for this particular method.
		// if pointer receiver but non-pointer t and
		// this is not an embedded pointer inside a struct,
		// method does not apply.
		this = getthisx(f-&gt;type)-&gt;type-&gt;type;
		if(isptr[this-&gt;etype] &amp;&amp; this-&gt;type == t)
			continue;
		if(isptr[this-&gt;etype] &amp;&amp; !isptr[t-&gt;etype]
		&amp;&amp; f-&gt;embedded != 2 &amp;&amp; !isifacemethod(f))
			continue;

		b = mal(sizeof(*b));
		b-&gt;link = a;
		a = b;

		a-&gt;name = method-&gt;name;
		a-&gt;hash = PRIME8*stringhash(a-&gt;name) + PRIME9*typehash(f-&gt;type);
		if(!exportname(a-&gt;name)) {
			a-&gt;package = method-&gt;package;
			a-&gt;hash += PRIME10*stringhash(a-&gt;package);
		}
		a-&gt;perm = o++;
		a-&gt;isym = methodsym(method, it);
		a-&gt;tsym = methodsym(method, t);
		a-&gt;type = methodfunc(f);

		if(!(a-&gt;isym-&gt;flags &amp; SymSiggen)) {
			a-&gt;isym-&gt;flags |= SymSiggen;
			if(!eqtype(this, it)) {
				if(oldlist == nil)
					oldlist = pc;
				// Is okay to call genwrapper here always,
				// but we can generate more efficient code
				// using genembedtramp if all that is necessary
				// is a pointer adjustment and a JMP.
				if(isptr[it-&gt;etype] &amp;&amp; isptr[this-&gt;etype]
				&amp;&amp; f-&gt;embedded &amp;&amp; !isifacemethod(f))
					genembedtramp(it, f, a-&gt;isym);
				else
					genwrapper(it, f, a-&gt;isym);
			}
		}

		if(!(a-&gt;tsym-&gt;flags &amp; SymSiggen)) {
			a-&gt;tsym-&gt;flags |= SymSiggen;
			if(!eqtype(this, t)) {
				if(oldlist == nil)
					oldlist = pc;
				if(isptr[t-&gt;etype] &amp;&amp; isptr[this-&gt;etype]
				&amp;&amp; f-&gt;embedded &amp;&amp; !isifacemethod(f))
					genembedtramp(t, f, a-&gt;tsym);
				else
					genwrapper(t, f, a-&gt;tsym);
			}
		}
	}

	// restore data output
	if(oldlist) {
		// old list ended with AEND; change to ANOP
		// so that the trampolines that follow can be found.
		nopout(oldlist);

		// start new data list
		newplist();
	}

	return lsort(a, sigcmp);
}

/*
 * return methods of interface type t, sorted by hash.
 */
Sig*
imethods(Type *t)
{
	Sig *a, *b;
	int o;
	Type *f;

	a = nil;
	o = 0;
	for(f=t-&gt;type; f; f=f-&gt;down) {
		if(f-&gt;etype != TFIELD)
			fatal(&#34;imethods: not field&#34;);
		if(f-&gt;type-&gt;etype != TFUNC || f-&gt;sym == nil)
			continue;
		b = mal(sizeof(*b));
		b-&gt;link = a;
		a = b;

		a-&gt;name = f-&gt;sym-&gt;name;
		a-&gt;hash = PRIME8*stringhash(a-&gt;name) + PRIME9*typehash(f-&gt;type);
		if(!exportname(a-&gt;name)) {
			a-&gt;package = f-&gt;sym-&gt;package;
			a-&gt;hash += PRIME10*stringhash(a-&gt;package);
		}
		a-&gt;perm = o++;
		a-&gt;offset = 0;
		a-&gt;type = methodfunc(f);
	}

	return lsort(a, sigcmp);
}

/*
 * uncommonType
 * ../../pkg/runtime/type.go:/uncommonType
 */
static Sym*
dextratype(Type *t)
{
	int ot, n;
	char *p;
	Sym *s;
	Sig *a, *m;

	m = methods(t);
	if(t-&gt;sym == nil &amp;&amp; m == nil)
		return nil;

	n = 0;
	for(a=m; a; a=a-&gt;link) {
		dtypesym(a-&gt;type);
		n++;
	}

	p = smprint(&#34;%#-T&#34;, t);
	s = pkglookup(p, &#34;extratype&#34;);
	ot = 0;
	if(t-&gt;sym) {
		ot = dgostringptr(s, ot, t-&gt;sym-&gt;name);
		if(t != types[t-&gt;etype])
			ot = dgostringptr(s, ot, t-&gt;sym-&gt;package);
		else
			ot = dgostringptr(s, ot, nil);
	} else {
		ot = dgostringptr(s, ot, nil);
		ot = dgostringptr(s, ot, nil);
	}

	// slice header
	ot = dsymptr(s, ot, s, ot + widthptr + 2*4);
	ot = duint32(s, ot, n);
	ot = duint32(s, ot, n);

	// methods
	for(a=m; a; a=a-&gt;link) {
		// method
		// ../../pkg/runtime/type.go:/method
		ot = duint32(s, ot, a-&gt;hash);
		ot = rnd(ot, widthptr);
		ot = dgostringptr(s, ot, a-&gt;name);
		ot = dgostringptr(s, ot, a-&gt;package);
		ot = dsymptr(s, ot, dtypesym(a-&gt;type), 0);
		if(a-&gt;isym)
			ot = dsymptr(s, ot, a-&gt;isym, 0);
		else
			ot = duintptr(s, ot, 0);
		if(a-&gt;tsym)
			ot = dsymptr(s, ot, a-&gt;tsym, 0);
		else
			ot = duintptr(s, ot, 0);
	}
	ggloblsym(s, ot, 0);

	return s;
}

static char*
structnames[] =
{
	[TINT]		= &#34;*runtime.IntType&#34;,
	[TUINT]		= &#34;*runtime.UintType&#34;,
	[TINT8]		= &#34;*runtime.Int8Type&#34;,
	[TUINT8]	= &#34;*runtime.Uint8Type&#34;,
	[TINT16]	= &#34;*runtime.Int16Type&#34;,
	[TUINT16]	= &#34;*runtime.Uint16Type&#34;,
	[TINT32]	= &#34;*runtime.Int32Type&#34;,
	[TUINT32]	= &#34;*runtime.Uint32Type&#34;,
	[TINT64]	= &#34;*runtime.Int64Type&#34;,
	[TUINT64]	= &#34;*runtime.Uint64Type&#34;,
	[TUINTPTR]	= &#34;*runtime.UintptrType&#34;,
	[TFLOAT]	= &#34;*runtime.FloatType&#34;,
	[TFLOAT32]	= &#34;*runtime.Float32Type&#34;,
	[TFLOAT64]	= &#34;*runtime.Float64Type&#34;,
	[TBOOL]		= &#34;*runtime.BoolType&#34;,
	[TSTRING]		= &#34;*runtime.StringType&#34;,
	[TDDD]		= &#34;*runtime.DotDotDotType&#34;,

	[TPTR32]		= &#34;*runtime.PtrType&#34;,
	[TPTR64]		= &#34;*runtime.PtrType&#34;,
	[TSTRUCT]	= &#34;*runtime.StructType&#34;,
	[TINTER]		= &#34;*runtime.InterfaceType&#34;,
	[TCHAN]		= &#34;*runtime.ChanType&#34;,
	[TMAP]		= &#34;*runtime.MapType&#34;,
	[TARRAY]		= &#34;*runtime.ArrayType&#34;,
	[TFUNC]		= &#34;*runtime.FuncType&#34;,
};

static Sym*
typestruct(Type *t)
{
	char *name;
	int et;

	et = t-&gt;etype;
	if(et &lt; 0 || et &gt;= nelem(structnames) || (name = structnames[et]) == nil) {
		fatal(&#34;typestruct %lT&#34;, t);
		return nil;	// silence gcc
	}

	if(isslice(t))
		name = &#34;*runtime.SliceType&#34;;

	if(isptr[et] &amp;&amp; t-&gt;type-&gt;etype == TANY)
		name = &#34;*runtime.UnsafePointerType&#34;;

	return pkglookup(name, &#34;type&#34;);
}

/*
 * commonType
 * ../../pkg/runtime/type.go:/commonType
 */
static int
dcommontype(Sym *s, int ot, Type *t)
{
	int i;
	Sym *s1;
	Type *elem;
	char *p;

	dowidth(t);
	s1 = dextratype(t);

	// empty interface pointing at this type.
	// all the references that we emit are *interface{};
	// they point here.
	ot = rnd(ot, widthptr);
	ot = dsymptr(s, ot, typestruct(t), 0);
	ot = dsymptr(s, ot, s, 2*widthptr);

	// ../../pkg/runtime/type.go:/commonType
	// actual type structure
	//	type commonType struct {
	//		size uintptr;
	//		hash uint32;
	//		alg uint8;
	//		align uint8;
	//		fieldAlign uint8;
	//		string *string;
	//		*nameInfo;
	//	}
	ot = duintptr(s, ot, t-&gt;width);
	ot = duint32(s, ot, typehash(t));
	ot = duint8(s, ot, algtype(t));
	elem = t;
	while(elem-&gt;etype == TARRAY &amp;&amp; elem-&gt;bound &gt;= 0)
		elem = elem-&gt;type;
	i = elem-&gt;width;
	if(i &gt; maxround)
		i = maxround;
	ot = duint8(s, ot, i);	// align
	ot = duint8(s, ot, i);	// fieldAlign
	p = smprint(&#34;%#-T&#34;, t);
	ot = dgostringptr(s, ot, p);	// string
	free(p);
	if(s1)
		ot = dsymptr(s, ot, s1, 0);	// extraType
	else
		ot = duintptr(s, ot, 0);

	return ot;
}

Sym*
typesym(Type *t)
{
	char *p;
	Sym *s;

	p = smprint(&#34;%#-T&#34;, t);
	s = pkglookup(p, &#34;type&#34;);
	free(p);
	return s;
}

Node*
typename(Type *t)
{
	Sym *s;
	Node *n;

	if((isptr[t-&gt;etype] &amp;&amp; t-&gt;type == T) || isideal(t))
		fatal(&#34;typename %T&#34;, t);
	s = typesym(t);
	if(s-&gt;def == N) {
		n = nod(ONAME, N, N);
		n-&gt;sym = s;
		n-&gt;type = types[TUINT8];
		n-&gt;addable = 1;
		n-&gt;ullman = 1;
		n-&gt;class = PEXTERN;
		n-&gt;xoffset = 0;
		s-&gt;def = n;

		signatlist = list(signatlist, typenod(t));
	}

	n = nod(OADDR, s-&gt;def, N);
	n-&gt;type = ptrto(s-&gt;def-&gt;type);
	n-&gt;addable = 1;
	n-&gt;ullman = 2;
	return n;
}

Sym*
dtypesym(Type *t)
{
	int ot, n;
	Sym *s, *s1, *s2;
	Sig *a, *m;
	Type *t1;
	Sym *tsym;

	if(isideal(t))
		fatal(&#34;dtypesym %T&#34;, t);

	s = typesym(t);
	if(s-&gt;flags &amp; SymSiggen)
		return s;
	s-&gt;flags |= SymSiggen;

	// special case (look for runtime below):
	// when compiling package runtime,
	// emit the type structures for int, float, etc.
	t1 = T;
	if(isptr[t-&gt;etype])
		t1 = t-&gt;type;
	tsym = S;
	if(t1)
		tsym = t1-&gt;sym;
	else
		tsym = t-&gt;sym;

	if(strcmp(package, &#34;runtime&#34;) == 0) {
		if(t == types[t-&gt;etype])
			goto ok;
		if(t1 &amp;&amp; t1 == types[t1-&gt;etype])
			goto ok;
		if(t1 &amp;&amp; t1-&gt;etype == tptr &amp;&amp; t1-&gt;type-&gt;etype == TANY)
			goto ok;
	}

	// named types from other files are defined in those files
	if(t-&gt;sym &amp;&amp; !t-&gt;local)
		return s;
	if(!t-&gt;sym &amp;&amp; t1 &amp;&amp; t1-&gt;sym &amp;&amp; !t1-&gt;local)
		return s;
	if(isforw[t-&gt;etype] || (t1 &amp;&amp; isforw[t1-&gt;etype]))
		return s;

ok:
	ot = 0;
	switch(t-&gt;etype) {
	default:
		ot = dcommontype(s, ot, t);
		break;

	case TARRAY:
		// ../../pkg/runtime/type.go:/ArrayType
		s1 = dtypesym(t-&gt;type);
		ot = dcommontype(s, ot, t);
		ot = dsymptr(s, ot, s1, 0);
		if(t-&gt;bound &lt; 0)
			ot = duintptr(s, ot, -1);
		else
			ot = duintptr(s, ot, t-&gt;bound);
		break;

	case TCHAN:
		// ../../pkg/runtime/type.go:/ChanType
		s1 = dtypesym(t-&gt;type);
		ot = dcommontype(s, ot, t);
		ot = dsymptr(s, ot, s1, 0);
		ot = duintptr(s, ot, t-&gt;chan);
		break;

	case TFUNC:
		for(t1=getthisx(t)-&gt;type; t1; t1=t1-&gt;down)
			dtypesym(t1-&gt;type);
		for(t1=getinargx(t)-&gt;type; t1; t1=t1-&gt;down)
			dtypesym(t1-&gt;type);
		for(t1=getoutargx(t)-&gt;type; t1; t1=t1-&gt;down)
			dtypesym(t1-&gt;type);

		ot = dcommontype(s, ot, t);

		// two slice headers: in and out.
		ot = dsymptr(s, ot, s, ot+2*(widthptr+2*4));
		n = t-&gt;thistuple + t-&gt;intuple;
		ot = duint32(s, ot, n);
		ot = duint32(s, ot, n);
		ot = dsymptr(s, ot, s, ot+1*(widthptr+2*4)+n*widthptr);
		ot = duint32(s, ot, t-&gt;outtuple);
		ot = duint32(s, ot, t-&gt;outtuple);

		// slice data
		for(t1=getthisx(t)-&gt;type; t1; t1=t1-&gt;down, n++)
			ot = dsymptr(s, ot, dtypesym(t1-&gt;type), 0);
		for(t1=getinargx(t)-&gt;type; t1; t1=t1-&gt;down, n++)
			ot = dsymptr(s, ot, dtypesym(t1-&gt;type), 0);
		for(t1=getoutargx(t)-&gt;type; t1; t1=t1-&gt;down, n++)
			ot = dsymptr(s, ot, dtypesym(t1-&gt;type), 0);
		break;

	case TINTER:
		m = imethods(t);
		n = 0;
		for(a=m; a; a=a-&gt;link) {
			dtypesym(a-&gt;type);
			n++;
		}

		// ../../pkg/runtime/type.go:/InterfaceType
		ot = dcommontype(s, ot, t);
		ot = dsymptr(s, ot, s, ot+widthptr+2*4);
		ot = duint32(s, ot, n);
		ot = duint32(s, ot, n);
		for(a=m; a; a=a-&gt;link) {
			// ../../pkg/runtime/type.go:/imethod
			ot = duint32(s, ot, a-&gt;hash);
			ot = duint32(s, ot, a-&gt;perm);
			ot = dgostringptr(s, ot, a-&gt;name);
			ot = dgostringptr(s, ot, a-&gt;package);
			ot = dsymptr(s, ot, dtypesym(a-&gt;type), 0);
		}
		break;

	case TMAP:
		// ../../pkg/runtime/type.go:/MapType
		s1 = dtypesym(t-&gt;down);
		s2 = dtypesym(t-&gt;type);
		ot = dcommontype(s, ot, t);
		ot = dsymptr(s, ot, s1, 0);
		ot = dsymptr(s, ot, s2, 0);
		break;

	case TPTR32:
	case TPTR64:
		if(t-&gt;type-&gt;etype == TANY) {
			ot = dcommontype(s, ot, t);
			break;
		}
		// ../../pkg/runtime/type.go:/PtrType
		s1 = dtypesym(t-&gt;type);
		ot = dcommontype(s, ot, t);
		ot = dsymptr(s, ot, s1, 0);
		break;

	case TSTRUCT:
		// ../../pkg/runtime/type.go:/StructType
		// for security, only the exported fields.
		n = 0;
		for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down) {
			dtypesym(t1-&gt;type);
			n++;
		}
		ot = dcommontype(s, ot, t);
		ot = dsymptr(s, ot, s, ot+widthptr+2*4);
		ot = duint32(s, ot, n);
		ot = duint32(s, ot, n);
		for(t1=t-&gt;type; t1!=T; t1=t1-&gt;down) {
			// ../../pkg/runtime/type.go:/structField
			if(t1-&gt;sym &amp;&amp; !t1-&gt;embedded) {
				ot = dgostringptr(s, ot, t1-&gt;sym-&gt;name);
				if(exportname(t1-&gt;sym-&gt;name))
					ot = dgostringptr(s, ot, nil);
				else
					ot = dgostringptr(s, ot, t1-&gt;sym-&gt;package);
			} else {
				ot = dgostringptr(s, ot, nil);
				ot = dgostringptr(s, ot, nil);
			}
			ot = dsymptr(s, ot, dtypesym(t1-&gt;type), 0);
			ot = dgostrlitptr(s, ot, t1-&gt;note);
			ot = duintptr(s, ot, t1-&gt;width);	// field offset
		}
		break;
	}

	ggloblsym(s, ot, tsym == nil);
	return s;
}

void
dumptypestructs(void)
{
	int i;
	NodeList *l;
	Node *n;
	Type *t;

	// copy types from externdcl list to signatlist
	for(l=externdcl; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op != OTYPE)
			continue;
		signatlist = list(signatlist, n);
	}

	// process signatlist
	for(l=signatlist; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op != OTYPE)
			continue;
		t = n-&gt;type;
		dtypesym(t);
		if(t-&gt;sym &amp;&amp; !isptr[t-&gt;etype])
			dtypesym(ptrto(t));
	}

	// do basic types if compiling package runtime.
	// they have to be in at least one package,
	// and reflect is always loaded implicitly,
	// so this is as good as any.
	// another possible choice would be package main,
	// but using runtime means fewer copies in .6 files.
	if(strcmp(package, &#34;runtime&#34;) == 0) {
		for(i=1; i&lt;=TBOOL; i++)
			dtypesym(ptrto(types[i]));
		dtypesym(ptrto(types[TSTRING]));
		dtypesym(typ(TDDD));
		dtypesym(ptrto(pkglookup(&#34;Pointer&#34;, &#34;unsafe&#34;)-&gt;def-&gt;type));
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
