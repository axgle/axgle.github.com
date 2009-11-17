<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/base64/base64_test.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/base64/base64_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package base64

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

<a id="L20"></a>var pairs = []testpair{
    <a id="L21"></a><span class="comment">// RFC 3548 examples</span>
    <a id="L22"></a>testpair{&#34;\x14\xfb\x9c\x03\xd9\x7e&#34;, &#34;FPucA9l+&#34;},
    <a id="L23"></a>testpair{&#34;\x14\xfb\x9c\x03\xd9&#34;, &#34;FPucA9k=&#34;},
    <a id="L24"></a>testpair{&#34;\x14\xfb\x9c\x03&#34;, &#34;FPucAw==&#34;},

    <a id="L26"></a><span class="comment">// RFC 4648 examples</span>
    <a id="L27"></a>testpair{&#34;&#34;, &#34;&#34;},
    <a id="L28"></a>testpair{&#34;f&#34;, &#34;Zg==&#34;},
    <a id="L29"></a>testpair{&#34;fo&#34;, &#34;Zm8=&#34;},
    <a id="L30"></a>testpair{&#34;foo&#34;, &#34;Zm9v&#34;},
    <a id="L31"></a>testpair{&#34;foob&#34;, &#34;Zm9vYg==&#34;},
    <a id="L32"></a>testpair{&#34;fooba&#34;, &#34;Zm9vYmE=&#34;},
    <a id="L33"></a>testpair{&#34;foobar&#34;, &#34;Zm9vYmFy&#34;},

    <a id="L35"></a><span class="comment">// Wikipedia examples</span>
    <a id="L36"></a>testpair{&#34;sure.&#34;, &#34;c3VyZS4=&#34;},
    <a id="L37"></a>testpair{&#34;sure&#34;, &#34;c3VyZQ==&#34;},
    <a id="L38"></a>testpair{&#34;sur&#34;, &#34;c3Vy&#34;},
    <a id="L39"></a>testpair{&#34;su&#34;, &#34;c3U=&#34;},
    <a id="L40"></a>testpair{&#34;leasure.&#34;, &#34;bGVhc3VyZS4=&#34;},
    <a id="L41"></a>testpair{&#34;easure.&#34;, &#34;ZWFzdXJlLg==&#34;},
    <a id="L42"></a>testpair{&#34;asure.&#34;, &#34;YXN1cmUu&#34;},
    <a id="L43"></a>testpair{&#34;sure.&#34;, &#34;c3VyZS4=&#34;},
<a id="L44"></a>}

<a id="L46"></a>var bigtest = testpair{
    <a id="L47"></a>&#34;Twas brillig, and the slithy toves&#34;,
    <a id="L48"></a>&#34;VHdhcyBicmlsbGlnLCBhbmQgdGhlIHNsaXRoeSB0b3Zlcw==&#34;,
<a id="L49"></a>}

<a id="L51"></a>func testEqual(t *testing.T, msg string, args ...) bool {
    <a id="L52"></a>v := reflect.NewValue(args).(*reflect.StructValue);
    <a id="L53"></a>v1 := v.Field(v.NumField() - 2);
    <a id="L54"></a>v2 := v.Field(v.NumField() - 1);
    <a id="L55"></a>if v1.Interface() != v2.Interface() {
        <a id="L56"></a>t.Errorf(msg, args);
        <a id="L57"></a>return false;
    <a id="L58"></a>}
    <a id="L59"></a>return true;
<a id="L60"></a>}

<a id="L62"></a>func TestEncode(t *testing.T) {
    <a id="L63"></a>for _, p := range pairs {
        <a id="L64"></a>buf := make([]byte, StdEncoding.EncodedLen(len(p.decoded)));
        <a id="L65"></a>StdEncoding.Encode(buf, strings.Bytes(p.decoded));
        <a id="L66"></a>testEqual(t, &#34;Encode(%q) = %q, want %q&#34;, p.decoded, string(buf), p.encoded);
    <a id="L67"></a>}
<a id="L68"></a>}

<a id="L70"></a>func TestEncoder(t *testing.T) {
    <a id="L71"></a>for _, p := range pairs {
        <a id="L72"></a>bb := &amp;bytes.Buffer{};
        <a id="L73"></a>encoder := NewEncoder(StdEncoding, bb);
        <a id="L74"></a>encoder.Write(strings.Bytes(p.decoded));
        <a id="L75"></a>encoder.Close();
        <a id="L76"></a>testEqual(t, &#34;Encode(%q) = %q, want %q&#34;, p.decoded, bb.String(), p.encoded);
    <a id="L77"></a>}
<a id="L78"></a>}

<a id="L80"></a>func TestEncoderBuffering(t *testing.T) {
    <a id="L81"></a>input := strings.Bytes(bigtest.decoded);
    <a id="L82"></a>for bs := 1; bs &lt;= 12; bs++ {
        <a id="L83"></a>bb := &amp;bytes.Buffer{};
        <a id="L84"></a>encoder := NewEncoder(StdEncoding, bb);
        <a id="L85"></a>for pos := 0; pos &lt; len(input); pos += bs {
            <a id="L86"></a>end := pos + bs;
            <a id="L87"></a>if end &gt; len(input) {
                <a id="L88"></a>end = len(input)
            <a id="L89"></a>}
            <a id="L90"></a>n, err := encoder.Write(input[pos:end]);
            <a id="L91"></a>testEqual(t, &#34;Write(%q) gave error %v, want %v&#34;, input[pos:end], err, os.Error(nil));
            <a id="L92"></a>testEqual(t, &#34;Write(%q) gave length %v, want %v&#34;, input[pos:end], n, end-pos);
        <a id="L93"></a>}
        <a id="L94"></a>err := encoder.Close();
        <a id="L95"></a>testEqual(t, &#34;Close gave error %v, want %v&#34;, err, os.Error(nil));
        <a id="L96"></a>testEqual(t, &#34;Encoding/%d of %q = %q, want %q&#34;, bs, bigtest.decoded, bb.String(), bigtest.encoded);
    <a id="L97"></a>}
<a id="L98"></a>}

<a id="L100"></a>func TestDecode(t *testing.T) {
    <a id="L101"></a>for _, p := range pairs {
        <a id="L102"></a>dbuf := make([]byte, StdEncoding.DecodedLen(len(p.encoded)));
        <a id="L103"></a>count, end, err := StdEncoding.decode(dbuf, strings.Bytes(p.encoded));
        <a id="L104"></a>testEqual(t, &#34;Decode(%q) = error %v, want %v&#34;, p.encoded, err, os.Error(nil));
        <a id="L105"></a>testEqual(t, &#34;Decode(%q) = length %v, want %v&#34;, p.encoded, count, len(p.decoded));
        <a id="L106"></a>if len(p.encoded) &gt; 0 {
            <a id="L107"></a>testEqual(t, &#34;Decode(%q) = end %v, want %v&#34;, p.encoded, end, (p.encoded[len(p.encoded)-1] == &#39;=&#39;))
        <a id="L108"></a>}
        <a id="L109"></a>testEqual(t, &#34;Decode(%q) = %q, want %q&#34;, p.encoded, string(dbuf[0:count]), p.decoded);
    <a id="L110"></a>}
<a id="L111"></a>}

<a id="L113"></a>func TestDecoder(t *testing.T) {
    <a id="L114"></a>for _, p := range pairs {
        <a id="L115"></a>decoder := NewDecoder(StdEncoding, bytes.NewBufferString(p.encoded));
        <a id="L116"></a>dbuf := make([]byte, StdEncoding.DecodedLen(len(p.encoded)));
        <a id="L117"></a>count, err := decoder.Read(dbuf);
        <a id="L118"></a>if err != nil &amp;&amp; err != os.EOF {
            <a id="L119"></a>t.Fatal(&#34;Read failed&#34;, err)
        <a id="L120"></a>}
        <a id="L121"></a>testEqual(t, &#34;Read from %q = length %v, want %v&#34;, p.encoded, count, len(p.decoded));
        <a id="L122"></a>testEqual(t, &#34;Decoding of %q = %q, want %q&#34;, p.encoded, string(dbuf[0:count]), p.decoded);
        <a id="L123"></a>if err != os.EOF {
            <a id="L124"></a>count, err = decoder.Read(dbuf)
        <a id="L125"></a>}
        <a id="L126"></a>testEqual(t, &#34;Read from %q = %v, want %v&#34;, p.encoded, err, os.EOF);
    <a id="L127"></a>}
<a id="L128"></a>}

<a id="L130"></a>func TestDecoderBuffering(t *testing.T) {
    <a id="L131"></a>for bs := 1; bs &lt;= 12; bs++ {
        <a id="L132"></a>decoder := NewDecoder(StdEncoding, bytes.NewBufferString(bigtest.encoded));
        <a id="L133"></a>buf := make([]byte, len(bigtest.decoded)+12);
        <a id="L134"></a>var total int;
        <a id="L135"></a>for total = 0; total &lt; len(bigtest.decoded); {
            <a id="L136"></a>n, err := decoder.Read(buf[total : total+bs]);
            <a id="L137"></a>testEqual(t, &#34;Read from %q at pos %d = %d, %v, want _, %v&#34;, bigtest.encoded, total, n, err, os.Error(nil));
            <a id="L138"></a>total += n;
        <a id="L139"></a>}
        <a id="L140"></a>testEqual(t, &#34;Decoding/%d of %q = %q, want %q&#34;, bs, bigtest.encoded, string(buf[0:total]), bigtest.decoded);
    <a id="L141"></a>}
<a id="L142"></a>}

<a id="L144"></a>func TestDecodeCorrupt(t *testing.T) {
    <a id="L145"></a>type corrupt struct {
        <a id="L146"></a>e   string;
        <a id="L147"></a>p   int;
    <a id="L148"></a>}
    <a id="L149"></a>examples := []corrupt{
        <a id="L150"></a>corrupt{&#34;!!!!&#34;, 0},
        <a id="L151"></a>corrupt{&#34;x===&#34;, 1},
        <a id="L152"></a>corrupt{&#34;AA=A&#34;, 2},
        <a id="L153"></a>corrupt{&#34;AAA=AAAA&#34;, 3},
        <a id="L154"></a>corrupt{&#34;AAAAA&#34;, 4},
        <a id="L155"></a>corrupt{&#34;AAAAAA&#34;, 4},
    <a id="L156"></a>};

    <a id="L158"></a>for _, e := range examples {
        <a id="L159"></a>dbuf := make([]byte, StdEncoding.DecodedLen(len(e.e)));
        <a id="L160"></a>_, err := StdEncoding.Decode(dbuf, strings.Bytes(e.e));
        <a id="L161"></a>switch err := err.(type) {
        <a id="L162"></a>case CorruptInputError:
            <a id="L163"></a>testEqual(t, &#34;Corruption in %q at offset %v, want %v&#34;, e.e, int(err), e.p)
        <a id="L164"></a>default:
            <a id="L165"></a>t.Error(&#34;Decoder failed to detect corruption in&#34;, e)
        <a id="L166"></a>}
    <a id="L167"></a>}
<a id="L168"></a>}

<a id="L170"></a>func TestBig(t *testing.T) {
    <a id="L171"></a>n := 3*1000 + 1;
    <a id="L172"></a>raw := make([]byte, n);
    <a id="L173"></a>const alpha = &#34;0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ&#34;;
    <a id="L174"></a>for i := 0; i &lt; n; i++ {
        <a id="L175"></a>raw[i] = alpha[i%len(alpha)]
    <a id="L176"></a>}
    <a id="L177"></a>encoded := new(bytes.Buffer);
    <a id="L178"></a>w := NewEncoder(StdEncoding, encoded);
    <a id="L179"></a>nn, err := w.Write(raw);
    <a id="L180"></a>if nn != n || err != nil {
        <a id="L181"></a>t.Fatalf(&#34;Encoder.Write(raw) = %d, %v want %d, nil&#34;, nn, err, n)
    <a id="L182"></a>}
    <a id="L183"></a>err = w.Close();
    <a id="L184"></a>if err != nil {
        <a id="L185"></a>t.Fatalf(&#34;Encoder.Close() = %v want nil&#34;, err)
    <a id="L186"></a>}
    <a id="L187"></a>decoded, err := io.ReadAll(NewDecoder(StdEncoding, encoded));
    <a id="L188"></a>if err != nil {
        <a id="L189"></a>t.Fatalf(&#34;io.ReadAll(NewDecoder(...)): %v&#34;, err)
    <a id="L190"></a>}

    <a id="L192"></a>if !bytes.Equal(raw, decoded) {
        <a id="L193"></a>var i int;
        <a id="L194"></a>for i = 0; i &lt; len(decoded) &amp;&amp; i &lt; len(raw); i++ {
            <a id="L195"></a>if decoded[i] != raw[i] {
                <a id="L196"></a>break
            <a id="L197"></a>}
        <a id="L198"></a>}
        <a id="L199"></a>t.Errorf(&#34;Decode(Encode(%d-byte string)) failed at offset %d&#34;, n, i);
    <a id="L200"></a>}
<a id="L201"></a>}
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
