<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/pgen.c</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/pgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/6c/sgen.c
// http://code.google.com/p/inferno-os/source/browse/utils/6c/sgen.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

#include &#34;gc.h&#34;

vlong
argsize(void)
{
	Type *t;
	int32 s;

//print(&#34;t=%T\n&#34;, thisfn);
	s = align(0, thisfn-&gt;link, Aarg0);
	for(t=thisfn-&gt;down; t!=T; t=t-&gt;down) {
		switch(t-&gt;etype) {
		case TVOID:
			break;
		case TDOT:
			s += 64;
			break;
		default:
			s = align(s, t, Aarg1);
			s = align(s, t, Aarg2);
			break;
		}
//print(&#34;	%d %T\n&#34;, s, t);
	}
	if(thechar == &#39;6&#39;)
		s = (s+7) &amp; ~7;
	else
		s = (s+3) &amp; ~3;
	return s;
}

void
codgen(Node *n, Node *nn)
{
	Prog *sp;
	Node *n1, nod, nod1;

	cursafe = 0;
	curarg = 0;
	maxargsafe = 0;

	/*
	 * isolate name
	 */
	for(n1 = nn;; n1 = n1-&gt;left) {
		if(n1 == Z) {
			diag(nn, &#34;cant find function name&#34;);
			return;
		}
		if(n1-&gt;op == ONAME)
			break;
	}
	nearln = nn-&gt;lineno;

	p = gtext(n1-&gt;sym, stkoff);
	sp = p;

	/*
	 * isolate first argument
	 */
	if(REGARG &gt;= 0) {
		if(typesuv[thisfn-&gt;link-&gt;etype]) {
			nod1 = *nodret-&gt;left;
			nodreg(&amp;nod, &amp;nod1, REGARG);
			gmove(&amp;nod, &amp;nod1);
		} else
		if(firstarg &amp;&amp; typechlp[firstargtype-&gt;etype]) {
			nod1 = *nodret-&gt;left;
			nod1.sym = firstarg;
			nod1.type = firstargtype;
			nod1.xoffset = align(0, firstargtype, Aarg1);
			nod1.etype = firstargtype-&gt;etype;
			nodreg(&amp;nod, &amp;nod1, REGARG);
			gmove(&amp;nod, &amp;nod1);
		}
	}

	retok = 0;

	canreach = 1;
	warnreach = 1;
	gen(n);
	if(canreach &amp;&amp; thisfn-&gt;link-&gt;etype != TVOID)
		warn(Z, &#34;no return at end of function: %s&#34;, n1-&gt;sym-&gt;name);
	noretval(3);
	gbranch(ORETURN);

	if(!debug[&#39;N&#39;] || debug[&#39;R&#39;] || debug[&#39;P&#39;])
		regopt(sp);

	if(thechar==&#39;6&#39; || thechar==&#39;7&#39;)	/* [sic] */
		maxargsafe = xround(maxargsafe, 8);
	sp-&gt;to.offset += maxargsafe;
}

void
supgen(Node *n)
{
	int owarn;
	long spc;
	Prog *sp;

	if(n == Z)
		return;
	suppress++;
	owarn = warnreach;
	warnreach = 0;
	spc = pc;
	sp = lastp;
	gen(n);
	lastp = sp;
	pc = spc;
	sp-&gt;link = nil;
	suppress--;
	warnreach = owarn;
}

void
gen(Node *n)
{
	Node *l, nod;
	Prog *sp, *spc, *spb;
	Case *cn;
	long sbc, scc;
	int snbreak, sncontin;
	int f, o, oldreach;

loop:
	if(n == Z)
		return;
	nearln = n-&gt;lineno;
	o = n-&gt;op;
	if(debug[&#39;G&#39;])
		if(o != OLIST)
			print(&#34;%L %O\n&#34;, nearln, o);

	if(!canreach) {
		switch(o) {
		case OLABEL:
		case OCASE:
		case OLIST:
		case OBREAK:
		case OFOR:
		case OWHILE:
		case ODWHILE:
			/* all handled specially - see switch body below */
			break;
		default:
			if(warnreach) {
				warn(n, &#34;unreachable code %O&#34;, o);
				warnreach = 0;
			}
		}
	}

	switch(o) {

	default:
		complex(n);
		cgen(n, Z);
		break;

	case OLIST:
		gen(n-&gt;left);

	rloop:
		n = n-&gt;right;
		goto loop;

	case ORETURN:
		canreach = 0;
		warnreach = !suppress;
		complex(n);
		if(n-&gt;type == T)
			break;
		l = n-&gt;left;
		if(l == Z) {
			noretval(3);
			gbranch(ORETURN);
			break;
		}
		if(typecmplx[n-&gt;type-&gt;etype]) {
			sugen(l, nodret, n-&gt;type-&gt;width);
			noretval(3);
			gbranch(ORETURN);
			break;
		}
		regret(&amp;nod, n);
		cgen(l, &amp;nod);
		regfree(&amp;nod);
		if(typefd[n-&gt;type-&gt;etype])
			noretval(1);
		else
			noretval(2);
		gbranch(ORETURN);
		break;

	case OLABEL:
		canreach = 1;
		l = n-&gt;left;
		if(l) {
			l-&gt;pc = pc;
			if(l-&gt;label)
				patch(l-&gt;label, pc);
		}
		gbranch(OGOTO);	/* prevent self reference in reg */
		patch(p, pc);
		goto rloop;

	case OGOTO:
		canreach = 0;
		warnreach = !suppress;
		n = n-&gt;left;
		if(n == Z)
			return;
		if(n-&gt;complex == 0) {
			diag(Z, &#34;label undefined: %s&#34;, n-&gt;sym-&gt;name);
			return;
		}
		if(suppress)
			return;
		gbranch(OGOTO);
		if(n-&gt;pc) {
			patch(p, n-&gt;pc);
			return;
		}
		if(n-&gt;label)
			patch(n-&gt;label, pc-1);
		n-&gt;label = p;
		return;

	case OCASE:
		canreach = 1;
		l = n-&gt;left;
		if(cases == C)
			diag(n, &#34;case/default outside a switch&#34;);
		if(l == Z) {
			cas();
			cases-&gt;val = 0;
			cases-&gt;def = 1;
			cases-&gt;label = pc;
			cases-&gt;isv = 0;
			goto rloop;
		}
		complex(l);
		if(l-&gt;type == T)
			goto rloop;
		if(l-&gt;op == OCONST)
		if(typeword[l-&gt;type-&gt;etype] &amp;&amp; l-&gt;type-&gt;etype != TIND) {
			cas();
			cases-&gt;val = l-&gt;vconst;
			cases-&gt;def = 0;
			cases-&gt;label = pc;
			cases-&gt;isv = typev[l-&gt;type-&gt;etype];
			goto rloop;
		}
		diag(n, &#34;case expression must be integer constant&#34;);
		goto rloop;

	case OSWITCH:
		l = n-&gt;left;
		complex(l);
		if(l-&gt;type == T)
			break;
		if(!typeword[l-&gt;type-&gt;etype] || l-&gt;type-&gt;etype == TIND) {
			diag(n, &#34;switch expression must be integer&#34;);
			break;
		}

		gbranch(OGOTO);		/* entry */
		sp = p;

		cn = cases;
		cases = C;
		cas();

		sbc = breakpc;
		breakpc = pc;
		snbreak = nbreak;
		nbreak = 0;
		gbranch(OGOTO);
		spb = p;

		gen(n-&gt;right);		/* body */
		if(canreach){
			gbranch(OGOTO);
			patch(p, breakpc);
			nbreak++;
		}

		patch(sp, pc);
		regalloc(&amp;nod, l, Z);
		/* always signed */
		if(typev[l-&gt;type-&gt;etype])
			nod.type = types[TVLONG];
		else
			nod.type = types[TLONG];
		cgen(l, &amp;nod);
		doswit(&amp;nod);
		regfree(&amp;nod);
		patch(spb, pc);

		cases = cn;
		breakpc = sbc;
		canreach = nbreak!=0;
		if(canreach == 0)
			warnreach = !suppress;
		nbreak = snbreak;
		break;

	case OWHILE:
	case ODWHILE:
		l = n-&gt;left;
		gbranch(OGOTO);		/* entry */
		sp = p;

		scc = continpc;
		continpc = pc;
		gbranch(OGOTO);
		spc = p;

		sbc = breakpc;
		breakpc = pc;
		snbreak = nbreak;
		nbreak = 0;
		gbranch(OGOTO);
		spb = p;

		patch(spc, pc);
		if(n-&gt;op == OWHILE)
			patch(sp, pc);
		bcomplex(l, Z);		/* test */
		patch(p, breakpc);
		if(l-&gt;op != OCONST || vconst(l) == 0)
			nbreak++;

		if(n-&gt;op == ODWHILE)
			patch(sp, pc);
		gen(n-&gt;right);		/* body */
		gbranch(OGOTO);
		patch(p, continpc);

		patch(spb, pc);
		continpc = scc;
		breakpc = sbc;
		canreach = nbreak!=0;
		if(canreach == 0)
			warnreach = !suppress;
		nbreak = snbreak;
		break;

	case OFOR:
		l = n-&gt;left;
		if(!canreach &amp;&amp; l-&gt;right-&gt;left &amp;&amp; warnreach) {
			warn(n, &#34;unreachable code FOR&#34;);
			warnreach = 0;
		}
		gen(l-&gt;right-&gt;left);	/* init */
		gbranch(OGOTO);		/* entry */
		sp = p;

		/*
		 * if there are no incoming labels in the
		 * body and the top&#39;s not reachable, warn
		 */
		if(!canreach &amp;&amp; warnreach &amp;&amp; deadheads(n)) {
			warn(n, &#34;unreachable code %O&#34;, o);
			warnreach = 0;
		}

		scc = continpc;
		continpc = pc;
		gbranch(OGOTO);
		spc = p;

		sbc = breakpc;
		breakpc = pc;
		snbreak = nbreak;
		nbreak = 0;
		sncontin = ncontin;
		ncontin = 0;
		gbranch(OGOTO);
		spb = p;

		patch(spc, pc);
		gen(l-&gt;right-&gt;right);	/* inc */
		patch(sp, pc);
		if(l-&gt;left != Z) {	/* test */
			bcomplex(l-&gt;left, Z);
			patch(p, breakpc);
			if(l-&gt;left-&gt;op != OCONST || vconst(l-&gt;left) == 0)
				nbreak++;
		}
		canreach = 1;
		gen(n-&gt;right);		/* body */
		if(canreach){
			gbranch(OGOTO);
			patch(p, continpc);
			ncontin++;
		}
		if(!ncontin &amp;&amp; l-&gt;right-&gt;right &amp;&amp; warnreach) {
			warn(l-&gt;right-&gt;right, &#34;unreachable FOR inc&#34;);
			warnreach = 0;
		}

		patch(spb, pc);
		continpc = scc;
		breakpc = sbc;
		canreach = nbreak!=0;
		if(canreach == 0)
			warnreach = !suppress;
		nbreak = snbreak;
		ncontin = sncontin;
		break;

	case OCONTINUE:
		if(continpc &lt; 0) {
			diag(n, &#34;continue not in a loop&#34;);
			break;
		}
		gbranch(OGOTO);
		patch(p, continpc);
		ncontin++;
		canreach = 0;
		warnreach = !suppress;
		break;

	case OBREAK:
		if(breakpc &lt; 0) {
			diag(n, &#34;break not in a loop&#34;);
			break;
		}
		/*
		 * Don&#39;t complain about unreachable break statements.
		 * There are breaks hidden in yacc&#39;s output and some people
		 * write return; break; in their switch statements out of habit.
		 * However, don&#39;t confuse the analysis by inserting an
		 * unreachable reference to breakpc either.
		 */
		if(!canreach)
			break;
		gbranch(OGOTO);
		patch(p, breakpc);
		nbreak++;
		canreach = 0;
		warnreach = !suppress;
		break;

	case OIF:
		l = n-&gt;left;
		if(bcomplex(l, n-&gt;right)) {
			if(typefd[l-&gt;type-&gt;etype])
				f = !l-&gt;fconst;
			else
				f = !l-&gt;vconst;
			if(debug[&#39;c&#39;])
				print(&#34;%L const if %s\n&#34;, nearln, f ? &#34;false&#34; : &#34;true&#34;);
			if(f) {
				canreach = 1;
				supgen(n-&gt;right-&gt;left);
				oldreach = canreach;
				canreach = 1;
				gen(n-&gt;right-&gt;right);
				/*
				 * treat constant ifs as regular ifs for
				 * reachability warnings.
				 */
				if(!canreach &amp;&amp; oldreach &amp;&amp; debug[&#39;w&#39;] &lt; 2)
					warnreach = 0;
			}
			else {
				canreach = 1;
				gen(n-&gt;right-&gt;left);
				oldreach = canreach;
				canreach = 1;
				supgen(n-&gt;right-&gt;right);
				/*
				 * treat constant ifs as regular ifs for
				 * reachability warnings.
				 */
				if(!oldreach &amp;&amp; canreach &amp;&amp; debug[&#39;w&#39;] &lt; 2)
					warnreach = 0;
				canreach = oldreach;
			}
		}
		else {
			sp = p;
			canreach = 1;
			if(n-&gt;right-&gt;left != Z)
				gen(n-&gt;right-&gt;left);
			oldreach = canreach;
			canreach = 1;
			if(n-&gt;right-&gt;right != Z) {
				gbranch(OGOTO);
				patch(sp, pc);
				sp = p;
				gen(n-&gt;right-&gt;right);
			}
			patch(sp, pc);
			canreach = canreach || oldreach;
			if(canreach == 0)
				warnreach = !suppress;
		}
		break;

	case OSET:
	case OUSED:
		usedset(n-&gt;left, o);
		break;
	}
}

void
usedset(Node *n, int o)
{
	if(n-&gt;op == OLIST) {
		usedset(n-&gt;left, o);
		usedset(n-&gt;right, o);
		return;
	}
	complex(n);
	switch(n-&gt;op) {
	case OADDR:	/* volatile */
		gins(ANOP, n, Z);
		break;
	case ONAME:
		if(o == OSET)
			gins(ANOP, Z, n);
		else
			gins(ANOP, n, Z);
		break;
	}
}

int
bcomplex(Node *n, Node *c)
{
	Node *b, nod;

	complex(n);
	if(n-&gt;type != T)
	if(tcompat(n, T, n-&gt;type, tnot))
		n-&gt;type = T;
	if(n-&gt;type == T) {
		gbranch(OGOTO);
		return 0;
	}
	if(c != Z &amp;&amp; n-&gt;op == OCONST &amp;&amp; deadheads(c))
		return 1;
	if(typev[n-&gt;type-&gt;etype] &amp;&amp; machcap(Z)) {
		b = &amp;nod;
		b-&gt;op = ONE;
		b-&gt;left = n;
		b-&gt;right = new(0, Z, Z);
		*b-&gt;right = *nodconst(0);
		b-&gt;right-&gt;type = n-&gt;type;
		b-&gt;type = types[TLONG];
		cgen(b, Z);
		return 0;
	}
	bool64(n);
	boolgen(n, 1, Z);
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
