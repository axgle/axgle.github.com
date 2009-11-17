<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/typecheck.c</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/gc/typecheck.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * type check the whole tree of an expression.
 * calculates expression types.
 * evaluates compile time constants.
 * marks variables that escape the local frame.
 * rewrites n-&gt;op to be more specific in some cases.
 * sets n-&gt;walk to walking function.
 *
 * TODO:
 *	trailing ... section of function calls
 */

#include &#34;go.h&#34;

static void	implicitstar(Node**);
static int	onearg(Node*);
static int	lookdot(Node*, Type*);
static void	typecheckaste(int, Type*, NodeList*);
static int	exportassignok(Type*);
static Type*	lookdot1(Sym *s, Type *t, Type *f);
static int	nokeys(NodeList*);
static void	typecheckcomplit(Node**);
static void	addrescapes(Node*);
static void	typecheckas2(Node*);
static void	typecheckas(Node*);
static void	typecheckfunc(Node*);
static void	checklvalue(Node*, char*);
static void	checkassign(Node*);
static void	checkassignlist(NodeList*);
static int	islvalue(Node*);

void
typechecklist(NodeList *l, int top)
{
	for(; l; l=l-&gt;next)
		typecheck(&amp;l-&gt;n, top);
}

/*
 * type check node *np.
 * replaces *np with a new pointer in some cases.
 * returns the final value of *np as a convenience.
 */
