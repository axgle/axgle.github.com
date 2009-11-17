<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/gc/select.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/gc/select.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * select
 */

#include &#34;go.h&#34;

void
typecheckselect(Node *sel)
{
	Node *ncase, *n, *def;
	NodeList *l;
	int lno, count;

	def = nil;
	lno = setlineno(sel);
	count = 0;
	typechecklist(sel-&gt;ninit, Etop);
	for(l=sel-&gt;list; l; l=l-&gt;next) {
		count++;
		ncase = l-&gt;n;
		setlineno(ncase);
		if(ncase-&gt;op != OXCASE)
			fatal(&#34;typecheckselect %O&#34;, ncase-&gt;op);

		if(ncase-&gt;list == nil) {
			// default
			if(def != N)
				yyerror(&#34;multiple defaults in select (first at %L)&#34;, def-&gt;lineno);
			else
				def = ncase;
		} else if(ncase-&gt;list-&gt;next) {
			yyerror(&#34;select cases cannot be lists&#34;);
		} else {
			n = typecheck(&amp;ncase-&gt;list-&gt;n, Etop);
			ncase-&gt;left = n;
			ncase-&gt;list = nil;
			setlineno(n);
			switch(n-&gt;op) {
			default:
				yyerror(&#34;select case must be receive, send or assign recv&#34;);;
				break;

			case OAS:
				// convert x = &lt;-c into OSELRECV(x, c)
				if(n-&gt;right-&gt;op != ORECV) {
					yyerror(&#34;select assignment must have receive on right hand side&#34;);
					break;
				}
				n-&gt;op = OSELRECV;
				n-&gt;right = n-&gt;right-&gt;left;
				break;

			case ORECV:
				// convert &lt;-c into OSELRECV(N, c)
				n-&gt;op = OSELRECV;
				n-&gt;right = n-&gt;left;
				n-&gt;left = N;
				break;

			case OSEND:
				break;
			}
		}
		typechecklist(ncase-&gt;nbody, Etop);
	}
	sel-&gt;xoffset = count;
	if(count == 0)
		yyerror(&#34;empty select&#34;);
	lineno = lno;
}

void
walkselect(Node *sel)
{
	int lno;
	Node *n, *ncase, *r, *a, *tmp, *var;
	NodeList *l, *init;

	lno = setlineno(sel);
	init = sel-&gt;ninit;
	sel-&gt;ninit = nil;

	// generate sel-struct
	var = nod(OXXX, N, N);
	tempname(var, ptrto(types[TUINT8]));
	r = nod(OAS, var, mkcall(&#34;newselect&#34;, var-&gt;type, nil, nodintconst(sel-&gt;xoffset)));
	typecheck(&amp;r, Etop);
	init = list(init, r);

	if(sel-&gt;list == nil)
		fatal(&#34;double walkselect&#34;);	// already rewrote

	// register cases
	for(l=sel-&gt;list; l; l=l-&gt;next) {
		ncase = l-&gt;n;
		n = ncase-&gt;left;
		r = nod(OIF, N, N);
		r-&gt;nbody = ncase-&gt;ninit;
		ncase-&gt;ninit = nil;
		if(n != nil) {
			r-&gt;nbody = concat(r-&gt;nbody, n-&gt;ninit);
			n-&gt;ninit = nil;
		}
		if(n == nil) {
			// selectdefault(sel *byte);
			r-&gt;ntest = mkcall(&#34;selectdefault&#34;, types[TBOOL], &amp;init, var);
		} else if(n-&gt;op == OSEND) {
			// selectsend(sel *byte, hchan *chan any, elem any) (selected bool);
			r-&gt;ntest = mkcall1(chanfn(&#34;selectsend&#34;, 2, n-&gt;left-&gt;type), types[TBOOL], &amp;init, var, n-&gt;left, n-&gt;right);
		} else if(n-&gt;op == OSELRECV) {
			tmp = N;
			if(n-&gt;left == N)
				a = nodnil();
			else {
				// introduce temporary until we&#39;re sure this will succeed.
				tmp = nod(OXXX, N, N);
				tempname(tmp, n-&gt;left-&gt;type);
				a = nod(OADDR, tmp, N);
			}
			// selectrecv(sel *byte, hchan *chan any, elem *any) (selected bool);
			r-&gt;ntest = mkcall1(chanfn(&#34;selectrecv&#34;, 2, n-&gt;right-&gt;type), types[TBOOL], &amp;init, var, n-&gt;right, a);
			if(tmp != N) {
				a = nod(OAS, n-&gt;left, tmp);
				typecheck(&amp;a, Etop);
				r-&gt;nbody = list(r-&gt;nbody, a);
			}
		} else
			fatal(&#34;select %O&#34;, n-&gt;op);
		r-&gt;nbody = concat(r-&gt;nbody, ncase-&gt;nbody);
		r-&gt;nbody = list(r-&gt;nbody, nod(OBREAK, N, N));
		init = list(init, r);
	}

	// run the select
	init = list(init, mkcall(&#34;selectgo&#34;, T, nil, var));
	sel-&gt;nbody = init;
	sel-&gt;list = nil;
	walkstmtlist(init);

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
