<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8g/ggen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8g/ggen.c</h1>

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

	if(!debug[&#39;N&#39;] || debug[&#39;R&#39;] || debug[&#39;P&#39;]) {
		regopt(ptxt);
	}
	// fill in argument size
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

void
clearfat(Node *nl)
{
	uint32 w, c, q;
	Node n1;

	/* clear a fat object */
	if(debug[&#39;g&#39;])
		dump(&#34;\nclearfat&#34;, nl);

	w = nl-&gt;type-&gt;width;
	c = w % 4;	// bytes
	q = w / 4;	// quads

	gconreg(AMOVL, 0, D_AX);
	nodreg(&amp;n1, types[tptr], D_DI);
	agen(nl, &amp;n1);

	if(q &gt;= 4) {
		gconreg(AMOVL, q, D_CX);
		gins(AREP, N, N);	// repeat
		gins(ASTOSL, N, N);	// STOL AL,*(DI)+
	} else
	while(q &gt; 0) {
		gins(ASTOSL, N, N);	// STOL AL,*(DI)+
		q--;
	}

	if(c &gt;= 4) {
		gconreg(AMOVL, c, D_CX);
		gins(AREP, N, N);	// repeat
		gins(ASTOSB, N, N);	// STOB AL,*(DI)+
	} else
	while(c &gt; 0) {
		gins(ASTOSB, N, N);	// STOB AL,*(DI)+
		c--;
	}
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
		nodreg(&amp;reg, types[TINT32], D_AX);
		gins(APUSHL, f, N);
		nodconst(&amp;con, types[TINT32], argsize(f-&gt;type));
		gins(APUSHL, &amp;con, N);
		if(proc == 1)
			ginscall(newproc, 0);
		else
			ginscall(deferproc, 0);
		gins(APOPL, N, &amp;reg);
		gins(APOPL, N, &amp;reg);
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
		tempalloc(&amp;tmpi, i-&gt;type);
		cgen(i, &amp;tmpi);
		i = &amp;tmpi;
	}

	genlist(n-&gt;list);		// assign the args

	// Can regalloc now; i is known to be addable,
	// so the agen will be easy.
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

	if(i == &amp;tmpi)
		tempfree(i);

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
		tempalloc(&amp;afun, types[tptr]);
		cgen(n-&gt;left, &amp;afun);
	}

	genlist(n-&gt;list);		// assign the args
	t = n-&gt;left-&gt;type;

	setmaxarg(t);

	// call tempname pointer
	if(n-&gt;left-&gt;ullman &gt;= UINF) {
		regalloc(&amp;nod, types[tptr], N);
		cgen_as(&amp;nod, &amp;afun);
		tempfree(&amp;afun);
		nod.type = t;
		ginscall(&amp;nod, proc);
		regfree(&amp;nod);
		return;
	}

	// call pointer
	if(n-&gt;left-&gt;op != ONAME || n-&gt;left-&gt;class != PFUNC) {
		regalloc(&amp;nod, types[tptr], N);
		cgen_as(&amp;nod, n-&gt;left);
		nod.type = t;
		ginscall(&amp;nod, proc);
		regfree(&amp;nod);
		return;
	}

	// call direct
	n-&gt;left-&gt;method = 1;
	ginscall(n-&gt;left, proc);
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
		gins(ALEAL, &amp;nod1, &amp;nod2);
		gins(AMOVL, &amp;nod2, res);
		regfree(&amp;nod2);
	} else
		gins(ALEAL, &amp;nod1, res);
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
		tempalloc(&amp;n1, nr-&gt;type);
		cgen(nr, &amp;n1);
		n2 = *n;
		n2.right = &amp;n1;
		cgen_asop(&amp;n2);
		tempfree(&amp;n1);
		goto ret;
	}

	if(!isint[nl-&gt;type-&gt;etype])
		goto hard;
	if(!isint[nr-&gt;type-&gt;etype])
		goto hard;
	if(is64(nl-&gt;type) || is64(nr-&gt;type))
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
		tempalloc(&amp;n2, nr-&gt;type);
		cgen(nr, &amp;n2);
		igen(nl, &amp;n1, N);
	} else {
		igen(nl, &amp;n1, N);
		tempalloc(&amp;n2, nr-&gt;type);
		cgen(nr, &amp;n2);
	}

	n3 = *n;
	n3.left = &amp;n1;
	n3.right = &amp;n2;
	n3.op = n-&gt;etype;

	tempalloc(&amp;n4, nl-&gt;type);
	cgen(&amp;n3, &amp;n4);
	gmove(&amp;n4, &amp;n1);

	regfree(&amp;n1);
	tempfree(&amp;n4);
	tempfree(&amp;n2);

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
dodiv(int op, Type *t, Node *nl, Node *nr, Node *res, Node *ax, Node *dx)
{
	Node n1, t1, t2, nz;

	tempalloc(&amp;t1, nl-&gt;type);
	tempalloc(&amp;t2, nr-&gt;type);
	cgen(nl, &amp;t1);
	cgen(nr, &amp;t2);

	if(!samereg(ax, res) &amp;&amp; !samereg(dx, res))
		regalloc(&amp;n1, t, res);
	else
		regalloc(&amp;n1, t, N);
	gmove(&amp;t2, &amp;n1);
	gmove(&amp;t1, ax);
	if(!issigned[t-&gt;etype]) {
		nodconst(&amp;nz, t, 0);
		gmove(&amp;nz, dx);
	} else
		gins(optoas(OEXTEND, t), N, N);
	gins(optoas(op, t), &amp;n1, N);
	regfree(&amp;n1);
	tempfree(&amp;t2);
	tempfree(&amp;t1);

	if(op == ODIV)
		gmove(ax, res);
	else
		gmove(dx, res);
}

