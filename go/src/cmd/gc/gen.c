<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/gen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/gen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * portable half of code generator.
 * mainly statements and control flow.
 */

#include &#34;go.h&#34;

Node*
sysfunc(char *name)
{
	Node *n;

	n = newname(pkglookup(name, &#34;runtime&#34;));
	n-&gt;class = PFUNC;
	return n;
}

void
allocparams(void)
{
	NodeList *l;
	Node *n;
	uint32 w;
	Sym *s;
	int lno;

	if(stksize &lt; 0)
		fatal(&#34;allocparams not during code generation&#34;);

	/*
	 * allocate (set xoffset) the stack
	 * slots for all automatics.
	 * allocated starting at -w down.
	 */
	lno = lineno;
	for(l=curfn-&gt;dcl; l; l=l-&gt;next) {
		n = l-&gt;n;
		if(n-&gt;op == ONAME &amp;&amp; n-&gt;class == PHEAP-1) {
			// heap address variable; finish the job
			// started in addrescapes.
			s = n-&gt;sym;
			tempname(n, n-&gt;type);
			n-&gt;sym = s;
		}
		if(n-&gt;op != ONAME || n-&gt;class != PAUTO)
			continue;
		if(n-&gt;type == T)
			continue;
		dowidth(n-&gt;type);
		w = n-&gt;type-&gt;width;
		if(w &gt;= 100000000)
			fatal(&#34;bad width&#34;);
		stksize += w;
		stksize = rnd(stksize, w);
		n-&gt;xoffset = -stksize;
	}
	lineno = lno;
}

void
newlab(int op, Sym *s, Node *stmt)
{
	Label *lab;

	lab = mal(sizeof(*lab));
	lab-&gt;link = labellist;
	labellist = lab;

	lab-&gt;sym = s;
	lab-&gt;op = op;
	lab-&gt;label = pc;
	lab-&gt;stmt = stmt;
}

void
checklabels(void)
{
	Label *l, *m;
	Sym *s;

//	// print the label list
//	for(l=labellist; l!=L; l=l-&gt;link) {
//		print(&#34;lab %O %S\n&#34;, l-&gt;op, l-&gt;sym);
//	}

	for(l=labellist; l!=L; l=l-&gt;link) {
	switch(l-&gt;op) {
		case OLABEL:
			// these are definitions -
			s = l-&gt;sym;
			for(m=labellist; m!=L; m=m-&gt;link) {
				if(m-&gt;sym != s)
					continue;
				switch(m-&gt;op) {
				case OLABEL:
					// these are definitions -
					// look for redefinitions
					if(l != m)
						yyerror(&#34;label %S redefined&#34;, s);
					break;
				case OGOTO:
					// these are references -
					// patch to definition
					patch(m-&gt;label, l-&gt;label);
					m-&gt;sym = S;	// mark done
					break;
				}
			}
		}
	}

	// diagnostic for all undefined references
	for(l=labellist; l!=L; l=l-&gt;link)
		if(l-&gt;op == OGOTO &amp;&amp; l-&gt;sym != S)
			yyerror(&#34;label %S not defined&#34;, l-&gt;sym);
}

/*
 * compile statements
 */
void
genlist(NodeList *l)
{
	for(; l; l=l-&gt;next)
		gen(l-&gt;n);
}

void
gen(Node *n)
{
	int32 lno;
	Prog *scontin, *sbreak;
	Prog *p1, *p2, *p3;
	Label *lab;

	lno = setlineno(n);

	if(n == N)
		goto ret;

	p3 = pc;	// save pc for loop labels
	if(n-&gt;ninit)
		genlist(n-&gt;ninit);

	setlineno(n);

	switch(n-&gt;op) {
	default:
		fatal(&#34;gen: unknown op %N&#34;, n);
		break;

	case OCASE:
	case OFALL:
	case OXCASE:
	case OXFALL:
	case ODCLCONST:
	case ODCLFUNC:
	case ODCLTYPE:
		break;

	case OEMPTY:
		// insert no-op so that
		//	L:; for { }
		// does not treat L as a label for the loop.
		if(labellist &amp;&amp; labellist-&gt;label == p3)
			gused(N);
		break;

	case OBLOCK:
		genlist(n-&gt;list);
		break;

	case OLABEL:
		newlab(OLABEL, n-&gt;left-&gt;sym, n-&gt;right);
		break;

	case OGOTO:
		newlab(OGOTO, n-&gt;left-&gt;sym, N);
		gjmp(P);
		break;

	case OBREAK:
		if(n-&gt;left != N) {
			for(lab=labellist; lab!=L; lab=lab-&gt;link) {
				if(lab-&gt;sym == n-&gt;left-&gt;sym) {
					if(lab-&gt;breakpc == P)
						yyerror(&#34;invalid break label %S&#34;, n-&gt;left-&gt;sym);
					gjmp(lab-&gt;breakpc);
					goto donebreak;
				}
			}
			if(lab == L)
				yyerror(&#34;break label not defined: %S&#34;, n-&gt;left-&gt;sym);
			break;
		}
		if(breakpc == P) {
			yyerror(&#34;break is not in a loop&#34;);
			break;
		}
		gjmp(breakpc);
	donebreak:
		break;

	case OCONTINUE:
		if(n-&gt;left != N) {
			for(lab=labellist; lab!=L; lab=lab-&gt;link) {
				if(lab-&gt;sym == n-&gt;left-&gt;sym) {
					if(lab-&gt;continpc == P)
						yyerror(&#34;invalid continue label %S&#34;, n-&gt;left-&gt;sym);
					gjmp(lab-&gt;continpc);
					goto donecont;
				}
			}
			if(lab == L)
				yyerror(&#34;continue label not defined: %S&#34;, n-&gt;left-&gt;sym);
			break;
		}

		if(continpc == P) {
			yyerror(&#34;continue is not in a loop&#34;);
			break;
		}
		gjmp(continpc);
	donecont:
		break;

	case OFOR:
		sbreak = breakpc;
		p1 = gjmp(P);			//		goto test
		breakpc = gjmp(P);		// break:	goto done
		scontin = continpc;
		continpc = pc;

		// define break and continue labels
		if((lab = labellist) != L &amp;&amp; lab-&gt;label == p3 &amp;&amp; lab-&gt;op == OLABEL &amp;&amp; lab-&gt;stmt == n) {
			lab-&gt;breakpc = breakpc;
			lab-&gt;continpc = continpc;
		}

		gen(n-&gt;nincr);				// contin:	incr
		patch(p1, pc);				// test:
		if(n-&gt;ntest != N)
			if(n-&gt;ntest-&gt;ninit != nil)
				genlist(n-&gt;ntest-&gt;ninit);
		bgen(n-&gt;ntest, 0, breakpc);		//		if(!test) goto break
		genlist(n-&gt;nbody);				//		body
		gjmp(continpc);
		patch(breakpc, pc);			// done:
		continpc = scontin;
		breakpc = sbreak;
		break;

	case OIF:
		p1 = gjmp(P);			//		goto test
		p2 = gjmp(P);			// p2:		goto else
		patch(p1, pc);				// test:
		if(n-&gt;ntest != N)
			if(n-&gt;ntest-&gt;ninit != nil)
				genlist(n-&gt;ntest-&gt;ninit);
		bgen(n-&gt;ntest, 0, p2);			//		if(!test) goto p2
		genlist(n-&gt;nbody);				//		then
		p3 = gjmp(P);			//		goto done
		patch(p2, pc);				// else:
		genlist(n-&gt;nelse);				//		else
		patch(p3, pc);				// done:
		break;

	case OSWITCH:
		sbreak = breakpc;
		p1 = gjmp(P);			//		goto test
		breakpc = gjmp(P);		// break:	goto done

		// define break label
		if((lab = labellist) != L &amp;&amp; lab-&gt;label == p3 &amp;&amp; lab-&gt;op == OLABEL &amp;&amp; lab-&gt;stmt == n)
			lab-&gt;breakpc = breakpc;

		patch(p1, pc);				// test:
		genlist(n-&gt;nbody);				//		switch(test) body
		patch(breakpc, pc);			// done:
		breakpc = sbreak;
		break;

	case OSELECT:
		sbreak = breakpc;
		p1 = gjmp(P);			//		goto test
		breakpc = gjmp(P);		// break:	goto done

		// define break label
		if((lab = labellist) != L &amp;&amp; lab-&gt;label == p3 &amp;&amp; lab-&gt;op == OLABEL &amp;&amp; lab-&gt;stmt == n)
			lab-&gt;breakpc = breakpc;

		patch(p1, pc);				// test:
		genlist(n-&gt;nbody);				//		select() body
		patch(breakpc, pc);			// done:
		breakpc = sbreak;
		break;

	case OASOP:
		cgen_asop(n);
		break;

	case ODCL:
		cgen_dcl(n-&gt;left);
		break;

	case OAS:
		if(gen_as_init(n))
			break;
		cgen_as(n-&gt;left, n-&gt;right);
		break;

	case OCALLMETH:
		cgen_callmeth(n, 0);
		break;

	case OCALLINTER:
		cgen_callinter(n, N, 0);
		break;

	case OCALLFUNC:
		cgen_call(n, 0);
		break;

	case OPROC:
		cgen_proc(n, 1);
		break;

	case ODEFER:
		cgen_proc(n, 2);
		break;

	case ORETURN:
		cgen_ret(n);
		break;
	}

ret:
	lineno = lno;
}

/*
 * generate call to non-interface method
 *	proc=0	normal call
 *	proc=1	goroutine run in new proc
 *	proc=2	defer call save away stack
 */
void
cgen_callmeth(Node *n, int proc)
{
	Node *l;

	// generate a rewrite for method call
	// (p.f)(...) goes to (f)(p,...)

	l = n-&gt;left;
	if(l-&gt;op != ODOTMETH)
		fatal(&#34;cgen_callmeth: not dotmethod: %N&#34;);

	n-&gt;op = OCALLFUNC;
	n-&gt;left = n-&gt;left-&gt;right;
	n-&gt;left-&gt;type = l-&gt;type;

	if(n-&gt;left-&gt;op == ONAME)
		n-&gt;left-&gt;class = PFUNC;
	cgen_call(n, proc);
}

/*
 * generate code to start new proc running call n.
 */
void
cgen_proc(Node *n, int proc)
{
	switch(n-&gt;left-&gt;op) {
	default:
		fatal(&#34;cgen_proc: unknown call %O&#34;, n-&gt;left-&gt;op);

	case OCALLMETH:
		cgen_callmeth(n-&gt;left, proc);
		break;

	case OCALLINTER:
		cgen_callinter(n-&gt;left, N, proc);
		break;

	case OCALLFUNC:
		cgen_call(n-&gt;left, proc);
		break;
	}

}

/*
 * generate declaration.
 * nothing to do for on-stack automatics,
 * but might have to allocate heap copy
 * for escaped variables.
 */
void
cgen_dcl(Node *n)
{
	if(debug[&#39;g&#39;])
		dump(&#34;\ncgen-dcl&#34;, n);
	if(n-&gt;op != ONAME) {
		dump(&#34;cgen_dcl&#34;, n);
		fatal(&#34;cgen_dcl&#34;);
	}
	if(!(n-&gt;class &amp; PHEAP))
		return;
	cgen_as(n-&gt;heapaddr, n-&gt;alloc);
}

/*
 * generate discard of value
 */
void
cgen_discard(Node *nr)
{
	Node tmp;

	if(nr == N)
		return;

	switch(nr-&gt;op) {
	case ONAME:
		if(!(nr-&gt;class &amp; PHEAP) &amp;&amp; nr-&gt;class != PEXTERN &amp;&amp; nr-&gt;class != PFUNC)
			gused(nr);
		break;

	// unary
	case OADD:
	case OAND:
	case ODIV:
	case OEQ:
	case OGE:
	case OGT:
	case OLE:
	case OLSH:
	case OLT:
	case OMOD:
	case OMUL:
	case ONE:
	case OOR:
	case ORSH:
	case OSUB:
	case OXOR:
		cgen_discard(nr-&gt;left);
		cgen_discard(nr-&gt;right);
		break;

	// binary
	case OCAP:
	case OCOM:
	case OLEN:
	case OMINUS:
	case ONOT:
	case OPLUS:
		cgen_discard(nr-&gt;left);
		break;

	// special enough to just evaluate
	default:
		tempname(&amp;tmp, nr-&gt;type);
		cgen_as(&amp;tmp, nr);
		gused(&amp;tmp);
	}
}

/*
 * generate assignment:
 *	nl = nr
 * nr == N means zero nl.
 */
void
cgen_as(Node *nl, Node *nr)
{
	Node nc;
	Type *tl;
	int iszer;

	if(nl == N)
		return;

	if(debug[&#39;g&#39;]) {
		dump(&#34;cgen_as&#34;, nl);
		dump(&#34;cgen_as = &#34;, nr);
	}

	if(isblank(nl)) {
		cgen_discard(nr);
		return;
	}

	iszer = 0;
	if(nr == N || isnil(nr)) {
		// externals and heaps should already be clear
		if(nr == N) {
			if(nl-&gt;class == PEXTERN)
				return;
			if(nl-&gt;class &amp; PHEAP)
				return;
		}

		tl = nl-&gt;type;
		if(tl == T)
			return;
		if(isfat(tl)) {
			clearfat(nl);
			goto ret;
		}

		/* invent a &#34;zero&#34; for the rhs */
		iszer = 1;
		nr = &amp;nc;
		memset(nr, 0, sizeof(*nr));
		switch(simtype[tl-&gt;etype]) {
		default:
			fatal(&#34;cgen_as: tl %T&#34;, tl);
			break;

		case TINT8:
		case TUINT8:
		case TINT16:
		case TUINT16:
		case TINT32:
		case TUINT32:
		case TINT64:
		case TUINT64:
			nr-&gt;val.u.xval = mal(sizeof(*nr-&gt;val.u.xval));
			mpmovecfix(nr-&gt;val.u.xval, 0);
			nr-&gt;val.ctype = CTINT;
			break;

		case TFLOAT32:
		case TFLOAT64:
			nr-&gt;val.u.fval = mal(sizeof(*nr-&gt;val.u.fval));
			mpmovecflt(nr-&gt;val.u.fval, 0.0);
			nr-&gt;val.ctype = CTFLT;
			break;

		case TBOOL:
			nr-&gt;val.u.bval = 0;
			nr-&gt;val.ctype = CTBOOL;
			break;

		case TPTR32:
		case TPTR64:
			nr-&gt;val.ctype = CTNIL;
			break;

		}
		nr-&gt;op = OLITERAL;
		nr-&gt;type = tl;
		nr-&gt;addable = 1;
		ullmancalc(nr);
	}

	tl = nl-&gt;type;
	if(tl == T)
		return;

	cgen(nr, nl);
	if(iszer &amp;&amp; nl-&gt;addable)
		gused(nl);

ret:
	;
}

/*
 * gather series of offsets
 * &gt;=0 is direct addressed field
 * &lt;0 is pointer to next field (+1)
 */
int
dotoffset(Node *n, int *oary, Node **nn)
{
	int i;

	switch(n-&gt;op) {
	case ODOT:
		if(n-&gt;xoffset == BADWIDTH) {
			dump(&#34;bad width in dotoffset&#34;, n);
			fatal(&#34;bad width in dotoffset&#34;);
		}
		i = dotoffset(n-&gt;left, oary, nn);
		if(i &gt; 0) {
			if(oary[i-1] &gt;= 0)
				oary[i-1] += n-&gt;xoffset;
			else
				oary[i-1] -= n-&gt;xoffset;
			break;
		}
		if(i &lt; 10)
			oary[i++] = n-&gt;xoffset;
		break;

	case ODOTPTR:
		if(n-&gt;xoffset == BADWIDTH) {
			dump(&#34;bad width in dotoffset&#34;, n);
			fatal(&#34;bad width in dotoffset&#34;);
		}
		i = dotoffset(n-&gt;left, oary, nn);
		if(i &lt; 10)
			oary[i++] = -(n-&gt;xoffset+1);
		break;

	default:
		*nn = n;
		return 0;
	}
	if(i &gt;= 10)
		*nn = N;
	return i;
}

/*
 * make a new off the books
 */
void
tempname(Node *n, Type *t)
{
	Sym *s;
	uint32 w;

	if(stksize &lt; 0)
		fatal(&#34;tempname not during code generation&#34;);

	if(t == T) {
		yyerror(&#34;tempname called with nil type&#34;);
		t = types[TINT32];
	}

	// give each tmp a different name so that there
	// a chance to registerizer them
	snprint(namebuf, sizeof(namebuf), &#34;autotmp_%.4d&#34;, statuniqgen);
	statuniqgen++;
	s = lookup(namebuf);

	memset(n, 0, sizeof(*n));
	n-&gt;op = ONAME;
	n-&gt;sym = s;
	n-&gt;type = t;
	n-&gt;class = PAUTO;
	n-&gt;addable = 1;
	n-&gt;ullman = 1;
	n-&gt;noescape = 1;

	dowidth(t);
	w = t-&gt;width;
	stksize += w;
	stksize = rnd(stksize, w);
	n-&gt;xoffset = -stksize;
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
