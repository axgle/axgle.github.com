<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/ascii85/ascii85_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/encoding/ascii85/ascii85_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ascii85

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
    <a id="L21"></a><span class="comment">// Wikipedia example</span>
    <a id="L22"></a>testpair{
        <a id="L23"></a>&#34;Man is distinguished, not only by his reason, but by this singular passion from &#34;
            <a id="L24"></a>&#34;other animals, which is a lust of the mind, that by a perseverance of delight in &#34;
            <a id="L25"></a>&#34;the continued and indefatigable generation of knowledge, exceeds the short &#34;
            <a id="L26"></a>&#34;vehemence of any carnal pleasure.&#34;,
        <a id="L27"></a>&#34;9jqo^BlbD-BleB1DJ+*+F(f,q/0JhKF&lt;GL&gt;Cj@.4Gp$d7F!,L7@&lt;6@)/0JDEF&lt;G%&lt;+EV:2F!,\n&#34;
            <a id="L28"></a>&#34;O&lt;DJ+*.@&lt;*K0@&lt;6L(Df-\\0Ec5e;DffZ(EZee.Bl.9pF\&#34;AGXBPCsi+DGm&gt;@3BB/F*&amp;OCAfu2/AKY\n&#34;
            <a id="L29"></a>&#34;i(DIb:@FD,*)+C]U=@3BN#EcYf8ATD3s@q?d$AftVqCh[NqF&lt;G:8+EV:.+Cf&gt;-FD5W8ARlolDIa\n&#34;
            <a id="L30"></a>&#34;l(DId&lt;j@&lt;?3r@:F%a+D58&#39;ATD4$Bl@l3De:,-DJs`8ARoFb/0JMK@qB4^F!,R&lt;AKZ&amp;-DfTqBG%G\n&#34;
            <a id="L31"></a>&#34;&gt;uD.RTpAKYo&#39;+CT/5+Cei#DII?(E,9)oF*2M7/c\n&#34;,
    <a id="L32"></a>},
<a id="L33"></a>}

<a id="L35"></a>var bigtest = pairs[len(pairs)-1]

<a id="L37"></a>func testEqual(t *testing.T, msg string, args ...) bool {
    <a id="L38"></a>v := reflect.NewValue(args).(*reflect.StructValue);
    <a id="L39"></a>v1 := v.Field(v.NumField() - 2);
    <a id="L40"></a>v2 := v.Field(v.NumField() - 1);
    <a id="L41"></a>if v1.Interface() != v2.Interface() {
        <a id="L42"></a>t.Errorf(msg, args);
        <a id="L43"></a>return false;
    <a id="L44"></a>}
    <a id="L45"></a>return true;
<a id="L46"></a>}

<a id="L48"></a>func strip85(s string) string {
    <a id="L49"></a>t := make([]byte, len(s));
    <a id="L50"></a>w := 0;
    <a id="L51"></a>for r := 0; r &lt; len(s); r++ {
        <a id="L52"></a>c := s[r];
        <a id="L53"></a>if c &gt; &#39; &#39; {
            <a id="L54"></a>t[w] = c;
            <a id="L55"></a>w++;
        <a id="L56"></a>}
    <a id="L57"></a>}
    <a id="L58"></a>return string(t[0:w]);
<a id="L59"></a>}

