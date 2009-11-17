<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/sha1/sha1_test.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/sha1/sha1_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// SHA1 hash algorithm.  See RFC 3174.</span>

<a id="L7"></a>package sha1

<a id="L9"></a>import (
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>type sha1Test struct {
    <a id="L16"></a>out string;
    <a id="L17"></a>in  string;
<a id="L18"></a>}

<a id="L20"></a>var golden = []sha1Test{
    <a id="L21"></a>sha1Test{&#34;da39a3ee5e6b4b0d3255bfef95601890afd80709&#34;, &#34;&#34;},
    <a id="L22"></a>sha1Test{&#34;86f7e437faa5a7fce15d1ddcb9eaeaea377667b8&#34;, &#34;a&#34;},
    <a id="L23"></a>sha1Test{&#34;da23614e02469a0d7c7bd1bdab5c9c474b1904dc&#34;, &#34;ab&#34;},
    <a id="L24"></a>sha1Test{&#34;a9993e364706816aba3e25717850c26c9cd0d89d&#34;, &#34;abc&#34;},
    <a id="L25"></a>sha1Test{&#34;81fe8bfe87576c3ecb22426f8e57847382917acf&#34;, &#34;abcd&#34;},
    <a id="L26"></a>sha1Test{&#34;03de6c570bfe24bfc328ccd7ca46b76eadaf4334&#34;, &#34;abcde&#34;},
    <a id="L27"></a>sha1Test{&#34;1f8ac10f23c5b5bc1167bda84b833e5c057a77d2&#34;, &#34;abcdef&#34;},
    <a id="L28"></a>sha1Test{&#34;2fb5e13419fc89246865e7a324f476ec624e8740&#34;, &#34;abcdefg&#34;},
    <a id="L29"></a>sha1Test{&#34;425af12a0743502b322e93a015bcf868e324d56a&#34;, &#34;abcdefgh&#34;},
    <a id="L30"></a>sha1Test{&#34;c63b19f1e4c8b5f76b25c49b8b87f57d8e4872a1&#34;, &#34;abcdefghi&#34;},
    <a id="L31"></a>sha1Test{&#34;d68c19a0a345b7eab78d5e11e991c026ec60db63&#34;, &#34;abcdefghij&#34;},
    <a id="L32"></a>sha1Test{&#34;ebf81ddcbe5bf13aaabdc4d65354fdf2044f38a7&#34;, &#34;Discard medicine more than two years old.&#34;},
    <a id="L33"></a>sha1Test{&#34;e5dea09392dd886ca63531aaa00571dc07554bb6&#34;, &#34;He who has a shady past knows that nice guys finish last.&#34;},
    <a id="L34"></a>sha1Test{&#34;45988f7234467b94e3e9494434c96ee3609d8f8f&#34;, &#34;I wouldn&#39;t marry him with a ten foot pole.&#34;},
    <a id="L35"></a>sha1Test{&#34;55dee037eb7460d5a692d1ce11330b260e40c988&#34;, &#34;Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave&#34;},
    <a id="L36"></a>sha1Test{&#34;b7bc5fb91080c7de6b582ea281f8a396d7c0aee8&#34;, &#34;The days of the digital watch are numbered.  -Tom Stoppard&#34;},
    <a id="L37"></a>sha1Test{&#34;c3aed9358f7c77f523afe86135f06b95b3999797&#34;, &#34;Nepal premier won&#39;t resign.&#34;},
    <a id="L38"></a>sha1Test{&#34;6e29d302bf6e3a5e4305ff318d983197d6906bb9&#34;, &#34;For every action there is an equal and opposite government program.&#34;},
    <a id="L39"></a>sha1Test{&#34;597f6a540010f94c15d71806a99a2c8710e747bd&#34;, &#34;His money is twice tainted: &#39;taint yours and &#39;taint mine.&#34;},
    <a id="L40"></a>sha1Test{&#34;6859733b2590a8a091cecf50086febc5ceef1e80&#34;, &#34;There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977&#34;},
    <a id="L41"></a>sha1Test{&#34;514b2630ec089b8aee18795fc0cf1f4860cdacad&#34;, &#34;It&#39;s a tiny change to the code and not completely disgusting. - Bob Manchek&#34;},
    <a id="L42"></a>sha1Test{&#34;c5ca0d4a7b6676fc7aa72caa41cc3d5df567ed69&#34;, &#34;size:  a.out:  bad magic&#34;},
    <a id="L43"></a>sha1Test{&#34;74c51fa9a04eadc8c1bbeaa7fc442f834b90a00a&#34;, &#34;The major problem is with sendmail.  -Mark Horton&#34;},
    <a id="L44"></a>sha1Test{&#34;0b4c4ce5f52c3ad2821852a8dc00217fa18b8b66&#34;, &#34;Give me a rock, paper and scissors and I will move the world.  CCFestoon&#34;},
    <a id="L45"></a>sha1Test{&#34;3ae7937dd790315beb0f48330e8642237c61550a&#34;, &#34;If the enemy is within range, then so are you.&#34;},
    <a id="L46"></a>sha1Test{&#34;410a2b296df92b9a47412b13281df8f830a9f44b&#34;, &#34;It&#39;s well we cannot hear the screams/That we create in others&#39; dreams.&#34;},
    <a id="L47"></a>sha1Test{&#34;841e7c85ca1adcddbdd0187f1289acb5c642f7f5&#34;, &#34;You remind me of a TV show, but that&#39;s all right: I watch it anyway.&#34;},
    <a id="L48"></a>sha1Test{&#34;163173b825d03b952601376b25212df66763e1db&#34;, &#34;C is as portable as Stonehedge!!&#34;},
    <a id="L49"></a>sha1Test{&#34;32b0377f2687eb88e22106f133c586ab314d5279&#34;, &#34;Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley&#34;},
    <a id="L50"></a>sha1Test{&#34;0885aaf99b569542fd165fa44e322718f4a984e0&#34;, &#34;The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule&#34;},
    <a id="L51"></a>sha1Test{&#34;6627d6904d71420b0bf3886ab629623538689f45&#34;, &#34;How can you write a big system without C++?  -Paul Glick&#34;},
<a id="L52"></a>}

<a id="L54"></a>func TestGolden(t *testing.T) {
    <a id="L55"></a>for i := 0; i &lt; len(golden); i++ {
        <a id="L56"></a>g := golden[i];
        <a id="L57"></a>c := New();
        <a id="L58"></a>for j := 0; j &lt; 2; j++ {
            <a id="L59"></a>io.WriteString(c, g.in);
            <a id="L60"></a>s := fmt.Sprintf(&#34;%x&#34;, c.Sum());
            <a id="L61"></a>if s != g.out {
                <a id="L62"></a>t.Errorf(&#34;sha1[%d](%s) = %s want %s&#34;, j, g.in, s, g.out);
                <a id="L63"></a>t.FailNow();
            <a id="L64"></a>}
            <a id="L65"></a>c.Reset();
        <a id="L66"></a>}
    <a id="L67"></a>}
<a id="L68"></a>}
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
