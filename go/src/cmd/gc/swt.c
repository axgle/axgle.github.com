<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/swt.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/swt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include	&#34;go.h&#34;

enum
{
	Snorm		= 0,
	Strue,
	Sfalse,
	Stype,

	Tdefault,	// default case
	Texprconst,	// normal constant case
	Texprvar,	// normal variable case
	Ttypenil,	// case nil
	Ttypeconst,	// type hashes
	Ttypevar,	// interface type

	Ncase	= 4,	// count needed to split
};

typedef	struct	Case	Case;
struct	Case
{
	Node*	node;		// points at case statement
	uint32	hash;		// hash of a type switch
	uint8	type;		// type of case
	uint8	diag;		// suppress multiple diagnostics
	uint16	ordinal;	// position in switch
	Case*	link;		// linked list to link
};
#define	C	((Case*)nil)

Type*
notideal(Type *t)
{
	if(t != T &amp;&amp; t-&gt;etype == TIDEAL)
		return T;
	return t;
}

void
dumpcase(Case *c0)
{
	Case *c;

	for(c=c0; c!=C; c=c-&gt;link) {
		switch(c-&gt;type) {
		case Tdefault:
			print(&#34;case-default\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			break;
		case Texprconst:
			print(&#34;case-exprconst\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			break;
		case Texprvar:
			print(&#34;case-exprvar\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			print(&#34;	op=%O\n&#34;, c-&gt;node-&gt;left-&gt;op);
			break;
		case Ttypenil:
			print(&#34;case-typenil\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			break;
		case Ttypeconst:
			print(&#34;case-typeconst\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			print(&#34;	hash=%ux\n&#34;, c-&gt;hash);
			break;
		case Ttypevar:
			print(&#34;case-typevar\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			break;
		default:
			print(&#34;case-???\n&#34;);
			print(&#34;	ord=%d\n&#34;, c-&gt;ordinal);
			print(&#34;	op=%O\n&#34;, c-&gt;node-&gt;left-&gt;op);
			print(&#34;	hash=%ux\n&#34;, c-&gt;hash);
			break;
		}
	}
	print(&#34;\n&#34;);
}

static int
ordlcmp(Case *c1, Case *c2)
{
	// sort default first
	if(c1-&gt;type == Tdefault)
		return -1;
	if(c2-&gt;type == Tdefault)
		return +1;

	// sort nil second
	if(c1-&gt;type == Ttypenil)
		return -1;
	if(c2-&gt;type == Ttypenil)
		return +1;

	// sort by ordinal
	if(c1-&gt;ordinal &gt; c2-&gt;ordinal)
		return +1;
	if(c1-&gt;ordinal &lt; c2-&gt;ordinal)
		return -1;
	return 0;
}

static int
exprcmp(Case *c1, Case *c2)
{
	int ct, n;
	Node *n1, *n2;

	// sort non-constants last
	if(c1-&gt;type != Texprconst)
		return +1;
	if(c2-&gt;type != Texprconst)
		return -1;

	n1 = c1-&gt;node-&gt;left;
	n2 = c2-&gt;node-&gt;left;

	ct = n1-&gt;val.ctype;
	if(ct != n2-&gt;val.ctype) {
		// invalid program, but return a sort
		// order so that we can give a better
		// error later.
		return ct - n2-&gt;val.ctype;
	}

	// sort by constant value
	n = 0;
	switch(ct) {
	case CTFLT:
		n = mpcmpfltflt(n1-&gt;val.u.fval, n2-&gt;val.u.fval);
		break;
	case CTINT:
		n = mpcmpfixfix(n1-&gt;val.u.xval, n2-&gt;val.u.xval);
		break;
	case CTSTR:
		n = cmpslit(n1, n2);
		break;
	}

	return n;
}

static int
typecmp(Case *c1, Case *c2)
{

	// sort non-constants last
	if(c1-&gt;type != Ttypeconst)
		return +1;
	if(c2-&gt;type != Ttypeconst)
		return -1;

	// sort by hash code
	if(c1-&gt;hash &gt; c2-&gt;hash)
		return +1;
	if(c1-&gt;hash &lt; c2-&gt;hash)
		return -1;
	return 0;
}

static Case*
csort(Case *l, int(*f)(Case*, Case*))
{
	Case *l1, *l2, *le;

	if(l == C || l-&gt;link == C)
		return l;

	l1 = l;
	l2 = l;
	for(;;) {
		l2 = l2-&gt;link;
		if(l2 == C)
			break;
		l2 = l2-&gt;link;
		if(l2 == C)
			break;
		l1 = l1-&gt;link;
	}

	l2 = l1-&gt;link;
	l1-&gt;link = C;
	l1 = csort(l, f);
	l2 = csort(l2, f);

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
		if(l1 == C) {
			while(l2) {
				le-&gt;link = l2;
				le = l2;
				l2 = l2-&gt;link;
			}
			le-&gt;link = C;
			break;
		}
		if(l2 == C) {
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
	le-&gt;link = C;
	return l;
}

Node*
newlabel(void)
{
	static int label;

	label++;
	snprint(namebuf, sizeof(namebuf), &#34;%.6d&#34;, label);
	return newname(lookup(namebuf));
}

/*
 * build separate list of statements and cases
 * make labels between cases and statements
 * deal with fallthrough, break, unreachable statements
 */
void
casebody(Node *sw, Node *typeswvar)
{
	Node *os, *oc, *n, *c, *last;
	Node *def;
	NodeList *cas, *stat, *l, *lc;
	Node *go, *br;
	int32 lno, needvar;

	lno = setlineno(sw);
	if(sw-&gt;list == nil)
		return;

	cas = nil;	// cases
	stat = nil;	// statements
	def = N;	// defaults
	os = N;		// last statement
	oc = N;		// last case
	br = nod(OBREAK, N, N);

	for(l=sw-&gt;list; l; l=l-&gt;next) {
		n = l-&gt;n;
		lno = setlineno(n);
		if(n-&gt;op != OXCASE)
			fatal(&#34;casebody %O&#34;, n-&gt;op);
		n-&gt;op = OCASE;
		needvar = count(n-&gt;list) != 1 || n-&gt;list-&gt;n-&gt;op == OLITERAL;

		go = nod(OGOTO, newlabel(), N);
		if(n-&gt;list == nil) {
			if(def != N)
				yyerror(&#34;more than one default case&#34;);
			// reuse original default case
			n-&gt;right = go;
			def = n;
		}

		if(n-&gt;list != nil &amp;&amp; n-&gt;list-&gt;next == nil) {
			// one case - reuse OCASE node.
			c = n-&gt;list-&gt;n;
			n-&gt;left = c;
			n-&gt;right = go;
			n-&gt;list = nil;
			cas = list(cas, n);
		} else {
			// expand multi-valued cases
			for(lc=n-&gt;list; lc; lc=lc-&gt;next) {
				c = lc-&gt;n;
				cas = list(cas, nod(OCASE, c, go));
			}
		}

		stat = list(stat, nod(OLABEL, go-&gt;left, N));
		if(typeswvar &amp;&amp; needvar &amp;&amp; n-&gt;nname != N) {
			NodeList *l;

			l = list1(nod(ODCL, n-&gt;nname, N));
			l = list(l, nod(OAS, n-&gt;nname, typeswvar));
			typechecklist(l, Etop);
			stat = concat(stat, l);
		}
		stat = concat(stat, n-&gt;nbody);

		// botch - shouldnt fall thru declaration
		last = stat-&gt;end-&gt;n;
		if(last-&gt;op == OXFALL) {
			if(typeswvar) {
				setlineno(last);
				yyerror(&#34;cannot fallthrough in type switch&#34;);
			}
			last-&gt;op = OFALL;
		} else
			stat = list(stat, br);
	}

	stat = list(stat, br);
	if(def)
		cas = list(cas, def);

	sw-&gt;list = cas;
	sw-&gt;nbody = stat;
	lineno = lno;
}

Case*
mkcaselist(Node *sw, int arg)
{
	Node *n;
	Case *c, *c1;
	NodeList *l;
	int ord;

	c = C;
	ord = 0;

	for(l=sw-&gt;list; l; l=l-&gt;next) {
		n = l-&gt;n;
		c1 = mal(sizeof(*c1));
		c1-&gt;link = c;
		c = c1;

		ord++;
		c-&gt;ordinal = ord;
		c-&gt;node = n;

		if(n-&gt;left == N) {
			c-&gt;type = Tdefault;
			continue;
		}

		switch(arg) {
		case Stype:
			c-&gt;hash = 0;
			if(n-&gt;left-&gt;op == OLITERAL) {
				c-&gt;type = Ttypenil;
				continue;
			}
			if(istype(n-&gt;left-&gt;type, TINTER)) {
				c-&gt;type = Ttypevar;
				continue;
			}

			c-&gt;hash = typehash(n-&gt;left-&gt;type);
			c-&gt;type = Ttypeconst;
			continue;

		case Snorm:
		case Strue:
		case Sfalse:
			c-&gt;type = Texprvar;
			switch(consttype(n-&gt;left)) {
			case CTFLT:
			case CTINT:
			case CTSTR:
				c-&gt;type = Texprconst;
			}
			continue;
		}
	}

	if(c == C)
		return C;

	// sort by value and diagnose duplicate cases
	switch(arg) {
	case Stype:
		c = csort(c, typecmp);
		for(c1=c; c1-&gt;link!=C; c1=c1-&gt;link) {
			if(typecmp(c1, c1-&gt;link) != 0)
				continue;
			setlineno(c1-&gt;link-&gt;node);
			yyerror(&#34;duplicate case in switch\n\tprevious case at %L&#34;, c1-&gt;node-&gt;lineno);
		}
		break;
	case Snorm:
	case Strue:
	case Sfalse:
		c = csort(c, exprcmp);
		for(c1=c; c1-&gt;link!=C; c1=c1-&gt;link) {
			if(exprcmp(c1, c1-&gt;link) != 0)
				continue;
			setlineno(c1-&gt;link-&gt;node);
			yyerror(&#34;duplicate case in switch\n\tprevious case at %L&#34;, c1-&gt;node-&gt;lineno);
		}
		break;
	}

	// put list back in processing order
	c = csort(c, ordlcmp);
	return c;
}

static	Node*	exprname;

Node*
exprbsw(Case *c0, int ncase, int arg)
{
	NodeList *cas;
	Node *a, *n;
	Case *c;
	int i, half, lno;

	cas = nil;
	if(ncase &lt; Ncase) {
		for(i=0; i&lt;ncase; i++) {
			n = c0-&gt;node;
			lno = setlineno(n);

			switch(arg) {
			case Strue:
				a = nod(OIF, N, N);
				a-&gt;ntest = n-&gt;left;			// if val
				a-&gt;nbody = list1(n-&gt;right);			// then goto l
				break;

			case Sfalse:
				a = nod(OIF, N, N);
				a-&gt;ntest = nod(ONOT, n-&gt;left, N);	// if !val
				typecheck(&amp;a-&gt;ntest, Erv);
				a-&gt;nbody = list1(n-&gt;right);			// then goto l
				break;

			default:
				a = nod(OIF, N, N);
				a-&gt;ntest = nod(OEQ, exprname, n-&gt;left);	// if name == val
				typecheck(&amp;a-&gt;ntest, Erv);
				a-&gt;nbody = list1(n-&gt;right);			// then goto l
				break;
			}

			cas = list(cas, a);
			c0 = c0-&gt;link;
			lineno = lno;
		}
		return liststmt(cas);
	}

	// find the middle and recur
	c = c0;
	half = ncase&gt;&gt;1;
	for(i=1; i&lt;half; i++)
		c = c-&gt;link;
	a = nod(OIF, N, N);
	a-&gt;ntest = nod(OLE, exprname, c-&gt;node-&gt;left);
	typecheck(&amp;a-&gt;ntest, Erv);
	a-&gt;nbody = list1(exprbsw(c0, half, arg));
	a-&gt;nelse = list1(exprbsw(c-&gt;link, ncase-half, arg));
	return a;
}

/*
 * normal (expression) switch.
 * rebulid case statements into if .. goto
 */
void
exprswitch(Node *sw)
{
	Node *def;
	NodeList *cas;
	Node *a;
	Case *c0, *c, *c1;
	Type *t;
	int arg, ncase;

	casebody(sw, N);

	arg = Snorm;
	if(isconst(sw-&gt;ntest, CTBOOL)) {
		arg = Strue;
		if(sw-&gt;ntest-&gt;val.u.bval == 0)
			arg = Sfalse;
	}
	walkexpr(&amp;sw-&gt;ntest, &amp;sw-&gt;ninit);
	t = sw-&gt;type;
	if(t == T)
		return;

	/*
	 * convert the switch into OIF statements
	 */
	exprname = N;
	cas = nil;
	if(arg != Strue &amp;&amp; arg != Sfalse) {
		exprname = nod(OXXX, N, N);
		tempname(exprname, sw-&gt;ntest-&gt;type);
		cas = list1(nod(OAS, exprname, sw-&gt;ntest));
		typechecklist(cas, Etop);
	}

	c0 = mkcaselist(sw, arg);
	if(c0 != C &amp;&amp; c0-&gt;type == Tdefault) {
		def = c0-&gt;node-&gt;right;
		c0 = c0-&gt;link;
	} else {
		def = nod(OBREAK, N, N);
	}

loop:
	if(c0 == C) {
		cas = list(cas, def);
		sw-&gt;nbody = concat(cas, sw-&gt;nbody);
		sw-&gt;list = nil;
		walkstmtlist(sw-&gt;nbody);
		return;
	}

	// deal with the variables one-at-a-time
	if(c0-&gt;type != Texprconst) {
		a = exprbsw(c0, 1, arg);
		cas = list(cas, a);
		c0 = c0-&gt;link;
		goto loop;
	}

	// do binary search on run of constants
	ncase = 1;
	for(c=c0; c-&gt;link!=C; c=c-&gt;link) {
		if(c-&gt;link-&gt;type != Texprconst)
			break;
		ncase++;
	}

	// break the chain at the count
	c1 = c-&gt;link;
	c-&gt;link = C;

	// sort and compile constants
	c0 = csort(c0, exprcmp);
	a = exprbsw(c0, ncase, arg);
	cas = list(cas, a);

	c0 = c1;
	goto loop;

}

static	Node*	hashname;
static	Node*	facename;
static	Node*	boolname;

Node*
typeone(Node *t)
{
	NodeList *init;
	Node *a, *b, *var;

	var = t-&gt;nname;
	init = nil;
	if(var == N) {
		typecheck(&amp;nblank, Erv | Easgn);
		var = nblank;
	} else
		init = list1(nod(ODCL, var, N));

	a = nod(OAS2, N, N);
	a-&gt;list = list(list1(var), boolname);	// var,bool =
	b = nod(ODOTTYPE, facename, N);
	b-&gt;type = t-&gt;left-&gt;type;		// interface.(type)
	a-&gt;rlist = list1(b);
	typecheck(&amp;a, Etop);
	init = list(init, a);

	b = nod(OIF, N, N);
	b-&gt;ntest = boolname;
	b-&gt;nbody = list1(t-&gt;right);		// if bool { goto l }
	a = liststmt(list(init, b));
	return a;
}

Node*
typebsw(Case *c0, int ncase)
{
	NodeList *cas;
	Node *a, *n;
	Case *c;
	int i, half;
	Val v;

	cas = nil;

	if(ncase &lt; Ncase) {
		for(i=0; i&lt;ncase; i++) {
			n = c0-&gt;node;

			switch(c0-&gt;type) {

			case Ttypenil:
				v.ctype = CTNIL;
				a = nod(OIF, N, N);
				a-&gt;ntest = nod(OEQ, facename, nodlit(v));
				typecheck(&amp;a-&gt;ntest, Erv);
				a-&gt;nbody = list1(n-&gt;right);		// if i==nil { goto l }
				cas = list(cas, a);
				break;

			case Ttypevar:
				a = typeone(n);
				cas = list(cas, a);
				break;

			case Ttypeconst:
				a = nod(OIF, N, N);
				a-&gt;ntest = nod(OEQ, hashname, nodintconst(c0-&gt;hash));
				typecheck(&amp;a-&gt;ntest, Erv);
				a-&gt;nbody = list1(typeone(n));
				cas = list(cas, a);
				break;
			}
			c0 = c0-&gt;link;
		}
		return liststmt(cas);
	}

	// find the middle and recur
	c = c0;
	half = ncase&gt;&gt;1;
	for(i=1; i&lt;half; i++)
		c = c-&gt;link;
	a = nod(OIF, N, N);
	a-&gt;ntest = nod(OLE, hashname, nodintconst(c-&gt;hash));
	typecheck(&amp;a-&gt;ntest, Erv);
	a-&gt;nbody = list1(typebsw(c0, half));
	a-&gt;nelse = list1(typebsw(c-&gt;link, ncase-half));
	return a;
}

/*
 * convert switch of the form
 *	switch v := i.(type) { case t1: ..; case t2: ..; }
 * into if statements
 */
void
typeswitch(Node *sw)
{
	Node *def;
	NodeList *cas;
	Node *a;
	Case *c, *c0, *c1;
	int ncase;
	Type *t;

	if(sw-&gt;ntest == nil)
		return;
	if(sw-&gt;ntest-&gt;right == nil) {
		setlineno(sw);
		yyerror(&#34;type switch must have an assignment&#34;);
		return;
	}
	walkexpr(&amp;sw-&gt;ntest-&gt;right, &amp;sw-&gt;ninit);
	if(!istype(sw-&gt;ntest-&gt;right-&gt;type, TINTER)) {
		yyerror(&#34;type switch must be on an interface&#34;);
		return;
	}
	cas = nil;

	/*
	 * predeclare temporary variables
	 * and the boolean var
	 */
	facename = nod(OXXX, N, N);
	tempname(facename, sw-&gt;ntest-&gt;right-&gt;type);
	a = nod(OAS, facename, sw-&gt;ntest-&gt;right);
	typecheck(&amp;a, Etop);
	cas = list(cas, a);

	casebody(sw, facename);

	boolname = nod(OXXX, N, N);
	tempname(boolname, types[TBOOL]);
	typecheck(&amp;boolname, Erv);

	hashname = nod(OXXX, N, N);
	tempname(hashname, types[TUINT32]);
	typecheck(&amp;hashname, Erv);

	t = sw-&gt;ntest-&gt;right-&gt;type;
	if(isnilinter(t))
		a = syslook(&#34;efacethash&#34;, 1);
	else
		a = syslook(&#34;ifacethash&#34;, 1);
	argtype(a, t);
	a = nod(OCALL, a, N);
	a-&gt;list = list1(facename);
	a = nod(OAS, hashname, a);
	typecheck(&amp;a, Etop);
	cas = list(cas, a);

	c0 = mkcaselist(sw, Stype);
	if(c0 != C &amp;&amp; c0-&gt;type == Tdefault) {
		def = c0-&gt;node-&gt;right;
		c0 = c0-&gt;link;
	} else {
		def = nod(OBREAK, N, N);
	}

loop:
	if(c0 == C) {
		cas = list(cas, def);
		sw-&gt;nbody = concat(cas, sw-&gt;nbody);
		sw-&gt;list = nil;
		walkstmtlist(sw-&gt;nbody);
		return;
	}

	// deal with the variables one-at-a-time
	if(c0-&gt;type != Ttypeconst) {
		a = typebsw(c0, 1);
		cas = list(cas, a);
		c0 = c0-&gt;link;
		goto loop;
	}

	// do binary search on run of constants
	ncase = 1;
	for(c=c0; c-&gt;link!=C; c=c-&gt;link) {
		if(c-&gt;link-&gt;type != Ttypeconst)
			break;
		ncase++;
	}

	// break the chain at the count
	c1 = c-&gt;link;
	c-&gt;link = C;

	// sort and compile constants
	c0 = csort(c0, typecmp);
	a = typebsw(c0, ncase);
	cas = list(cas, a);

	c0 = c1;
	goto loop;
}

void
walkswitch(Node *sw)
{

	/*
	 * reorder the body into (OLIST, cases, statements)
	 * cases have OGOTO into statements.
	 * both have inserted OBREAK statements
	 */
	walkstmtlist(sw-&gt;ninit);
	if(sw-&gt;ntest == N) {
		sw-&gt;ntest = nodbool(1);
		typecheck(&amp;sw-&gt;ntest, Erv);
	}

	if(sw-&gt;ntest-&gt;op == OTYPESW) {
		typeswitch(sw);
//dump(&#34;sw&#34;, sw);
		return;
	}
	exprswitch(sw);
}

/*
 * type check switch statement
 */
void
typecheckswitch(Node *n)
{
	int top, lno;
	Type *t;
	NodeList *l, *ll;
	Node *ncase, *nvar;
	Node *def;

	lno = lineno;
	typechecklist(n-&gt;ninit, Etop);

	if(n-&gt;ntest != N &amp;&amp; n-&gt;ntest-&gt;op == OTYPESW) {
		// type switch
		top = Etype;
		typecheck(&amp;n-&gt;ntest-&gt;right, Erv);
		t = n-&gt;ntest-&gt;right-&gt;type;
		if(t != T &amp;&amp; t-&gt;etype != TINTER)
			yyerror(&#34;cannot type switch on non-interface value %+N&#34;, n-&gt;ntest-&gt;right);
	} else {
		// value switch
		top = Erv;
		if(n-&gt;ntest) {
			typecheck(&amp;n-&gt;ntest, Erv);
			defaultlit(&amp;n-&gt;ntest, T);
			t = n-&gt;ntest-&gt;type;
		} else
			t = types[TBOOL];
	}
	n-&gt;type = t;

	def = N;
	for(l=n-&gt;list; l; l=l-&gt;next) {
		ncase = l-&gt;n;
		setlineno(n);
		if(ncase-&gt;list == nil) {
			// default
			if(def != N)
				yyerror(&#34;multiple defaults in switch (first at %L)&#34;, def-&gt;lineno);
			else
				def = ncase;
		} else {
			for(ll=ncase-&gt;list; ll; ll=ll-&gt;next) {
				setlineno(ll-&gt;n);
				typecheck(&amp;ll-&gt;n, Erv | Etype);
				if(ll-&gt;n-&gt;type == T || t == T)
					continue;
				switch(top) {
				case Erv:	// expression switch
					defaultlit(&amp;ll-&gt;n, t);
					if(ll-&gt;n-&gt;op == OTYPE)
						yyerror(&#34;type %T is not an expression&#34;, ll-&gt;n-&gt;type);
					else if(ll-&gt;n-&gt;type != T &amp;&amp; !eqtype(ll-&gt;n-&gt;type, t))
						yyerror(&#34;case %+N in %T switch&#34;, ll-&gt;n, t);
					break;
				case Etype:	// type switch
					if(ll-&gt;n-&gt;op == OLITERAL &amp;&amp; istype(ll-&gt;n-&gt;type, TNIL))
						;
					else if(ll-&gt;n-&gt;op != OTYPE &amp;&amp; ll-&gt;n-&gt;type != T)
						yyerror(&#34;%#N is not a type&#34;, ll-&gt;n);
					break;
				}
			}
		}
		if(top == Etype &amp;&amp; n-&gt;type != T) {
			ll = ncase-&gt;list;
			nvar = ncase-&gt;nname;
			if(nvar != N) {
				if(ll &amp;&amp; ll-&gt;next == nil &amp;&amp; ll-&gt;n-&gt;type != T &amp;&amp; !istype(ll-&gt;n-&gt;type, TNIL)) {
					// single entry type switch
					nvar-&gt;ntype = typenod(ll-&gt;n-&gt;type);
				} else {
					// multiple entry type switch or default
					nvar-&gt;ntype = typenod(n-&gt;type);
				}
			}
		}
		typechecklist(ncase-&gt;nbody, Etop);
	}

	lineno = lno;
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
