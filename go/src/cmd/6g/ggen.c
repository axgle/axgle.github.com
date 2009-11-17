<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6g/ggen.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/6g/ggen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#undef	EXTERN
#define	EXTERN
#include &#34;gg.h&#34;
#include &#34;opt.h&#34;

void
compile(Node *fn)
{
	Plist *pl;
	Node nod1;
	Prog *ptxt;
	int32 lno;
	Type *t;
	Iter save;

	if(newproc == N) {
		newproc = sysfunc(&#34;newproc&#34;);
		deferproc = sysfunc(&#34;deferproc&#34;);
		deferreturn = sysfunc(&#34;deferreturn&#34;);
		throwindex = sysfunc(&#34;throwindex&#34;);
		throwslice = sysfunc(&#34;throwslice&#34;);
		throwreturn = sysfunc(&#34;throwreturn&#34;);
	}

	if(fn-&gt;nbody == nil)
		return;

	// set up domain for labels
	labellist = L;

	lno = setlineno(fn);

	curfn = fn;
	dowidth(curfn-&gt;type);

	if(curfn-&gt;type-&gt;outnamed) {
		// add clearing of the output parameters
		t = structfirst(&amp;save, getoutarg(curfn-&gt;type));
		while(t != T) {
			if(t-&gt;nname != N)
				curfn-&gt;nbody = concat(list1(nod(OAS, t-&gt;nname, N)), curfn-&gt;nbody);
			t = structnext(&amp;save);
		}
	}

	hasdefer = 0;
	walk(curfn);
	if(nerrors != 0 || isblank(curfn-&gt;nname))
		goto ret;

	allocparams();

	continpc = P;
	breakpc = P;

	pl = newplist();
	pl-&gt;name = curfn-&gt;nname;

	nodconst(&amp;nod1, types[TINT32], 0);
	ptxt = gins(ATEXT, curfn-&gt;nname, &amp;nod1);
	afunclit(&amp;ptxt-&gt;from);

	ginit();
	genlist(curfn-&gt;enter);
	genlist(curfn-&gt;nbody);
	gclean();
	checklabels();
	if(nerrors != 0)
		goto ret;

	if(curfn-&gt;type-&gt;outtuple != 0)
		ginscall(throwreturn, 0);

	if(hasdefer)
		ginscall(deferreturn, 0);
	pc-&gt;as = ARET;	// overwrite AEND
	pc-&gt;lineno = lineno;

	if(!debug[&#39;N&#39;] || debug[&#39;R&#39;] || debug[&#39;P&#39;])
		regopt(ptxt);

	// fill in argument size
	ptxt-&gt;to.offset = rnd(curfn-&gt;type-&gt;argwid, maxround);

	// fill in final stack size
	ptxt-&gt;to.offset &lt;&lt;= 32;
	ptxt-&gt;to.offset |= rnd(stksize+maxarg, maxround);

	if(debug[&#39;f&#39;])
		frame(0);

ret:
	lineno = lno;
}

/*
 * generate:
 *	call f
 *	proc=0	normal call
 *	proc=1	goroutine run in new proc
 *	proc=2	defer call save away stack
 */
void
ginscall(Node *f, int proc)
{
	Prog *p;
	Node reg, con;

	switch(proc) {
	default:
		fatal(&#34;ginscall: bad proc %d&#34;, proc);
		break;

	case 0:	// normal call
		p = gins(ACALL, N, f);
		afunclit(&amp;p-&gt;to);
		break;

	case 1:	// call in new proc (go)
	case 2:	// defered call (defer)
		nodreg(&amp;reg, types[TINT64], D_AX);
		gins(APUSHQ, f, N);
		nodconst(&amp;con, types[TINT32], argsize(f-&gt;type));
		gins(APUSHQ, &amp;con, N);
		if(proc == 1)
			ginscall(newproc, 0);
		else {
			if(!hasdefer)
				fatal(&#34;hasdefer=0 but has defer&#34;);
			ginscall(deferproc, 0);
		}
		gins(APOPQ, N, &amp;reg);
		gins(APOPQ, N, &amp;reg);
		break;
	}
}

/*
 * n is call to interface method.
 * generate res = n.
 */
void
cgen_callinter(Node *n, Node *res, int proc)
{
	Node *i, *f;
	Node tmpi, nodo, nodr, nodsp;

	i = n-&gt;left;
	if(i-&gt;op != ODOTINTER)
		fatal(&#34;cgen_callinter: not ODOTINTER %O&#34;, i-&gt;op);

	f = i-&gt;right;		// field
	if(f-&gt;op != ONAME)
		fatal(&#34;cgen_callinter: not ONAME %O&#34;, f-&gt;op);

	i = i-&gt;left;		// interface

	if(!i-&gt;addable) {
		tempname(&amp;tmpi, i-&gt;type);
		cgen(i, &amp;tmpi);
		i = &amp;tmpi;
	}

	genlist(n-&gt;list);		// assign the args

	regalloc(&amp;nodr, types[tptr], res);
	regalloc(&amp;nodo, types[tptr], &amp;nodr);
	nodo.op = OINDREG;

	agen(i, &amp;nodr);		// REG = &amp;inter

	nodindreg(&amp;nodsp, types[tptr], D_SP);
	nodo.xoffset += widthptr;
	cgen(&amp;nodo, &amp;nodsp);	// 0(SP) = 8(REG) -- i.s

	nodo.xoffset -= widthptr;
	cgen(&amp;nodo, &amp;nodr);	// REG = 0(REG) -- i.m

	nodo.xoffset = n-&gt;left-&gt;xoffset + 3*widthptr + 8;
	cgen(&amp;nodo, &amp;nodr);	// REG = 32+offset(REG) -- i.m-&gt;fun[f]

	// BOTCH nodr.type = fntype;
	nodr.type = n-&gt;left-&gt;type;
	ginscall(&amp;nodr, proc);

	regfree(&amp;nodr);
	regfree(&amp;nodo);

	setmaxarg(n-&gt;left-&gt;type);
}

/*
 * generate function call;
 *	proc=0	normal call
 *	proc=1	goroutine run in new proc
 *	proc=2	defer call save away stack
 */
void
cgen_call(Node *n, int proc)
{
	Type *t;
	Node nod, afun;

	if(n == N)
		return;

	if(n-&gt;left-&gt;ullman &gt;= UINF) {
		// if name involves a fn call
		// precompute the address of the fn
		tempname(&amp;afun, types[tptr]);
		cgen(n-&gt;left, &amp;afun);
	}

	genlist(n-&gt;list);		// assign the args
	t = n-&gt;left-&gt;type;

	setmaxarg(t);

	// call tempname pointer
	if(n-&gt;left-&gt;ullman &gt;= UINF) {
		regalloc(&amp;nod, types[tptr], N);
		cgen_as(&amp;nod, &amp;afun);
		nod.type = t;
		ginscall(&amp;nod, proc);
		regfree(&amp;nod);
		goto ret;
	}

	// call pointer
	if(n-&gt;left-&gt;op != ONAME || n-&gt;left-&gt;class != PFUNC) {
		regalloc(&amp;nod, types[tptr], N);
		cgen_as(&amp;nod, n-&gt;left);
		nod.type = t;
		ginscall(&amp;nod, proc);
		regfree(&amp;nod);
		goto ret;
	}

	// call direct
	n-&gt;left-&gt;method = 1;
	ginscall(n-&gt;left, proc);


ret:
	;
}

/*
 * call to n has already been generated.
 * generate:
 *	res = return value from call.
 */
void
cgen_callret(Node *n, Node *res)
{
	Node nod;
	Type *fp, *t;
	Iter flist;

	t = n-&gt;left-&gt;type;
	if(t-&gt;etype == TPTR32 || t-&gt;etype == TPTR64)
		t = t-&gt;type;

	fp = structfirst(&amp;flist, getoutarg(t));
	if(fp == T)
		fatal(&#34;cgen_callret: nil&#34;);

	memset(&amp;nod, 0, sizeof(nod));
	nod.op = OINDREG;
	nod.val.u.reg = D_SP;
	nod.addable = 1;

	nod.xoffset = fp-&gt;width;
	nod.type = fp-&gt;type;
	cgen_as(res, &amp;nod);
}

/*
 * call to n has already been generated.
 * generate:
 *	res = &amp;return value from call.
 */
void
cgen_aret(Node *n, Node *res)
{
	Node nod1, nod2;
	Type *fp, *t;
	Iter flist;

	t = n-&gt;left-&gt;type;
	if(isptr[t-&gt;etype])
		t = t-&gt;type;

	fp = structfirst(&amp;flist, getoutarg(t));
	if(fp == T)
		fatal(&#34;cgen_aret: nil&#34;);

	memset(&amp;nod1, 0, sizeof(nod1));
	nod1.op = OINDREG;
	nod1.val.u.reg = D_SP;
	nod1.addable = 1;

	nod1.xoffset = fp-&gt;width;
	nod1.type = fp-&gt;type;

	if(res-&gt;op != OREGISTER) {
		regalloc(&amp;nod2, types[tptr], res);
		gins(ALEAQ, &amp;nod1, &amp;nod2);
		gins(AMOVQ, &amp;nod2, res);
		regfree(&amp;nod2);
	} else
		gins(ALEAQ, &amp;nod1, res);
}

/*
 * generate return.
 * n-&gt;left is assignments to return values.
 */
void
cgen_ret(Node *n)
{
	genlist(n-&gt;list);		// copy out args
	if(hasdefer)
		ginscall(deferreturn, 0);
	gins(ARET, N, N);
}

/*
 * generate += *= etc.
 */
void
cgen_asop(Node *n)
{
	Node n1, n2, n3, n4;
	Node *nl, *nr;
	Prog *p1;
	Addr addr;
	int a;

	nl = n-&gt;left;
	nr = n-&gt;right;

	if(nr-&gt;ullman &gt;= UINF &amp;&amp; nl-&gt;ullman &gt;= UINF) {
		tempname(&amp;n1, nr-&gt;type);
		cgen(nr, &amp;n1);
		n2 = *n;
		n2.right = &amp;n1;
		cgen_asop(&amp;n2);
		goto ret;
	}

	if(!isint[nl-&gt;type-&gt;etype])
		goto hard;
	if(!isint[nr-&gt;type-&gt;etype])
		goto hard;

	switch(n-&gt;etype) {
	case OADD:
		if(smallintconst(nr))
		if(mpgetfix(nr-&gt;val.u.xval) == 1) {
			a = optoas(OINC, nl-&gt;type);
			if(nl-&gt;addable) {
				gins(a, N, nl);
				goto ret;
			}
			if(sudoaddable(a, nl, &amp;addr)) {
				p1 = gins(a, N, N);
				p1-&gt;to = addr;
				sudoclean();
				goto ret;
			}
		}
		break;

	case OSUB:
		if(smallintconst(nr))
		if(mpgetfix(nr-&gt;val.u.xval) == 1) {
			a = optoas(ODEC, nl-&gt;type);
			if(nl-&gt;addable) {
				gins(a, N, nl);
				goto ret;
			}
			if(sudoaddable(a, nl, &amp;addr)) {
				p1 = gins(a, N, N);
				p1-&gt;to = addr;
				sudoclean();
				goto ret;
			}
		}
		break;
	}

	switch(n-&gt;etype) {
	case OADD:
	case OSUB:
	case OXOR:
	case OAND:
	case OOR:
		a = optoas(n-&gt;etype, nl-&gt;type);
		if(nl-&gt;addable) {
			if(smallintconst(nr)) {
				gins(a, nr, nl);
				goto ret;
			}
			regalloc(&amp;n2, nr-&gt;type, N);
			cgen(nr, &amp;n2);
			gins(a, &amp;n2, nl);
			regfree(&amp;n2);
			goto ret;
		}
		if(nr-&gt;ullman &lt; UINF)
		if(sudoaddable(a, nl, &amp;addr)) {
			if(smallintconst(nr)) {
				p1 = gins(a, nr, N);
				p1-&gt;to = addr;
				sudoclean();
				goto ret;
			}
			regalloc(&amp;n2, nr-&gt;type, N);
			cgen(nr, &amp;n2);
			p1 = gins(a, &amp;n2, N);
			p1-&gt;to = addr;
			regfree(&amp;n2);
			sudoclean();
			goto ret;
		}
	}

hard:
	if(nr-&gt;ullman &gt; nl-&gt;ullman) {
		regalloc(&amp;n2, nr-&gt;type, N);
		cgen(nr, &amp;n2);
		igen(nl, &amp;n1, N);
	} else {
		igen(nl, &amp;n1, N);
		regalloc(&amp;n2, nr-&gt;type, N);
		cgen(nr, &amp;n2);
	}

	n3 = *n;
	n3.left = &amp;n1;
	n3.right = &amp;n2;
	n3.op = n-&gt;etype;

	regalloc(&amp;n4, nl-&gt;type, N);
	cgen(&amp;n3, &amp;n4);
	gmove(&amp;n4, &amp;n1);

	regfree(&amp;n1);
	regfree(&amp;n2);
	regfree(&amp;n4);

ret:
	;
}

int
samereg(Node *a, Node *b)
{
	if(a == N || b == N)
		return 0;
	if(a-&gt;op != OREGISTER)
		return 0;
	if(b-&gt;op != OREGISTER)
		return 0;
	if(a-&gt;val.u.reg != b-&gt;val.u.reg)
		return 0;
	return 1;
}

/*
 * generate division.
 * caller must set:
 *	ax = allocated AX register
 *	dx = allocated DX register
 * generates one of:
 *	res = nl / nr
 *	res = nl % nr
 * according to op.
 */
void
dodiv(int op, Node *nl, Node *nr, Node *res)
{
	int a;
	Node n3, n4;
	Type *t;
	Node ax, dx, oldax, olddx;

	t = nl-&gt;type;
	if(t-&gt;width == 1) {
		if(issigned[t-&gt;etype])
			t = types[TINT32];
		else
			t = types[TUINT32];
	}
	a = optoas(op, t);

	regalloc(&amp;n3, t, N);
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		savex(D_AX, &amp;ax, &amp;oldax, res, t);
		cgen(nl, &amp;ax);
		regalloc(&amp;ax, t, &amp;ax);	// mark ax live during cgen
		cgen(nr, &amp;n3);
		regfree(&amp;ax);
	} else {
		cgen(nr, &amp;n3);
		savex(D_AX, &amp;ax, &amp;oldax, res, t);
		cgen(nl, &amp;ax);
	}
	savex(D_DX, &amp;dx, &amp;olddx, res, t);
	if(!issigned[t-&gt;etype]) {
		nodconst(&amp;n4, t, 0);
		gmove(&amp;n4, &amp;dx);
	} else
		gins(optoas(OEXTEND, t), N, N);
	gins(a, &amp;n3, N);
	regfree(&amp;n3);

	if(op == ODIV)
		gmove(&amp;ax, res);
	else
		gmove(&amp;dx, res);
	restx(&amp;ax, &amp;oldax);
	restx(&amp;dx, &amp;olddx);
}

/*
 * register dr is one of the special ones (AX, CX, DI, SI, etc.).
 * we need to use it.  if it is already allocated as a temporary
 * (r &gt; 1; can only happen if a routine like sgen passed a
 * special as cgen&#39;s res and then cgen used regalloc to reuse
 * it as its own temporary), then move it for now to another
 * register.  caller must call restx to move it back.
 * the move is not necessary if dr == res, because res is
 * known to be dead.
 */
void
savex(int dr, Node *x, Node *oldx, Node *res, Type *t)
{
	int r;

	r = reg[dr];

	// save current ax and dx if they are live
	// and not the destination
	memset(oldx, 0, sizeof *oldx);
	nodreg(x, t, dr);
	if(r &gt; 1 &amp;&amp; !samereg(x, res)) {
		regalloc(oldx, types[TINT64], N);
		x-&gt;type = types[TINT64];
		gmove(x, oldx);
		x-&gt;type = t;
		oldx-&gt;ostk = r;	// squirrel away old r value
		reg[dr] = 1;
	}
}

void
restx(Node *x, Node *oldx)
{
	if(oldx-&gt;op != 0) {
		x-&gt;type = types[TINT64];
		reg[x-&gt;val.u.reg] = oldx-&gt;ostk;
		gmove(oldx, x);
		regfree(oldx);
	}
}

/*
 * generate division according to op, one of:
 *	res = nl / nr
 *	res = nl % nr
 */
void
cgen_div(int op, Node *nl, Node *nr, Node *res)
{
	Node n1, n2, n3, savl, savr;
	Node ax, dx, oldax, olddx;
	int n, w, s, a;
	Magic m;

	if(nl-&gt;ullman &gt;= UINF) {
		tempname(&amp;savl, nl-&gt;type);
		cgen(nl, &amp;savl);
		nl = &amp;savl;
	}
	if(nr-&gt;ullman &gt;= UINF) {
		tempname(&amp;savr, nr-&gt;type);
		cgen(nr, &amp;savr);
		nr = &amp;savr;
	}

	if(nr-&gt;op != OLITERAL)
		goto longdiv;

	// special cases of mod/div
	// by a constant
	w = nl-&gt;type-&gt;width*8;
	s = 0;
	n = powtwo(nr);
	if(n &gt;= 1000) {
		// negative power of 2
		s = 1;
		n -= 1000;
	}

	if(n+1 &gt;= w) {
		// just sign bit
		goto longdiv;
	}

	if(n &lt; 0)
		goto divbymul;
	switch(n) {
	case 0:
		// divide by 1
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		if(op == OMOD) {
			gins(optoas(OXOR, nl-&gt;type), &amp;n1, &amp;n1);
		} else
		if(s)
			gins(optoas(OMINUS, nl-&gt;type), N, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		return;
	case 1:
		// divide by 2
		if(op == OMOD) {
			if(issigned[nl-&gt;type-&gt;etype])
				goto longmod;
			regalloc(&amp;n1, nl-&gt;type, res);
			cgen(nl, &amp;n1);
			nodconst(&amp;n2, nl-&gt;type, 1);
			gins(optoas(OAND, nl-&gt;type), &amp;n2, &amp;n1);
			gmove(&amp;n1, res);
			regfree(&amp;n1);
			return;
		}
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		if(!issigned[nl-&gt;type-&gt;etype])
			break;

		// develop -1 iff nl is negative
		regalloc(&amp;n2, nl-&gt;type, N);
		gmove(&amp;n1, &amp;n2);
		nodconst(&amp;n3, nl-&gt;type, w-1);
		gins(optoas(ORSH, nl-&gt;type), &amp;n3, &amp;n2);
		gins(optoas(OSUB, nl-&gt;type), &amp;n2, &amp;n1);
		regfree(&amp;n2);
		break;
	default:
		if(op == OMOD) {
			if(issigned[nl-&gt;type-&gt;etype])
				goto longmod;
			regalloc(&amp;n1, nl-&gt;type, res);
			cgen(nl, &amp;n1);
			nodconst(&amp;n2, nl-&gt;type, mpgetfix(nr-&gt;val.u.xval)-1);
			if(!smallintconst(&amp;n2)) {
				regalloc(&amp;n3, nl-&gt;type, N);
				gmove(&amp;n2, &amp;n3);
				gins(optoas(OAND, nl-&gt;type), &amp;n3, &amp;n1);
				regfree(&amp;n3);
			} else
				gins(optoas(OAND, nl-&gt;type), &amp;n2, &amp;n1);
			gmove(&amp;n1, res);
			regfree(&amp;n1);
			return;
		}
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		if(!issigned[nl-&gt;type-&gt;etype])
			break;

		// develop (2^k)-1 iff nl is negative
		regalloc(&amp;n2, nl-&gt;type, N);
		gmove(&amp;n1, &amp;n2);
		nodconst(&amp;n3, nl-&gt;type, w-1);
		gins(optoas(ORSH, nl-&gt;type), &amp;n3, &amp;n2);
		nodconst(&amp;n3, nl-&gt;type, w-n);
		gins(optoas(ORSH, tounsigned(nl-&gt;type)), &amp;n3, &amp;n2);
		gins(optoas(OADD, nl-&gt;type), &amp;n2, &amp;n1);
		regfree(&amp;n2);
		break;
	}
	nodconst(&amp;n2, nl-&gt;type, n);
	gins(optoas(ORSH, nl-&gt;type), &amp;n2, &amp;n1);
	if(s)
		gins(optoas(OMINUS, nl-&gt;type), N, &amp;n1);
	gmove(&amp;n1, res);
	regfree(&amp;n1);
	return;

divbymul:
	// try to do division by multiply by (2^w)/d
	// see hacker&#39;s delight chapter 10
	switch(simtype[nl-&gt;type-&gt;etype]) {
	default:
		goto longdiv;

	case TUINT8:
	case TUINT16:
	case TUINT32:
	case TUINT64:
		m.w = w;
		m.ud = mpgetfix(nr-&gt;val.u.xval);
		umagic(&amp;m);
		if(m.bad)
			break;
		if(op == OMOD)
			goto longmod;

		regalloc(&amp;n1, nl-&gt;type, N);
		cgen(nl, &amp;n1);				// num -&gt; reg(n1)

		savex(D_AX, &amp;ax, &amp;oldax, res, nl-&gt;type);
		savex(D_DX, &amp;dx, &amp;olddx, res, nl-&gt;type);

		nodconst(&amp;n2, nl-&gt;type, m.um);
		gmove(&amp;n2, &amp;ax);			// const-&gt;ax

		gins(optoas(OHMUL, nl-&gt;type), &amp;n1, N);	// imul reg
		if(w == 8) {
			// fix up 8-bit multiply
			Node ah, dl;
			nodreg(&amp;ah, types[TUINT8], D_AH);
			nodreg(&amp;dl, types[TUINT8], D_DL);
			gins(AMOVB, &amp;ah, &amp;dl);
		}

		if(m.ua) {
			// need to add numerator accounting for overflow
			gins(optoas(OADD, nl-&gt;type), &amp;n1, &amp;dx);
			nodconst(&amp;n2, nl-&gt;type, 1);
			gins(optoas(ORRC, nl-&gt;type), &amp;n2, &amp;dx);
			nodconst(&amp;n2, nl-&gt;type, m.s-1);
			gins(optoas(ORSH, nl-&gt;type), &amp;n2, &amp;dx);
		} else {
			nodconst(&amp;n2, nl-&gt;type, m.s);
			gins(optoas(ORSH, nl-&gt;type), &amp;n2, &amp;dx);	// shift dx
		}


		regfree(&amp;n1);
		gmove(&amp;dx, res);

		restx(&amp;ax, &amp;oldax);
		restx(&amp;dx, &amp;olddx);
		return;

	case TINT8:
	case TINT16:
	case TINT32:
	case TINT64:
		m.w = w;
		m.sd = mpgetfix(nr-&gt;val.u.xval);
		smagic(&amp;m);
		if(m.bad)
			break;
		if(op == OMOD)
			goto longmod;

		regalloc(&amp;n1, nl-&gt;type, N);
		cgen(nl, &amp;n1);				// num -&gt; reg(n1)

		savex(D_AX, &amp;ax, &amp;oldax, res, nl-&gt;type);
		savex(D_DX, &amp;dx, &amp;olddx, res, nl-&gt;type);

		nodconst(&amp;n2, nl-&gt;type, m.sm);
		gmove(&amp;n2, &amp;ax);			// const-&gt;ax

		gins(optoas(OHMUL, nl-&gt;type), &amp;n1, N);	// imul reg
		if(w == 8) {
			// fix up 8-bit multiply
			Node ah, dl;
			nodreg(&amp;ah, types[TUINT8], D_AH);
			nodreg(&amp;dl, types[TUINT8], D_DL);
			gins(AMOVB, &amp;ah, &amp;dl);
		}

		if(m.sm &lt; 0) {
			// need to add numerator
			gins(optoas(OADD, nl-&gt;type), &amp;n1, &amp;dx);
		}

		nodconst(&amp;n2, nl-&gt;type, m.s);
		gins(optoas(ORSH, nl-&gt;type), &amp;n2, &amp;dx);	// shift dx

		nodconst(&amp;n2, nl-&gt;type, w-1);
		gins(optoas(ORSH, nl-&gt;type), &amp;n2, &amp;n1);	// -1 iff num is neg
		gins(optoas(OSUB, nl-&gt;type), &amp;n1, &amp;dx);	// added

		if(m.sd &lt; 0) {
			// this could probably be removed
			// by factoring it into the multiplier
			gins(optoas(OMINUS, nl-&gt;type), N, &amp;dx);
		}

		regfree(&amp;n1);
		gmove(&amp;dx, res);

		restx(&amp;ax, &amp;oldax);
		restx(&amp;dx, &amp;olddx);
		return;
	}
	goto longdiv;

longdiv:
	// division and mod using (slow) hardware instruction
	dodiv(op, nl, nr, res);
	return;

longmod:
	// mod using formula A%B = A-(A/B*B) but
	// we know that there is a fast algorithm for A/B
	regalloc(&amp;n1, nl-&gt;type, res);
	cgen(nl, &amp;n1);
	regalloc(&amp;n2, nl-&gt;type, N);
	cgen_div(ODIV, &amp;n1, nr, &amp;n2);
	a = optoas(OMUL, nl-&gt;type);
	if(w == 8) {
		// use 2-operand 16-bit multiply
		// because there is no 2-operand 8-bit multiply
		a = AIMULW;
	}
	if(!smallintconst(nr)) {
		regalloc(&amp;n3, nl-&gt;type, N);
		cgen(nr, &amp;n3);
		gins(a, &amp;n3, &amp;n2);
		regfree(&amp;n3);
	} else
		gins(a, nr, &amp;n2);
	gins(optoas(OSUB, nl-&gt;type), &amp;n2, &amp;n1);
	gmove(&amp;n1, res);
	regfree(&amp;n1);
	regfree(&amp;n2);
}

/*
 * generate shift according to op, one of:
 *	res = nl &lt;&lt; nr
 *	res = nl &gt;&gt; nr
 */
void
cgen_shift(int op, Node *nl, Node *nr, Node *res)
{
	Node n1, n2, n3, n4, n5, cx, oldcx;
	int a, rcx;
	Prog *p1;
	uvlong sc;

	a = optoas(op, nl-&gt;type);

	if(nr-&gt;op == OLITERAL) {
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		sc = mpgetfix(nr-&gt;val.u.xval);
		if(sc &gt;= nl-&gt;type-&gt;width*8) {
			// large shift gets 2 shifts by width
			nodconst(&amp;n3, types[TUINT32], nl-&gt;type-&gt;width*8-1);
			gins(a, &amp;n3, &amp;n1);
			gins(a, &amp;n3, &amp;n1);
		} else
			gins(a, nr, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		goto ret;
	}

	if(nl-&gt;ullman &gt;= UINF) {
		tempname(&amp;n4, nl-&gt;type);
		cgen(nl, &amp;n4);
		nl = &amp;n4;
	}
	if(nr-&gt;ullman &gt;= UINF) {
		tempname(&amp;n5, nr-&gt;type);
		cgen(nr, &amp;n5);
		nr = &amp;n5;
	}

	rcx = reg[D_CX];
	nodreg(&amp;n1, types[TUINT32], D_CX);
	regalloc(&amp;n1, nr-&gt;type, &amp;n1);		// to hold the shift type in CX
	regalloc(&amp;n3, types[TUINT64], &amp;n1);	// to clear high bits of CX

	nodreg(&amp;cx, types[TUINT64], D_CX);
	memset(&amp;oldcx, 0, sizeof oldcx);
	if(rcx &gt; 0 &amp;&amp; !samereg(&amp;cx, res)) {
		regalloc(&amp;oldcx, types[TUINT64], N);
		gmove(&amp;cx, &amp;oldcx);
	}

	if(samereg(&amp;cx, res))
		regalloc(&amp;n2, nl-&gt;type, N);
	else
		regalloc(&amp;n2, nl-&gt;type, res);
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		cgen(nl, &amp;n2);
		cgen(nr, &amp;n1);
		gmove(&amp;n1, &amp;n3);
	} else {
		cgen(nr, &amp;n1);
		gmove(&amp;n1, &amp;n3);
		cgen(nl, &amp;n2);
	}
	regfree(&amp;n3);

	// test and fix up large shifts
	nodconst(&amp;n3, types[TUINT64], nl-&gt;type-&gt;width*8);
	gins(optoas(OCMP, types[TUINT64]), &amp;n1, &amp;n3);
	p1 = gbranch(optoas(OLT, types[TUINT64]), T);
	if(op == ORSH &amp;&amp; issigned[nl-&gt;type-&gt;etype]) {
		nodconst(&amp;n3, types[TUINT32], nl-&gt;type-&gt;width*8-1);
		gins(a, &amp;n3, &amp;n2);
	} else {
		nodconst(&amp;n3, nl-&gt;type, 0);
		gmove(&amp;n3, &amp;n2);
	}
	patch(p1, pc);
	gins(a, &amp;n1, &amp;n2);

	if(oldcx.op != 0) {
		gmove(&amp;oldcx, &amp;cx);
		regfree(&amp;oldcx);
	}

	gmove(&amp;n2, res);

	regfree(&amp;n1);
	regfree(&amp;n2);

ret:
	;
}

/*
 * generate byte multiply:
 *	res = nl * nr
 * no 2-operand byte multiply instruction so have to do
 * 16-bit multiply and take bottom half.
 */
void
cgen_bmul(int op, Node *nl, Node *nr, Node *res)
{
	Node n1b, n2b, n1w, n2w;
	Type *t;
	int a;

	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		regalloc(&amp;n1b, nl-&gt;type, res);
		cgen(nl, &amp;n1b);
		regalloc(&amp;n2b, nr-&gt;type, N);
		cgen(nr, &amp;n2b);
	} else {
		regalloc(&amp;n2b, nr-&gt;type, N);
		cgen(nr, &amp;n2b);
		regalloc(&amp;n1b, nl-&gt;type, res);
		cgen(nl, &amp;n1b);
	}

	// copy from byte to short registers
	t = types[TUINT16];
	if(issigned[nl-&gt;type-&gt;etype])
		t = types[TINT16];

	regalloc(&amp;n2w, t, &amp;n2b);
	cgen(&amp;n2b, &amp;n2w);

	regalloc(&amp;n1w, t, &amp;n1b);
	cgen(&amp;n1b, &amp;n1w);

	a = optoas(op, t);
	gins(a, &amp;n2w, &amp;n1w);
	cgen(&amp;n1w, &amp;n1b);
	cgen(&amp;n1b, res);

	regfree(&amp;n1w);
	regfree(&amp;n2w);
	regfree(&amp;n1b);
	regfree(&amp;n2b);
}

void
clearfat(Node *nl)
{
	uint32 w, c, q;
	Node n1, oldn1, ax, oldax;

	/* clear a fat object */
	if(debug[&#39;g&#39;])
		dump(&#34;\nclearfat&#34;, nl);

	w = nl-&gt;type-&gt;width;
	c = w % 8;	// bytes
	q = w / 8;	// quads

	savex(D_DI, &amp;n1, &amp;oldn1, N, types[tptr]);
	agen(nl, &amp;n1);

	savex(D_AX, &amp;ax, &amp;oldax, N, types[tptr]);
	gconreg(AMOVQ, 0, D_AX);

	if(q &gt;= 4) {
		gconreg(AMOVQ, q, D_CX);
		gins(AREP, N, N);	// repeat
		gins(ASTOSQ, N, N);	// STOQ AL,*(DI)+
	} else
	while(q &gt; 0) {
		gins(ASTOSQ, N, N);	// STOQ AL,*(DI)+
		q--;
	}

	if(c &gt;= 4) {
		gconreg(AMOVQ, c, D_CX);
		gins(AREP, N, N);	// repeat
		gins(ASTOSB, N, N);	// STOB AL,*(DI)+
	} else
	while(c &gt; 0) {
		gins(ASTOSB, N, N);	// STOB AL,*(DI)+
		c--;
	}

	restx(&amp;n1, &amp;oldn1);
	restx(&amp;ax, &amp;oldax);
}

static int
regcmp(const void *va, const void *vb)
{
	Node *ra, *rb;

	ra = (Node*)va;
	rb = (Node*)vb;
	return ra-&gt;local - rb-&gt;local;
}

static	Prog*	throwpc;

void
getargs(NodeList *nn, Node *reg, int n)
{
	NodeList *l;
	int i;

	throwpc = nil;

	l = nn;
	for(i=0; i&lt;n; i++) {
		if(!smallintconst(l-&gt;n-&gt;right) &amp;&amp; !isslice(l-&gt;n-&gt;right-&gt;type)) {
			regalloc(reg+i, l-&gt;n-&gt;right-&gt;type, N);
			cgen(l-&gt;n-&gt;right, reg+i);
		} else
			reg[i] = *l-&gt;n-&gt;right;
		if(reg[i].local != 0)
			yyerror(&#34;local used&#34;);
		reg[i].local = l-&gt;n-&gt;left-&gt;xoffset;
		l = l-&gt;next;
	}
	qsort((void*)reg, n, sizeof(*reg), regcmp);
	for(i=0; i&lt;n; i++)
		reg[i].local = 0;
}

void
cmpandthrow(Node *nl, Node *nr)
{
	vlong cl, cr;
	Prog *p1;
	int op;
	Node *c;

	op = OLE;
	if(smallintconst(nl)) {
		cl = mpgetfix(nl-&gt;val.u.xval);
		if(cl == 0)
			return;
		if(smallintconst(nr)) {
			cr = mpgetfix(nr-&gt;val.u.xval);
			if(cl &gt; cr) {
				if(throwpc == nil) {
					throwpc = pc;
					ginscall(throwslice, 0);
				} else
					patch(gbranch(AJMP, T), throwpc);
			}
			return;
		}

		// put the constant on the right
		op = brrev(op);
		c = nl;
		nl = nr;
		nr = c;
	}

	gins(optoas(OCMP, types[TUINT32]), nl, nr);
	if(throwpc == nil) {
		p1 = gbranch(optoas(op, types[TUINT32]), T);
		throwpc = pc;
		ginscall(throwslice, 0);
		patch(p1, pc);
	} else {
		op = brcom(op);
		p1 = gbranch(optoas(op, types[TUINT32]), T);
		patch(p1, throwpc);
	}
}

int
sleasy(Node *n)
{
	if(n-&gt;op != ONAME)
		return 0;
	if(!n-&gt;addable)
		return 0;
	return 1;
}

// generate inline code for
//	slicearray
//	sliceslice
//	arraytoslice
int
cgen_inline(Node *n, Node *res)
{
	Node nodes[5];
	Node n1, n2, nres, nnode0, ntemp;
	vlong v;
	int i, bad;

	if(n-&gt;op != OCALLFUNC)
		goto no;
	if(!n-&gt;left-&gt;addable)
		goto no;
	if(strcmp(n-&gt;left-&gt;sym-&gt;package, &#34;runtime&#34;) != 0)
		goto no;
	if(strcmp(n-&gt;left-&gt;sym-&gt;name, &#34;slicearray&#34;) == 0)
		goto slicearray;
	if(strcmp(n-&gt;left-&gt;sym-&gt;name, &#34;sliceslice&#34;) == 0)
		goto sliceslice;
	if(strcmp(n-&gt;left-&gt;sym-&gt;name, &#34;arraytoslice&#34;) == 0)
		goto arraytoslice;
	goto no;

slicearray:
	if(!sleasy(res))
		goto no;
	getargs(n-&gt;list, nodes, 5);

	// if(hb[3] &gt; nel[1]) goto throw
	cmpandthrow(&amp;nodes[3], &amp;nodes[1]);

	// if(lb[2] &gt; hb[3]) goto throw
	cmpandthrow(&amp;nodes[2], &amp;nodes[3]);

	// len = hb[3] - lb[2] (destroys hb)
	n2 = *res;
	n2.xoffset += Array_nel;

	if(smallintconst(&amp;nodes[3]) &amp;&amp; smallintconst(&amp;nodes[2])) {
		v = mpgetfix(nodes[3].val.u.xval) -
			mpgetfix(nodes[2].val.u.xval);
		nodconst(&amp;n1, types[TUINT32], v);
		gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
	} else {
		regalloc(&amp;n1, types[TUINT32], &amp;nodes[3]);
		gmove(&amp;nodes[3], &amp;n1);
		if(!smallintconst(&amp;nodes[2]) || mpgetfix(nodes[2].val.u.xval) != 0)
			gins(optoas(OSUB, types[TUINT32]), &amp;nodes[2], &amp;n1);
		gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
		regfree(&amp;n1);
	}

	// cap = nel[1] - lb[2] (destroys nel)
	n2 = *res;
	n2.xoffset += Array_cap;

	if(smallintconst(&amp;nodes[1]) &amp;&amp; smallintconst(&amp;nodes[2])) {
		v = mpgetfix(nodes[1].val.u.xval) -
			mpgetfix(nodes[2].val.u.xval);
		nodconst(&amp;n1, types[TUINT32], v);
		gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
	} else {
		regalloc(&amp;n1, types[TUINT32], &amp;nodes[1]);
		gmove(&amp;nodes[1], &amp;n1);
		if(!smallintconst(&amp;nodes[2]) || mpgetfix(nodes[2].val.u.xval) != 0)
			gins(optoas(OSUB, types[TUINT32]), &amp;nodes[2], &amp;n1);
		gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
		regfree(&amp;n1);
	}

	// if slice could be too big, dereference to
	// catch nil array pointer.
	if(nodes[0].op == OREGISTER &amp;&amp; nodes[0].type-&gt;type-&gt;width &gt;= unmappedzero) {
		n2 = nodes[0];
		n2.xoffset = 0;
		n2.op = OINDREG;
		n2.type = types[TUINT8];
		gins(ATESTB, nodintconst(0), &amp;n2);
	}

	// ary = old[0] + (lb[2] * width[4]) (destroys old)
	n2 = *res;
	n2.xoffset += Array_array;

	if(smallintconst(&amp;nodes[2]) &amp;&amp; smallintconst(&amp;nodes[4])) {
		v = mpgetfix(nodes[2].val.u.xval) *
			mpgetfix(nodes[4].val.u.xval);
		if(v != 0) {
			nodconst(&amp;n1, types[tptr], v);
			gins(optoas(OADD, types[tptr]), &amp;n1, &amp;nodes[0]);
		}
	} else {
		regalloc(&amp;n1, types[tptr], &amp;nodes[2]);
		gmove(&amp;nodes[2], &amp;n1);
		if(!smallintconst(&amp;nodes[4]) || mpgetfix(nodes[4].val.u.xval) != 1)
			gins(optoas(OMUL, types[tptr]), &amp;nodes[4], &amp;n1);
		gins(optoas(OADD, types[tptr]), &amp;n1, &amp;nodes[0]);
		regfree(&amp;n1);
	}
	gins(optoas(OAS, types[tptr]), &amp;nodes[0], &amp;n2);

	for(i=0; i&lt;5; i++) {
		if(nodes[i].op == OREGISTER)
			regfree(&amp;nodes[i]);
	}
	return 1;

arraytoslice:
	if(!sleasy(res))
		goto no;
	getargs(n-&gt;list, nodes, 2);

	// ret.len = nel[1];
	n2 = *res;
	n2.xoffset += Array_nel;
	gins(optoas(OAS, types[TUINT32]), &amp;nodes[1], &amp;n2);

	// ret.cap = nel[1];
	n2 = *res;
	n2.xoffset += Array_cap;
	gins(optoas(OAS, types[TUINT32]), &amp;nodes[1], &amp;n2);

	// ret.array = old[0];
	n2 = *res;
	n2.xoffset += Array_array;
	gins(optoas(OAS, types[tptr]), &amp;nodes[0], &amp;n2);

	// if slice could be too big, dereference to
	// catch nil array pointer.
	if(nodes[0].op == OREGISTER &amp;&amp; nodes[0].type-&gt;type-&gt;width &gt;= unmappedzero) {
		n2 = nodes[0];
		n2.xoffset = 0;
		n2.op = OINDREG;
		n2.type = types[TUINT8];
		gins(ATESTB, nodintconst(0), &amp;n2);
	}

	for(i=0; i&lt;2; i++) {
		if(nodes[i].op == OREGISTER)
			regfree(&amp;nodes[i]);
	}
	return 1;

sliceslice:
	getargs(n-&gt;list, nodes, 4);

	nres = *res;		// result
	nnode0 = nodes[0];	// input slice
	if(!sleasy(res) || !sleasy(&amp;nodes[0])) {
		bad = 0;
		if(res-&gt;ullman &gt;= UINF)
			bad = 1;
		for(i=0; i&lt;4; i++) {
			if(nodes[i].ullman &gt;= UINF)
				bad = 1;
			if(nodes[i].op == OREGISTER)
				regfree(&amp;nodes[i]);
		}

		if(bad)
			goto no;

		tempname(&amp;ntemp, res-&gt;type);
		if(!sleasy(&amp;nodes[0])) {
			cgen(&amp;nodes[0], &amp;ntemp);
			nnode0 = ntemp;
		}
		getargs(n-&gt;list, nodes, 4);
		if(!sleasy(res))
			nres = ntemp;
	}

	// if(hb[2] &gt; old.cap[0]) goto throw;
	n2 = nnode0;
	n2.xoffset += Array_cap;
	cmpandthrow(&amp;nodes[2], &amp;n2);

	// if(lb[1] &gt; hb[2]) goto throw;
	cmpandthrow(&amp;nodes[1], &amp;nodes[2]);

	// ret.len = hb[2]-lb[1]; (destroys hb[2])
	n2 = nres;
	n2.xoffset += Array_nel;

	if(smallintconst(&amp;nodes[2]) &amp;&amp; smallintconst(&amp;nodes[1])) {
		v = mpgetfix(nodes[2].val.u.xval) -
			mpgetfix(nodes[1].val.u.xval);
		nodconst(&amp;n1, types[TUINT32], v);
		gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
	} else {
		regalloc(&amp;n1, types[TUINT32], &amp;nodes[2]);
		gmove(&amp;nodes[2], &amp;n1);
		if(!smallintconst(&amp;nodes[1]) || mpgetfix(nodes[1].val.u.xval) != 0)
			gins(optoas(OSUB, types[TUINT32]), &amp;nodes[1], &amp;n1);
		gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
		regfree(&amp;n1);
	}

	// ret.cap = old.cap[0]-lb[1]; (uses hb[2])
	n2 = nnode0;
	n2.xoffset += Array_cap;

	regalloc(&amp;n1, types[TUINT32], &amp;nodes[2]);
	gins(optoas(OAS, types[TUINT32]), &amp;n2, &amp;n1);
	if(!smallintconst(&amp;nodes[1]) || mpgetfix(nodes[1].val.u.xval) != 0)
		gins(optoas(OSUB, types[TUINT32]), &amp;nodes[1], &amp;n1);

	n2 = nres;
	n2.xoffset += Array_cap;
	gins(optoas(OAS, types[TUINT32]), &amp;n1, &amp;n2);
	regfree(&amp;n1);

	// ret.array = old.array[0]+lb[1]*width[3]; (uses lb[1])
	n2 = nnode0;
	n2.xoffset += Array_array;

	regalloc(&amp;n1, types[tptr], &amp;nodes[1]);
	if(smallintconst(&amp;nodes[1]) &amp;&amp; smallintconst(&amp;nodes[3])) {
		gins(optoas(OAS, types[tptr]), &amp;n2, &amp;n1);
		v = mpgetfix(nodes[1].val.u.xval) *
			mpgetfix(nodes[3].val.u.xval);
		if(v != 0) {
			nodconst(&amp;n2, types[tptr], v);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n1);
		}
	} else {
		gmove(&amp;nodes[1], &amp;n1);
		if(!smallintconst(&amp;nodes[3]) || mpgetfix(nodes[3].val.u.xval) != 1)
			gins(optoas(OMUL, types[tptr]), &amp;nodes[3], &amp;n1);
		gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n1);
	}

	n2 = nres;
	n2.xoffset += Array_array;
	gins(optoas(OAS, types[tptr]), &amp;n1, &amp;n2);
	regfree(&amp;n1);

	for(i=0; i&lt;4; i++) {
		if(nodes[i].op == OREGISTER)
			regfree(&amp;nodes[i]);
	}

	if(!sleasy(res)) {
		cgen(&amp;nres, res);
	}
	return 1;

no:
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
