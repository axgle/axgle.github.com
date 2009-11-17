<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/md5/md5_test.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/md5/md5_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package md5

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)

<a id="L13"></a>type md5Test struct {
    <a id="L14"></a>out string;
    <a id="L15"></a>in  string;
<a id="L16"></a>}

<a id="L18"></a>var golden = []md5Test{
    <a id="L19"></a>md5Test{&#34;d41d8cd98f00b204e9800998ecf8427e&#34;, &#34;&#34;},
    <a id="L20"></a>md5Test{&#34;0cc175b9c0f1b6a831c399e269772661&#34;, &#34;a&#34;},
    <a id="L21"></a>md5Test{&#34;187ef4436122d1cc2f40dc2b92f0eba0&#34;, &#34;ab&#34;},
    <a id="L22"></a>md5Test{&#34;900150983cd24fb0d6963f7d28e17f72&#34;, &#34;abc&#34;},
    <a id="L23"></a>md5Test{&#34;e2fc714c4727ee9395f324cd2e7f331f&#34;, &#34;abcd&#34;},
    <a id="L24"></a>md5Test{&#34;ab56b4d92b40713acc5af89985d4b786&#34;, &#34;abcde&#34;},
    <a id="L25"></a>md5Test{&#34;e80b5017098950fc58aad83c8c14978e&#34;, &#34;abcdef&#34;},
    <a id="L26"></a>md5Test{&#34;7ac66c0f148de9519b8bd264312c4d64&#34;, &#34;abcdefg&#34;},
    <a id="L27"></a>md5Test{&#34;e8dc4081b13434b45189a720b77b6818&#34;, &#34;abcdefgh&#34;},
    <a id="L28"></a>md5Test{&#34;8aa99b1f439ff71293e95357bac6fd94&#34;, &#34;abcdefghi&#34;},
    <a id="L29"></a>md5Test{&#34;a925576942e94b2ef57a066101b48876&#34;, &#34;abcdefghij&#34;},
    <a id="L30"></a>md5Test{&#34;d747fc1719c7eacb84058196cfe56d57&#34;, &#34;Discard medicine more than two years old.&#34;},
    <a id="L31"></a>md5Test{&#34;bff2dcb37ef3a44ba43ab144768ca837&#34;, &#34;He who has a shady past knows that nice guys finish last.&#34;},
    <a id="L32"></a>md5Test{&#34;0441015ecb54a7342d017ed1bcfdbea5&#34;, &#34;I wouldn&#39;t marry him with a ten foot pole.&#34;},
    <a id="L33"></a>md5Test{&#34;9e3cac8e9e9757a60c3ea391130d3689&#34;, &#34;Free! Free!/A trip/to Mars/for 900/empty jars/Burma Shave&#34;},
    <a id="L34"></a>md5Test{&#34;a0f04459b031f916a59a35cc482dc039&#34;, &#34;The days of the digital watch are numbered.  -Tom Stoppard&#34;},
    <a id="L35"></a>md5Test{&#34;e7a48e0fe884faf31475d2a04b1362cc&#34;, &#34;Nepal premier won&#39;t resign.&#34;},
    <a id="L36"></a>md5Test{&#34;637d2fe925c07c113800509964fb0e06&#34;, &#34;For every action there is an equal and opposite government program.&#34;},
    <a id="L37"></a>md5Test{&#34;834a8d18d5c6562119cf4c7f5086cb71&#34;, &#34;His money is twice tainted: &#39;taint yours and &#39;taint mine.&#34;},
    <a id="L38"></a>md5Test{&#34;de3a4d2fd6c73ec2db2abad23b444281&#34;, &#34;There is no reason for any individual to have a computer in their home. -Ken Olsen, 1977&#34;},
    <a id="L39"></a>md5Test{&#34;acf203f997e2cf74ea3aff86985aefaf&#34;, &#34;It&#39;s a tiny change to the code and not completely disgusting. - Bob Manchek&#34;},
    <a id="L40"></a>md5Test{&#34;e1c1384cb4d2221dfdd7c795a4222c9a&#34;, &#34;size:  a.out:  bad magic&#34;},
    <a id="L41"></a>md5Test{&#34;c90f3ddecc54f34228c063d7525bf644&#34;, &#34;The major problem is with sendmail.  -Mark Horton&#34;},
    <a id="L42"></a>md5Test{&#34;cdf7ab6c1fd49bd9933c43f3ea5af185&#34;, &#34;Give me a rock, paper and scissors and I will move the world.  CCFestoon&#34;},
    <a id="L43"></a>md5Test{&#34;83bc85234942fc883c063cbd7f0ad5d0&#34;, &#34;If the enemy is within range, then so are you.&#34;},
    <a id="L44"></a>md5Test{&#34;277cbe255686b48dd7e8f389394d9299&#34;, &#34;It&#39;s well we cannot hear the screams/That we create in others&#39; dreams.&#34;},
    <a id="L45"></a>md5Test{&#34;fd3fb0a7ffb8af16603f3d3af98f8e1f&#34;, &#34;You remind me of a TV show, but that&#39;s all right: I watch it anyway.&#34;},
    <a id="L46"></a>md5Test{&#34;469b13a78ebf297ecda64d4723655154&#34;, &#34;C is as portable as Stonehedge!!&#34;},
    <a id="L47"></a>md5Test{&#34;63eb3a2f466410104731c4b037600110&#34;, &#34;Even if I could be Shakespeare, I think I should still choose to be Faraday. - A. Huxley&#34;},
    <a id="L48"></a>md5Test{&#34;72c2ed7592debca1c90fc0100f931a2f&#34;, &#34;The fugacity of a constituent in a mixture of gases at a given temperature is proportional to its mole fraction.  Lewis-Randall Rule&#34;},
    <a id="L49"></a>md5Test{&#34;132f7619d33b523b1d9e5bd8e0928355&#34;, &#34;How can you write a big system without C++?  -Paul Glick&#34;},
<a id="L50"></a>}

<a id="L52"></a>func TestGolden(t *testing.T) {
    <a id="L53"></a>for i := 0; i &lt; len(golden); i++ {
        <a id="L54"></a>g := golden[i];
        <a id="L55"></a>c := New();
        <a id="L56"></a>for j := 0; j &lt; 2; j++ {
            <a id="L57"></a>io.WriteString(c, g.in);
            <a id="L58"></a>s := fmt.Sprintf(&#34;%x&#34;, c.Sum());
            <a id="L59"></a>if s != g.out {
                <a id="L60"></a>t.Errorf(&#34;md5[%d](%s) = %s want %s&#34;, j, g.in, s, g.out);
                <a id="L61"></a>t.FailNow();
            <a id="L62"></a>}
            <a id="L63"></a>c.Reset();
        <a id="L64"></a>}
    <a id="L65"></a>}
<a id="L66"></a>}
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
