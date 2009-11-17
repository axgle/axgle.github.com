<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/4s/data.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/4s/data.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// games/4s - a tetris clone</span>
<a id="L2"></a><span class="comment">//</span>
<a id="L3"></a><span class="comment">// Derived from Plan 9&#39;s /sys/src/games/xs.c</span>
<a id="L4"></a><span class="comment">// http://plan9.bell-labs.com/sources/plan9/sys/src/games/xs.c</span>
<a id="L5"></a><span class="comment">//</span>
<a id="L6"></a><span class="comment">// Copyright (C) 2003, Lucent Technologies Inc. and others. All Rights Reserved.</span>
<a id="L7"></a><span class="comment">// Portions Copyright 2009 The Go Authors.  All Rights Reserved.</span>
<a id="L8"></a><span class="comment">// Distributed under the terms of the Lucent Public License Version 1.02</span>
<a id="L9"></a><span class="comment">// See http://plan9.bell-labs.com/plan9/license.html</span>

<a id="L11"></a>package main

<a id="L13"></a>import . &#34;exp/draw&#34;

<a id="L15"></a>var pieces4 = []Piece{
    <a id="L16"></a>Piece{0, 0, Point{4, 1}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L17"></a>Piece{1, 0, Point{1, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}}, nil, nil},
    <a id="L18"></a>Piece{2, 0, Point{4, 1}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L19"></a>Piece{3, 0, Point{1, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}}, nil, nil},

    <a id="L21"></a>Piece{0, 3, Point{2, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}, nil, nil},
    <a id="L22"></a>Piece{1, 3, Point{2, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}, nil, nil},
    <a id="L23"></a>Piece{2, 3, Point{2, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}, nil, nil},
    <a id="L24"></a>Piece{3, 3, Point{2, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}, nil, nil},

    <a id="L26"></a>Piece{0, 1, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L27"></a>Piece{1, 1, Point{2, 3}, []Point{Point{1, 0}, Point{0, 1}, Point{0, 1}, Point{-1, 0}}, nil, nil},
    <a id="L28"></a>Piece{2, 1, Point{3, 2}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L29"></a>Piece{3, 1, Point{2, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{-1, 1}, Point{0, 1}}, nil, nil},

    <a id="L31"></a>Piece{0, 2, Point{3, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{0, -1}}, nil, nil},
    <a id="L32"></a>Piece{1, 2, Point{2, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L33"></a>Piece{2, 2, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{-2, 1}}, nil, nil},
    <a id="L34"></a>Piece{3, 2, Point{2, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{0, 1}}, nil, nil},

    <a id="L36"></a>Piece{0, 4, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L37"></a>Piece{1, 4, Point{2, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L38"></a>Piece{2, 4, Point{3, 2}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L39"></a>Piece{3, 4, Point{2, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{1, -1}}, nil, nil},

    <a id="L41"></a>Piece{0, 5, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L42"></a>Piece{1, 5, Point{2, 3}, []Point{Point{1, 0}, Point{0, 1}, Point{-1, 0}, Point{0, 1}}, nil, nil},
    <a id="L43"></a>Piece{2, 5, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L44"></a>Piece{3, 5, Point{2, 3}, []Point{Point{1, 0}, Point{0, 1}, Point{-1, 0}, Point{0, 1}}, nil, nil},

    <a id="L46"></a>Piece{0, 6, Point{3, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{1, 0}}, nil, nil},
    <a id="L47"></a>Piece{1, 6, Point{2, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L48"></a>Piece{2, 6, Point{3, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{1, 0}}, nil, nil},
    <a id="L49"></a>Piece{3, 6, Point{2, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
<a id="L50"></a>}

<a id="L52"></a>var pieces5 = []Piece{
    <a id="L53"></a>Piece{0, 1, Point{5, 1}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L54"></a>Piece{1, 1, Point{1, 5}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}, Point{0, 1}}, nil, nil},
    <a id="L55"></a>Piece{2, 1, Point{5, 1}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L56"></a>Piece{3, 1, Point{1, 5}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}, Point{0, 1}}, nil, nil},

    <a id="L58"></a>Piece{0, 0, Point{4, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L59"></a>Piece{1, 0, Point{2, 4}, []Point{Point{1, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}, Point{-1, 0}}, nil, nil},
    <a id="L60"></a>Piece{2, 0, Point{4, 2}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L61"></a>Piece{3, 0, Point{2, 4}, []Point{Point{0, 0}, Point{1, 0}, Point{-1, 1}, Point{0, 1}, Point{0, 1}}, nil, nil},

    <a id="L63"></a>Piece{0, 2, Point{4, 2}, []Point{Point{0, 0}, Point{0, 1}, Point{1, -1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L64"></a>Piece{1, 2, Point{2, 4}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}}, nil, nil},
    <a id="L65"></a>Piece{2, 2, Point{4, 2}, []Point{Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{0, -1}}, nil, nil},
    <a id="L66"></a>Piece{3, 2, Point{2, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{0, 1}, Point{1, 0}}, nil, nil},

    <a id="L68"></a>Piece{0, 7, Point{3, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{0, 1}, Point{0, 1}}, nil, nil},
    <a id="L69"></a>Piece{1, 7, Point{3, 3}, []Point{Point{0, 2}, Point{1, 0}, Point{1, 0}, Point{0, -1}, Point{0, -1}}, nil, nil},
    <a id="L70"></a>Piece{2, 7, Point{3, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L71"></a>Piece{3, 7, Point{3, 3}, []Point{Point{0, 2}, Point{0, -1}, Point{0, -1}, Point{1, 0}, Point{1, 0}}, nil, nil},

    <a id="L73"></a>Piece{0, 3, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{-2, 1}, Point{1, 0}}, nil, nil},
    <a id="L74"></a>Piece{1, 3, Point{2, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L75"></a>Piece{2, 3, Point{3, 2}, []Point{Point{1, 0}, Point{1, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L76"></a>Piece{3, 3, Point{2, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{-1, 1}, Point{1, 0}}, nil, nil},

    <a id="L78"></a>Piece{0, 4, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{-1, 1}, Point{1, 0}}, nil, nil},
    <a id="L79"></a>Piece{1, 4, Point{2, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{-1, 1}, Point{1, 0}}, nil, nil},
    <a id="L80"></a>Piece{2, 4, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L81"></a>Piece{3, 4, Point{2, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{-1, 1}}, nil, nil},

    <a id="L83"></a>Piece{0, 7, Point{3, 2}, []Point{Point{0, 0}, Point{2, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L84"></a>Piece{1, 7, Point{2, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{-1, 1}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L85"></a>Piece{2, 7, Point{3, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{-2, 1}, Point{2, 0}}, nil, nil},
    <a id="L86"></a>Piece{3, 7, Point{2, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{-1, 1}, Point{1, 0}}, nil, nil},

    <a id="L88"></a>Piece{0, 5, Point{3, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L89"></a>Piece{1, 5, Point{3, 3}, []Point{Point{2, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L90"></a>Piece{2, 5, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L91"></a>Piece{3, 5, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{-2, 1}}, nil, nil},

    <a id="L93"></a>Piece{0, 6, Point{3, 3}, []Point{Point{1, 0}, Point{1, 0}, Point{-2, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L94"></a>Piece{1, 6, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L95"></a>Piece{2, 6, Point{3, 3}, []Point{Point{1, 0}, Point{0, 1}, Point{1, 0}, Point{-2, 1}, Point{1, 0}}, nil, nil},
    <a id="L96"></a>Piece{3, 6, Point{3, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},

    <a id="L98"></a>Piece{0, 0, Point{4, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{-2, 1}}, nil, nil},
    <a id="L99"></a>Piece{1, 0, Point{2, 4}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{0, 1}, Point{0, 1}}, nil, nil},
    <a id="L100"></a>Piece{2, 0, Point{4, 2}, []Point{Point{2, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L101"></a>Piece{3, 0, Point{2, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{1, 0}, Point{-1, 1}}, nil, nil},

    <a id="L103"></a>Piece{0, 2, Point{4, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L104"></a>Piece{1, 2, Point{2, 4}, []Point{Point{1, 0}, Point{0, 1}, Point{-1, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L105"></a>Piece{2, 2, Point{4, 2}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L106"></a>Piece{3, 2, Point{2, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{-1, 1}, Point{0, 1}}, nil, nil},

    <a id="L108"></a>Piece{0, 1, Point{3, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L109"></a>Piece{1, 1, Point{3, 3}, []Point{Point{2, 0}, Point{-1, 1}, Point{1, 0}, Point{-2, 1}, Point{1, 0}}, nil, nil},
    <a id="L110"></a>Piece{2, 1, Point{3, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L111"></a>Piece{3, 1, Point{3, 3}, []Point{Point{1, 0}, Point{1, 0}, Point{-2, 1}, Point{1, 0}, Point{-1, 1}}, nil, nil},

    <a id="L113"></a>Piece{0, 3, Point{3, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{-1, 1}, Point{0, 1}}, nil, nil},
    <a id="L114"></a>Piece{1, 3, Point{3, 3}, []Point{Point{2, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L115"></a>Piece{2, 3, Point{3, 3}, []Point{Point{1, 0}, Point{0, 1}, Point{-1, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L116"></a>Piece{3, 3, Point{3, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{-2, 1}}, nil, nil},

    <a id="L118"></a>Piece{0, 4, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L119"></a>Piece{1, 4, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L120"></a>Piece{2, 4, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},
    <a id="L121"></a>Piece{3, 4, Point{3, 3}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{1, 0}, Point{-1, 1}}, nil, nil},

    <a id="L123"></a>Piece{0, 8, Point{4, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L124"></a>Piece{1, 8, Point{2, 4}, []Point{Point{1, 0}, Point{-1, 1}, Point{1, 0}, Point{-1, 1}, Point{0, 1}}, nil, nil},
    <a id="L125"></a>Piece{2, 8, Point{4, 2}, []Point{Point{0, 0}, Point{1, 0}, Point{1, 0}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L126"></a>Piece{3, 8, Point{2, 4}, []Point{Point{1, 0}, Point{0, 1}, Point{-1, 1}, Point{1, 0}, Point{-1, 1}}, nil, nil},

    <a id="L128"></a>Piece{0, 9, Point{4, 2}, []Point{Point{2, 0}, Point{1, 0}, Point{-3, 1}, Point{1, 0}, Point{1, 0}}, nil, nil},
    <a id="L129"></a>Piece{1, 9, Point{2, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{0, 1}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L130"></a>Piece{2, 9, Point{4, 2}, []Point{Point{1, 0}, Point{1, 0}, Point{1, 0}, Point{-3, 1}, Point{1, 0}}, nil, nil},
    <a id="L131"></a>Piece{3, 9, Point{2, 4}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{0, 1}, Point{0, 1}}, nil, nil},

    <a id="L133"></a>Piece{0, 5, Point{3, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L134"></a>Piece{1, 5, Point{3, 3}, []Point{Point{1, 0}, Point{1, 0}, Point{-1, 1}, Point{-1, 1}, Point{1, 0}}, nil, nil},
    <a id="L135"></a>Piece{2, 5, Point{3, 3}, []Point{Point{0, 0}, Point{0, 1}, Point{1, 0}, Point{1, 0}, Point{0, 1}}, nil, nil},
    <a id="L136"></a>Piece{3, 5, Point{3, 3}, []Point{Point{1, 0}, Point{1, 0}, Point{-1, 1}, Point{-1, 1}, Point{1, 0}}, nil, nil},

    <a id="L138"></a>Piece{0, 6, Point{3, 3}, []Point{Point{2, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}, Point{-2, 1}}, nil, nil},
    <a id="L139"></a>Piece{1, 6, Point{3, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{0, 1}, Point{1, 0}}, nil, nil},
    <a id="L140"></a>Piece{2, 6, Point{3, 3}, []Point{Point{2, 0}, Point{-2, 1}, Point{1, 0}, Point{1, 0}, Point{-2, 1}}, nil, nil},
    <a id="L141"></a>Piece{3, 6, Point{3, 3}, []Point{Point{0, 0}, Point{1, 0}, Point{0, 1}, Point{0, 1}, Point{1, 0}}, nil, nil},
<a id="L142"></a>}
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