Node*
typecheck(Node **np, int top)
{
	int et, op;
	Node *n, *l, *r;
	NodeList *args;
	int lno, ok, ntop;
	Type *t;

	// cannot type check until all the source has been parsed
	if(!typecheckok)
		fatal(&#34;early typecheck&#34;);

	n = *np;
	if(n == N)
		return N;

	// Skip typecheck if already done.
	// But re-typecheck ONAME/OTYPE/OLITERAL/OPACK node in case context has changed.
	if(n-&gt;typecheck == 1) {
		switch(n-&gt;op) {
		case ONAME:
		case OTYPE:
		case OLITERAL:
		case OPACK:
			break;
		default:
			return n;
		}
	}

	if(n-&gt;typecheck == 2)
		fatal(&#34;typecheck loop&#34;);
	n-&gt;typecheck = 2;

redo:
	lno = setlineno(n);
	if(n-&gt;sym) {
		walkdef(n);
		if(n-&gt;op == ONONAME)
			goto error;
	}

reswitch:
	ok = 0;
	switch(n-&gt;op) {
	default:
		// until typecheck is complete, do nothing.
		dump(&#34;typecheck&#34;, n);
		fatal(&#34;typecheck %O&#34;, n-&gt;op);

	/*
	 * names
	 */
	case OLITERAL:
		ok |= Erv;
		if(n-&gt;val.ctype == CTSTR)
			n-&gt;type = idealstring;
		goto ret;

	case ONONAME:
		ok |= Erv;
		goto ret;

	case ONAME:
		if(n-&gt;etype != 0) {
			ok |= Ecall;
			goto ret;
		}
		if(!(top &amp; Easgn)) {
			// not a write to the variable
			if(isblank(n)) {
				yyerror(&#34;cannot use _ as value&#34;);
				goto error;
			}
			n-&gt;used = 1;
		}
		ok |= Erv;
		goto ret;

	case OPACK:
		yyerror(&#34;use of package %S not in selector&#34;, n-&gt;sym);
		goto error;

	case OIOTA:
		// looked like iota during parsing but might
		// have been redefined.  decide.
		if(n-&gt;left-&gt;op != ONONAME)
			n = n-&gt;left;
		else
			n = n-&gt;right;
		goto redo;

	/*
	 * types (OIND is with exprs)
	 */
	case OTYPE:
		ok |= Etype;
		if(n-&gt;type == T)
			goto error;
		break;

	case OTARRAY:
		ok |= Etype;
		t = typ(TARRAY);
		l = n-&gt;left;
		r = n-&gt;right;
		if(l == nil) {
			t-&gt;bound = -1;
		} else {
			typecheck(&amp;l, Erv | Etype);
			switch(l-&gt;op) {
			default:
				yyerror(&#34;invalid array bound %#N&#34;, l);
				goto error;

			case OLITERAL:
				if(consttype(l) == CTINT) {
					t-&gt;bound = mpgetfix(l-&gt;val.u.xval);
					if(t-&gt;bound &lt; 0) {
						yyerror(&#34;array bound must be non-negative&#34;);
						goto error;
					}
				}
				break;

			case OTYPE:
				if(l-&gt;type == T)
					goto error;
				if(l-&gt;type-&gt;etype != TDDD) {
					yyerror(&#34;invalid array bound %T&#34;, l-&gt;type);
					goto error;
				}
				t-&gt;bound = -100;
				break;
			}
		}
		typecheck(&amp;r, Etype);
		if(r-&gt;type == T)
			goto error;
		t-&gt;type = r-&gt;type;
		n-&gt;op = OTYPE;
		n-&gt;type = t;
		n-&gt;left = N;
		n-&gt;right = N;
		if(t-&gt;bound != -100)
			checkwidth(t);
		break;

	case OTMAP:
		ok |= Etype;
		l = typecheck(&amp;n-&gt;left, Etype);
		r = typecheck(&amp;n-&gt;right, Etype);
		if(l-&gt;type == T || r-&gt;type == T)
			goto error;
		n-&gt;op = OTYPE;
		n-&gt;type = maptype(l-&gt;type, r-&gt;type);
		n-&gt;left = N;
		n-&gt;right = N;
		break;

	case OTCHAN:
		ok |= Etype;
		l = typecheck(&amp;n-&gt;left, Etype);
		if(l-&gt;type == T)
			goto error;
		t = typ(TCHAN);
		t-&gt;type = l-&gt;type;
		t-&gt;chan = n-&gt;etype;
		n-&gt;op = OTYPE;
		n-&gt;type = t;
		n-&gt;left = N;
		n-&gt;etype = 0;
		break;

	case OTSTRUCT:
		ok |= Etype;
		n-&gt;op = OTYPE;
		n-&gt;type = dostruct(n-&gt;list, TSTRUCT);
		if(n-&gt;type == T)
			goto error;
		n-&gt;list = nil;
		break;

	case OTINTER:
		ok |= Etype;
		n-&gt;op = OTYPE;
		n-&gt;type = dostruct(n-&gt;list, TINTER);
		if(n-&gt;type == T)
			goto error;
		n-&gt;type = sortinter(n-&gt;type);
		break;

	case OTFUNC:
		ok |= Etype;
		n-&gt;op = OTYPE;
		n-&gt;type = functype(n-&gt;left, n-&gt;list, n-&gt;rlist);
		if(n-&gt;type == T)
			goto error;
		break;

	/*
	 * type or expr
	 */
	case OIND:
		ntop = Erv | Etype;
		if(!(top &amp; Eaddr))
			ntop |= Eindir;
		l = typecheck(&amp;n-&gt;left, ntop);
		if((t = l-&gt;type) == T)
			goto error;
		if(l-&gt;op == OTYPE) {
			ok |= Etype;
			n-&gt;op = OTYPE;
			n-&gt;type = ptrto(l-&gt;type);
			n-&gt;left = N;
			goto ret;
		}
		if(!isptr[t-&gt;etype]) {
			yyerror(&#34;invalid indirect of %+N&#34;, n-&gt;left);
			goto error;
		}
		ok |= Erv;
		n-&gt;type = t-&gt;type;
		goto ret;

	/*
	 * arithmetic exprs
	 */
	case OASOP:
		ok |= Etop;
		l = typecheck(&amp;n-&gt;left, Erv);
		checkassign(n-&gt;left);
		r = typecheck(&amp;n-&gt;right, Erv);
		if(l-&gt;type == T || r-&gt;type == T)
			goto error;
		op = n-&gt;etype;
		goto arith;

	case OADD:
	case OAND:
	case OANDAND:
	case OANDNOT:
	case ODIV:
	case OEQ:
	case OGE:
	case OGT:
	case OLE:
	case OLT:
	case OLSH:
	case ORSH:
	case OMOD:
	case OMUL:
	case ONE:
	case OOR:
	case OOROR:
	case OSUB:
	case OXOR:
		ok |= Erv;
		l = typecheck(&amp;n-&gt;left, Erv | (top &amp; Eiota));
		r = typecheck(&amp;n-&gt;right, Erv | (top &amp; Eiota));
		if(l-&gt;type == T || r-&gt;type == T)
			goto error;
		op = n-&gt;op;
	arith:
		if(op == OLSH || op == ORSH)
			goto shift;
		// ideal mixed with non-ideal
		defaultlit2(&amp;l, &amp;r, 0);
		n-&gt;left = l;
		n-&gt;right = r;
		if(l-&gt;type == T || r-&gt;type == T)
			goto error;
		t = l-&gt;type;
		if(t-&gt;etype == TIDEAL)
			t = r-&gt;type;
		et = t-&gt;etype;
		if(et == TIDEAL)
			et = TINT;
		if(t-&gt;etype != TIDEAL &amp;&amp; !eqtype(l-&gt;type, r-&gt;type)) {
		badbinary:
			defaultlit2(&amp;l, &amp;r, 1);
			yyerror(&#34;invalid operation: %#N (type %T %#O %T)&#34;, n, l-&gt;type, op, r-&gt;type);
			goto error;
		}
		if(!okfor[op][et])
			goto badbinary;
		// okfor allows any array == array;
		// restrict to slice == nil and nil == slice.
		if(l-&gt;type-&gt;etype == TARRAY &amp;&amp; !isslice(l-&gt;type))
			goto badbinary;
		if(r-&gt;type-&gt;etype == TARRAY &amp;&amp; !isslice(r-&gt;type))
			goto badbinary;
		if(isslice(l-&gt;type) &amp;&amp; !isnil(l) &amp;&amp; !isnil(r))
			goto badbinary;
		t = l-&gt;type;
		if(iscmp[n-&gt;op]) {
			evconst(n);
			t = types[TBOOL];
			if(n-&gt;op != OLITERAL) {
				defaultlit2(&amp;l, &amp;r, 1);
				n-&gt;left = l;
				n-&gt;right = r;
			}
		}
		if(et == TSTRING) {
			if(iscmp[n-&gt;op]) {
				n-&gt;etype = n-&gt;op;
				n-&gt;op = OCMPSTR;
			} else if(n-&gt;op == OASOP)
				n-&gt;op = OAPPENDSTR;
			else if(n-&gt;op == OADD)
				n-&gt;op = OADDSTR;
		}
		if(et == TINTER) {
			if(l-&gt;op == OLITERAL &amp;&amp; l-&gt;val.ctype == CTNIL) {
				// swap for back end
				n-&gt;left = r;
				n-&gt;right = l;
			} else if(r-&gt;op == OLITERAL &amp;&amp; r-&gt;val.ctype == CTNIL) {
				// leave alone for back end
			} else {
				n-&gt;etype = n-&gt;op;
				n-&gt;op = OCMPIFACE;
			}
		}
		n-&gt;type = t;
		goto ret;

	shift:
		defaultlit(&amp;r, types[TUINT]);
		n-&gt;right = r;
		t = r-&gt;type;
		if(!isint[t-&gt;etype] || issigned[t-&gt;etype]) {
			yyerror(&#34;invalid operation: %#N (shift count type %T)&#34;, n, r-&gt;type);
			goto error;
		}
		t = l-&gt;type;
		if(t != T &amp;&amp; t-&gt;etype != TIDEAL &amp;&amp; !isint[t-&gt;etype]) {
			yyerror(&#34;invalid operation: %#N (shift of type %T)&#34;, n, t);
			goto error;
		}
		// no defaultlit for left
		// the outer context gives the type
		n-&gt;type = l-&gt;type;
		goto ret;

	case OCOM:
	case OMINUS:
	case ONOT:
	case OPLUS:
		ok |= Erv;
		l = typecheck(&amp;n-&gt;left, Erv | (top &amp; Eiota));
		if((t = l-&gt;type) == T)
			goto error;
		if(!okfor[n-&gt;op][t-&gt;etype]) {
			yyerror(&#34;invalid operation: %#O %T&#34;, n-&gt;op, t);
			goto error;
		}
		n-&gt;type = t;
		goto ret;

	/*
	 * exprs
	 */
	case OADDR:
		ok |= Erv;
		typecheck(&amp;n-&gt;left, Erv | Eaddr);
		if(n-&gt;left-&gt;type == T)
			goto error;
		switch(n-&gt;left-&gt;op) {
		case OMAPLIT:
		case OSTRUCTLIT:
		case OARRAYLIT:
			break;
		default:
			checklvalue(n-&gt;left, &#34;take the address of&#34;);
		}
		defaultlit(&amp;n-&gt;left, T);
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		if(!(top &amp; Eindir))
			addrescapes(n-&gt;left);
		n-&gt;type = ptrto(t);
		goto ret;

	case OCOMPLIT:
		ok |= Erv;
		typecheckcomplit(&amp;n);
		if(n-&gt;type == T)
			goto error;
		goto ret;

	case OXDOT:
		n = adddot(n);
		n-&gt;op = ODOT;
		// fall through
	case ODOT:
		l = typecheck(&amp;n-&gt;left, Erv);
		if((t = l-&gt;type) == T)
			goto error;
		if(n-&gt;right-&gt;op != ONAME) {
			yyerror(&#34;rhs of . must be a name&#34;);	// impossible
			goto error;
		}
		if(isptr[t-&gt;etype]) {
			t = t-&gt;type;
			if(t == T)
				goto error;
			n-&gt;op = ODOTPTR;
			checkwidth(t);
		}
		if(!lookdot(n, t)) {
			yyerror(&#34;%#N undefined (type %T has no field %S)&#34;, n, t, n-&gt;right-&gt;sym);
			goto error;
		}
		switch(n-&gt;op) {
		case ODOTINTER:
		case ODOTMETH:
			ok |= Ecall;
			break;
		default:
			ok |= Erv;
			break;
		}
		goto ret;

	case ODOTTYPE:
		ok |= Erv;
		typecheck(&amp;n-&gt;left, Erv);
		defaultlit(&amp;n-&gt;left, T);
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		if(!isinter(t)) {
			yyerror(&#34;invalid type assertion: %#N (non-interface type %T on left)&#34;, n, t);
			goto error;
		}
		if(n-&gt;right != N) {
			typecheck(&amp;n-&gt;right, Etype);
			n-&gt;type = n-&gt;right-&gt;type;
			n-&gt;right = N;
			if(n-&gt;type == T)
				goto error;
		}
		goto ret;

	case OINDEX:
		ok |= Erv;
		typecheck(&amp;n-&gt;left, Erv);
		defaultlit(&amp;n-&gt;left, T);
		implicitstar(&amp;n-&gt;left);
		l = n-&gt;left;
		typecheck(&amp;n-&gt;right, Erv);
		r = n-&gt;right;
		if((t = l-&gt;type) == T || r-&gt;type == T)
			goto error;
		switch(t-&gt;etype) {
		default:
			yyerror(&#34;invalid operation: %#N (index of type %T)&#34;, n, t);
			goto error;

		case TARRAY:
			defaultlit(&amp;n-&gt;right, types[TUINT]);
			if(n-&gt;right-&gt;type != T &amp;&amp; !isint[n-&gt;right-&gt;type-&gt;etype])
				yyerror(&#34;non-integer array index %#N&#34;, n-&gt;right);
			n-&gt;type = t-&gt;type;
			break;

		case TMAP:
			n-&gt;etype = 0;
			defaultlit(&amp;n-&gt;right, t-&gt;down);
			if(n-&gt;right-&gt;type != T &amp;&amp; !eqtype(n-&gt;right-&gt;type, t-&gt;down))
				yyerror(&#34;invalid map index %#N - need type %T&#34;, n-&gt;right, t-&gt;down);
			n-&gt;type = t-&gt;type;
			n-&gt;op = OINDEXMAP;
			break;

		case TSTRING:
			defaultlit(&amp;n-&gt;right, types[TUINT]);
			if(n-&gt;right-&gt;type != T &amp;&amp; !isint[n-&gt;right-&gt;type-&gt;etype])
				yyerror(&#34;non-integer string index %#N&#34;, n-&gt;right);
			n-&gt;type = types[TUINT8];
			n-&gt;op = OINDEXSTR;
			break;
		}
		goto ret;

	case ORECV:
		ok |= Etop | Erv;
		typecheck(&amp;n-&gt;left, Erv);
		defaultlit(&amp;n-&gt;left, T);
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		if(t-&gt;etype != TCHAN) {
			yyerror(&#34;invalid operation: %#N (receive from non-chan type %T)&#34;, n, t);
			goto error;
		}
		if(!(t-&gt;chan &amp; Crecv)) {
			yyerror(&#34;invalid operation: %#N (receive from send-only type %T)&#34;, n, t);
			goto error;
		}
		n-&gt;type = t-&gt;type;
		goto ret;

	case OSEND:
		ok |= Etop | Erv;
		l = typecheck(&amp;n-&gt;left, Erv);
		typecheck(&amp;n-&gt;right, Erv);
		defaultlit(&amp;n-&gt;left, T);
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		if(!(t-&gt;chan &amp; Csend)) {
			yyerror(&#34;invalid operation: %#N (send to receive-only type %T)&#34;, n, t);
			goto error;
		}
		defaultlit(&amp;n-&gt;right, t-&gt;type);
		r = n-&gt;right;
		if((t = r-&gt;type) == T)
			goto error;
		// TODO: more aggressive
		n-&gt;etype = 0;
		n-&gt;type = T;
		if(top &amp; Erv) {
			n-&gt;op = OSENDNB;
			n-&gt;type = types[TBOOL];
		}
		goto ret;

	case OSLICE:
		ok |= Erv;
		typecheck(&amp;n-&gt;left, top);
		typecheck(&amp;n-&gt;right-&gt;left, Erv);
		typecheck(&amp;n-&gt;right-&gt;right, Erv);
		defaultlit(&amp;n-&gt;left, T);
		defaultlit(&amp;n-&gt;right-&gt;left, types[TUINT]);
		defaultlit(&amp;n-&gt;right-&gt;right, types[TUINT]);
		implicitstar(&amp;n-&gt;left);
		if(n-&gt;right-&gt;left == N || n-&gt;right-&gt;right == N) {
			yyerror(&#34;missing slice bounds?&#34;);
			goto error;
		}
		if((t = n-&gt;right-&gt;left-&gt;type) == T)
			goto error;
		if(!isint[t-&gt;etype]) {
			yyerror(&#34;invalid slice index %#N (type %T)&#34;, n-&gt;right-&gt;left, t);
			goto error;
		}
		if((t = n-&gt;right-&gt;right-&gt;type) == T)
			goto error;
		if(!isint[t-&gt;etype]) {
			yyerror(&#34;invalid slice index %#N (type %T)&#34;, n-&gt;right-&gt;right, t);
			goto error;
		}
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		// TODO(rsc): 64-bit slice index needs to be checked
		// for overflow in generated code
		if(istype(t, TSTRING)) {
			n-&gt;type = t;
			n-&gt;op = OSLICESTR;
			goto ret;
		}
		if(isfixedarray(t)) {
			n-&gt;type = typ(TARRAY);
			n-&gt;type-&gt;type = t-&gt;type;
			n-&gt;type-&gt;bound = -1;
			dowidth(n-&gt;type);
			n-&gt;op = OSLICEARR;
			goto ret;
		}
		if(isslice(t)) {
			n-&gt;type = t;
			goto ret;
		}
		yyerror(&#34;cannot slice %#N (type %T)&#34;, l, t);
		goto error;

	/*
	 * call and call like
	 */
	case OCALL:
		l = n-&gt;left;
		if(l-&gt;op == ONAME &amp;&amp; l-&gt;etype != 0) {
			// builtin: OLEN, OCAP, etc.
			n-&gt;op = l-&gt;etype;
			n-&gt;left = n-&gt;right;
			n-&gt;right = N;
			goto reswitch;
		}
		if(l-&gt;op == ONAME &amp;&amp; (r = unsafenmagic(l, n-&gt;list)) != N) {
			n = r;
			goto reswitch;
		}
		typecheck(&amp;n-&gt;left, Erv | Etype | Ecall);
		defaultlit(&amp;n-&gt;left, T);
		l = n-&gt;left;
		if(l-&gt;op == OTYPE) {
			// pick off before type-checking arguments
			ok |= Erv;
			// turn CALL(type, arg) into CONV(arg) w/ type
			n-&gt;left = N;
			if(onearg(n) &lt; 0)
				goto error;
			n-&gt;op = OCONV;
			n-&gt;type = l-&gt;type;
			goto doconv;
		}

		if(count(n-&gt;list) == 1)
			typecheck(&amp;n-&gt;list-&gt;n, Erv | Efnstruct);
		else
			typechecklist(n-&gt;list, Erv);
		if((t = l-&gt;type) == T)
			goto error;
		checkwidth(t);

		switch(l-&gt;op) {
		case ODOTINTER:
			n-&gt;op = OCALLINTER;
			break;

		case ODOTMETH:
			n-&gt;op = OCALLMETH;
			typecheckaste(OCALL, getthisx(t), list1(l-&gt;left));
			break;

		default:
			n-&gt;op = OCALLFUNC;
			if(t-&gt;etype != TFUNC) {
				yyerror(&#34;cannot call non-function %#N (type %T)&#34;, l, t);
				goto error;
			}
			break;
		}
		typecheckaste(OCALL, getinargx(t), n-&gt;list);
		ok |= Etop;
		if(t-&gt;outtuple == 0)
			goto ret;
		ok |= Erv;
		if(t-&gt;outtuple == 1) {
			t = getoutargx(l-&gt;type)-&gt;type;
			if(t == T)
				goto error;
			if(t-&gt;etype == TFIELD)
				t = t-&gt;type;
			n-&gt;type = t;
			goto ret;
		}
		// multiple return
		if(!(top &amp; (Efnstruct | Etop))) {
			yyerror(&#34;multiple-value %#N() in single-value context&#34;, l);
			goto ret;
		}
		n-&gt;type = getoutargx(l-&gt;type);
		goto ret;

	case OCAP:
	case OLEN:
		ok |= Erv;
		if(onearg(n) &lt; 0)
			goto error;
		typecheck(&amp;n-&gt;left, Erv);
		defaultlit(&amp;n-&gt;left, T);
		implicitstar(&amp;n-&gt;left);
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		switch(n-&gt;op) {
		case OCAP:
			if(!okforcap[t-&gt;etype])
				goto badcall1;
			break;
		case OLEN:
			if(!okforlen[t-&gt;etype])
				goto badcall1;
			break;
		}
		// might be constant
		switch(t-&gt;etype) {
		case TSTRING:
			if(isconst(l, CTSTR))
				nodconst(n, types[TINT], l-&gt;val.u.sval-&gt;len);
			break;
		case TARRAY:
			if(t-&gt;bound &gt;= 0)
				nodconst(n, types[TINT], t-&gt;bound);
			break;
		}
		n-&gt;type = types[TINT];
		goto ret;

	case OCLOSED:
	case OCLOSE:
		if(onearg(n) &lt; 0)
			goto error;
		typecheck(&amp;n-&gt;left, Erv);
		defaultlit(&amp;n-&gt;left, T);
		l = n-&gt;left;
		if((t = l-&gt;type) == T)
			goto error;
		if(t-&gt;etype != TCHAN) {
			yyerror(&#34;invalid operation: %#N (non-chan type %T)&#34;, n, t);
			goto error;
		}
		if(n-&gt;op == OCLOSED) {
			n-&gt;type = types[TBOOL];
			ok |= Erv;
		} else
			ok |= Etop;
		goto ret;

	case OCONV:
	doconv:
		ok |= Erv;
		typecheck(&amp;n-&gt;left, Erv | (top &amp; Eindir));
		convlit1(&amp;n-&gt;left, n-&gt;type, 1);
		if((t = n-&gt;left-&gt;type) == T || n-&gt;type == T)
			goto error;
		n = typecheckconv(n, n-&gt;left, n-&gt;type, 1);
		if(n-&gt;type == T)
			goto error;
		goto ret;

	case OMAKE:
		ok |= Erv;
		args = n-&gt;list;
		if(args == nil) {
			yyerror(&#34;missing argument to make&#34;);
			goto error;
		}
		l = args-&gt;n;
		args = args-&gt;next;
		typecheck(&amp;l, Etype);
		if((t = l-&gt;type) == T)
			goto error;

		switch(t-&gt;etype) {
		default:
		badmake:
			yyerror(&#34;cannot make type %T&#34;, t);
			goto error;

		case TARRAY:
			if(!isslice(t))
				goto badmake;
			if(args == nil) {
				yyerror(&#34;missing len argument to make(%T)&#34;, t);
				goto error;
			}
			l = args-&gt;n;
			args = args-&gt;next;
			typecheck(&amp;l, Erv);
			defaultlit(&amp;l, types[TUINT]);
			r = N;
			if(args != nil) {
				r = args-&gt;n;
				args = args-&gt;next;
				typecheck(&amp;r, Erv);
				defaultlit(&amp;r, types[TUINT]);
			}
			if(l-&gt;type == T || (r &amp;&amp; r-&gt;type == T))
				goto error;
			if(!isint[l-&gt;type-&gt;etype]) {
				yyerror(&#34;non-integer len argument to make(%T)&#34;, t);
				goto error;
			}
			if(r &amp;&amp; !isint[r-&gt;type-&gt;etype]) {
				yyerror(&#34;non-integer cap argument to make(%T)&#34;, t);
				goto error;
			}
			if(r == N)
				r = nodintconst(0);
			n-&gt;left = l;
			n-&gt;right = r;
			n-&gt;op = OMAKESLICE;
			break;

		case TMAP:
			if(args != nil) {
				l = args-&gt;n;
				args = args-&gt;next;
				typecheck(&amp;l, Erv);
				defaultlit(&amp;l, types[TUINT]);
				if(l-&gt;type == T)
					goto error;
				if(!isint[l-&gt;type-&gt;etype]) {
					yyerror(&#34;non-integer size argument to make(%T)&#34;, t);
					goto error;
				}
				n-&gt;left = l;
			} else
				n-&gt;left = nodintconst(0);
			n-&gt;op = OMAKEMAP;
			break;

		case TCHAN:
			l = N;
			if(args != nil) {
				l = args-&gt;n;
				args = args-&gt;next;
				typecheck(&amp;l, Erv);
				defaultlit(&amp;l, types[TUINT]);
				if(l-&gt;type == T)
					goto error;
				if(!isint[l-&gt;type-&gt;etype]) {
					yyerror(&#34;non-integer buffer argument to make(%T)&#34;, t);
					goto error;
				}
				n-&gt;left = l;
			} else
				n-&gt;left = nodintconst(0);
			n-&gt;op = OMAKECHAN;
			break;
		}
		if(args != nil) {
			yyerror(&#34;too many arguments to make(%T)&#34;, t);
			n-&gt;op = OMAKE;
			goto error;
		}
		n-&gt;type = t;
		goto ret;

	case ONEW:
		ok |= Erv;
		args = n-&gt;list;
		if(args == nil) {
			yyerror(&#34;missing argument to new&#34;);
			goto error;
		}
		l = args-&gt;n;
		typecheck(&amp;l, Etype);
		if((t = l-&gt;type) == T)
			goto error;
		if(args-&gt;next != nil) {
			yyerror(&#34;too many arguments to new(%T)&#34;, t);
			goto error;
		}
		n-&gt;left = l;
		n-&gt;type = ptrto(t);
		goto ret;

	case OPANIC:
	case OPANICN:
	case OPRINT:
	case OPRINTN:
		ok |= Etop;
		typechecklist(n-&gt;list, Erv);
		goto ret;

	case OCLOSURE:
		ok |= Erv;
		typecheckclosure(n);
		if(n-&gt;type == T)
			goto error;
		goto ret;

	/*
	 * statements
	 */
	case OAS:
		ok |= Etop;
		typecheckas(n);
		goto ret;

	case OAS2:
		ok |= Etop;
		typecheckas2(n);
		goto ret;

	case OBREAK:
	case OCONTINUE:
	case ODCL:
	case OEMPTY:
	case OGOTO:
	case OLABEL:
	case OXFALL:
		ok |= Etop;
		goto ret;

	case ODEFER:
	case OPROC:
		ok |= Etop;
		typecheck(&amp;n-&gt;left, Etop);
		goto ret;

	case OFOR:
		ok |= Etop;
		typechecklist(n-&gt;ninit, Etop);
		typecheck(&amp;n-&gt;ntest, Erv);
		if(n-&gt;ntest != N &amp;&amp; (t = n-&gt;ntest-&gt;type) != T &amp;&amp; t-&gt;etype != TBOOL)
			yyerror(&#34;non-bool %+N used as for condition&#34;, n-&gt;ntest);
		typecheck(&amp;n-&gt;nincr, Etop);
		typechecklist(n-&gt;nbody, Etop);
		goto ret;

	case OIF:
		ok |= Etop;
		typechecklist(n-&gt;ninit, Etop);
		typecheck(&amp;n-&gt;ntest, Erv);
		if(n-&gt;ntest != N &amp;&amp; (t = n-&gt;ntest-&gt;type) != T &amp;&amp; t-&gt;etype != TBOOL)
			yyerror(&#34;non-bool %+N used as if condition&#34;, n-&gt;ntest);
		typechecklist(n-&gt;nbody, Etop);
		typechecklist(n-&gt;nelse, Etop);
		goto ret;

	case ORETURN:
		ok |= Etop;
		typechecklist(n-&gt;list, Erv | Efnstruct);
		if(curfn-&gt;type-&gt;outnamed &amp;&amp; n-&gt;list == nil)
			goto ret;
		typecheckaste(ORETURN, getoutargx(curfn-&gt;type), n-&gt;list);
		goto ret;

	case OSELECT:
		ok |= Etop;
		typecheckselect(n);
		goto ret;

	case OSWITCH:
		ok |= Etop;
		typecheckswitch(n);
		goto ret;

	case ORANGE:
		ok |= Etop;
		typecheckrange(n);
		goto ret;

	case OTYPECASE:
		ok |= Etop | Erv;
		typecheck(&amp;n-&gt;left, Erv);
		goto ret;

	case OTYPESW:
		yyerror(&#34;use of .(type) outside type switch&#34;);
		goto error;

	case OXCASE:
		ok |= Etop;
		typechecklist(n-&gt;list, Erv);
		typechecklist(n-&gt;nbody, Etop);
		goto ret;

	case ODCLFUNC:
		ok |= Etop;
		typecheckfunc(n);
		goto ret;

	case ODCLCONST:
		ok |= Etop;
		typecheck(&amp;n-&gt;left, Erv);
		goto ret;

	case ODCLTYPE:
		ok |= Etop;
		typecheck(&amp;n-&gt;left, Etype);
		goto ret;
	}

ret:
	t = n-&gt;type;
	if(t &amp;&amp; !t-&gt;funarg &amp;&amp; n-&gt;op != OTYPE) {
		switch(t-&gt;etype) {
		case TFUNC:	// might have TANY; wait until its called
		case TANY:
		case TFORW:
		case TIDEAL:
		case TNIL:
		case TBLANK:
			break;
		case TARRAY:
			if(t-&gt;bound == -100) {
				yyerror(&#34;use of [...] array outside of array literal&#34;);
				t-&gt;bound = 1;
			}
		default:
			checkwidth(t);
		}
	}

	evconst(n);
	if(n-&gt;op == OTYPE &amp;&amp; !(top &amp; Etype)) {
		yyerror(&#34;type %T is not an expression&#34;, n-&gt;type);
		goto error;
	}
	if((top &amp; (Erv|Etype)) == Etype &amp;&amp; n-&gt;op != OTYPE) {
		yyerror(&#34;%#N is not a type&#34;, n);
		goto error;
	}
	if((ok &amp; Ecall) &amp;&amp; !(top &amp; Ecall)) {
		yyerror(&#34;must call %#N&#34;, n);
		goto error;
	}
	// TODO(rsc): simplify
	if((top &amp; (Ecall|Erv|Etype)) &amp;&amp; !(top &amp; Etop) &amp;&amp; !(ok &amp; (Erv|Etype|Ecall))) {
		yyerror(&#34;%#N used as value&#34;, n);
		goto error;
	}
	if((top &amp; Etop) &amp;&amp; !(top &amp; (Ecall|Erv|Etype)) &amp;&amp; !(ok &amp; Etop)) {
		yyerror(&#34;%#N not used&#34;, n);
		goto error;
	}

	/* TODO
	if(n-&gt;type == T)
		fatal(&#34;typecheck nil type&#34;);
	*/
	goto out;

badcall1:
	yyerror(&#34;invalid argument %#N (type %T) for %#O&#34;, n-&gt;left, n-&gt;left-&gt;type, n-&gt;op);
	goto error;

error:
	n-&gt;type = T;

out:
	lineno = lno;
	n-&gt;typecheck = 1;
	*np = n;
	return n;
}

static void
implicitstar(Node **nn)
{
	Type *t;
	Node *n;

	// insert implicit * if needed
	n = *nn;
	t = n-&gt;type;
	if(t == T || !isptr[t-&gt;etype])
		return;
	t = t-&gt;type;
	if(t == T)
		return;
	if(!isfixedarray(t))
		return;
	n = nod(OIND, n, N);
	typecheck(&amp;n, Erv);
	*nn = n;
}

static int
onearg(Node *n)
{
	if(n-&gt;left != N)
		return 0;
	if(n-&gt;list == nil) {
		yyerror(&#34;missing argument to %#O - %#N&#34;, n-&gt;op, n);
		return -1;
	}
	n-&gt;left = n-&gt;list-&gt;n;
	if(n-&gt;list-&gt;next != nil) {
		yyerror(&#34;too many arguments to %#O&#34;, n-&gt;op);
		n-&gt;list = nil;
		return -1;
	}
	n-&gt;list = nil;
	return 0;
}

static Type*
lookdot1(Sym *s, Type *t, Type *f)
{
	Type *r;

	r = T;
	for(; f!=T; f=f-&gt;down) {
		if(f-&gt;sym != s)
			continue;
		if(r != T) {
			yyerror(&#34;ambiguous DOT reference %T.%S&#34;, t, s);
			break;
		}
		r = f;
	}
	return r;
}

static int
lookdot(Node *n, Type *t)
{
	Type *f1, *f2, *tt, *rcvr;
	Sym *s;

	s = n-&gt;right-&gt;sym;

	dowidth(t);
	f1 = T;
	if(t-&gt;etype == TSTRUCT || t-&gt;etype == TINTER)
		f1 = lookdot1(s, t, t-&gt;type);

	f2 = methtype(n-&gt;left-&gt;type);
	if(f2 != T)
		f2 = lookdot1(s, f2, f2-&gt;method);

	if(f1 != T) {
		if(f2 != T)
			yyerror(&#34;ambiguous DOT reference %S as both field and method&#34;,
				n-&gt;right-&gt;sym);
		n-&gt;xoffset = f1-&gt;width;
		n-&gt;type = f1-&gt;type;
		if(t-&gt;etype == TINTER) {
			if(isptr[n-&gt;left-&gt;type-&gt;etype]) {
				n-&gt;left = nod(OIND, n-&gt;left, N);	// implicitstar
				typecheck(&amp;n-&gt;left, Erv);
			}
			n-&gt;op = ODOTINTER;
		}
		return 1;
	}

	if(f2 != T) {
		tt = n-&gt;left-&gt;type;
		dowidth(tt);
		rcvr = getthisx(f2-&gt;type)-&gt;type-&gt;type;
		if(!eqtype(rcvr, tt)) {
			if(rcvr-&gt;etype == tptr &amp;&amp; eqtype(rcvr-&gt;type, tt)) {
				typecheck(&amp;n-&gt;left, Erv);
				checklvalue(n-&gt;left, &#34;call pointer method on&#34;);
				addrescapes(n-&gt;left);
				n-&gt;left = nod(OADDR, n-&gt;left, N);
				typecheck(&amp;n-&gt;left, Erv);
			} else if(tt-&gt;etype == tptr &amp;&amp; eqtype(tt-&gt;type, rcvr)) {
				n-&gt;left = nod(OIND, n-&gt;left, N);
				typecheck(&amp;n-&gt;left, Erv);
			} else {
				// method is attached to wrong type?
				fatal(&#34;method mismatch: %T for %T&#34;, rcvr, tt);
			}
		}
		n-&gt;right = methodname(n-&gt;right, n-&gt;left-&gt;type);
		n-&gt;xoffset = f2-&gt;width;
		n-&gt;type = f2-&gt;type;
		n-&gt;op = ODOTMETH;
		return 1;
	}

	return 0;
}

static int
nokeys(NodeList *l)
{
	for(; l; l=l-&gt;next)
		if(l-&gt;n-&gt;op == OKEY)
			return 0;
	return 1;
}

/*
 * check implicit or explicit conversion from node type nt to type t.
 */
int
checkconv(Type *nt, Type *t, int explicit, int *op, int *et)
{
	*op = OCONV;
	*et = 0;



	// preexisting error
	if(t == T || t-&gt;etype == TFORW)
		return 0;

	/*
	 * implicit conversions
	 */
	if(nt == T)
		return 0;

	if(t-&gt;etype == TBLANK) {
		*op = OCONVNOP;
		return 0;
	}

	if(eqtype(t, nt)) {
		exportassignok(t);
		*op = OCONVNOP;
		if(!explicit || t == nt)
			return 0;
		return 1;
	}

	// interfaces are not subject to the name restrictions below.
	// accept anything involving interfaces and let ifacecvt
	// generate a good message.  some messages have to be
	// delayed anyway.
	// TODO(rsc): now that everything is delayed for whole-package
	// compilation, the messages could be generated right here.
	if(isnilinter(t) || isnilinter(nt) || isinter(t) || isinter(nt)) {
		*et = ifaceas1(t, nt, 0);
		*op = OCONVIFACE;
		return 1;
	}

	// otherwise, if concrete types have names, they must match.
	if(!explicit &amp;&amp; t-&gt;sym &amp;&amp; nt-&gt;sym &amp;&amp; t != nt)
		return -1;

	// channel must not lose directionality
	if(t-&gt;etype == TCHAN &amp;&amp; nt-&gt;etype == TCHAN) {
		if(t-&gt;chan &amp; ~nt-&gt;chan)
			return -1;
		if(eqtype(t-&gt;type, nt-&gt;type)) {
			*op = OCONVNOP;
			return 1;
		}
	}

	// array to slice
	if(isslice(t) &amp;&amp; isptr[nt-&gt;etype] &amp;&amp; isfixedarray(nt-&gt;type)
	&amp;&amp; eqtype(t-&gt;type, nt-&gt;type-&gt;type)) {
		*op = OCONVSLICE;
		return 1;
	}

	/*
	 * explicit conversions
	 */
	if(!explicit)
		return -1;

	// same representation
	if(cvttype(t, nt)) {
		*op = OCONVNOP;
		return 1;
	}

	// simple fix-float
	if(isint[t-&gt;etype] || isfloat[t-&gt;etype])
	if(isint[nt-&gt;etype] || isfloat[nt-&gt;etype])
		return 1;

	// to string
	if(istype(t, TSTRING)) {
		// integer rune
		if(isint[nt-&gt;etype]) {
			*op = ORUNESTR;
			return 1;
		}

		// *[10]byte -&gt; string
		// in preparation for next step
		if(isptr[nt-&gt;etype] &amp;&amp; isfixedarray(nt-&gt;type)) {
			switch(nt-&gt;type-&gt;type-&gt;etype) {
			case TUINT8:
				*op = OARRAYBYTESTR;
				return 1;
			case TINT:
				*op = OARRAYRUNESTR;
				return 1;
			}
		}

		// []byte -&gt; string
		if(isslice(nt)) {
			switch(nt-&gt;type-&gt;etype) {
			case TUINT8:
				*op = OARRAYBYTESTR;
				return 1;
			case TINT:
				*op = OARRAYRUNESTR;
				return 1;
			}
		}
	}

	// convert to unsafe pointer
	if(isptrto(t, TANY)
	&amp;&amp; (isptr[nt-&gt;etype] || nt-&gt;etype == TUINTPTR))
		return 1;

	// convert from unsafe pointer
	if(isptrto(nt, TANY)
	&amp;&amp; (isptr[t-&gt;etype] || t-&gt;etype == TUINTPTR))
		return 1;

	return -1;
}

Node*
typecheckconv(Node *nconv, Node *n, Type *t, int explicit)
{
	int et, op;
	Node *n1;

	convlit1(&amp;n, t, explicit);
	if(n-&gt;type == T)
		return n;

	if(n-&gt;op == OLITERAL)
	if(explicit || isideal(n-&gt;type))
	if(cvttype(t, n-&gt;type)) {
		// can convert literal in place
		// TODO(rsc) is this needed?
		n1 = nod(OXXX, N, N);
		*n1 = *n;
		n1-&gt;type = t;
		return n1;
	}

	switch(checkconv(n-&gt;type, t, explicit, &amp;op, &amp;et)) {
	case -1:
		if(explicit)
			yyerror(&#34;cannot convert %+N to type %T&#34;, n, t);
		else
			yyerror(&#34;cannot use %+N as type %T&#34;, n, t);
		return n;

	case 0:
		if(nconv) {
			nconv-&gt;op = OCONVNOP;
			return nconv;
		}
		return n;
	}

	if(op == OCONVIFACE)
		defaultlit(&amp;n, T);

	if(nconv == N)
		nconv = nod(OCONV, n, N);
	nconv-&gt;op = op;
	nconv-&gt;etype = et;
	nconv-&gt;type = t;
	nconv-&gt;typecheck = 1;
	return nconv;
}

/*
 * typecheck assignment: type list = expression list
 */
static void
typecheckaste(int op, Type *tstruct, NodeList *nl)
{
	Type *t, *tl, *tn;
	Node *n;
	int lno;

	lno = lineno;

	if(tstruct-&gt;broke)
		goto out;

	if(nl != nil &amp;&amp; nl-&gt;next == nil &amp;&amp; (n = nl-&gt;n)-&gt;type != T)
	if(n-&gt;type-&gt;etype == TSTRUCT &amp;&amp; n-&gt;type-&gt;funarg) {
		setlineno(n);
		tn = n-&gt;type-&gt;type;
		for(tl=tstruct-&gt;type; tl; tl=tl-&gt;down) {
			int xx, yy;
			if(tn == T) {
				yyerror(&#34;not enough arguments to %#O&#34;, op);
				goto out;
			}
			if(isddd(tl-&gt;type))
				goto out;
			if(checkconv(tn-&gt;type, tl-&gt;type, 0, &amp;xx, &amp;yy) &lt; 0)
				yyerror(&#34;cannot use type %T as type %T&#34;, tn-&gt;type, tl-&gt;type);
			tn = tn-&gt;down;
		}
		if(tn != T)
			yyerror(&#34;too many arguments to %#O&#34;, op);
		goto out;
	}

	for(tl=tstruct-&gt;type; tl; tl=tl-&gt;down) {
		t = tl-&gt;type;
		if(isddd(t)) {
			for(; nl; nl=nl-&gt;next) {
				setlineno(nl-&gt;n);
				defaultlit(&amp;nl-&gt;n, T);
			}
			goto out;
		}
		if(nl == nil) {
			yyerror(&#34;not enough arguments to %#O&#34;, op);
			goto out;
		}
		n = nl-&gt;n;
		setlineno(nl-&gt;n);
		if(n-&gt;type != T)
			nl-&gt;n = typecheckconv(nil, n, t, 0);
		nl = nl-&gt;next;
	}
	if(nl != nil) {
		yyerror(&#34;too many arguments to %#O&#34;, op);
		goto out;
	}

out:
	lineno = lno;
}

/*
 * do the export rules allow writing to this type?
 * cannot be implicitly assigning to any type with
 * an unavailable field.
 */
static int
exportassignok(Type *t)
{
	Type *f;
	Sym *s;

	if(t == T)
		return 1;
	switch(t-&gt;etype) {
	default:
		// most types can&#39;t contain others; they&#39;re all fine.
		break;
	case TSTRUCT:
		for(f=t-&gt;type; f; f=f-&gt;down) {
			if(f-&gt;etype != TFIELD)
				fatal(&#34;structas: not field&#34;);
			s = f-&gt;sym;
			// s == nil doesn&#39;t happen for embedded fields (they get the type symbol).
			// it only happens for fields in a ... struct.
			if(s != nil &amp;&amp; !exportname(s-&gt;name) &amp;&amp; strcmp(package, s-&gt;package) != 0) {
				yyerror(&#34;implicit assignment of %T field &#39;%s&#39;&#34;, t, s-&gt;name);
				return 0;
			}
			if(!exportassignok(f-&gt;type))
				return 0;
		}
		break;

	case TARRAY:
		if(t-&gt;bound &lt; 0)	// slices are pointers; that&#39;s fine
			break;
		if(!exportassignok(t-&gt;type))
			return 0;
		break;
	}
	return 1;
}


/*
 * type check composite
 */

static void
fielddup(Node *n, Node *hash[], ulong nhash)
{
	uint h;
	char *s;
	Node *a;

	if(n-&gt;op != ONAME)
		fatal(&#34;fielddup: not ONAME&#34;);
	s = n-&gt;sym-&gt;name;
	h = stringhash(s)%nhash;
	for(a=hash[h]; a!=N; a=a-&gt;ntest) {
		if(strcmp(a-&gt;sym-&gt;name, s) == 0) {
			yyerror(&#34;duplicate field name in struct literal: %s&#34;, s);
			return;
		}
	}
	n-&gt;ntest = hash[h];
	hash[h] = n;
}

static void
keydup(Node *n, Node *hash[], ulong nhash)
{
	uint h;
	ulong b;
	double d;
	int i;
	Node *a;
	Node cmp;
	char *s;

	evconst(n);
	if(n-&gt;op != OLITERAL)
		return;	// we dont check variables

	switch(n-&gt;val.ctype) {
	default:	// unknown, bool, nil
		b = 23;
		break;
	case CTINT:
		b = mpgetfix(n-&gt;val.u.xval);
		break;
	case CTFLT:
		d = mpgetflt(n-&gt;val.u.fval);
		s = (char*)&amp;d;
		b = 0;
		for(i=sizeof(d); i&gt;0; i--)
			b = b*PRIME1 + *s++;
		break;
	case CTSTR:
		b = 0;
		s = n-&gt;val.u.sval-&gt;s;
		for(i=n-&gt;val.u.sval-&gt;len; i&gt;0; i--)
			b = b*PRIME1 + *s++;
		break;
	}

	h = b%nhash;
	memset(&amp;cmp, 0, sizeof(cmp));
	for(a=hash[h]; a!=N; a=a-&gt;ntest) {
		cmp.op = OEQ;
		cmp.left = n;
		cmp.right = a;
		evconst(&amp;cmp);
		b = cmp.val.u.bval;
		if(b) {
			// too lazy to print the literal
			yyerror(&#34;duplicate key in map literal&#34;);
			return;
		}
	}
	n-&gt;ntest = hash[h];
	hash[h] = n;
}

static void
indexdup(Node *n, Node *hash[], ulong nhash)
{
	uint h;
	Node *a;
	ulong b, c;

	if(n-&gt;op != OLITERAL)
		fatal(&#34;indexdup: not OLITERAL&#34;);

	b = mpgetfix(n-&gt;val.u.xval);
	h = b%nhash;
	for(a=hash[h]; a!=N; a=a-&gt;ntest) {
		c = mpgetfix(a-&gt;val.u.xval);
		if(b == c) {
			yyerror(&#34;duplicate index in array literal: %ld&#34;, b);
			return;
		}
	}
	n-&gt;ntest = hash[h];
	hash[h] = n;
}

static void
typecheckcomplit(Node **np)
{
	int bad, i, len, nerr;
	Node *l, *n, *hash[101];
	NodeList *ll;
	Type *t, *f;

	n = *np;

	memset(hash, 0, sizeof hash);

	l = typecheck(&amp;n-&gt;right /* sic */, Etype /* TODO | Edotarray */);
	if((t = l-&gt;type) == T)
		goto error;
	nerr = nerrors;
	switch(t-&gt;etype) {
	default:
		yyerror(&#34;invalid type for composite literal: %T&#34;, t);
		n-&gt;type = T;
		break;

	case TARRAY:
		len = 0;
		i = 0;
		for(ll=n-&gt;list; ll; ll=ll-&gt;next) {
			l = ll-&gt;n;
			if(l-&gt;op == OKEY) {
				typecheck(&amp;l-&gt;left, Erv);
				evconst(l-&gt;left);
				i = nonnegconst(l-&gt;left);
				if(i &lt; 0) {
					yyerror(&#34;array index must be non-negative integer constant&#34;);
					i = -(1&lt;&lt;30);	// stay negative for a while
				}
				typecheck(&amp;l-&gt;right, Erv);
				defaultlit(&amp;l-&gt;right, t-&gt;type);
				l-&gt;right = typecheckconv(nil, l-&gt;right, t-&gt;type, 0);
			} else {
				typecheck(&amp;ll-&gt;n, Erv);
				defaultlit(&amp;ll-&gt;n, t-&gt;type);
				ll-&gt;n = typecheckconv(nil, ll-&gt;n, t-&gt;type, 0);
				ll-&gt;n = nod(OKEY, nodintconst(i), ll-&gt;n);
				ll-&gt;n-&gt;left-&gt;type = types[TINT];
				ll-&gt;n-&gt;left-&gt;typecheck = 1;
			}
			if(i &gt;= 0)
				indexdup(ll-&gt;n-&gt;left, hash, nelem(hash));
			i++;
			if(i &gt; len) {
				len = i;
				if(t-&gt;bound &gt;= 0 &amp;&amp; len &gt; t-&gt;bound) {
					setlineno(l);
					yyerror(&#34;array index %d out of bounds [0:%d]&#34;, len, t-&gt;bound);
					t-&gt;bound = -1;	// no more errors
				}
			}
		}
		if(t-&gt;bound == -100)
			t-&gt;bound = len;
		if(t-&gt;bound &lt; 0)
			n-&gt;right = nodintconst(len);
		n-&gt;op = OARRAYLIT;
		break;

	case TMAP:
		for(ll=n-&gt;list; ll; ll=ll-&gt;next) {
			l = ll-&gt;n;
			if(l-&gt;op != OKEY) {
				typecheck(&amp;ll-&gt;n, Erv);
				yyerror(&#34;missing key in map literal&#34;);
				continue;
			}
			typecheck(&amp;l-&gt;left, Erv);
			typecheck(&amp;l-&gt;right, Erv);
			defaultlit(&amp;l-&gt;left, t-&gt;down);
			defaultlit(&amp;l-&gt;right, t-&gt;type);
			l-&gt;left = typecheckconv(nil, l-&gt;left, t-&gt;down, 0);
			l-&gt;right = typecheckconv(nil, l-&gt;right, t-&gt;type, 0);
			keydup(l-&gt;left, hash, nelem(hash));
		}
		n-&gt;op = OMAPLIT;
		break;

	case TSTRUCT:
		bad = 0;
		if(n-&gt;list != nil &amp;&amp; nokeys(n-&gt;list)) {
			// simple list of variables
			f = t-&gt;type;
			for(ll=n-&gt;list; ll; ll=ll-&gt;next) {
				typecheck(&amp;ll-&gt;n, Erv);
				if(f == nil) {
					if(!bad++)
						yyerror(&#34;too many values in struct initializer&#34;);
					continue;
				}
				ll-&gt;n = typecheckconv(nil, ll-&gt;n, f-&gt;type, 0);
				ll-&gt;n = nod(OKEY, newname(f-&gt;sym), ll-&gt;n);
				ll-&gt;n-&gt;left-&gt;typecheck = 1;
				f = f-&gt;down;
			}
			if(f != nil)
				yyerror(&#34;too few values in struct initializer&#34;);
		} else {
			// keyed list
			for(ll=n-&gt;list; ll; ll=ll-&gt;next) {
				l = ll-&gt;n;
				if(l-&gt;op != OKEY) {
					if(!bad++)
						yyerror(&#34;mixture of field:value and value initializers&#34;);
					typecheck(&amp;ll-&gt;n, Erv);
					continue;
				}
				if(l-&gt;left-&gt;sym == S) {
					yyerror(&#34;invalid field name %#N in struct initializer&#34;, l-&gt;left);
					typecheck(&amp;l-&gt;right, Erv);
					continue;
				}
				l-&gt;left = newname(l-&gt;left-&gt;sym);
				l-&gt;left-&gt;typecheck = 1;
				f = lookdot1(l-&gt;left-&gt;sym, t, t-&gt;type);
				typecheck(&amp;l-&gt;right, Erv);
				if(f == nil)
					continue;
				fielddup(newname(f-&gt;sym), hash, nelem(hash));
				l-&gt;right = typecheckconv(nil, l-&gt;right, f-&gt;type, 0);
			}
		}
		n-&gt;op = OSTRUCTLIT;
		break;
	}
	if(nerr != nerrors)
		goto error;
	n-&gt;type = t;

	*np = n;
	return;

error:
	n-&gt;type = T;
	*np = n;
}

/*
 * the address of n has been taken and might be used after
 * the current function returns.  mark any local vars
 * as needing to move to the heap.
 */
static void
addrescapes(Node *n)
{
	char buf[100];
	switch(n-&gt;op) {
	default:
		// probably a type error already.
		// dump(&#34;addrescapes&#34;, n);
		break;

	case ONAME:
		if(n-&gt;noescape)
			break;
		switch(n-&gt;class) {
		case PPARAMOUT:
			yyerror(&#34;cannot take address of out parameter %s&#34;, n-&gt;sym-&gt;name);
			break;
		case PAUTO:
		case PPARAM:
			// if func param, need separate temporary
			// to hold heap pointer.
			// the function type has already been checked
			// (we&#39;re in the function body)
			// so the param already has a valid xoffset.
			if(n-&gt;class == PPARAM) {
				// expression to refer to stack copy
				n-&gt;stackparam = nod(OPARAM, n, N);
				n-&gt;stackparam-&gt;type = n-&gt;type;
				n-&gt;stackparam-&gt;addable = 1;
				if(n-&gt;xoffset == BADWIDTH)
					fatal(&#34;addrescapes before param assignment&#34;);
				n-&gt;stackparam-&gt;xoffset = n-&gt;xoffset;
				n-&gt;xoffset = 0;
			}

			n-&gt;class |= PHEAP;
			n-&gt;addable = 0;
			n-&gt;ullman = 2;
			n-&gt;alloc = callnew(n-&gt;type);
			n-&gt;xoffset = 0;

			// create stack variable to hold pointer to heap
			n-&gt;heapaddr = nod(ONAME, N, N);
			n-&gt;heapaddr-&gt;type = ptrto(n-&gt;type);
			snprint(buf, sizeof buf, &#34;&amp;%S&#34;, n-&gt;sym);
			n-&gt;heapaddr-&gt;sym = lookup(buf);
			n-&gt;heapaddr-&gt;class = PHEAP-1;	// defer tempname to allocparams
			curfn-&gt;dcl = list(curfn-&gt;dcl, n-&gt;heapaddr);
			break;
		}
		break;

	case OIND:
	case ODOTPTR:
		break;

	case ODOT:
	case OINDEX:
		// ODOTPTR has already been introduced,
		// so these are the non-pointer ODOT and OINDEX.
		// In &amp;x[0], if x is a slice, then x does not
		// escape--the pointer inside x does, but that
		// is always a heap pointer anyway.
		if(!isslice(n-&gt;left-&gt;type))
			addrescapes(n-&gt;left);
		break;
	}
}

/*
 * lvalue etc
 */
static int
islvalue(Node *n)
{
	switch(n-&gt;op) {
	case OINDEX:
	case OIND:
	case ODOTPTR:
		return 1;
	case ODOT:
		return islvalue(n-&gt;left);
	case ONAME:
		if(n-&gt;class == PFUNC)
			return 0;
		return 1;
	}
	return 0;
}

static void
checklvalue(Node *n, char *verb)
{
	if(!islvalue(n))
		yyerror(&#34;cannot %s %#N&#34;, verb, n);
}

static void
checkassign(Node *n)
{
	if(islvalue(n))
		return;
	if(n-&gt;op == OINDEXMAP) {
		n-&gt;etype = 1;
		return;
	}
	yyerror(&#34;cannot assign to %#N&#34;, n);
}

static void
checkassignlist(NodeList *l)
{
	for(; l; l=l-&gt;next)
		checkassign(l-&gt;n);
}

/*
 * type check assignment.
 * if this assignment is the definition of a var on the left side,
 * fill in the var&#39;s type.
 */

static void
typecheckas(Node *n)
{
	// delicate little dance.
	// the definition of n may refer to this assignment
	// as its definition, in which case it will call typecheckas.
	// in that case, do not call typecheck back, or it will cycle.
	// if the variable has a type (ntype) then typechecking
	// will not look at defn, so it is okay (and desirable,
	// so that the conversion below happens).
	if(n-&gt;left-&gt;defn != n || n-&gt;left-&gt;ntype)
		typecheck(&amp;n-&gt;left, Erv | Easgn);

	checkassign(n-&gt;left);
	typecheck(&amp;n-&gt;right, Erv);
	if(n-&gt;left-&gt;type != T &amp;&amp; n-&gt;right &amp;&amp; n-&gt;right-&gt;type != T)
		n-&gt;right = typecheckconv(nil, n-&gt;right, n-&gt;left-&gt;type, 0);
	if(n-&gt;left-&gt;defn == n &amp;&amp; n-&gt;left-&gt;ntype == N) {
		defaultlit(&amp;n-&gt;right, T);
		n-&gt;left-&gt;type = n-&gt;right-&gt;type;
	}

	// second half of dance.
	// now that right is done, typecheck the left
	// just to get it over with.  see dance above.
	n-&gt;typecheck = 1;
	if(n-&gt;left-&gt;typecheck == 0)
		typecheck(&amp;n-&gt;left, Erv | Easgn);
}

static void
typecheckas2(Node *n)
{
	int cl, cr, op, et;
	NodeList *ll, *lr;
	Node *l, *r;
	Iter s;
	Type *t;

	for(ll=n-&gt;list; ll; ll=ll-&gt;next) {
		// delicate little dance.
		if(ll-&gt;n-&gt;defn != n || ll-&gt;n-&gt;ntype)
			typecheck(&amp;ll-&gt;n, Erv | Easgn);
	}
	cl = count(n-&gt;list);
	cr = count(n-&gt;rlist);
	checkassignlist(n-&gt;list);
	if(cl &gt; 1 &amp;&amp; cr == 1)
		typecheck(&amp;n-&gt;rlist-&gt;n, Erv | Efnstruct);
	else
		typechecklist(n-&gt;rlist, Erv);

	if(cl == cr) {
		// easy
		for(ll=n-&gt;list, lr=n-&gt;rlist; ll; ll=ll-&gt;next, lr=lr-&gt;next) {
			if(ll-&gt;n-&gt;type != T &amp;&amp; lr-&gt;n-&gt;type != T)
				lr-&gt;n = typecheckconv(nil, lr-&gt;n, ll-&gt;n-&gt;type, 0);
			if(ll-&gt;n-&gt;defn == n &amp;&amp; ll-&gt;n-&gt;ntype == N) {
				defaultlit(&amp;lr-&gt;n, T);
				ll-&gt;n-&gt;type = lr-&gt;n-&gt;type;
			}
		}
		goto out;
	}


	l = n-&gt;list-&gt;n;
	r = n-&gt;rlist-&gt;n;

	// m[i] = x, ok
	if(cl == 1 &amp;&amp; cr == 2 &amp;&amp; l-&gt;op == OINDEXMAP) {
		if(l-&gt;type == T)
			goto out;
		n-&gt;op = OAS2MAPW;
		n-&gt;rlist-&gt;n = typecheckconv(nil, r, l-&gt;type-&gt;down, 0);
		r = n-&gt;rlist-&gt;next-&gt;n;
		n-&gt;rlist-&gt;next-&gt;n = typecheckconv(nil, r, types[TBOOL], 0);
		goto out;
	}

	// x,y,z = f()
	if(cr == 1) {
		if(r-&gt;type == T)
			goto out;
		switch(r-&gt;op) {
		case OCALLMETH:
		case OCALLINTER:
		case OCALLFUNC:
			if(r-&gt;type-&gt;etype != TSTRUCT || r-&gt;type-&gt;funarg == 0)
				break;
			cr = structcount(r-&gt;type);
			if(cr != cl)
				goto mismatch;
			n-&gt;op = OAS2FUNC;
			t = structfirst(&amp;s, &amp;r-&gt;type);
			for(ll=n-&gt;list; ll; ll=ll-&gt;next) {
				if(ll-&gt;n-&gt;type != T)
					if(checkconv(t-&gt;type, ll-&gt;n-&gt;type, 0, &amp;op, &amp;et) &lt; 0)
						yyerror(&#34;cannot assign type %T to %+N&#34;, t-&gt;type, ll-&gt;n);
				if(ll-&gt;n-&gt;defn == n &amp;&amp; ll-&gt;n-&gt;ntype == N)
					ll-&gt;n-&gt;type = t-&gt;type;
				t = structnext(&amp;s);
			}
			goto out;
		}
	}

	// x, ok = y
	if(cl == 2 &amp;&amp; cr == 1) {
		if(r-&gt;type == T)
			goto out;
		switch(r-&gt;op) {
		case OINDEXMAP:
			n-&gt;op = OAS2MAPR;
			goto common;
		case ORECV:
			n-&gt;op = OAS2RECV;
			goto common;
		case ODOTTYPE:
			n-&gt;op = OAS2DOTTYPE;
		common:
			if(l-&gt;type != T &amp;&amp; checkconv(r-&gt;type, l-&gt;type, 0, &amp;op, &amp;et) &lt; 0)
				yyerror(&#34;cannot assign %+N to %+N&#34;, r, l);
			if(l-&gt;defn == n)
				l-&gt;type = r-&gt;type;
			l = n-&gt;list-&gt;next-&gt;n;
			if(l-&gt;type != T &amp;&amp; checkconv(types[TBOOL], l-&gt;type, 0, &amp;op, &amp;et) &lt; 0)
				yyerror(&#34;cannot assign bool value to %+N&#34;, l);
			if(l-&gt;defn == n &amp;&amp; l-&gt;ntype == N)
				l-&gt;type = types[TBOOL];
			goto out;
		}
	}

mismatch:
	yyerror(&#34;assignment count mismatch: %d = %d&#34;, cl, cr);

out:
	// second half of dance
	n-&gt;typecheck = 1;
	for(ll=n-&gt;list; ll; ll=ll-&gt;next)
		if(ll-&gt;n-&gt;typecheck == 0)
			typecheck(&amp;ll-&gt;n, Erv | Easgn);
}

/*
 * type check function definition
 */
static void
typecheckfunc(Node *n)
{
	Type *t, *rcvr;

//dump(&#34;nname&#34;, n-&gt;nname);
	typecheck(&amp;n-&gt;nname, Erv | Easgn);
	if((t = n-&gt;nname-&gt;type) == T)
		return;
	n-&gt;type = t;

	rcvr = getthisx(t)-&gt;type;
	if(rcvr != nil &amp;&amp; n-&gt;shortname != N &amp;&amp; !isblank(n-&gt;shortname))
		addmethod(n-&gt;shortname-&gt;sym, t, 1);
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
