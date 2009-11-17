<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bytes/buffer_test.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bytes/buffer_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package bytes_test

<a id="L7"></a>import (
    <a id="L8"></a>. &#34;bytes&#34;;
    <a id="L9"></a>&#34;rand&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)


<a id="L14"></a>const N = 10000  <span class="comment">// make this bigger for a larger (and slower) test</span>
<a id="L15"></a>var data string  <span class="comment">// test data for write tests</span>
<a id="L16"></a>var bytes []byte <span class="comment">// test data; same as data but as a slice.</span>


<a id="L19"></a>func init() {
    <a id="L20"></a>bytes = make([]byte, N);
    <a id="L21"></a>for i := 0; i &lt; N; i++ {
        <a id="L22"></a>bytes[i] = &#39;a&#39; + byte(i%26)
    <a id="L23"></a>}
    <a id="L24"></a>data = string(bytes);
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// Verify that contents of buf match the string s.</span>
<a id="L28"></a>func check(t *testing.T, testname string, buf *Buffer, s string) {
    <a id="L29"></a>bytes := buf.Bytes();
    <a id="L30"></a>str := buf.String();
    <a id="L31"></a>if buf.Len() != len(bytes) {
        <a id="L32"></a>t.Errorf(&#34;%s: buf.Len() == %d, len(buf.Bytes()) == %d\n&#34;, testname, buf.Len(), len(bytes))
    <a id="L33"></a>}

    <a id="L35"></a>if buf.Len() != len(str) {
        <a id="L36"></a>t.Errorf(&#34;%s: buf.Len() == %d, len(buf.String()) == %d\n&#34;, testname, buf.Len(), len(str))
    <a id="L37"></a>}

    <a id="L39"></a>if buf.Len() != len(s) {
        <a id="L40"></a>t.Errorf(&#34;%s: buf.Len() == %d, len(s) == %d\n&#34;, testname, buf.Len(), len(s))
    <a id="L41"></a>}

    <a id="L43"></a>if string(bytes) != s {
        <a id="L44"></a>t.Errorf(&#34;%s: string(buf.Bytes()) == %q, s == %q\n&#34;, testname, string(bytes), s)
    <a id="L45"></a>}
<a id="L46"></a>}


<a id="L49"></a><span class="comment">// Fill buf through n writes of string fus.</span>
<a id="L50"></a><span class="comment">// The initial contents of buf corresponds to the string s;</span>
<a id="L51"></a><span class="comment">// the result is the final contents of buf returned as a string.</span>
<a id="L52"></a>func fillString(t *testing.T, testname string, buf *Buffer, s string, n int, fus string) string {
    <a id="L53"></a>check(t, testname+&#34; (fill 1)&#34;, buf, s);
    <a id="L54"></a>for ; n &gt; 0; n-- {
        <a id="L55"></a>m, err := buf.WriteString(fus);
        <a id="L56"></a>if m != len(fus) {
            <a id="L57"></a>t.Errorf(testname+&#34; (fill 2): m == %d, expected %d\n&#34;, m, len(fus))
        <a id="L58"></a>}
        <a id="L59"></a>if err != nil {
            <a id="L60"></a>t.Errorf(testname+&#34; (fill 3): err should always be nil, found err == %s\n&#34;, err)
        <a id="L61"></a>}
        <a id="L62"></a>s += fus;
        <a id="L63"></a>check(t, testname+&#34; (fill 4)&#34;, buf, s);
    <a id="L64"></a>}
    <a id="L65"></a>return s;
<a id="L66"></a>}


<a id="L69"></a><span class="comment">// Fill buf through n writes of byte slice fub.</span>
<a id="L70"></a><span class="comment">// The initial contents of buf corresponds to the string s;</span>
<a id="L71"></a><span class="comment">// the result is the final contents of buf returned as a string.</span>
<a id="L72"></a>func fillBytes(t *testing.T, testname string, buf *Buffer, s string, n int, fub []byte) string {
    <a id="L73"></a>check(t, testname+&#34; (fill 1)&#34;, buf, s);
    <a id="L74"></a>for ; n &gt; 0; n-- {
        <a id="L75"></a>m, err := buf.Write(fub);
        <a id="L76"></a>if m != len(fub) {
            <a id="L77"></a>t.Errorf(testname+&#34; (fill 2): m == %d, expected %d\n&#34;, m, len(fub))
        <a id="L78"></a>}
        <a id="L79"></a>if err != nil {
            <a id="L80"></a>t.Errorf(testname+&#34; (fill 3): err should always be nil, found err == %s\n&#34;, err)
        <a id="L81"></a>}
        <a id="L82"></a>s += string(fub);
        <a id="L83"></a>check(t, testname+&#34; (fill 4)&#34;, buf, s);
    <a id="L84"></a>}
    <a id="L85"></a>return s;
<a id="L86"></a>}


<a id="L89"></a>func TestNewBuffer(t *testing.T) {
    <a id="L90"></a>buf := NewBuffer(bytes);
    <a id="L91"></a>check(t, &#34;NewBuffer&#34;, buf, data);
<a id="L92"></a>}


<a id="L95"></a>func TestNewBufferString(t *testing.T) {
    <a id="L96"></a>buf := NewBufferString(data);
    <a id="L97"></a>check(t, &#34;NewBufferString&#34;, buf, data);
<a id="L98"></a>}


<a id="L101"></a><span class="comment">// Empty buf through repeated reads into fub.</span>
<a id="L102"></a><span class="comment">// The initial contents of buf corresponds to the string s.</span>
<a id="L103"></a>func empty(t *testing.T, testname string, buf *Buffer, s string, fub []byte) {
    <a id="L104"></a>check(t, testname+&#34; (empty 1)&#34;, buf, s);

    <a id="L106"></a>for {
        <a id="L107"></a>n, err := buf.Read(fub);
        <a id="L108"></a>if n == 0 {
            <a id="L109"></a>break
        <a id="L110"></a>}
        <a id="L111"></a>if err != nil {
            <a id="L112"></a>t.Errorf(testname+&#34; (empty 2): err should always be nil, found err == %s\n&#34;, err)
        <a id="L113"></a>}
        <a id="L114"></a>s = s[n:len(s)];
        <a id="L115"></a>check(t, testname+&#34; (empty 3)&#34;, buf, s);
    <a id="L116"></a>}

    <a id="L118"></a>check(t, testname+&#34; (empty 4)&#34;, buf, &#34;&#34;);
<a id="L119"></a>}


<a id="L122"></a>func TestBasicOperations(t *testing.T) {
    <a id="L123"></a>var buf Buffer;

    <a id="L125"></a>for i := 0; i &lt; 5; i++ {
        <a id="L126"></a>check(t, &#34;TestBasicOperations (1)&#34;, &amp;buf, &#34;&#34;);

        <a id="L128"></a>buf.Reset();
        <a id="L129"></a>check(t, &#34;TestBasicOperations (2)&#34;, &amp;buf, &#34;&#34;);

        <a id="L131"></a>buf.Truncate(0);
        <a id="L132"></a>check(t, &#34;TestBasicOperations (3)&#34;, &amp;buf, &#34;&#34;);

        <a id="L134"></a>n, err := buf.Write(Bytes(data[0:1]));
        <a id="L135"></a>if n != 1 {
            <a id="L136"></a>t.Errorf(&#34;wrote 1 byte, but n == %d\n&#34;, n)
        <a id="L137"></a>}
        <a id="L138"></a>if err != nil {
            <a id="L139"></a>t.Errorf(&#34;err should always be nil, but err == %s\n&#34;, err)
        <a id="L140"></a>}
        <a id="L141"></a>check(t, &#34;TestBasicOperations (4)&#34;, &amp;buf, &#34;a&#34;);

        <a id="L143"></a>buf.WriteByte(data[1]);
        <a id="L144"></a>check(t, &#34;TestBasicOperations (5)&#34;, &amp;buf, &#34;ab&#34;);

        <a id="L146"></a>n, err = buf.Write(Bytes(data[2:26]));
        <a id="L147"></a>if n != 24 {
            <a id="L148"></a>t.Errorf(&#34;wrote 25 bytes, but n == %d\n&#34;, n)
        <a id="L149"></a>}
        <a id="L150"></a>check(t, &#34;TestBasicOperations (6)&#34;, &amp;buf, string(data[0:26]));

        <a id="L152"></a>buf.Truncate(26);
        <a id="L153"></a>check(t, &#34;TestBasicOperations (7)&#34;, &amp;buf, string(data[0:26]));

        <a id="L155"></a>buf.Truncate(20);
        <a id="L156"></a>check(t, &#34;TestBasicOperations (8)&#34;, &amp;buf, string(data[0:20]));

        <a id="L158"></a>empty(t, &#34;TestBasicOperations (9)&#34;, &amp;buf, string(data[0:20]), make([]byte, 5));
        <a id="L159"></a>empty(t, &#34;TestBasicOperations (10)&#34;, &amp;buf, &#34;&#34;, make([]byte, 100));

        <a id="L161"></a>buf.WriteByte(data[1]);
        <a id="L162"></a>c, err := buf.ReadByte();
        <a id="L163"></a>if err != nil {
            <a id="L164"></a>t.Errorf(&#34;ReadByte unexpected eof\n&#34;)
        <a id="L165"></a>}
        <a id="L166"></a>if c != data[1] {
            <a id="L167"></a>t.Errorf(&#34;ReadByte wrong value c=%v\n&#34;, c)
        <a id="L168"></a>}
        <a id="L169"></a>c, err = buf.ReadByte();
        <a id="L170"></a>if err == nil {
            <a id="L171"></a>t.Errorf(&#34;ReadByte unexpected not eof\n&#34;)
        <a id="L172"></a>}
    <a id="L173"></a>}
<a id="L174"></a>}


<a id="L177"></a>func TestLargeStringWrites(t *testing.T) {
    <a id="L178"></a>var buf Buffer;
    <a id="L179"></a>for i := 3; i &lt; 30; i += 3 {
        <a id="L180"></a>s := fillString(t, &#34;TestLargeWrites (1)&#34;, &amp;buf, &#34;&#34;, 5, data);
        <a id="L181"></a>empty(t, &#34;TestLargeStringWrites (2)&#34;, &amp;buf, s, make([]byte, len(data)/i));
    <a id="L182"></a>}
    <a id="L183"></a>check(t, &#34;TestLargeStringWrites (3)&#34;, &amp;buf, &#34;&#34;);
<a id="L184"></a>}


<a id="L187"></a>func TestLargeByteWrites(t *testing.T) {
    <a id="L188"></a>var buf Buffer;
    <a id="L189"></a>for i := 3; i &lt; 30; i += 3 {
        <a id="L190"></a>s := fillBytes(t, &#34;TestLargeWrites (1)&#34;, &amp;buf, &#34;&#34;, 5, bytes);
        <a id="L191"></a>empty(t, &#34;TestLargeByteWrites (2)&#34;, &amp;buf, s, make([]byte, len(data)/i));
    <a id="L192"></a>}
    <a id="L193"></a>check(t, &#34;TestLargeByteWrites (3)&#34;, &amp;buf, &#34;&#34;);
<a id="L194"></a>}


<a id="L197"></a>func TestLargeStringReads(t *testing.T) {
    <a id="L198"></a>var buf Buffer;
    <a id="L199"></a>for i := 3; i &lt; 30; i += 3 {
        <a id="L200"></a>s := fillString(t, &#34;TestLargeReads (1)&#34;, &amp;buf, &#34;&#34;, 5, data[0:len(data)/i]);
        <a id="L201"></a>empty(t, &#34;TestLargeReads (2)&#34;, &amp;buf, s, make([]byte, len(data)));
    <a id="L202"></a>}
    <a id="L203"></a>check(t, &#34;TestLargeStringReads (3)&#34;, &amp;buf, &#34;&#34;);
<a id="L204"></a>}


<a id="L207"></a>func TestLargeByteReads(t *testing.T) {
    <a id="L208"></a>var buf Buffer;
    <a id="L209"></a>for i := 3; i &lt; 30; i += 3 {
        <a id="L210"></a>s := fillBytes(t, &#34;TestLargeReads (1)&#34;, &amp;buf, &#34;&#34;, 5, bytes[0:len(bytes)/i]);
        <a id="L211"></a>empty(t, &#34;TestLargeReads (2)&#34;, &amp;buf, s, make([]byte, len(data)));
    <a id="L212"></a>}
    <a id="L213"></a>check(t, &#34;TestLargeByteReads (3)&#34;, &amp;buf, &#34;&#34;);
<a id="L214"></a>}


<a id="L217"></a>func TestMixedReadsAndWrites(t *testing.T) {
    <a id="L218"></a>var buf Buffer;
    <a id="L219"></a>s := &#34;&#34;;
    <a id="L220"></a>for i := 0; i &lt; 50; i++ {
        <a id="L221"></a>wlen := rand.Intn(len(data));
        <a id="L222"></a>if i%2 == 0 {
            <a id="L223"></a>s = fillString(t, &#34;TestMixedReadsAndWrites (1)&#34;, &amp;buf, s, 1, data[0:wlen])
        <a id="L224"></a>} else {
            <a id="L225"></a>s = fillBytes(t, &#34;TestMixedReadsAndWrites (1)&#34;, &amp;buf, s, 1, bytes[0:wlen])
        <a id="L226"></a>}

        <a id="L228"></a>rlen := rand.Intn(len(data));
        <a id="L229"></a>fub := make([]byte, rlen);
        <a id="L230"></a>n, _ := buf.Read(fub);
        <a id="L231"></a>s = s[n:len(s)];
    <a id="L232"></a>}
    <a id="L233"></a>empty(t, &#34;TestMixedReadsAndWrites (2)&#34;, &amp;buf, s, make([]byte, buf.Len()));
<a id="L234"></a>}


<a id="L237"></a>func TestNil(t *testing.T) {
    <a id="L238"></a>var b *Buffer;
    <a id="L239"></a>if b.String() != &#34;&lt;nil&gt;&#34; {
        <a id="L240"></a>t.Error(&#34;expcted &lt;nil&gt;; got %q&#34;, b.String())
    <a id="L241"></a>}
<a id="L242"></a>}
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