<a id="L61"></a>func TestEncode(t *testing.T) {
    <a id="L62"></a>for _, p := range pairs {
        <a id="L63"></a>buf := make([]byte, MaxEncodedLen(len(p.decoded)));
        <a id="L64"></a>n := Encode(buf, strings.Bytes(p.decoded));
        <a id="L65"></a>buf = buf[0:n];
        <a id="L66"></a>testEqual(t, &#34;Encode(%q) = %q, want %q&#34;, p.decoded, strip85(string(buf)), strip85(p.encoded));
    <a id="L67"></a>}
<a id="L68"></a>}

<a id="L70"></a>func TestEncoder(t *testing.T) {
    <a id="L71"></a>for _, p := range pairs {
        <a id="L72"></a>bb := &amp;bytes.Buffer{};
        <a id="L73"></a>encoder := NewEncoder(bb);
        <a id="L74"></a>encoder.Write(strings.Bytes(p.decoded));
        <a id="L75"></a>encoder.Close();
        <a id="L76"></a>testEqual(t, &#34;Encode(%q) = %q, want %q&#34;, p.decoded, strip85(bb.String()), strip85(p.encoded));
    <a id="L77"></a>}
<a id="L78"></a>}

<a id="L80"></a>func TestEncoderBuffering(t *testing.T) {
    <a id="L81"></a>input := strings.Bytes(bigtest.decoded);
    <a id="L82"></a>for bs := 1; bs &lt;= 12; bs++ {
        <a id="L83"></a>bb := &amp;bytes.Buffer{};
        <a id="L84"></a>encoder := NewEncoder(bb);
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
        <a id="L96"></a>testEqual(t, &#34;Encoding/%d of %q = %q, want %q&#34;, bs, bigtest.decoded, strip85(bb.String()), strip85(bigtest.encoded));
    <a id="L97"></a>}
<a id="L98"></a>}

<a id="L100"></a>func TestDecode(t *testing.T) {
    <a id="L101"></a>for _, p := range pairs {
        <a id="L102"></a>dbuf := make([]byte, 4*len(p.encoded));
        <a id="L103"></a>ndst, nsrc, err := Decode(dbuf, strings.Bytes(p.encoded), true);
        <a id="L104"></a>testEqual(t, &#34;Decode(%q) = error %v, want %v&#34;, p.encoded, err, os.Error(nil));
        <a id="L105"></a>testEqual(t, &#34;Decode(%q) = nsrc %v, want %v&#34;, p.encoded, nsrc, len(p.encoded));
        <a id="L106"></a>testEqual(t, &#34;Decode(%q) = ndst %v, want %v&#34;, p.encoded, ndst, len(p.decoded));
        <a id="L107"></a>testEqual(t, &#34;Decode(%q) = %q, want %q&#34;, p.encoded, string(dbuf[0:ndst]), p.decoded);
    <a id="L108"></a>}
<a id="L109"></a>}

<a id="L111"></a>func TestDecoder(t *testing.T) {
    <a id="L112"></a>for _, p := range pairs {
        <a id="L113"></a>decoder := NewDecoder(bytes.NewBufferString(p.encoded));
        <a id="L114"></a>dbuf, err := io.ReadAll(decoder);
        <a id="L115"></a>if err != nil {
            <a id="L116"></a>t.Fatal(&#34;Read failed&#34;, err)
        <a id="L117"></a>}
        <a id="L118"></a>testEqual(t, &#34;Read from %q = length %v, want %v&#34;, p.encoded, len(dbuf), len(p.decoded));
        <a id="L119"></a>testEqual(t, &#34;Decoding of %q = %q, want %q&#34;, p.encoded, string(dbuf), p.decoded);
        <a id="L120"></a>if err != nil {
            <a id="L121"></a>testEqual(t, &#34;Read from %q = %v, want %v&#34;, p.encoded, err, os.EOF)
        <a id="L122"></a>}
    <a id="L123"></a>}
<a id="L124"></a>}

<a id="L126"></a>func TestDecoderBuffering(t *testing.T) {
    <a id="L127"></a>for bs := 1; bs &lt;= 12; bs++ {
        <a id="L128"></a>decoder := NewDecoder(bytes.NewBufferString(bigtest.encoded));
        <a id="L129"></a>buf := make([]byte, len(bigtest.decoded)+12);
        <a id="L130"></a>var total int;
        <a id="L131"></a>for total = 0; total &lt; len(bigtest.decoded); {
            <a id="L132"></a>n, err := decoder.Read(buf[total : total+bs]);
            <a id="L133"></a>testEqual(t, &#34;Read from %q at pos %d = %d, %v, want _, %v&#34;, bigtest.encoded, total, n, err, os.Error(nil));
            <a id="L134"></a>total += n;
        <a id="L135"></a>}
        <a id="L136"></a>testEqual(t, &#34;Decoding/%d of %q = %q, want %q&#34;, bs, bigtest.encoded, string(buf[0:total]), bigtest.decoded);
    <a id="L137"></a>}
<a id="L138"></a>}

<a id="L140"></a>func TestDecodeCorrupt(t *testing.T) {
    <a id="L141"></a>type corrupt struct {
        <a id="L142"></a>e   string;
        <a id="L143"></a>p   int;
    <a id="L144"></a>}
    <a id="L145"></a>examples := []corrupt{
        <a id="L146"></a>corrupt{&#34;v&#34;, 0},
        <a id="L147"></a>corrupt{&#34;!z!!!!!!!!!&#34;, 1},
    <a id="L148"></a>};

    <a id="L150"></a>for _, e := range examples {
        <a id="L151"></a>dbuf := make([]byte, 4*len(e.e));
        <a id="L152"></a>_, _, err := Decode(dbuf, strings.Bytes(e.e), true);
        <a id="L153"></a>switch err := err.(type) {
        <a id="L154"></a>case CorruptInputError:
            <a id="L155"></a>testEqual(t, &#34;Corruption in %q at offset %v, want %v&#34;, e.e, int(err), e.p)
        <a id="L156"></a>default:
            <a id="L157"></a>t.Error(&#34;Decoder failed to detect corruption in&#34;, e)
        <a id="L158"></a>}
    <a id="L159"></a>}
<a id="L160"></a>}

<a id="L162"></a>func TestBig(t *testing.T) {
    <a id="L163"></a>n := 3*1000 + 1;
    <a id="L164"></a>raw := make([]byte, n);
    <a id="L165"></a>const alpha = &#34;0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ&#34;;
    <a id="L166"></a>for i := 0; i &lt; n; i++ {
        <a id="L167"></a>raw[i] = alpha[i%len(alpha)]
    <a id="L168"></a>}
    <a id="L169"></a>encoded := new(bytes.Buffer);
    <a id="L170"></a>w := NewEncoder(encoded);
    <a id="L171"></a>nn, err := w.Write(raw);
    <a id="L172"></a>if nn != n || err != nil {
        <a id="L173"></a>t.Fatalf(&#34;Encoder.Write(raw) = %d, %v want %d, nil&#34;, nn, err, n)
    <a id="L174"></a>}
    <a id="L175"></a>err = w.Close();
    <a id="L176"></a>if err != nil {
        <a id="L177"></a>t.Fatalf(&#34;Encoder.Close() = %v want nil&#34;, err)
    <a id="L178"></a>}
    <a id="L179"></a>decoded, err := io.ReadAll(NewDecoder(encoded));
    <a id="L180"></a>if err != nil {
        <a id="L181"></a>t.Fatalf(&#34;io.ReadAll(NewDecoder(...)): %v&#34;, err)
    <a id="L182"></a>}

    <a id="L184"></a>if !bytes.Equal(raw, decoded) {
        <a id="L185"></a>var i int;
        <a id="L186"></a>for i = 0; i &lt; len(decoded) &amp;&amp; i &lt; len(raw); i++ {
            <a id="L187"></a>if decoded[i] != raw[i] {
                <a id="L188"></a>break
            <a id="L189"></a>}
        <a id="L190"></a>}
        <a id="L191"></a>t.Errorf(&#34;Decode(Encode(%d-byte string)) failed at offset %d&#34;, n, i);
    <a id="L192"></a>}
<a id="L193"></a>}
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
