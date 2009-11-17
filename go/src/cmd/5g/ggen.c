<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5g/ggen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5g/ggen.c</h1>

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

	genlist(curfn-&gt;enter);
	genlist(curfn-&gt;nbody);
	checklabels();
	if(nerrors != 0)
		goto ret;

	if(curfn-&gt;type-&gt;outtuple != 0)
		ginscall(throwreturn, 0);

	if(hasdefer)
		ginscall(deferreturn, 0);
	pc-&gt;as = ARET;	// overwrite AEND
	pc-&gt;lineno = lineno;

	/* TODO(kaib): Add back register optimizations
	if(!debug[&#39;N&#39;] || debug[&#39;R&#39;] || debug[&#39;P&#39;])
		regopt(ptxt);
	*/

	// fill in argument size
	ptxt-&gt;to.type = D_CONST2;
	ptxt-&gt;reg = 0; // flags
	ptxt-&gt;to.offset2 = rnd(curfn-&gt;type-&gt;argwid, maxround);

	// fill in final stack size
	if(stksize &gt; maxstksize)
		maxstksize = stksize;
	ptxt-&gt;to.offset = rnd(maxstksize+maxarg, maxround);
	maxstksize = 0;

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
	Node n1, r, con;

	switch(proc) {
	default:
		fatal(&#34;ginscall: bad proc %d&#34;, proc);
		break;

	case 0:	// normal call
		p = gins(ABL, N, f);
		afunclit(&amp;p-&gt;to);
		break;

	// TODO(kaib): unify newproc and defer if you can figure out how not to break things
	case 1:	// call in new proc (go)
		regalloc(&amp;r, types[tptr], N);
		p = gins(AMOVW, N, &amp;r);
		p-&gt;from.type = D_OREG;
		p-&gt;from.reg = REGSP;
		
		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = -12;
		p-&gt;scond |= C_WBIT;

		memset(&amp;n1, 0, sizeof n1);
		n1.op = OADDR;
		n1.left = f;
		gins(AMOVW, &amp;n1, &amp;r);

		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = 8;

		nodconst(&amp;con, types[TINT32], argsize(f-&gt;type) + 4);
		gins(AMOVW, &amp;con, &amp;r);
		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = 4;
		regfree(&amp;r);

		ginscall(newproc, 0);

		regalloc(&amp;r, types[tptr], N);
		p = gins(AMOVW, N, &amp;r);
		p-&gt;from.type = D_OREG;
		p-&gt;from.reg = REGSP;
		p-&gt;from.offset = 0;

		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = 12;
		p-&gt;scond |= C_WBIT;
		regfree(&amp;r);

		break;

	case 2:	// defered call (defer)
		regalloc(&amp;r, types[tptr], N);
		p = gins(AMOVW, N, &amp;r);
		p-&gt;from.type = D_OREG;
		p-&gt;from.reg = REGSP;
		
		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = -8;
		p-&gt;scond |= C_WBIT;

		memset(&amp;n1, 0, sizeof n1);
		n1.op = OADDR;
		n1.left = f;
		gins(AMOVW, &amp;n1, &amp;r);

		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = 8;

		nodconst(&amp;con, types[TINT32], argsize(f-&gt;type));
		gins(AMOVW, &amp;con, &amp;r);
		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = 4;
		regfree(&amp;r);

		ginscall(deferproc, 0);

		regalloc(&amp;r, types[tptr], N);
		p = gins(AMOVW, N, &amp;r);
		p-&gt;from.type = D_OREG;
		p-&gt;from.reg = REGSP;
		p-&gt;from.offset = 0;

		p = gins(AMOVW, &amp;r, N);
		p-&gt;to.type = D_OREG;
		p-&gt;to.reg = REGSP;
		p-&gt;to.offset = 8;
		p-&gt;scond |= C_WBIT;
		regfree(&amp;r);

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

	genlist(n-&gt;list);			// args

	regalloc(&amp;nodr, types[tptr], res);
	regalloc(&amp;nodo, types[tptr], &amp;nodr);
	nodo.op = OINDREG;

	agen(i, &amp;nodr);		// REG = &amp;inter

	nodindreg(&amp;nodsp, types[tptr], REGSP);
	nodsp.xoffset = 4;
	nodo.xoffset += widthptr;
	cgen(&amp;nodo, &amp;nodsp);	// 4(SP) = 8(REG) -- i.s

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
	nod.val.u.reg = REGSP;
	nod.addable = 1;

	nod.xoffset = fp-&gt;width + 4; // +4: saved lr at 0(SP)
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
	nod1.val.u.reg = REGSP;
	nod1.addable = 1;

	nod1.xoffset = fp-&gt;width + 4; // +4: saved lr at 0(SP)
	nod1.type = fp-&gt;type;

	if(res-&gt;op != OREGISTER) {
		regalloc(&amp;nod2, types[tptr], res);
		agen(&amp;nod1, &amp;nod2);
		gins(AMOVW, &amp;nod2, res);
		regfree(&amp;nod2);
	} else
		agen(&amp;nod1, res);
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
	int a, w;

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
	if(is64(nl-&gt;type) || is64(nr-&gt;type))
		goto hard64;

	switch(n-&gt;etype) {
	case OADD:
	case OSUB:
	case OXOR:
	case OAND:
	case OOR:
		a = optoas(n-&gt;etype, nl-&gt;type);
		if(nl-&gt;addable) {
			regalloc(&amp;n2, nl-&gt;type, N);
			regalloc(&amp;n3, nr-&gt;type, N);
			cgen(nl, &amp;n2);
			cgen(nr, &amp;n3);
			gins(a, &amp;n3, &amp;n2);
			cgen(&amp;n2, nl);
			regfree(&amp;n2);
			regfree(&amp;n3);
			goto ret;
		}
		if(nr-&gt;ullman &lt; UINF)
		if(sudoaddable(a, nl, &amp;addr, &amp;w)) {
			regalloc(&amp;n2, nl-&gt;type, N);
			regalloc(&amp;n3, nr-&gt;type, N);
			p1 = gins(AMOVW, N, &amp;n2);
			p1-&gt;from = addr;
			cgen(nr, &amp;n3);
			gins(a, &amp;n3, &amp;n2);
			p1 = gins(AMOVW, &amp;n2, N);
			p1-&gt;to = addr;
			regfree(&amp;n2);
			regfree(&amp;n3);
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
	goto ret;

hard64:
	if(nr-&gt;ullman &gt; nl-&gt;ullman) {
		tempname(&amp;n2, nr-&gt;type);
		cgen(nr, &amp;n2);
		igen(nl, &amp;n1, N);
	} else {
		igen(nl, &amp;n1, N);
		tempname(&amp;n2, nr-&gt;type);
		cgen(nr, &amp;n2);
	}

	n3 = *n;
	n3.left = &amp;n1;
	n3.right = &amp;n2;
	n3.op = n-&gt;etype;

	cgen(&amp;n3, &amp;n1);

ret:
	;
}

int
samereg(Node *a, Node *b)
{
	if(a-&gt;op != OREGISTER)
		return 0;
	if(b-&gt;op != OREGISTER)
		return 0;
	if(a-&gt;val.u.reg != b-&gt;val.u.reg)
		return 0;
	return 1;
}

/*
 * generate shift according to op, one of:
 *	res = nl &lt;&lt; nr
 *	res = nl &gt;&gt; nr
 */
void
cgen_shift(int op, Node *nl, Node *nr, Node *res)
{
	Node n1, n2, n3, t;
	int w;
	Prog *p1, *p2, *p3;
	uvlong sc;

	if(nl-&gt;type-&gt;width &gt; 4)
		fatal(&#34;cgen_shift %T&#34;, nl-&gt;type);

	w = nl-&gt;type-&gt;width * 8;

	if(nr-&gt;op == OLITERAL) {
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		sc = mpgetfix(nr-&gt;val.u.xval);
		if(sc == 0) {
			return;
		} else if(sc &gt;= nl-&gt;type-&gt;width*8) {
			if(op == ORSH &amp;&amp; issigned[nl-&gt;type-&gt;etype])
				gshift(AMOVW, &amp;n1, SHIFT_AR, w, &amp;n1);
			else
				gins(AEOR, &amp;n1, &amp;n1);
		} else {
			if(op == ORSH &amp;&amp; issigned[nl-&gt;type-&gt;etype])
				gshift(AMOVW, &amp;n1, SHIFT_AR, sc, &amp;n1);
			else if(op == ORSH)
				gshift(AMOVW, &amp;n1, SHIFT_LR, sc, &amp;n1);
			else // OLSH
				gshift(AMOVW, &amp;n1, SHIFT_LL, sc, &amp;n1);
		}
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		return;
	}

	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		regalloc(&amp;n2, nl-&gt;type, res);
		cgen(nl, &amp;n2);
		regalloc(&amp;n1, nr-&gt;type, N);
		cgen(nr, &amp;n1);
	} else {
		regalloc(&amp;n1, nr-&gt;type, N);
		cgen(nr, &amp;n1);
		regalloc(&amp;n2, nl-&gt;type, res);
		cgen(nl, &amp;n2);
	}

	// test for shift being 0
	p1 = gins(AMOVW, &amp;n1, &amp;n1);
	p1-&gt;scond |= C_SBIT;
	p3 = gbranch(ABEQ, T);

	// test and fix up large shifts
	regalloc(&amp;n3, nr-&gt;type, N);
	nodconst(&amp;t, types[TUINT32], w);
	gmove(&amp;t, &amp;n3);
	gcmp(ACMP, &amp;n1, &amp;n3);
	if(op == ORSH) {
		if(issigned[nl-&gt;type-&gt;etype]) {
			p1 = gshift(AMOVW, &amp;n2, SHIFT_AR, w-1, &amp;n2);
			p2 = gregshift(AMOVW, &amp;n2, SHIFT_AR, &amp;n1, &amp;n2);
		} else {
			p1 = gins(AEOR, &amp;n2, &amp;n2);
			p2 = gregshift(AMOVW, &amp;n2, SHIFT_LR, &amp;n1, &amp;n2);
		}
		p1-&gt;scond = C_SCOND_HS;
		p2-&gt;scond = C_SCOND_LO;
	} else {
		p1 = gins(AEOR, &amp;n2, &amp;n2);
		p2 = gregshift(AMOVW, &amp;n2, SHIFT_LL, &amp;n1, &amp;n2);
		p1-&gt;scond = C_SCOND_HS;
		p2-&gt;scond = C_SCOND_LO;
	}
	regfree(&amp;n3);

	patch(p3, pc);
	gmove(&amp;n2, res);

	regfree(&amp;n1);
	regfree(&amp;n2);
}

void
clearfat(Node *nl)
{
	uint32 w, c, q;
	Node dst, nc, nz, end;
	Prog *p, *pl;

	/* clear a fat object */
	if(debug[&#39;g&#39;])
		dump(&#34;\nclearfat&#34;, nl);

	w = nl-&gt;type-&gt;width;
	c = w % 4;	// bytes
	q = w / 4;	// quads

	regalloc(&amp;dst, types[tptr], N);
	agen(nl, &amp;dst);
	nodconst(&amp;nc, types[TUINT32], 0);
	regalloc(&amp;nz, types[TUINT32], 0);
	cgen(&amp;nc, &amp;nz);

	if(q &gt;= 4) {
		regalloc(&amp;end, types[tptr], N);
		p = gins(AMOVW, &amp;dst, &amp;end);
		p-&gt;from.type = D_CONST;
		p-&gt;from.offset = q*4;

		p = gins(AMOVW, &amp;nz, &amp;dst);
		p-&gt;to.type = D_OREG;
		p-&gt;to.offset = 4;
		p-&gt;scond |= C_PBIT;
		pl = p;

		p = gins(ACMP, &amp;dst, N);
		raddr(&amp;end, p);
		patch(gbranch(ABNE, T), pl);

		regfree(&amp;end);
	} else
	while(q &gt; 0) {
		p = gins(AMOVW, &amp;nz, &amp;dst);
		p-&gt;to.type = D_OREG;
		p-&gt;to.offset = 4;
 		p-&gt;scond |= C_PBIT;
//print(&#34;1. %P\n&#34;, p);
		q--;
	}

	while(c &gt; 0) {
		p = gins(AMOVBU, &amp;nz, &amp;dst);
		p-&gt;to.type = D_OREG;
		p-&gt;to.offset = 1;
 		p-&gt;scond |= C_PBIT;
//print(&#34;2. %P\n&#34;, p);
		c--;
	}
	regfree(&amp;dst);
	regfree(&amp;nz);
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
