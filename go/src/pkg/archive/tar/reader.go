<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/archive/tar/reader.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/archive/tar/reader.go</h1>

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
<a id="L8"></a><span class="comment">//   - pax extensions</span>

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strconv&#34;;
<a id="L15"></a>)

<a id="L17"></a>var (
    <a id="L18"></a>HeaderError os.Error = os.ErrorString(&#34;invalid tar header&#34;);
<a id="L19"></a>)

<a id="L21"></a><span class="comment">// A Reader provides sequential access to the contents of a tar archive.</span>
<a id="L22"></a><span class="comment">// A tar archive consists of a sequence of files.</span>
<a id="L23"></a><span class="comment">// The Next method advances to the next file in the archive (including the first),</span>
<a id="L24"></a><span class="comment">// and then it can be treated as an io.Reader to access the file&#39;s data.</span>
<a id="L25"></a><span class="comment">//</span>
<a id="L26"></a><span class="comment">// Example:</span>
<a id="L27"></a><span class="comment">//	tr := tar.NewReader(r);</span>
<a id="L28"></a><span class="comment">//	for {</span>
<a id="L29"></a><span class="comment">//		hdr, err := tr.Next();</span>
<a id="L30"></a><span class="comment">//		if err != nil {</span>
<a id="L31"></a><span class="comment">//			// handle error</span>
<a id="L32"></a><span class="comment">//		}</span>
<a id="L33"></a><span class="comment">//		if hdr == nil {</span>
<a id="L34"></a><span class="comment">//			// end of tar archive</span>
<a id="L35"></a><span class="comment">//			break</span>
<a id="L36"></a><span class="comment">//		}</span>
<a id="L37"></a><span class="comment">//		io.Copy(data, tr);</span>
<a id="L38"></a><span class="comment">//	}</span>
<a id="L39"></a>type Reader struct {
    <a id="L40"></a>r   io.Reader;
    <a id="L41"></a>err os.Error;
    <a id="L42"></a>nb  int64; <span class="comment">// number of unread bytes for current file entry</span>
    <a id="L43"></a>pad int64; <span class="comment">// amount of padding (ignored) after current file entry</span>
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// NewReader creates a new Reader reading from r.</span>
<a id="L47"></a>func NewReader(r io.Reader) *Reader { return &amp;Reader{r: r} }

<a id="L49"></a><span class="comment">// Next advances to the next entry in the tar archive.</span>
<a id="L50"></a>func (tr *Reader) Next() (*Header, os.Error) {
    <a id="L51"></a>var hdr *Header;
    <a id="L52"></a>if tr.err == nil {
        <a id="L53"></a>tr.skipUnread()
    <a id="L54"></a>}
    <a id="L55"></a>if tr.err == nil {
        <a id="L56"></a>hdr = tr.readHeader()
    <a id="L57"></a>}
    <a id="L58"></a>return hdr, tr.err;
<a id="L59"></a>}

<a id="L61"></a><span class="comment">// Parse bytes as a NUL-terminated C-style string.</span>
<a id="L62"></a><span class="comment">// If a NUL byte is not found then the whole slice is returned as a string.</span>
<a id="L63"></a>func cString(b []byte) string {
    <a id="L64"></a>n := 0;
    <a id="L65"></a>for n &lt; len(b) &amp;&amp; b[n] != 0 {
        <a id="L66"></a>n++
    <a id="L67"></a>}
    <a id="L68"></a>return string(b[0:n]);
<a id="L69"></a>}

<a id="L71"></a>func (tr *Reader) octal(b []byte) int64 {
    <a id="L72"></a><span class="comment">// Removing leading spaces.</span>
    <a id="L73"></a>for len(b) &gt; 0 &amp;&amp; b[0] == &#39; &#39; {
        <a id="L74"></a>b = b[1:len(b)]
    <a id="L75"></a>}
    <a id="L76"></a><span class="comment">// Removing trailing NULs and spaces.</span>
    <a id="L77"></a>for len(b) &gt; 0 &amp;&amp; (b[len(b)-1] == &#39; &#39; || b[len(b)-1] == &#39;\x00&#39;) {
        <a id="L78"></a>b = b[0 : len(b)-1]
    <a id="L79"></a>}
    <a id="L80"></a>x, err := strconv.Btoui64(cString(b), 8);
    <a id="L81"></a>if err != nil {
        <a id="L82"></a>tr.err = err
    <a id="L83"></a>}
    <a id="L84"></a>return int64(x);
<a id="L85"></a>}

<a id="L87"></a>type ignoreWriter struct{}

<a id="L89"></a>func (ignoreWriter) Write(b []byte) (n int, err os.Error) {
    <a id="L90"></a>return len(b), nil
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// Skip any unread bytes in the existing file entry, as well as any alignment padding.</span>
<a id="L94"></a>func (tr *Reader) skipUnread() {
    <a id="L95"></a>nr := tr.nb + tr.pad; <span class="comment">// number of bytes to skip</span>

    <a id="L97"></a>if sr, ok := tr.r.(io.Seeker); ok {
        <a id="L98"></a>_, tr.err = sr.Seek(nr, 1)
    <a id="L99"></a>} else {
        <a id="L100"></a>_, tr.err = io.Copyn(ignoreWriter{}, tr.r, nr)
    <a id="L101"></a>}
    <a id="L102"></a>tr.nb, tr.pad = 0, 0;
<a id="L103"></a>}

<a id="L105"></a>func (tr *Reader) verifyChecksum(header []byte) bool {
    <a id="L106"></a>if tr.err != nil {
        <a id="L107"></a>return false
    <a id="L108"></a>}

    <a id="L110"></a>given := tr.octal(header[148:156]);
    <a id="L111"></a>unsigned, signed := checksum(header);
    <a id="L112"></a>return given == unsigned || given == signed;
<a id="L113"></a>}

<a id="L115"></a>func (tr *Reader) readHeader() *Header {
    <a id="L116"></a>header := make([]byte, blockSize);
    <a id="L117"></a>if _, tr.err = io.ReadFull(tr.r, header); tr.err != nil {
        <a id="L118"></a>return nil
    <a id="L119"></a>}

    <a id="L121"></a><span class="comment">// Two blocks of zero bytes marks the end of the archive.</span>
    <a id="L122"></a>if bytes.Equal(header, zeroBlock[0:blockSize]) {
        <a id="L123"></a>if _, tr.err = io.ReadFull(tr.r, header); tr.err != nil {
            <a id="L124"></a>return nil
        <a id="L125"></a>}
        <a id="L126"></a>if !bytes.Equal(header, zeroBlock[0:blockSize]) {
            <a id="L127"></a>tr.err = HeaderError
        <a id="L128"></a>}
        <a id="L129"></a>return nil;
    <a id="L130"></a>}

    <a id="L132"></a>if !tr.verifyChecksum(header) {
        <a id="L133"></a>tr.err = HeaderError;
        <a id="L134"></a>return nil;
    <a id="L135"></a>}

    <a id="L137"></a><span class="comment">// Unpack</span>
    <a id="L138"></a>hdr := new(Header);
    <a id="L139"></a>s := slicer(header);

    <a id="L141"></a>hdr.Name = cString(s.next(100));
    <a id="L142"></a>hdr.Mode = tr.octal(s.next(8));
    <a id="L143"></a>hdr.Uid = tr.octal(s.next(8));
    <a id="L144"></a>hdr.Gid = tr.octal(s.next(8));
    <a id="L145"></a>hdr.Size = tr.octal(s.next(12));
    <a id="L146"></a>hdr.Mtime = tr.octal(s.next(12));
    <a id="L147"></a>s.next(8); <span class="comment">// chksum</span>
    <a id="L148"></a>hdr.Typeflag = s.next(1)[0];
    <a id="L149"></a>hdr.Linkname = cString(s.next(100));

    <a id="L151"></a><span class="comment">// The remainder of the header depends on the value of magic.</span>
    <a id="L152"></a><span class="comment">// The original (v7) version of tar had no explicit magic field,</span>
    <a id="L153"></a><span class="comment">// so its magic bytes, like the rest of the block, are NULs.</span>
    <a id="L154"></a>magic := string(s.next(8)); <span class="comment">// contains version field as well.</span>
    <a id="L155"></a>var format string;
    <a id="L156"></a>switch magic {
    <a id="L157"></a>case &#34;ustar\x0000&#34;: <span class="comment">// POSIX tar (1003.1-1988)</span>
        <a id="L158"></a>if string(header[508:512]) == &#34;tar\x00&#34; {
            <a id="L159"></a>format = &#34;star&#34;
        <a id="L160"></a>} else {
            <a id="L161"></a>format = &#34;posix&#34;
        <a id="L162"></a>}
    <a id="L163"></a>case &#34;ustar  \x00&#34;: <span class="comment">// old GNU tar</span>
        <a id="L164"></a>format = &#34;gnu&#34;
    <a id="L165"></a>}

    <a id="L167"></a>switch format {
    <a id="L168"></a>case &#34;posix&#34;, &#34;gnu&#34;, &#34;star&#34;:
        <a id="L169"></a>hdr.Uname = cString(s.next(32));
        <a id="L170"></a>hdr.Gname = cString(s.next(32));
        <a id="L171"></a>devmajor := s.next(8);
        <a id="L172"></a>devminor := s.next(8);
        <a id="L173"></a>if hdr.Typeflag == TypeChar || hdr.Typeflag == TypeBlock {
            <a id="L174"></a>hdr.Devmajor = tr.octal(devmajor);
            <a id="L175"></a>hdr.Devminor = tr.octal(devminor);
        <a id="L176"></a>}
        <a id="L177"></a>var prefix string;
        <a id="L178"></a>switch format {
        <a id="L179"></a>case &#34;posix&#34;, &#34;gnu&#34;:
            <a id="L180"></a>prefix = cString(s.next(155))
        <a id="L181"></a>case &#34;star&#34;:
            <a id="L182"></a>prefix = cString(s.next(131));
            <a id="L183"></a>hdr.Atime = tr.octal(s.next(12));
            <a id="L184"></a>hdr.Ctime = tr.octal(s.next(12));
        <a id="L185"></a>}
        <a id="L186"></a>if len(prefix) &gt; 0 {
            <a id="L187"></a>hdr.Name = prefix + &#34;/&#34; + hdr.Name
        <a id="L188"></a>}
    <a id="L189"></a>}

    <a id="L191"></a>if tr.err != nil {
        <a id="L192"></a>tr.err = HeaderError;
        <a id="L193"></a>return nil;
    <a id="L194"></a>}

    <a id="L196"></a><span class="comment">// Maximum value of hdr.Size is 64 GB (12 octal digits),</span>
    <a id="L197"></a><span class="comment">// so there&#39;s no risk of int64 overflowing.</span>
    <a id="L198"></a>tr.nb = int64(hdr.Size);
    <a id="L199"></a>tr.pad = -tr.nb &amp; (blockSize - 1); <span class="comment">// blockSize is a power of two</span>

    <a id="L201"></a>return hdr;
<a id="L202"></a>}

<a id="L204"></a><span class="comment">// Read reads from the current entry in the tar archive.</span>
<a id="L205"></a><span class="comment">// It returns 0, nil when it reaches the end of that entry,</span>
<a id="L206"></a><span class="comment">// until Next is called to advance to the next entry.</span>
<a id="L207"></a>func (tr *Reader) Read(b []uint8) (n int, err os.Error) {
    <a id="L208"></a>if int64(len(b)) &gt; tr.nb {
        <a id="L209"></a>b = b[0:tr.nb]
    <a id="L210"></a>}
    <a id="L211"></a>n, err = tr.r.Read(b);
    <a id="L212"></a>tr.nb -= int64(n);
    <a id="L213"></a>tr.err = err;
    <a id="L214"></a>return;
<a id="L215"></a>}
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
