<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/range.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/range.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * range
 */

#include &#34;go.h&#34;

void
typecheckrange(Node *n)
{
	int op, et;
	Type *t, *t1, *t2;
	Node *v1, *v2;
	NodeList *ll;

	// delicate little dance.  see typecheckas2
	for(ll=n-&gt;list; ll; ll=ll-&gt;next)
		if(ll-&gt;n-&gt;defn != n)
			typecheck(&amp;ll-&gt;n, Erv | Easgn);

	typecheck(&amp;n-&gt;right, Erv);
	if((t = n-&gt;right-&gt;type) == T)
		goto out;
	n-&gt;type = t;

	switch(t-&gt;etype) {
	default:
		yyerror(&#34;cannot range over %+N&#34;, n-&gt;right);
		goto out;

	case TARRAY:
		t1 = types[TINT];
		t2 = t-&gt;type;
		break;

	case TMAP:
		t1 = t-&gt;down;
		t2 = t-&gt;type;
		break;

	case TCHAN:
		t1 = t-&gt;type;
		t2 = nil;
		if(count(n-&gt;list) == 2)
			goto toomany;
		break;

	case TSTRING:
		t1 = types[TINT];
		t2 = types[TINT];
		break;
	}

	if(count(n-&gt;list) &gt; 2) {
	toomany:
		yyerror(&#34;too many variables in range&#34;);
	}

	v1 = n-&gt;list-&gt;n;
	v2 = N;
	if(n-&gt;list-&gt;next)
		v2 = n-&gt;list-&gt;next-&gt;n;

	if(v1-&gt;defn == n)
		v1-&gt;type = t1;
	else if(v1-&gt;type != T &amp;&amp; checkconv(t1, v1-&gt;type, 0, &amp;op, &amp;et) &lt; 0)
		yyerror(&#34;cannot assign type %T to %+N&#34;, t1, v1);
	if(v2) {
		if(v2-&gt;defn == n)
			v2-&gt;type = t2;
		else if(v2-&gt;type != T &amp;&amp; checkconv(t2, v2-&gt;type, 0, &amp;op, &amp;et) &lt; 0)
			yyerror(&#34;cannot assign type %T to %+N&#34;, t1, v1);
	}

out:
	typechecklist(n-&gt;nbody, Etop);

	// second half of dance
	n-&gt;typecheck = 1;
	for(ll=n-&gt;list; ll; ll=ll-&gt;next)
		if(ll-&gt;n-&gt;typecheck == 0)
			typecheck(&amp;ll-&gt;n, Erv | Easgn);
}

void
walkrange(Node *n)
{
	Node *ohv1, *hv1, *hv2;	// hidden (old) val 1, 2
	Node *ha, *hit;	// hidden aggregate, iterator
	Node *a, *v1, *v2;	// not hidden aggregate, val 1, 2
	Node *fn;
	NodeList *body, *init;
	Type *th, *t;

	t = n-&gt;type;
	init = nil;

	a = n-&gt;right;
	if(t-&gt;etype == TSTRING &amp;&amp; !eqtype(t, types[TSTRING])) {
		a = nod(OCONV, n-&gt;right, N);
		a-&gt;type = types[TSTRING];
	}
	ha = nod(OXXX, N, N);
	tempname(ha, a-&gt;type);
	init = list(init, nod(OAS, ha, a));

	v1 = n-&gt;list-&gt;n;
	hv1 = N;

	v2 = N;
	if(n-&gt;list-&gt;next)
		v2 = n-&gt;list-&gt;next-&gt;n;
	hv2 = N;

	switch(t-&gt;etype) {
	default:
		fatal(&#34;walkrange&#34;);

	case TARRAY:
		hv1 = nod(OXXX, N, n);
		tempname(hv1, types[TINT]);

		init = list(init, nod(OAS, hv1, N));
		n-&gt;ntest = nod(OLT, hv1, nod(OLEN, ha, N));
		n-&gt;nincr = nod(OASOP, hv1, nodintconst(1));
		n-&gt;nincr-&gt;etype = OADD;
		body = list1(nod(OAS, v1, hv1));
		if(v2)
			body = list(body, nod(OAS, v2, nod(OINDEX, ha, hv1)));
		break;

	case TMAP:
		th = typ(TARRAY);
		th-&gt;type = ptrto(types[TUINT8]);
		th-&gt;bound = (sizeof(struct Hiter) + widthptr - 1) / widthptr;
		hit = nod(OXXX, N, N);
		tempname(hit, th);

		fn = syslook(&#34;mapiterinit&#34;, 1);
		argtype(fn, t-&gt;down);
		argtype(fn, t-&gt;type);
		argtype(fn, th);
		init = list(init, mkcall1(fn, T, nil, ha, nod(OADDR, hit, N)));
		n-&gt;ntest = nod(ONE, nod(OINDEX, hit, nodintconst(0)), nodnil());

		fn = syslook(&#34;mapiternext&#34;, 1);
		argtype(fn, th);
		n-&gt;nincr = mkcall1(fn, T, nil, nod(OADDR, hit, N));

		if(v2 == N) {
			fn = syslook(&#34;mapiter1&#34;, 1);
			argtype(fn, th);
			argtype(fn, t-&gt;down);
			a = nod(OAS, v1, mkcall1(fn, t-&gt;down, nil, nod(OADDR, hit, N)));
		} else {
			fn = syslook(&#34;mapiter2&#34;, 1);
			argtype(fn, th);
			argtype(fn, t-&gt;down);
			argtype(fn, t-&gt;type);
			a = nod(OAS2, N, N);
			a-&gt;list = list(list1(v1), v2);
			a-&gt;rlist = list1(mkcall1(fn, getoutargx(fn-&gt;type), nil, nod(OADDR, hit, N)));
		}
		body = list1(a);
		break;

	case TCHAN:
		hv1 = nod(OXXX, N, n);
		tempname(hv1, t-&gt;type);

		n-&gt;ntest = nod(ONOT, nod(OCLOSED, ha, N), N);
		n-&gt;ntest-&gt;ninit = list1(nod(OAS, hv1, nod(ORECV, ha, N)));
		body = list1(nod(OAS, v1, hv1));
		break;

	case TSTRING:
		ohv1 = nod(OXXX, N, N);
		tempname(ohv1, types[TINT]);

		hv1 = nod(OXXX, N, N);
		tempname(hv1, types[TINT]);
		init = list(init, nod(OAS, hv1, N));

		if(v2 == N)
			a = nod(OAS, hv1, mkcall(&#34;stringiter&#34;, types[TINT], nil, ha, hv1));
		else {
			hv2 = nod(OXXX, N, N);
			tempname(hv2, types[TINT]);
			a = nod(OAS2, N, N);
			a-&gt;list = list(list1(hv1), hv2);
			fn = syslook(&#34;stringiter2&#34;, 0);
			a-&gt;rlist = list1(mkcall1(fn, getoutargx(fn-&gt;type), nil, ha, hv1));
		}
		n-&gt;ntest = nod(ONE, hv1, nodintconst(0));
		n-&gt;ntest-&gt;ninit = list(list1(nod(OAS, ohv1, hv1)), a);

		body = list1(nod(OAS, v1, ohv1));
		if(v2 != N)
			body = list(body, nod(OAS, v2, hv2));
		break;
	}

	n-&gt;op = OFOR;
	typechecklist(init, Etop);
	n-&gt;ninit = concat(n-&gt;ninit, init);
	typechecklist(n-&gt;ntest-&gt;ninit, Etop);
	typecheck(&amp;n-&gt;ntest, Erv);
	typecheck(&amp;n-&gt;nincr, Etop);
	typechecklist(body, Etop);
	n-&gt;nbody = concat(body, n-&gt;nbody);
	walkstmt(&amp;n);
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