static void
savex(int dr, Node *x, Node *oldx, Node *res, Type *t)
{
	int r;

	r = reg[dr];
	nodreg(x, types[TINT32], dr);

	// save current ax and dx if they are live
	// and not the destination
	memset(oldx, 0, sizeof *oldx);
	if(r &gt; 0 &amp;&amp; !samereg(x, res)) {
		tempalloc(oldx, types[TINT32]);
		gmove(x, oldx);
	}

	regalloc(x, t, x);
}

static void
restx(Node *x, Node *oldx)
{
	regfree(x);

	if(oldx-&gt;op != 0) {
		x-&gt;type = types[TINT32];
		gmove(oldx, x);
		tempfree(oldx);
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
	Node ax, dx, oldax, olddx;
	int rax, rdx;
	Type *t;

	rax = reg[D_AX];
	rdx = reg[D_DX];

	if(is64(nl-&gt;type))
		fatal(&#34;cgen_div %T&#34;, nl-&gt;type);

	t = nl-&gt;type;
	if(t-&gt;width == 1)
		t = types[t-&gt;etype+2];	// int8 -&gt; int16, uint8 -&gt; uint16

	savex(D_AX, &amp;ax, &amp;oldax, res, t);
	savex(D_DX, &amp;dx, &amp;olddx, res, t);
	dodiv(op, t, nl, nr, res, &amp;ax, &amp;dx);
	restx(&amp;dx, &amp;olddx);
	restx(&amp;ax, &amp;oldax);
}

/*
 * generate shift according to op, one of:
 *	res = nl &lt;&lt; nr
 *	res = nl &gt;&gt; nr
 */
void
cgen_shift(int op, Node *nl, Node *nr, Node *res)
{
	Node n1, n2, cx, oldcx;
	int a, w;
	Prog *p1;
	uvlong sc;

	if(nl-&gt;type-&gt;width &gt; 4)
		fatal(&#34;cgen_shift %T&#34;, nl-&gt;type);

	w = nl-&gt;type-&gt;width * 8;

	a = optoas(op, nl-&gt;type);

	if(nr-&gt;op == OLITERAL) {
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		sc = mpgetfix(nr-&gt;val.u.xval);
		if(sc &gt;= nl-&gt;type-&gt;width*8) {
			// large shift gets 2 shifts by width
			gins(a, ncon(w-1), &amp;n1);
			gins(a, ncon(w-1), &amp;n1);
		} else
			gins(a, nr, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		return;
	}

	memset(&amp;oldcx, 0, sizeof oldcx);
	nodreg(&amp;cx, types[TUINT32], D_CX);
	if(reg[D_CX] &gt; 1 &amp;&amp; !samereg(&amp;cx, res)) {
		tempalloc(&amp;oldcx, types[TUINT32]);
		gmove(&amp;cx, &amp;oldcx);
	}

	nodreg(&amp;n1, types[TUINT32], D_CX);
	regalloc(&amp;n1, nr-&gt;type, &amp;n1);		// to hold the shift type in CX

	if(samereg(&amp;cx, res))
		regalloc(&amp;n2, nl-&gt;type, N);
	else
		regalloc(&amp;n2, nl-&gt;type, res);
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		cgen(nl, &amp;n2);
		cgen(nr, &amp;n1);
	} else {
		cgen(nr, &amp;n1);
		cgen(nl, &amp;n2);
	}

	// test and fix up large shifts
	gins(optoas(OCMP, nr-&gt;type), &amp;n1, ncon(w));
	p1 = gbranch(optoas(OLT, types[TUINT32]), T);
	if(op == ORSH &amp;&amp; issigned[nl-&gt;type-&gt;etype]) {
		gins(a, ncon(w-1), &amp;n2);
	} else {
		gmove(ncon(0), &amp;n2);
	}
	patch(p1, pc);
	gins(a, &amp;n1, &amp;n2);

	if(oldcx.op != 0) {
		gmove(&amp;oldcx, &amp;cx);
		tempfree(&amp;oldcx);
	}

	gmove(&amp;n2, res);

	regfree(&amp;n1);
	regfree(&amp;n2);
}

/*
 * generate byte multiply:
 *	res = nl * nr
 * no byte multiply instruction so have to do
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
