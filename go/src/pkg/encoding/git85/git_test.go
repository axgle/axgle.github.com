<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/git85/git_test.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/git85/git_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package git85

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;testing&#34;;
<a id="L14"></a>)

<a id="L16"></a>type testpair struct {
    <a id="L17"></a>decoded, encoded string;
<a id="L18"></a>}

<a id="L20"></a>func testEqual(t *testing.T, msg string, args ...) bool {
    <a id="L21"></a>v := reflect.NewValue(args).(*reflect.StructValue);
    <a id="L22"></a>v1 := v.Field(v.NumField() - 2);
    <a id="L23"></a>v2 := v.Field(v.NumField() - 1);
    <a id="L24"></a>if v1.Interface() != v2.Interface() {
        <a id="L25"></a>t.Errorf(msg, args);
        <a id="L26"></a>return false;
    <a id="L27"></a>}
    <a id="L28"></a>return true;
<a id="L29"></a>}

<a id="L31"></a>func TestGitTable(t *testing.T) {
    <a id="L32"></a>var saw [256]bool;
    <a id="L33"></a>for i, c := range encode {
        <a id="L34"></a>if decode[c] != uint8(i+1) {
            <a id="L35"></a>t.Errorf(&#34;decode[&#39;%c&#39;] = %d, want %d&#34;, c, decode[c], i+1)
        <a id="L36"></a>}
        <a id="L37"></a>saw[c] = true;
    <a id="L38"></a>}
    <a id="L39"></a>for i, b := range saw {
        <a id="L40"></a>if !b &amp;&amp; decode[i] != 0 {
            <a id="L41"></a>t.Errorf(&#34;decode[%d] = %d, want 0&#34;, i, decode[i])
        <a id="L42"></a>}
    <a id="L43"></a>}
<a id="L44"></a>}

<a id="L46"></a>var gitPairs = []testpair{
    <a id="L47"></a><span class="comment">// Wikipedia example, adapted.</span>
    <a id="L48"></a>testpair{
        <a id="L49"></a>&#34;Man is distinguished, not only by his reason, but by this singular passion from &#34;
            <a id="L50"></a>&#34;other animals, which is a lust of the mind, that by a perseverance of delight in &#34;
            <a id="L51"></a>&#34;the continued and indefatigable generation of knowledge, exceeds the short &#34;
            <a id="L52"></a>&#34;vehemence of any carnal pleasure.&#34;,

        <a id="L54"></a>&#34;zO&lt;`^zX&gt;%ZCX&gt;)XGZfA9Ab7*B`EFf-gbRchTY&lt;VDJc_3(Mb0BhMVRLV8EFfZabRc4R\n&#34;
            <a id="L55"></a>&#34;zAarPHb0BkRZfA9DVR9gFVRLh7Z*CxFa&amp;K)QZ**v7av))DX&gt;DO_b1WctXlY|;AZc?T\n&#34;
            <a id="L56"></a>&#34;zVIXXEb95kYW*~HEWgu;7Ze%PVbZB98AYyqSVIXj2a&amp;u*NWpZI|V`U(3W*}r`Y-wj`\n&#34;
            <a id="L57"></a>&#34;zbRcPNAarPDAY*TCbZKsNWn&gt;^&gt;Ze$&gt;7Ze(R&lt;VRUI{VPb4$AZKN6WpZJ3X&gt;V&gt;IZ)PBC\n&#34;
            <a id="L58"></a>&#34;zZf|#NWn^b%EFfigV`XJzb0BnRWgv5CZ*p`Xc4cT~ZDnp_Wgu^6AYpEKAY);2ZeeU7\n&#34;
            <a id="L59"></a>&#34;IaBO8^b9HiME&amp;u=k\n&#34;,
    <a id="L60"></a>},
<a id="L61"></a>}

<a id="L63"></a>var gitBigtest = gitPairs[len(gitPairs)-1]

<a id="L65"></a>func TestEncode(t *testing.T) {
    <a id="L66"></a>for _, p := range gitPairs {
        <a id="L67"></a>buf := make([]byte, EncodedLen(len(p.decoded)));
        <a id="L68"></a>n := Encode(buf, strings.Bytes(p.decoded));
        <a id="L69"></a>if n != len(buf) {
            <a id="L70"></a>t.Errorf(&#34;EncodedLen does not agree with Encode&#34;)
        <a id="L71"></a>}
        <a id="L72"></a>buf = buf[0:n];
        <a id="L73"></a>testEqual(t, &#34;Encode(%q) = %q, want %q&#34;, p.decoded, string(buf), p.encoded);
    <a id="L74"></a>}
<a id="L75"></a>}

<a id="L77"></a>func TestEncoder(t *testing.T) {
    <a id="L78"></a>for _, p := range gitPairs {
        <a id="L79"></a>bb := &amp;bytes.Buffer{};
        <a id="L80"></a>encoder := NewEncoder(bb);
        <a id="L81"></a>encoder.Write(strings.Bytes(p.decoded));
        <a id="L82"></a>encoder.Close();
        <a id="L83"></a>testEqual(t, &#34;Encode(%q) = %q, want %q&#34;, p.decoded, bb.String(), p.encoded);
    <a id="L84"></a>}
<a id="L85"></a>}

<a id="L87"></a>func TestEncoderBuffering(t *testing.T) {
    <a id="L88"></a>input := strings.Bytes(gitBigtest.decoded);
    <a id="L89"></a>for bs := 1; bs &lt;= 12; bs++ {
        <a id="L90"></a>bb := &amp;bytes.Buffer{};
        <a id="L91"></a>encoder := NewEncoder(bb);
        <a id="L92"></a>for pos := 0; pos &lt; len(input); pos += bs {
            <a id="L93"></a>end := pos + bs;
            <a id="L94"></a>if end &gt; len(input) {
                <a id="L95"></a>end = len(input)
            <a id="L96"></a>}
            <a id="L97"></a>n, err := encoder.Write(input[pos:end]);
            <a id="L98"></a>testEqual(t, &#34;Write(%q) gave error %v, want %v&#34;, input[pos:end], err, os.Error(nil));
            <a id="L99"></a>testEqual(t, &#34;Write(%q) gave length %v, want %v&#34;, input[pos:end], n, end-pos);
        <a id="L100"></a>}
        <a id="L101"></a>err := encoder.Close();
        <a id="L102"></a>testEqual(t, &#34;Close gave error %v, want %v&#34;, err, os.Error(nil));
        <a id="L103"></a>testEqual(t, &#34;Encoding/%d of %q = %q, want %q&#34;, bs, gitBigtest.decoded, bb.String(), gitBigtest.encoded);
    <a id="L104"></a>}
<a id="L105"></a>}

<a id="L107"></a>func TestDecode(t *testing.T) {
    <a id="L108"></a>for _, p := range gitPairs {
        <a id="L109"></a>dbuf := make([]byte, 4*len(p.encoded));
        <a id="L110"></a>ndst, err := Decode(dbuf, strings.Bytes(p.encoded));
        <a id="L111"></a>testEqual(t, &#34;Decode(%q) = error %v, want %v&#34;, p.encoded, err, os.Error(nil));
        <a id="L112"></a>testEqual(t, &#34;Decode(%q) = ndst %v, want %v&#34;, p.encoded, ndst, len(p.decoded));
        <a id="L113"></a>testEqual(t, &#34;Decode(%q) = %q, want %q&#34;, p.encoded, string(dbuf[0:ndst]), p.decoded);
    <a id="L114"></a>}
<a id="L115"></a>}

<a id="L117"></a>func TestDecoder(t *testing.T) {
    <a id="L118"></a>for _, p := range gitPairs {
        <a id="L119"></a>decoder := NewDecoder(bytes.NewBufferString(p.encoded));
        <a id="L120"></a>dbuf, err := io.ReadAll(decoder);
        <a id="L121"></a>if err != nil {
            <a id="L122"></a>t.Fatal(&#34;Read failed&#34;, err)
        <a id="L123"></a>}
        <a id="L124"></a>testEqual(t, &#34;Read from %q = length %v, want %v&#34;, p.encoded, len(dbuf), len(p.decoded));
        <a id="L125"></a>testEqual(t, &#34;Decoding of %q = %q, want %q&#34;, p.encoded, string(dbuf), p.decoded);
        <a id="L126"></a>if err != nil {
            <a id="L127"></a>testEqual(t, &#34;Read from %q = %v, want %v&#34;, p.encoded, err, os.EOF)
        <a id="L128"></a>}
    <a id="L129"></a>}
<a id="L130"></a>}

<a id="L132"></a>func TestDecoderBuffering(t *testing.T) {
    <a id="L133"></a>for bs := 1; bs &lt;= 12; bs++ {
        <a id="L134"></a>decoder := NewDecoder(bytes.NewBufferString(gitBigtest.encoded));
        <a id="L135"></a>buf := make([]byte, len(gitBigtest.decoded)+12);
        <a id="L136"></a>var total int;
        <a id="L137"></a>for total = 0; total &lt; len(gitBigtest.decoded); {
            <a id="L138"></a>n, err := decoder.Read(buf[total : total+bs]);
            <a id="L139"></a>testEqual(t, &#34;Read from %q at pos %d = %d, %v, want _, %v&#34;, gitBigtest.encoded, total, n, err, os.Error(nil));
            <a id="L140"></a>total += n;
        <a id="L141"></a>}
        <a id="L142"></a>testEqual(t, &#34;Decoding/%d of %q = %q, want %q&#34;, bs, gitBigtest.encoded, string(buf[0:total]), gitBigtest.decoded);
    <a id="L143"></a>}
<a id="L144"></a>}

<a id="L146"></a>func TestDecodeCorrupt(t *testing.T) {
    <a id="L147"></a>type corrupt struct {
        <a id="L148"></a>e   string;
        <a id="L149"></a>p   int;
    <a id="L150"></a>}
    <a id="L151"></a>examples := []corrupt{
        <a id="L152"></a>corrupt{&#34;v&#34;, 0},
        <a id="L153"></a>corrupt{&#34;!z!!!!!!!!!&#34;, 0},
    <a id="L154"></a>};

    <a id="L156"></a>for _, e := range examples {
        <a id="L157"></a>dbuf := make([]byte, 2*len(e.e));
        <a id="L158"></a>_, err := Decode(dbuf, strings.Bytes(e.e));
        <a id="L159"></a>switch err := err.(type) {
        <a id="L160"></a>case CorruptInputError:
            <a id="L161"></a>testEqual(t, &#34;Corruption in %q at offset %v, want %v&#34;, e.e, int(err), e.p)
        <a id="L162"></a>default:
            <a id="L163"></a>t.Error(&#34;Decoder failed to detect corruption in&#34;, e)
        <a id="L164"></a>}
    <a id="L165"></a>}
<a id="L166"></a>}

<a id="L168"></a>func TestGitBig(t *testing.T) {
    <a id="L169"></a>n := 3*1000 + 1;
    <a id="L170"></a>raw := make([]byte, n);
    <a id="L171"></a>const alpha = &#34;0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ&#34;;
    <a id="L172"></a>for i := 0; i &lt; n; i++ {
        <a id="L173"></a>raw[i] = alpha[i%len(alpha)]
    <a id="L174"></a>}
    <a id="L175"></a>encoded := new(bytes.Buffer);
    <a id="L176"></a>w := NewEncoder(encoded);
    <a id="L177"></a>nn, err := w.Write(raw);
    <a id="L178"></a>if nn != n || err != nil {
        <a id="L179"></a>t.Fatalf(&#34;Encoder.Write(raw) = %d, %v want %d, nil&#34;, nn, err, n)
    <a id="L180"></a>}
    <a id="L181"></a>err = w.Close();
    <a id="L182"></a>if err != nil {
        <a id="L183"></a>t.Fatalf(&#34;Encoder.Close() = %v want nil&#34;, err)
    <a id="L184"></a>}
    <a id="L185"></a>decoded, err := io.ReadAll(NewDecoder(encoded));
    <a id="L186"></a>if err != nil {
        <a id="L187"></a>t.Fatalf(&#34;io.ReadAll(NewDecoder(...)): %v&#34;, err)
    <a id="L188"></a>}

    <a id="L190"></a>if !bytes.Equal(raw, decoded) {
        <a id="L191"></a>var i int;
        <a id="L192"></a>for i = 0; i &lt; len(decoded) &amp;&amp; i &lt; len(raw); i++ {
            <a id="L193"></a>if decoded[i] != raw[i] {
                <a id="L194"></a>break
            <a id="L195"></a>}
        <a id="L196"></a>}
        <a id="L197"></a>t.Errorf(&#34;Decode(Encode(%d-byte string)) failed at offset %d&#34;, n, i);
    <a id="L198"></a>}
<a id="L199"></a>}
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
