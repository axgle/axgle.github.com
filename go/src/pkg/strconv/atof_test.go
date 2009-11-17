<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/atof_test.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/atof_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strconv_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;reflect&#34;;
    <a id="L10"></a>. &#34;strconv&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type atofTest struct {
    <a id="L15"></a>in  string;
    <a id="L16"></a>out string;
    <a id="L17"></a>err os.Error;
<a id="L18"></a>}

<a id="L20"></a>var atoftests = []atofTest{
    <a id="L21"></a>atofTest{&#34;&#34;, &#34;0&#34;, os.EINVAL},
    <a id="L22"></a>atofTest{&#34;1&#34;, &#34;1&#34;, nil},
    <a id="L23"></a>atofTest{&#34;+1&#34;, &#34;1&#34;, nil},
    <a id="L24"></a>atofTest{&#34;1x&#34;, &#34;0&#34;, os.EINVAL},
    <a id="L25"></a>atofTest{&#34;1.1.&#34;, &#34;0&#34;, os.EINVAL},
    <a id="L26"></a>atofTest{&#34;1e23&#34;, &#34;1e+23&#34;, nil},
    <a id="L27"></a>atofTest{&#34;100000000000000000000000&#34;, &#34;1e+23&#34;, nil},
    <a id="L28"></a>atofTest{&#34;1e-100&#34;, &#34;1e-100&#34;, nil},
    <a id="L29"></a>atofTest{&#34;123456700&#34;, &#34;1.234567e+08&#34;, nil},
    <a id="L30"></a>atofTest{&#34;99999999999999974834176&#34;, &#34;9.999999999999997e+22&#34;, nil},
    <a id="L31"></a>atofTest{&#34;100000000000000000000001&#34;, &#34;1.0000000000000001e+23&#34;, nil},
    <a id="L32"></a>atofTest{&#34;100000000000000008388608&#34;, &#34;1.0000000000000001e+23&#34;, nil},
    <a id="L33"></a>atofTest{&#34;100000000000000016777215&#34;, &#34;1.0000000000000001e+23&#34;, nil},
    <a id="L34"></a>atofTest{&#34;100000000000000016777216&#34;, &#34;1.0000000000000003e+23&#34;, nil},
    <a id="L35"></a>atofTest{&#34;-1&#34;, &#34;-1&#34;, nil},
    <a id="L36"></a>atofTest{&#34;-0&#34;, &#34;-0&#34;, nil},
    <a id="L37"></a>atofTest{&#34;1e-20&#34;, &#34;1e-20&#34;, nil},
    <a id="L38"></a>atofTest{&#34;625e-3&#34;, &#34;0.625&#34;, nil},

    <a id="L40"></a><span class="comment">// largest float64</span>
    <a id="L41"></a>atofTest{&#34;1.7976931348623157e308&#34;, &#34;1.7976931348623157e+308&#34;, nil},
    <a id="L42"></a>atofTest{&#34;-1.7976931348623157e308&#34;, &#34;-1.7976931348623157e+308&#34;, nil},
    <a id="L43"></a><span class="comment">// next float64 - too large</span>
    <a id="L44"></a>atofTest{&#34;1.7976931348623159e308&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L45"></a>atofTest{&#34;-1.7976931348623159e308&#34;, &#34;-Inf&#34;, os.ERANGE},
    <a id="L46"></a><span class="comment">// the border is ...158079</span>
    <a id="L47"></a><span class="comment">// borderline - okay</span>
    <a id="L48"></a>atofTest{&#34;1.7976931348623158e308&#34;, &#34;1.7976931348623157e+308&#34;, nil},
    <a id="L49"></a>atofTest{&#34;-1.7976931348623158e308&#34;, &#34;-1.7976931348623157e+308&#34;, nil},
    <a id="L50"></a><span class="comment">// borderline - too large</span>
    <a id="L51"></a>atofTest{&#34;1.797693134862315808e308&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L52"></a>atofTest{&#34;-1.797693134862315808e308&#34;, &#34;-Inf&#34;, os.ERANGE},

    <a id="L54"></a><span class="comment">// a little too large</span>
    <a id="L55"></a>atofTest{&#34;1e308&#34;, &#34;1e+308&#34;, nil},
    <a id="L56"></a>atofTest{&#34;2e308&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L57"></a>atofTest{&#34;1e309&#34;, &#34;+Inf&#34;, os.ERANGE},

    <a id="L59"></a><span class="comment">// way too large</span>
    <a id="L60"></a>atofTest{&#34;1e310&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L61"></a>atofTest{&#34;-1e310&#34;, &#34;-Inf&#34;, os.ERANGE},
    <a id="L62"></a>atofTest{&#34;1e400&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L63"></a>atofTest{&#34;-1e400&#34;, &#34;-Inf&#34;, os.ERANGE},
    <a id="L64"></a>atofTest{&#34;1e400000&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L65"></a>atofTest{&#34;-1e400000&#34;, &#34;-Inf&#34;, os.ERANGE},

    <a id="L67"></a><span class="comment">// denormalized</span>
    <a id="L68"></a>atofTest{&#34;1e-305&#34;, &#34;1e-305&#34;, nil},
    <a id="L69"></a>atofTest{&#34;1e-306&#34;, &#34;1e-306&#34;, nil},
    <a id="L70"></a>atofTest{&#34;1e-307&#34;, &#34;1e-307&#34;, nil},
    <a id="L71"></a>atofTest{&#34;1e-308&#34;, &#34;1e-308&#34;, nil},
    <a id="L72"></a>atofTest{&#34;1e-309&#34;, &#34;1e-309&#34;, nil},
    <a id="L73"></a>atofTest{&#34;1e-310&#34;, &#34;1e-310&#34;, nil},
    <a id="L74"></a>atofTest{&#34;1e-322&#34;, &#34;1e-322&#34;, nil},
    <a id="L75"></a><span class="comment">// smallest denormal</span>
    <a id="L76"></a>atofTest{&#34;5e-324&#34;, &#34;5e-324&#34;, nil},
    <a id="L77"></a><span class="comment">// too small</span>
    <a id="L78"></a>atofTest{&#34;4e-324&#34;, &#34;0&#34;, nil},
    <a id="L79"></a><span class="comment">// way too small</span>
    <a id="L80"></a>atofTest{&#34;1e-350&#34;, &#34;0&#34;, nil},
    <a id="L81"></a>atofTest{&#34;1e-400000&#34;, &#34;0&#34;, nil},

    <a id="L83"></a><span class="comment">// try to overflow exponent</span>
    <a id="L84"></a>atofTest{&#34;1e-4294967296&#34;, &#34;0&#34;, nil},
    <a id="L85"></a>atofTest{&#34;1e+4294967296&#34;, &#34;+Inf&#34;, os.ERANGE},
    <a id="L86"></a>atofTest{&#34;1e-18446744073709551616&#34;, &#34;0&#34;, nil},
    <a id="L87"></a>atofTest{&#34;1e+18446744073709551616&#34;, &#34;+Inf&#34;, os.ERANGE},

    <a id="L89"></a><span class="comment">// Parse errors</span>
    <a id="L90"></a>atofTest{&#34;1e&#34;, &#34;0&#34;, os.EINVAL},
    <a id="L91"></a>atofTest{&#34;1e-&#34;, &#34;0&#34;, os.EINVAL},
    <a id="L92"></a>atofTest{&#34;.e-1&#34;, &#34;0&#34;, os.EINVAL},
<a id="L93"></a>}

<a id="L95"></a>func init() {
    <a id="L96"></a><span class="comment">// The atof routines return NumErrors wrapping</span>
    <a id="L97"></a><span class="comment">// the error and the string.  Convert the table above.</span>
    <a id="L98"></a>for i := range atoftests {
        <a id="L99"></a>test := &amp;atoftests[i];
        <a id="L100"></a>if test.err != nil {
            <a id="L101"></a>test.err = &amp;NumError{test.in, test.err}
        <a id="L102"></a>}
    <a id="L103"></a>}
<a id="L104"></a>}

<a id="L106"></a>func testAtof(t *testing.T, opt bool) {
    <a id="L107"></a>oldopt := SetOptimize(opt);
    <a id="L108"></a>for i := 0; i &lt; len(atoftests); i++ {
        <a id="L109"></a>test := &amp;atoftests[i];
        <a id="L110"></a>out, err := Atof64(test.in);
        <a id="L111"></a>outs := Ftoa64(out, &#39;g&#39;, -1);
        <a id="L112"></a>if outs != test.out || !reflect.DeepEqual(err, test.err) {
            <a id="L113"></a>t.Errorf(&#34;Atof64(%v) = %v, %v want %v, %v\n&#34;,
                <a id="L114"></a>test.in, out, err, test.out, test.err)
        <a id="L115"></a>}

        <a id="L117"></a>if float64(float32(out)) == out {
            <a id="L118"></a>out32, err := Atof32(test.in);
            <a id="L119"></a>outs := Ftoa32(out32, &#39;g&#39;, -1);
            <a id="L120"></a>if outs != test.out || !reflect.DeepEqual(err, test.err) {
                <a id="L121"></a>t.Errorf(&#34;Atof32(%v) = %v, %v want %v, %v  # %v\n&#34;,
                    <a id="L122"></a>test.in, out32, err, test.out, test.err, out)
            <a id="L123"></a>}
        <a id="L124"></a>}

        <a id="L126"></a>if FloatSize == 64 || float64(float32(out)) == out {
            <a id="L127"></a>outf, err := Atof(test.in);
            <a id="L128"></a>outs := Ftoa(outf, &#39;g&#39;, -1);
            <a id="L129"></a>if outs != test.out || !reflect.DeepEqual(err, test.err) {
                <a id="L130"></a>t.Errorf(&#34;Ftoa(%v) = %v, %v want %v, %v  # %v\n&#34;,
                    <a id="L131"></a>test.in, outf, err, test.out, test.err, out)
            <a id="L132"></a>}
        <a id="L133"></a>}
    <a id="L134"></a>}
    <a id="L135"></a>SetOptimize(oldopt);
<a id="L136"></a>}

<a id="L138"></a>func TestAtof(t *testing.T) { testAtof(t, true) }

<a id="L140"></a>func TestAtofSlow(t *testing.T) { testAtof(t, false) }
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
