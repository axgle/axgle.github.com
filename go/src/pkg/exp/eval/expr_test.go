<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/expr_test.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/expr_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bignum&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>var undefined = &#34;undefined&#34;
<a id="L13"></a>var typeAsExpr = &#34;type .* used as expression&#34;
<a id="L14"></a>var badCharLit = &#34;character literal&#34;
<a id="L15"></a>var illegalEscape = &#34;illegal char escape&#34;
<a id="L16"></a>var opTypes = &#34;illegal (operand|argument) type|cannot index into&#34;
<a id="L17"></a>var badAddrOf = &#34;cannot take the address&#34;
<a id="L18"></a>var constantTruncated = &#34;constant [^ ]* truncated&#34;
<a id="L19"></a>var constantUnderflows = &#34;constant [^ ]* underflows&#34;
<a id="L20"></a>var constantOverflows = &#34;constant [^ ]* overflows&#34;
<a id="L21"></a>var implLimit = &#34;implementation limit&#34;
<a id="L22"></a>var mustBeUnsigned = &#34;must be unsigned&#34;
<a id="L23"></a>var divByZero = &#34;divide by zero&#34;

<a id="L25"></a>var hugeInteger = bignum.Int(1).Shl(64)

<a id="L27"></a>var exprTests = []test{
    <a id="L28"></a>Val(&#34;i&#34;, 1),
    <a id="L29"></a>CErr(&#34;zzz&#34;, undefined),
    <a id="L30"></a><span class="comment">// TODO(austin) Test variable in constant context</span>
    <a id="L31"></a><span class="comment">//CErr(&#34;t&#34;, typeAsExpr),</span>

    <a id="L33"></a>Val(&#34;&#39;a&#39;&#34;, bignum.Int(&#39;a&#39;)),
    <a id="L34"></a>Val(&#34;&#39;\\uffff&#39;&#34;, bignum.Int(&#39;\uffff&#39;)),
    <a id="L35"></a>Val(&#34;&#39;\\n&#39;&#34;, bignum.Int(&#39;\n&#39;)),
    <a id="L36"></a>CErr(&#34;&#39;&#39;+x&#34;, badCharLit),
    <a id="L37"></a><span class="comment">// Produces two parse errors</span>
    <a id="L38"></a><span class="comment">//CErr(&#34;&#39;&#39;&#39;&#34;, &#34;&#34;),</span>
    <a id="L39"></a>CErr(&#34;&#39;\n&#39;&#34;, badCharLit),
    <a id="L40"></a>CErr(&#34;&#39;\\z&#39;&#34;, illegalEscape),
    <a id="L41"></a>CErr(&#34;&#39;ab&#39;&#34;, badCharLit),

    <a id="L43"></a>Val(&#34;1.0&#34;, bignum.Rat(1, 1)),
    <a id="L44"></a>Val(&#34;1.&#34;, bignum.Rat(1, 1)),
    <a id="L45"></a>Val(&#34;.1&#34;, bignum.Rat(1, 10)),
    <a id="L46"></a>Val(&#34;1e2&#34;, bignum.Rat(100, 1)),

    <a id="L48"></a>Val(&#34;\&#34;abc\&#34;&#34;, &#34;abc&#34;),
    <a id="L49"></a>Val(&#34;\&#34;\&#34;&#34;, &#34;&#34;),
    <a id="L50"></a>Val(&#34;\&#34;\\n\\\&#34;\&#34;&#34;, &#34;\n\&#34;&#34;),
    <a id="L51"></a>CErr(&#34;\&#34;\\z\&#34;&#34;, illegalEscape),
    <a id="L52"></a>CErr(&#34;\&#34;abc&#34;, &#34;string not terminated&#34;),

    <a id="L54"></a>Val(&#34;\&#34;abc\&#34; \&#34;def\&#34;&#34;, &#34;abcdef&#34;),
    <a id="L55"></a>CErr(&#34;\&#34;abc\&#34; \&#34;\\z\&#34;&#34;, illegalEscape),

    <a id="L57"></a>Val(&#34;(i)&#34;, 1),

    <a id="L59"></a>Val(&#34;ai[0]&#34;, 1),
    <a id="L60"></a>Val(&#34;(&amp;ai)[0]&#34;, 1),
    <a id="L61"></a>Val(&#34;ai[1]&#34;, 2),
    <a id="L62"></a>Val(&#34;ai[i]&#34;, 2),
    <a id="L63"></a>Val(&#34;ai[u]&#34;, 2),
    <a id="L64"></a>CErr(&#34;ai[f]&#34;, opTypes),
    <a id="L65"></a>CErr(&#34;ai[0][0]&#34;, opTypes),
    <a id="L66"></a>CErr(&#34;ai[2]&#34;, &#34;index 2 exceeds&#34;),
    <a id="L67"></a>CErr(&#34;ai[1+1]&#34;, &#34;index 2 exceeds&#34;),
    <a id="L68"></a>CErr(&#34;ai[-1]&#34;, &#34;negative index&#34;),
    <a id="L69"></a>RErr(&#34;ai[i+i]&#34;, &#34;index 2 exceeds&#34;),
    <a id="L70"></a>RErr(&#34;ai[-i]&#34;, &#34;negative index&#34;),
    <a id="L71"></a>CErr(&#34;i[0]&#34;, opTypes),
    <a id="L72"></a>CErr(&#34;f[0]&#34;, opTypes),

    <a id="L74"></a>Val(&#34;aai[0][0]&#34;, 1),
    <a id="L75"></a>Val(&#34;aai[1][1]&#34;, 4),
    <a id="L76"></a>CErr(&#34;aai[2][0]&#34;, &#34;index 2 exceeds&#34;),
    <a id="L77"></a>CErr(&#34;aai[0][2]&#34;, &#34;index 2 exceeds&#34;),

    <a id="L79"></a>Val(&#34;sli[0]&#34;, 1),
    <a id="L80"></a>Val(&#34;sli[1]&#34;, 2),
    <a id="L81"></a>CErr(&#34;sli[-1]&#34;, &#34;negative index&#34;),
    <a id="L82"></a>RErr(&#34;sli[-i]&#34;, &#34;negative index&#34;),
    <a id="L83"></a>RErr(&#34;sli[2]&#34;, &#34;index 2 exceeds&#34;),

    <a id="L85"></a>Val(&#34;s[0]&#34;, uint8(&#39;a&#39;)),
    <a id="L86"></a>Val(&#34;s[1]&#34;, uint8(&#39;b&#39;)),
    <a id="L87"></a>CErr(&#34;s[-1]&#34;, &#34;negative index&#34;),
    <a id="L88"></a>RErr(&#34;s[-i]&#34;, &#34;negative index&#34;),
    <a id="L89"></a>RErr(&#34;s[3]&#34;, &#34;index 3 exceeds&#34;),

    <a id="L91"></a>CErr(&#34;1(2)&#34;, &#34;cannot call&#34;),
    <a id="L92"></a>CErr(&#34;fn(1,2)&#34;, &#34;too many&#34;),
    <a id="L93"></a>CErr(&#34;fn()&#34;, &#34;not enough&#34;),
    <a id="L94"></a>CErr(&#34;fn(true)&#34;, opTypes),
    <a id="L95"></a>CErr(&#34;fn(true)&#34;, &#34;function call&#34;),
    <a id="L96"></a><span class="comment">// Single argument functions don&#39;t say which argument.</span>
    <a id="L97"></a><span class="comment">//CErr(&#34;fn(true)&#34;, &#34;argument 1&#34;),</span>
    <a id="L98"></a>Val(&#34;fn(1)&#34;, 2),
    <a id="L99"></a>Val(&#34;fn(1.0)&#34;, 2),
    <a id="L100"></a>CErr(&#34;fn(1.5)&#34;, constantTruncated),
    <a id="L101"></a>Val(&#34;fn(i)&#34;, 2),
    <a id="L102"></a>CErr(&#34;fn(u)&#34;, opTypes),

    <a id="L104"></a>CErr(&#34;void()+2&#34;, opTypes),
    <a id="L105"></a>CErr(&#34;oneTwo()+2&#34;, opTypes),

    <a id="L107"></a>Val(&#34;cap(ai)&#34;, 2),
    <a id="L108"></a>Val(&#34;cap(&amp;ai)&#34;, 2),
    <a id="L109"></a>Val(&#34;cap(aai)&#34;, 2),
    <a id="L110"></a>Val(&#34;cap(sli)&#34;, 3),
    <a id="L111"></a>CErr(&#34;cap(0)&#34;, opTypes),
    <a id="L112"></a>CErr(&#34;cap(i)&#34;, opTypes),
    <a id="L113"></a>CErr(&#34;cap(s)&#34;, opTypes),

    <a id="L115"></a>Val(&#34;len(s)&#34;, 3),
    <a id="L116"></a>Val(&#34;len(ai)&#34;, 2),
    <a id="L117"></a>Val(&#34;len(&amp;ai)&#34;, 2),
    <a id="L118"></a>Val(&#34;len(aai)&#34;, 2),
    <a id="L119"></a>Val(&#34;len(sli)&#34;, 2),
    <a id="L120"></a><span class="comment">// TODO(austin) Test len of map</span>
    <a id="L121"></a>CErr(&#34;len(0)&#34;, opTypes),
    <a id="L122"></a>CErr(&#34;len(i)&#34;, opTypes),

    <a id="L124"></a>CErr(&#34;*i&#34;, opTypes),
    <a id="L125"></a>Val(&#34;*&amp;i&#34;, 1),
    <a id="L126"></a>Val(&#34;*&amp;(i)&#34;, 1),
    <a id="L127"></a>CErr(&#34;&amp;1&#34;, badAddrOf),
    <a id="L128"></a>CErr(&#34;&amp;c&#34;, badAddrOf),
    <a id="L129"></a>Val(&#34;*(&amp;ai[0])&#34;, 1),

    <a id="L131"></a>Val(&#34;+1&#34;, bignum.Int(+1)),
    <a id="L132"></a>Val(&#34;+1.0&#34;, bignum.Rat(1, 1)),
    <a id="L133"></a>CErr(&#34;+\&#34;x\&#34;&#34;, opTypes),

    <a id="L135"></a>Val(&#34;-42&#34;, bignum.Int(-42)),
    <a id="L136"></a>Val(&#34;-i&#34;, -1),
    <a id="L137"></a>Val(&#34;-f&#34;, -1.0),
    <a id="L138"></a><span class="comment">// 6g bug?</span>
    <a id="L139"></a><span class="comment">//Val(&#34;-(f-1)&#34;, -0.0),</span>
    <a id="L140"></a>CErr(&#34;-\&#34;x\&#34;&#34;, opTypes),

    <a id="L142"></a><span class="comment">// TODO(austin) Test unary !</span>

    <a id="L144"></a>Val(&#34;^2&#34;, bignum.Int(^2)),
    <a id="L145"></a>Val(&#34;^(-2)&#34;, bignum.Int(^(-2))),
    <a id="L146"></a>CErr(&#34;^2.0&#34;, opTypes),
    <a id="L147"></a>CErr(&#34;^2.5&#34;, opTypes),
    <a id="L148"></a>Val(&#34;^i&#34;, ^1),
    <a id="L149"></a>Val(&#34;^u&#34;, ^uint(1)),
    <a id="L150"></a>CErr(&#34;^f&#34;, opTypes),

    <a id="L152"></a>Val(&#34;1+i&#34;, 2),
    <a id="L153"></a>Val(&#34;1+u&#34;, uint(2)),
    <a id="L154"></a>Val(&#34;3.0+i&#34;, 4),
    <a id="L155"></a>Val(&#34;1+1&#34;, bignum.Int(2)),
    <a id="L156"></a>Val(&#34;f+f&#34;, 2.0),
    <a id="L157"></a>Val(&#34;1+f&#34;, 2.0),
    <a id="L158"></a>Val(&#34;1.0+1&#34;, bignum.Rat(2, 1)),
    <a id="L159"></a>Val(&#34;\&#34;abc\&#34; + \&#34;def\&#34;&#34;, &#34;abcdef&#34;),
    <a id="L160"></a>CErr(&#34;i+u&#34;, opTypes),
    <a id="L161"></a>CErr(&#34;-1+u&#34;, constantUnderflows),
    <a id="L162"></a><span class="comment">// TODO(austin) Test named types</span>

    <a id="L164"></a>Val(&#34;2-1&#34;, bignum.Int(1)),
    <a id="L165"></a>Val(&#34;2.0-1&#34;, bignum.Rat(1, 1)),
    <a id="L166"></a>Val(&#34;f-2&#34;, -1.0),
    <a id="L167"></a><span class="comment">// TOOD(austin) bignum can&#39;t do negative 0?</span>
    <a id="L168"></a><span class="comment">//Val(&#34;-0.0&#34;, XXX),</span>
    <a id="L169"></a>Val(&#34;2*2&#34;, bignum.Int(4)),
    <a id="L170"></a>Val(&#34;2*i&#34;, 2),
    <a id="L171"></a>Val(&#34;3/2&#34;, bignum.Int(1)),
    <a id="L172"></a>Val(&#34;3/i&#34;, 3),
    <a id="L173"></a>CErr(&#34;1/0&#34;, divByZero),
    <a id="L174"></a>CErr(&#34;1.0/0&#34;, divByZero),
    <a id="L175"></a>RErr(&#34;i/0&#34;, divByZero),
    <a id="L176"></a>Val(&#34;3%2&#34;, bignum.Int(1)),
    <a id="L177"></a>Val(&#34;i%2&#34;, 1),
    <a id="L178"></a>CErr(&#34;3%0&#34;, divByZero),
    <a id="L179"></a>CErr(&#34;3.0%0&#34;, opTypes),
    <a id="L180"></a>RErr(&#34;i%0&#34;, divByZero),

    <a id="L182"></a><span class="comment">// Examples from &#34;Arithmetic operators&#34;</span>
    <a id="L183"></a>Val(&#34;5/3&#34;, bignum.Int(1)),
    <a id="L184"></a>Val(&#34;(i+4)/(i+2)&#34;, 1),
    <a id="L185"></a>Val(&#34;5%3&#34;, bignum.Int(2)),
    <a id="L186"></a>Val(&#34;(i+4)%(i+2)&#34;, 2),
    <a id="L187"></a>Val(&#34;-5/3&#34;, bignum.Int(-1)),
    <a id="L188"></a>Val(&#34;(i-6)/(i+2)&#34;, -1),
    <a id="L189"></a>Val(&#34;-5%3&#34;, bignum.Int(-2)),
    <a id="L190"></a>Val(&#34;(i-6)%(i+2)&#34;, -2),
    <a id="L191"></a>Val(&#34;5/-3&#34;, bignum.Int(-1)),
    <a id="L192"></a>Val(&#34;(i+4)/(i-4)&#34;, -1),
    <a id="L193"></a>Val(&#34;5%-3&#34;, bignum.Int(2)),
    <a id="L194"></a>Val(&#34;(i+4)%(i-4)&#34;, 2),
    <a id="L195"></a>Val(&#34;-5/-3&#34;, bignum.Int(1)),
    <a id="L196"></a>Val(&#34;(i-6)/(i-4)&#34;, 1),
    <a id="L197"></a>Val(&#34;-5%-3&#34;, bignum.Int(-2)),
    <a id="L198"></a>Val(&#34;(i-6)%(i-4)&#34;, -2),

    <a id="L200"></a><span class="comment">// Examples from &#34;Arithmetic operators&#34;</span>
    <a id="L201"></a>Val(&#34;11/4&#34;, bignum.Int(2)),
    <a id="L202"></a>Val(&#34;(i+10)/4&#34;, 2),
    <a id="L203"></a>Val(&#34;11%4&#34;, bignum.Int(3)),
    <a id="L204"></a>Val(&#34;(i+10)%4&#34;, 3),
    <a id="L205"></a>Val(&#34;11&gt;&gt;2&#34;, bignum.Int(2)),
    <a id="L206"></a>Val(&#34;(i+10)&gt;&gt;2&#34;, 2),
    <a id="L207"></a>Val(&#34;11&amp;3&#34;, bignum.Int(3)),
    <a id="L208"></a>Val(&#34;(i+10)&amp;3&#34;, 3),
    <a id="L209"></a>Val(&#34;-11/4&#34;, bignum.Int(-2)),
    <a id="L210"></a>Val(&#34;(i-12)/4&#34;, -2),
    <a id="L211"></a>Val(&#34;-11%4&#34;, bignum.Int(-3)),
    <a id="L212"></a>Val(&#34;(i-12)%4&#34;, -3),
    <a id="L213"></a>Val(&#34;-11&gt;&gt;2&#34;, bignum.Int(-3)),
    <a id="L214"></a>Val(&#34;(i-12)&gt;&gt;2&#34;, -3),
    <a id="L215"></a>Val(&#34;-11&amp;3&#34;, bignum.Int(1)),
    <a id="L216"></a>Val(&#34;(i-12)&amp;3&#34;, 1),

    <a id="L218"></a><span class="comment">// TODO(austin) Test bit ops</span>

    <a id="L220"></a><span class="comment">// For shift, we try nearly every combination of positive</span>
    <a id="L221"></a><span class="comment">// ideal int, negative ideal int, big ideal int, ideal</span>
    <a id="L222"></a><span class="comment">// fractional float, ideal non-fractional float, int, uint,</span>
    <a id="L223"></a><span class="comment">// and float.</span>
    <a id="L224"></a>Val(&#34;2&lt;&lt;2&#34;, bignum.Int(2&lt;&lt;2)),
    <a id="L225"></a>CErr(&#34;2&lt;&lt;(-1)&#34;, constantUnderflows),
    <a id="L226"></a>CErr(&#34;2&lt;&lt;0x10000000000000000&#34;, constantOverflows),
    <a id="L227"></a>CErr(&#34;2&lt;&lt;2.5&#34;, constantTruncated),
    <a id="L228"></a>Val(&#34;2&lt;&lt;2.0&#34;, bignum.Int(2&lt;&lt;2.0)),
    <a id="L229"></a>CErr(&#34;2&lt;&lt;i&#34;, mustBeUnsigned),
    <a id="L230"></a>Val(&#34;2&lt;&lt;u&#34;, 2&lt;&lt;1),
    <a id="L231"></a>CErr(&#34;2&lt;&lt;f&#34;, opTypes),

    <a id="L233"></a>Val(&#34;-2&lt;&lt;2&#34;, bignum.Int(-2&lt;&lt;2)),
    <a id="L234"></a>CErr(&#34;-2&lt;&lt;(-1)&#34;, constantUnderflows),
    <a id="L235"></a>CErr(&#34;-2&lt;&lt;0x10000000000000000&#34;, constantOverflows),
    <a id="L236"></a>CErr(&#34;-2&lt;&lt;2.5&#34;, constantTruncated),
    <a id="L237"></a>Val(&#34;-2&lt;&lt;2.0&#34;, bignum.Int(-2&lt;&lt;2.0)),
    <a id="L238"></a>CErr(&#34;-2&lt;&lt;i&#34;, mustBeUnsigned),
    <a id="L239"></a>Val(&#34;-2&lt;&lt;u&#34;, -2&lt;&lt;1),
    <a id="L240"></a>CErr(&#34;-2&lt;&lt;f&#34;, opTypes),

    <a id="L242"></a>Val(&#34;0x10000000000000000&lt;&lt;2&#34;, hugeInteger.Shl(2)),
    <a id="L243"></a>CErr(&#34;0x10000000000000000&lt;&lt;(-1)&#34;, constantUnderflows),
    <a id="L244"></a>CErr(&#34;0x10000000000000000&lt;&lt;0x10000000000000000&#34;, constantOverflows),
    <a id="L245"></a>CErr(&#34;0x10000000000000000&lt;&lt;2.5&#34;, constantTruncated),
    <a id="L246"></a>Val(&#34;0x10000000000000000&lt;&lt;2.0&#34;, hugeInteger.Shl(2)),
    <a id="L247"></a>CErr(&#34;0x10000000000000000&lt;&lt;i&#34;, mustBeUnsigned),
    <a id="L248"></a>CErr(&#34;0x10000000000000000&lt;&lt;u&#34;, constantOverflows),
    <a id="L249"></a>CErr(&#34;0x10000000000000000&lt;&lt;f&#34;, opTypes),

    <a id="L251"></a>CErr(&#34;2.5&lt;&lt;2&#34;, opTypes),
    <a id="L252"></a>CErr(&#34;2.0&lt;&lt;2&#34;, opTypes),

    <a id="L254"></a>Val(&#34;i&lt;&lt;2&#34;, 1&lt;&lt;2),
    <a id="L255"></a>CErr(&#34;i&lt;&lt;(-1)&#34;, constantUnderflows),
    <a id="L256"></a>CErr(&#34;i&lt;&lt;0x10000000000000000&#34;, constantOverflows),
    <a id="L257"></a>CErr(&#34;i&lt;&lt;2.5&#34;, constantTruncated),
    <a id="L258"></a>Val(&#34;i&lt;&lt;2.0&#34;, 1&lt;&lt;2),
    <a id="L259"></a>CErr(&#34;i&lt;&lt;i&#34;, mustBeUnsigned),
    <a id="L260"></a>Val(&#34;i&lt;&lt;u&#34;, 1&lt;&lt;1),
    <a id="L261"></a>CErr(&#34;i&lt;&lt;f&#34;, opTypes),
    <a id="L262"></a>Val(&#34;i&lt;&lt;u&#34;, 1&lt;&lt;1),

    <a id="L264"></a>Val(&#34;u&lt;&lt;2&#34;, uint(1&lt;&lt;2)),
    <a id="L265"></a>CErr(&#34;u&lt;&lt;(-1)&#34;, constantUnderflows),
    <a id="L266"></a>CErr(&#34;u&lt;&lt;0x10000000000000000&#34;, constantOverflows),
    <a id="L267"></a>CErr(&#34;u&lt;&lt;2.5&#34;, constantTruncated),
    <a id="L268"></a>Val(&#34;u&lt;&lt;2.0&#34;, uint(1&lt;&lt;2)),
    <a id="L269"></a>CErr(&#34;u&lt;&lt;i&#34;, mustBeUnsigned),
    <a id="L270"></a>Val(&#34;u&lt;&lt;u&#34;, uint(1&lt;&lt;1)),
    <a id="L271"></a>CErr(&#34;u&lt;&lt;f&#34;, opTypes),
    <a id="L272"></a>Val(&#34;u&lt;&lt;u&#34;, uint(1&lt;&lt;1)),

    <a id="L274"></a>CErr(&#34;f&lt;&lt;2&#34;, opTypes),

    <a id="L276"></a><span class="comment">// &lt;, &lt;=, &gt;, &gt;=</span>
    <a id="L277"></a>Val(&#34;1&lt;2&#34;, 1 &lt; 2),
    <a id="L278"></a>Val(&#34;1&lt;=2&#34;, 1 &lt;= 2),
    <a id="L279"></a>Val(&#34;2&lt;=2&#34;, 2 &lt;= 2),
    <a id="L280"></a>Val(&#34;1&gt;2&#34;, 1 &gt; 2),
    <a id="L281"></a>Val(&#34;1&gt;=2&#34;, 1 &gt;= 2),
    <a id="L282"></a>Val(&#34;2&gt;=2&#34;, 2 &gt;= 2),

    <a id="L284"></a>Val(&#34;i&lt;2&#34;, 1 &lt; 2),
    <a id="L285"></a>Val(&#34;i&lt;=2&#34;, 1 &lt;= 2),
    <a id="L286"></a>Val(&#34;i+1&lt;=2&#34;, 2 &lt;= 2),
    <a id="L287"></a>Val(&#34;i&gt;2&#34;, 1 &gt; 2),
    <a id="L288"></a>Val(&#34;i&gt;=2&#34;, 1 &gt;= 2),
    <a id="L289"></a>Val(&#34;i+1&gt;=2&#34;, 2 &gt;= 2),

    <a id="L291"></a>Val(&#34;u&lt;2&#34;, 1 &lt; 2),
    <a id="L292"></a>Val(&#34;f&lt;2&#34;, 1 &lt; 2),

    <a id="L294"></a>Val(&#34;s&lt;\&#34;b\&#34;&#34;, true),
    <a id="L295"></a>Val(&#34;s&lt;\&#34;a\&#34;&#34;, false),
    <a id="L296"></a>Val(&#34;s&lt;=\&#34;abc\&#34;&#34;, true),
    <a id="L297"></a>Val(&#34;s&gt;\&#34;aa\&#34;&#34;, true),
    <a id="L298"></a>Val(&#34;s&gt;\&#34;ac\&#34;&#34;, false),
    <a id="L299"></a>Val(&#34;s&gt;=\&#34;abc\&#34;&#34;, true),

    <a id="L301"></a>CErr(&#34;i&lt;u&#34;, opTypes),
    <a id="L302"></a>CErr(&#34;i&lt;f&#34;, opTypes),
    <a id="L303"></a>CErr(&#34;i&lt;s&#34;, opTypes),
    <a id="L304"></a>CErr(&#34;&amp;i&lt;&amp;i&#34;, opTypes),
    <a id="L305"></a>CErr(&#34;ai&lt;ai&#34;, opTypes),

    <a id="L307"></a><span class="comment">// ==, !=</span>
    <a id="L308"></a>Val(&#34;1==1&#34;, true),
    <a id="L309"></a>Val(&#34;1!=1&#34;, false),
    <a id="L310"></a>Val(&#34;1==2&#34;, false),
    <a id="L311"></a>Val(&#34;1!=2&#34;, true),

    <a id="L313"></a>Val(&#34;1.0==1&#34;, true),
    <a id="L314"></a>Val(&#34;1.5==1&#34;, false),

    <a id="L316"></a>Val(&#34;i==1&#34;, true),
    <a id="L317"></a>Val(&#34;i!=1&#34;, false),
    <a id="L318"></a>Val(&#34;i==2&#34;, false),
    <a id="L319"></a>Val(&#34;i!=2&#34;, true),

    <a id="L321"></a>Val(&#34;u==1&#34;, true),
    <a id="L322"></a>Val(&#34;f==1&#34;, true),

    <a id="L324"></a>Val(&#34;s==\&#34;abc\&#34;&#34;, true),
    <a id="L325"></a>Val(&#34;s!=\&#34;abc\&#34;&#34;, false),
    <a id="L326"></a>Val(&#34;s==\&#34;abcd\&#34;&#34;, false),
    <a id="L327"></a>Val(&#34;s!=\&#34;abcd\&#34;&#34;, true),

    <a id="L329"></a>Val(&#34;&amp;i==&amp;i&#34;, true),
    <a id="L330"></a>Val(&#34;&amp;i==&amp;i2&#34;, false),

    <a id="L332"></a>Val(&#34;fn==fn&#34;, true),
    <a id="L333"></a>Val(&#34;fn==func(int)int{return 0}&#34;, false),

    <a id="L335"></a>CErr(&#34;i==u&#34;, opTypes),
    <a id="L336"></a>CErr(&#34;i==f&#34;, opTypes),
    <a id="L337"></a>CErr(&#34;&amp;i==&amp;f&#34;, opTypes),
    <a id="L338"></a>CErr(&#34;ai==ai&#34;, opTypes),
    <a id="L339"></a>CErr(&#34;t==t&#34;, opTypes),
    <a id="L340"></a>CErr(&#34;fn==oneTwo&#34;, opTypes),
<a id="L341"></a>}

<a id="L343"></a>func TestExpr(t *testing.T) { runTests(t, &#34;exprTests&#34;, exprTests) }
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
