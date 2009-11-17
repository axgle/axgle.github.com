<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/archive/tar/writer.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/archive/tar/writer.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tar

<a id="L7"></a><span class="comment">// TODO(dsymonds):</span>
<a id="L8"></a><span class="comment">// - catch more errors (no first header, write after close, etc.)</span>

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strconv&#34;;
    <a id="L15"></a>&#34;strings&#34;;
<a id="L16"></a>)

<a id="L18"></a>var (
    <a id="L19"></a>ErrWriteTooLong = os.NewError(&#34;write too long&#34;);
    <a id="L20"></a>ErrFieldTooLong = os.NewError(&#34;header field too long&#34;);
<a id="L21"></a>)

<a id="L23"></a><span class="comment">// A Writer provides sequential writing of a tar archive in POSIX.1 format.</span>
<a id="L24"></a><span class="comment">// A tar archive consists of a sequence of files.</span>
<a id="L25"></a><span class="comment">// Call WriteHeader to begin a new file, and then call Write to supply that file&#39;s data,</span>
<a id="L26"></a><span class="comment">// writing at most hdr.Size bytes in total.</span>
<a id="L27"></a><span class="comment">//</span>
<a id="L28"></a><span class="comment">// Example:</span>
<a id="L29"></a><span class="comment">//	tw := tar.NewWriter(w);</span>
<a id="L30"></a><span class="comment">//	hdr := new(Header);</span>
<a id="L31"></a><span class="comment">//	hdr.Size = length of data in bytes;</span>
<a id="L32"></a><span class="comment">//	// populate other hdr fields as desired</span>
<a id="L33"></a><span class="comment">//	if err := tw.WriteHeader(hdr); err != nil {</span>
<a id="L34"></a><span class="comment">//		// handle error</span>
<a id="L35"></a><span class="comment">//	}</span>
<a id="L36"></a><span class="comment">//	io.Copy(tw, data);</span>
<a id="L37"></a><span class="comment">//	tw.Close();</span>
<a id="L38"></a>type Writer struct {
    <a id="L39"></a>w          io.Writer;
    <a id="L40"></a>err        os.Error;
    <a id="L41"></a>nb         int64; <span class="comment">// number of unwritten bytes for current file entry</span>
    <a id="L42"></a>pad        int64; <span class="comment">// amount of padding to write after current file entry</span>
    <a id="L43"></a>closed     bool;
    <a id="L44"></a>usedBinary bool; <span class="comment">// whether the binary numeric field extension was used</span>
<a id="L45"></a>}

<a id="L47"></a><span class="comment">// NewWriter creates a new Writer writing to w.</span>
<a id="L48"></a>func NewWriter(w io.Writer) *Writer { return &amp;Writer{w: w} }

<a id="L50"></a><span class="comment">// Flush finishes writing the current file (optional).</span>
<a id="L51"></a>func (tw *Writer) Flush() os.Error {
    <a id="L52"></a>n := tw.nb + tw.pad;
    <a id="L53"></a>for n &gt; 0 &amp;&amp; tw.err == nil {
        <a id="L54"></a>nr := n;
        <a id="L55"></a>if nr &gt; blockSize {
            <a id="L56"></a>nr = blockSize
        <a id="L57"></a>}
        <a id="L58"></a>var nw int;
        <a id="L59"></a>nw, tw.err = tw.w.Write(zeroBlock[0:nr]);
        <a id="L60"></a>n -= int64(nw);
    <a id="L61"></a>}
    <a id="L62"></a>tw.nb = 0;
    <a id="L63"></a>tw.pad = 0;
    <a id="L64"></a>return tw.err;
<a id="L65"></a>}

<a id="L67"></a><span class="comment">// Write s into b, terminating it with a NUL if there is room.</span>
<a id="L68"></a>func (tw *Writer) cString(b []byte, s string) {
    <a id="L69"></a>if len(s) &gt; len(b) {
        <a id="L70"></a>if tw.err == nil {
            <a id="L71"></a>tw.err = ErrFieldTooLong
        <a id="L72"></a>}
        <a id="L73"></a>return;
    <a id="L74"></a>}
    <a id="L75"></a>for i, ch := range strings.Bytes(s) {
        <a id="L76"></a>b[i] = ch
    <a id="L77"></a>}
    <a id="L78"></a>if len(s) &lt; len(b) {
        <a id="L79"></a>b[len(s)] = 0
    <a id="L80"></a>}
<a id="L81"></a>}

<a id="L83"></a><span class="comment">// Encode x as an octal ASCII string and write it into b with leading zeros.</span>
<a id="L84"></a>func (tw *Writer) octal(b []byte, x int64) {
    <a id="L85"></a>s := strconv.Itob64(x, 8);
    <a id="L86"></a><span class="comment">// leading zeros, but leave room for a NUL.</span>
    <a id="L87"></a>for len(s)+1 &lt; len(b) {
        <a id="L88"></a>s = &#34;0&#34; + s
    <a id="L89"></a>}
    <a id="L90"></a>tw.cString(b, s);
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// Write x into b, either as octal or as binary (GNUtar/star extension).</span>
<a id="L94"></a>func (tw *Writer) numeric(b []byte, x int64) {
    <a id="L95"></a><span class="comment">// Try octal first.</span>
    <a id="L96"></a>s := strconv.Itob64(x, 8);
    <a id="L97"></a>if len(s) &lt; len(b) {
        <a id="L98"></a>tw.octal(b, x);
        <a id="L99"></a>return;
    <a id="L100"></a>}
    <a id="L101"></a><span class="comment">// Too big: use binary (big-endian).</span>
    <a id="L102"></a>tw.usedBinary = true;
    <a id="L103"></a>for i := len(b) - 1; x &gt; 0 &amp;&amp; i &gt;= 0; i-- {
        <a id="L104"></a>b[i] = byte(x);
        <a id="L105"></a>x &gt;&gt;= 8;
    <a id="L106"></a>}
    <a id="L107"></a>b[0] |= 0x80; <span class="comment">// highest bit indicates binary format</span>
<a id="L108"></a>}

<a id="L110"></a><span class="comment">// WriteHeader writes hdr and prepares to accept the file&#39;s contents.</span>
<a id="L111"></a><span class="comment">// WriteHeader calls Flush if it is not the first header.</span>
<a id="L112"></a>func (tw *Writer) WriteHeader(hdr *Header) os.Error {
    <a id="L113"></a>if tw.err == nil {
        <a id="L114"></a>tw.Flush()
    <a id="L115"></a>}
    <a id="L116"></a>if tw.err != nil {
        <a id="L117"></a>return tw.err
    <a id="L118"></a>}

    <a id="L120"></a>tw.nb = int64(hdr.Size);
    <a id="L121"></a>tw.pad = -tw.nb &amp; (blockSize - 1); <span class="comment">// blockSize is a power of two</span>

    <a id="L123"></a>header := make([]byte, blockSize);
    <a id="L124"></a>s := slicer(header);

    <a id="L126"></a><span class="comment">// TODO(dsymonds): handle names longer than 100 chars</span>
    <a id="L127"></a>bytes.Copy(s.next(100), strings.Bytes(hdr.Name));

    <a id="L129"></a>tw.octal(s.next(8), hdr.Mode);                               <span class="comment">// 100:108</span>
    <a id="L130"></a>tw.numeric(s.next(8), hdr.Uid);                              <span class="comment">// 108:116</span>
    <a id="L131"></a>tw.numeric(s.next(8), hdr.Gid);                              <span class="comment">// 116:124</span>
    <a id="L132"></a>tw.numeric(s.next(12), hdr.Size);                            <span class="comment">// 124:136</span>
    <a id="L133"></a>tw.numeric(s.next(12), hdr.Mtime);                           <span class="comment">// 136:148</span>
    <a id="L134"></a>s.next(8);                                                   <span class="comment">// chksum (148:156)</span>
    <a id="L135"></a>s.next(1)[0] = hdr.Typeflag;                                 <span class="comment">// 156:157</span>
    <a id="L136"></a>s.next(100);                                                 <span class="comment">// linkname (157:257)</span>
    <a id="L137"></a>bytes.Copy(s.next(8), strings.Bytes(&#34;ustar\x0000&#34;)); <span class="comment">// 257:265</span>
    <a id="L138"></a>tw.cString(s.next(32), hdr.Uname);                           <span class="comment">// 265:297</span>
    <a id="L139"></a>tw.cString(s.next(32), hdr.Gname);                           <span class="comment">// 297:329</span>
    <a id="L140"></a>tw.numeric(s.next(8), hdr.Devmajor);                         <span class="comment">// 329:337</span>
    <a id="L141"></a>tw.numeric(s.next(8), hdr.Devminor);                         <span class="comment">// 337:345</span>

    <a id="L143"></a><span class="comment">// Use the GNU magic instead of POSIX magic if we used any GNU extensions.</span>
    <a id="L144"></a>if tw.usedBinary {
        <a id="L145"></a>bytes.Copy(header[257:265], strings.Bytes(&#34;ustar  \x00&#34;))
    <a id="L146"></a>}

    <a id="L148"></a><span class="comment">// The chksum field is terminated by a NUL and a space.</span>
    <a id="L149"></a><span class="comment">// This is different from the other octal fields.</span>
    <a id="L150"></a>chksum, _ := checksum(header);
    <a id="L151"></a>tw.octal(header[148:155], chksum);
    <a id="L152"></a>header[155] = &#39; &#39;;

    <a id="L154"></a>if tw.err != nil {
        <a id="L155"></a><span class="comment">// problem with header; probably integer too big for a field.</span>
        <a id="L156"></a>return tw.err
    <a id="L157"></a>}

    <a id="L159"></a>_, tw.err = tw.w.Write(header);

    <a id="L161"></a>return tw.err;
<a id="L162"></a>}

<a id="L164"></a><span class="comment">// Write writes to the current entry in the tar archive.</span>
<a id="L165"></a><span class="comment">// Write returns the error ErrWriteTooLong if more than</span>
<a id="L166"></a><span class="comment">// hdr.Size bytes are written after WriteHeader.</span>
<a id="L167"></a>func (tw *Writer) Write(b []uint8) (n int, err os.Error) {
    <a id="L168"></a>overwrite := false;
    <a id="L169"></a>if int64(len(b)) &gt; tw.nb {
        <a id="L170"></a>b = b[0:tw.nb];
        <a id="L171"></a>overwrite = true;
    <a id="L172"></a>}
    <a id="L173"></a>n, err = tw.w.Write(b);
    <a id="L174"></a>tw.nb -= int64(n);
    <a id="L175"></a>if err == nil &amp;&amp; overwrite {
        <a id="L176"></a>err = ErrWriteTooLong
    <a id="L177"></a>}
    <a id="L178"></a>tw.err = err;
    <a id="L179"></a>return;
<a id="L180"></a>}

<a id="L182"></a>func (tw *Writer) Close() os.Error {
    <a id="L183"></a>if tw.err != nil || tw.closed {
        <a id="L184"></a>return tw.err
    <a id="L185"></a>}
    <a id="L186"></a>tw.Flush();
    <a id="L187"></a>tw.closed = true;

    <a id="L189"></a><span class="comment">// trailer: two zero blocks</span>
    <a id="L190"></a>for i := 0; i &lt; 2; i++ {
        <a id="L191"></a>_, tw.err = tw.w.Write(zeroBlock);
        <a id="L192"></a>if tw.err != nil {
            <a id="L193"></a>break
        <a id="L194"></a>}
    <a id="L195"></a>}
    <a id="L196"></a>return tw.err;
<a id="L197"></a>}
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
