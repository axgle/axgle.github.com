<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/nacl/srpc/msg.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../../doc/style.css">
  <script type="text/javascript" src="../../../../../doc/godocs.js"></script>

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
        <a href="../../../../../index.html"><img src="../../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/nacl/srpc/msg.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// SRPC constants, data structures, and parsing.</span>

<a id="L7"></a>package srpc

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bytes&#34;;
    <a id="L11"></a>&#34;math&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
    <a id="L14"></a>&#34;syscall&#34;;
    <a id="L15"></a>&#34;unsafe&#34;;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// An Errno is an SRPC status code.</span>
<a id="L19"></a>type Errno uint32

<a id="L21"></a>const (
    <a id="L22"></a>OK  Errno = 256 + iota;
    <a id="L23"></a>ErrBreak;
    <a id="L24"></a>ErrMessageTruncated;
    <a id="L25"></a>ErrNoMemory;
    <a id="L26"></a>ErrProtocolMismatch;
    <a id="L27"></a>ErrBadRPCNumber;
    <a id="L28"></a>ErrBadArgType;
    <a id="L29"></a>ErrTooFewArgs;
    <a id="L30"></a>ErrTooManyArgs;
    <a id="L31"></a>ErrInArgTypeMismatch;
    <a id="L32"></a>ErrOutArgTypeMismatch;
    <a id="L33"></a>ErrInternalError;
    <a id="L34"></a>ErrAppError;
<a id="L35"></a>)

<a id="L37"></a>var errstr = [...]string{
    <a id="L38"></a>OK - OK: &#34;ok&#34;,
    <a id="L39"></a>ErrBreak - OK: &#34;break&#34;,
    <a id="L40"></a>ErrMessageTruncated - OK: &#34;message truncated&#34;,
    <a id="L41"></a>ErrNoMemory - OK: &#34;out of memory&#34;,
    <a id="L42"></a>ErrProtocolMismatch - OK: &#34;protocol mismatch&#34;,
    <a id="L43"></a>ErrBadRPCNumber - OK: &#34;invalid RPC method number&#34;,
    <a id="L44"></a>ErrBadArgType - OK: &#34;unexpected argument type&#34;,
    <a id="L45"></a>ErrTooFewArgs - OK: &#34;too few arguments&#34;,
    <a id="L46"></a>ErrTooManyArgs - OK: &#34;too many arguments&#34;,
    <a id="L47"></a>ErrInArgTypeMismatch - OK: &#34;input argument type mismatch&#34;,
    <a id="L48"></a>ErrOutArgTypeMismatch - OK: &#34;output argument type mismatch&#34;,
    <a id="L49"></a>ErrInternalError - OK: &#34;internal error&#34;,
    <a id="L50"></a>ErrAppError - OK: &#34;application error&#34;,
<a id="L51"></a>}

<a id="L53"></a>func (e Errno) String() string {
    <a id="L54"></a>if e &lt; OK || int(e-OK) &gt;= len(errstr) {
        <a id="L55"></a>return &#34;Errno(&#34; + strconv.Itoa64(int64(e)) + &#34;)&#34;
    <a id="L56"></a>}
    <a id="L57"></a>return errstr[e-OK];
<a id="L58"></a>}

<a id="L60"></a><span class="comment">// A *msgHdr is the data argument to the imc_recvmsg</span>
<a id="L61"></a><span class="comment">// and imc_sendmsg system calls.  Because it contains unchecked</span>
<a id="L62"></a><span class="comment">// counts trusted by the system calls, the data structure is unsafe</span>
<a id="L63"></a><span class="comment">// to expose to package clients.</span>
<a id="L64"></a>type msgHdr struct {
    <a id="L65"></a>iov   *iov;
    <a id="L66"></a>niov  int32;
    <a id="L67"></a>desc  *int32;
    <a id="L68"></a>ndesc int32;
    <a id="L69"></a>flags uint32;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// A single region for I/O.  Just as unsafe as msgHdr.</span>
<a id="L73"></a>type iov struct {
    <a id="L74"></a>base *byte;
    <a id="L75"></a>len  int32;
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// A msg is the Go representation of a message.</span>
<a id="L79"></a>type msg struct {
    <a id="L80"></a>rdata []byte;  <span class="comment">// data being consumed during message parsing</span>
    <a id="L81"></a>rdesc []int32; <span class="comment">// file descriptors being consumed during message parsing</span>
    <a id="L82"></a>wdata []byte;  <span class="comment">// data being generated when replying</span>

    <a id="L84"></a><span class="comment">// parsed version of message</span>
    <a id="L85"></a>protocol  uint32;
    <a id="L86"></a>requestId uint64;
    <a id="L87"></a>isReq     bool;
    <a id="L88"></a>rpcNumber uint32;
    <a id="L89"></a>gotHeader bool;
    <a id="L90"></a>status    Errno;         <span class="comment">// error code sent in response</span>
    <a id="L91"></a>Arg       []interface{}; <span class="comment">// method arguments</span>
    <a id="L92"></a>Ret       []interface{}; <span class="comment">// method results</span>
    <a id="L93"></a>Size      []int;         <span class="comment">// max sizes for arrays in method results</span>
    <a id="L94"></a>fmt       string;        <span class="comment">// accumulated format string of arg+&#34;:&#34;+ret</span>
<a id="L95"></a>}

<a id="L97"></a><span class="comment">// A msgReceiver receives messages from a file descriptor.</span>
<a id="L98"></a>type msgReceiver struct {
    <a id="L99"></a>fd   int;
    <a id="L100"></a>data [128 * 1024]byte;
    <a id="L101"></a>desc [8]int32;
    <a id="L102"></a>hdr  msgHdr;
    <a id="L103"></a>iov  iov;
<a id="L104"></a>}

<a id="L106"></a>func (r *msgReceiver) recv() (*msg, os.Error) {
    <a id="L107"></a><span class="comment">// Init pointers to buffers where syscall recvmsg can write.</span>
    <a id="L108"></a>r.iov.base = &amp;r.data[0];
    <a id="L109"></a>r.iov.len = int32(len(r.data));
    <a id="L110"></a>r.hdr.iov = &amp;r.iov;
    <a id="L111"></a>r.hdr.niov = 1;
    <a id="L112"></a>r.hdr.desc = &amp;r.desc[0];
    <a id="L113"></a>r.hdr.ndesc = int32(len(r.desc));
    <a id="L114"></a>n, _, e := syscall.Syscall(syscall.SYS_IMC_RECVMSG, uintptr(r.fd), uintptr(unsafe.Pointer(&amp;r.hdr)), 0);
    <a id="L115"></a>if e != 0 {
        <a id="L116"></a>return nil, os.NewSyscallError(&#34;imc_recvmsg&#34;, int(e))
    <a id="L117"></a>}

    <a id="L119"></a><span class="comment">// Make a copy of the data so that the next recvmsg doesn&#39;t</span>
    <a id="L120"></a><span class="comment">// smash it.  The system call did not update r.iov.len.  Instead it</span>
    <a id="L121"></a><span class="comment">// returned the total byte count as n.</span>
    <a id="L122"></a>m := new(msg);
    <a id="L123"></a>m.rdata = make([]byte, n);
    <a id="L124"></a>bytes.Copy(m.rdata, &amp;r.data);

    <a id="L126"></a><span class="comment">// Make a copy of the desc too.</span>
    <a id="L127"></a><span class="comment">// The system call *did* update r.hdr.ndesc.</span>
    <a id="L128"></a>if r.hdr.ndesc &gt; 0 {
        <a id="L129"></a>m.rdesc = make([]int32, r.hdr.ndesc);
        <a id="L130"></a>for i := range m.rdesc {
            <a id="L131"></a>m.rdesc[i] = r.desc[i]
        <a id="L132"></a>}
    <a id="L133"></a>}

    <a id="L135"></a>return m, nil;
<a id="L136"></a>}

<a id="L138"></a><span class="comment">// A msgSender sends messages on a file descriptor.</span>
<a id="L139"></a>type msgSender struct {
    <a id="L140"></a>fd  int;
    <a id="L141"></a>hdr msgHdr;
    <a id="L142"></a>iov iov;
<a id="L143"></a>}

<a id="L145"></a>func (s *msgSender) send(m *msg) os.Error {
    <a id="L146"></a>if len(m.wdata) &gt; 0 {
        <a id="L147"></a>s.iov.base = &amp;m.wdata[0]
    <a id="L148"></a>}
    <a id="L149"></a>s.iov.len = int32(len(m.wdata));
    <a id="L150"></a>s.hdr.iov = &amp;s.iov;
    <a id="L151"></a>s.hdr.niov = 1;
    <a id="L152"></a>s.hdr.desc = nil;
    <a id="L153"></a>s.hdr.ndesc = 0;
    <a id="L154"></a>_, _, e := syscall.Syscall(syscall.SYS_IMC_SENDMSG, uintptr(s.fd), uintptr(unsafe.Pointer(&amp;s.hdr)), 0);
    <a id="L155"></a>if e != 0 {
        <a id="L156"></a>return os.NewSyscallError(&#34;imc_sendmsg&#34;, int(e))
    <a id="L157"></a>}
    <a id="L158"></a>return nil;
<a id="L159"></a>}

<a id="L161"></a><span class="comment">// Reading from msg.rdata.</span>
<a id="L162"></a>func (m *msg) uint8() uint8 {
    <a id="L163"></a>if m.status != OK {
        <a id="L164"></a>return 0
    <a id="L165"></a>}
    <a id="L166"></a>if len(m.rdata) &lt; 1 {
        <a id="L167"></a>m.status = ErrMessageTruncated;
        <a id="L168"></a>return 0;
    <a id="L169"></a>}
    <a id="L170"></a>x := m.rdata[0];
    <a id="L171"></a>m.rdata = m.rdata[1:len(m.rdata)];
    <a id="L172"></a>return x;
<a id="L173"></a>}

<a id="L175"></a>func (m *msg) uint32() uint32 {
    <a id="L176"></a>if m.status != OK {
        <a id="L177"></a>return 0
    <a id="L178"></a>}
    <a id="L179"></a>if len(m.rdata) &lt; 4 {
        <a id="L180"></a>m.status = ErrMessageTruncated;
        <a id="L181"></a>return 0;
    <a id="L182"></a>}
    <a id="L183"></a>b := m.rdata[0:4];
    <a id="L184"></a>x := uint32(b[0]) | uint32(b[1])&lt;&lt;8 | uint32(b[2])&lt;&lt;16 | uint32(b[3])&lt;&lt;24;
    <a id="L185"></a>m.rdata = m.rdata[4:len(m.rdata)];
    <a id="L186"></a>return x;
<a id="L187"></a>}

<a id="L189"></a>func (m *msg) uint64() uint64 {
    <a id="L190"></a>if m.status != OK {
        <a id="L191"></a>return 0
    <a id="L192"></a>}
    <a id="L193"></a>if len(m.rdata) &lt; 8 {
        <a id="L194"></a>m.status = ErrMessageTruncated;
        <a id="L195"></a>return 0;
    <a id="L196"></a>}
    <a id="L197"></a>b := m.rdata[0:8];
    <a id="L198"></a>x := uint64(uint32(b[0]) | uint32(b[1])&lt;&lt;8 | uint32(b[2])&lt;&lt;16 | uint32(b[3])&lt;&lt;24);
    <a id="L199"></a>x |= uint64(uint32(b[4])|uint32(b[5])&lt;&lt;8|uint32(b[6])&lt;&lt;16|uint32(b[7])&lt;&lt;24) &lt;&lt; 32;
    <a id="L200"></a>m.rdata = m.rdata[8:len(m.rdata)];
    <a id="L201"></a>return x;
<a id="L202"></a>}

<a id="L204"></a>func (m *msg) bytes(n int) []byte {
    <a id="L205"></a>if m.status != OK {
        <a id="L206"></a>return nil
    <a id="L207"></a>}
    <a id="L208"></a>if len(m.rdata) &lt; n {
        <a id="L209"></a>m.status = ErrMessageTruncated;
        <a id="L210"></a>return nil;
    <a id="L211"></a>}
    <a id="L212"></a>x := m.rdata[0:n];
    <a id="L213"></a>m.rdata = m.rdata[n:len(m.rdata)];
    <a id="L214"></a>return x;
<a id="L215"></a>}

<a id="L217"></a><span class="comment">// Writing to msg.wdata.</span>
<a id="L218"></a>func (m *msg) grow(n int) []byte {
    <a id="L219"></a>i := len(m.wdata);
    <a id="L220"></a>if i+n &gt; cap(m.wdata) {
        <a id="L221"></a>a := make([]byte, i, (i+n)*2);
        <a id="L222"></a>bytes.Copy(a, m.wdata);
        <a id="L223"></a>m.wdata = a;
    <a id="L224"></a>}
    <a id="L225"></a>m.wdata = m.wdata[0 : i+n];
    <a id="L226"></a>return m.wdata[i : i+n];
<a id="L227"></a>}

<a id="L229"></a>func (m *msg) wuint8(x uint8) { m.grow(1)[0] = x }

<a id="L231"></a>func (m *msg) wuint32(x uint32) {
    <a id="L232"></a>b := m.grow(4);
    <a id="L233"></a>b[0] = byte(x);
    <a id="L234"></a>b[1] = byte(x &gt;&gt; 8);
    <a id="L235"></a>b[2] = byte(x &gt;&gt; 16);
    <a id="L236"></a>b[3] = byte(x &gt;&gt; 24);
<a id="L237"></a>}

<a id="L239"></a>func (m *msg) wuint64(x uint64) {
    <a id="L240"></a>b := m.grow(8);
    <a id="L241"></a>lo := uint32(x);
    <a id="L242"></a>b[0] = byte(lo);
    <a id="L243"></a>b[1] = byte(lo &gt;&gt; 8);
    <a id="L244"></a>b[2] = byte(lo &gt;&gt; 16);
    <a id="L245"></a>b[3] = byte(lo &gt;&gt; 24);
    <a id="L246"></a>hi := uint32(x &gt;&gt; 32);
    <a id="L247"></a>b[4] = byte(hi);
    <a id="L248"></a>b[5] = byte(hi &gt;&gt; 8);
    <a id="L249"></a>b[6] = byte(hi &gt;&gt; 16);
    <a id="L250"></a>b[7] = byte(hi &gt;&gt; 24);
<a id="L251"></a>}

<a id="L253"></a>func (m *msg) wbytes(p []byte) { bytes.Copy(m.grow(len(p)), p) }

<a id="L255"></a>func (m *msg) wstring(s string) {
    <a id="L256"></a>b := m.grow(len(s));
    <a id="L257"></a>for i := range b {
        <a id="L258"></a>b[i] = s[i]
    <a id="L259"></a>}
<a id="L260"></a>}

<a id="L262"></a><span class="comment">// Parsing of RPC header and arguments.</span>
<a id="L263"></a><span class="comment">//</span>
<a id="L264"></a><span class="comment">// The header format is:</span>
<a id="L265"></a><span class="comment">//	protocol uint32;</span>
<a id="L266"></a><span class="comment">//	requestId uint64;</span>
<a id="L267"></a><span class="comment">//	isReq bool;</span>
<a id="L268"></a><span class="comment">//	rpcNumber uint32;</span>
<a id="L269"></a><span class="comment">//	status uint32;  // only for response</span>
<a id="L270"></a><span class="comment">//</span>
<a id="L271"></a><span class="comment">// Then a sequence of values follow, preceded by the length:</span>
<a id="L272"></a><span class="comment">//	nvalue uint32;</span>
<a id="L273"></a><span class="comment">//</span>
<a id="L274"></a><span class="comment">// Each value begins with a one-byte type followed by</span>
<a id="L275"></a><span class="comment">// type-specific data.</span>
<a id="L276"></a><span class="comment">//</span>
<a id="L277"></a><span class="comment">//	type uint8;</span>
<a id="L278"></a><span class="comment">//	&#39;b&#39;:	x bool;</span>
<a id="L279"></a><span class="comment">//	&#39;C&#39;:	len uint32; x [len]byte;</span>
<a id="L280"></a><span class="comment">//	&#39;d&#39;:	x float64;</span>
<a id="L281"></a><span class="comment">//	&#39;D&#39;:	len uint32; x [len]float64;</span>
<a id="L282"></a><span class="comment">//	&#39;h&#39;:	x int;	// handle aka file descriptor</span>
<a id="L283"></a><span class="comment">//	&#39;i&#39;:	x int32;</span>
<a id="L284"></a><span class="comment">//	&#39;I&#39;:	len uint32; x [len]int32;</span>
<a id="L285"></a><span class="comment">//	&#39;s&#39;:	len uint32; x [len]byte;</span>
<a id="L286"></a><span class="comment">//</span>
<a id="L287"></a><span class="comment">// If this is a request, a sequence of pseudo-values follows,</span>
<a id="L288"></a><span class="comment">// preceded by its length (nvalue uint32).</span>
<a id="L289"></a><span class="comment">//</span>
<a id="L290"></a><span class="comment">// Each pseudo-value is a one-byte type as above,</span>
<a id="L291"></a><span class="comment">// followed by a maximum length (len uint32)</span>
<a id="L292"></a><span class="comment">// for the &#39;C&#39;, &#39;D&#39;, &#39;I&#39;, and &#39;s&#39; types.</span>
<a id="L293"></a><span class="comment">//</span>
<a id="L294"></a><span class="comment">// In the Go msg, we represent each argument by</span>
<a id="L295"></a><span class="comment">// an empty interface containing the type of x in the</span>
<a id="L296"></a><span class="comment">// corresponding case.</span>

<a id="L298"></a><span class="comment">// The current protocol number.</span>
<a id="L299"></a>const protocol = 0xc0da0002

<a id="L301"></a>func (m *msg) unpackHeader() {
    <a id="L302"></a>m.protocol = m.uint32();
    <a id="L303"></a>m.requestId = m.uint64();
    <a id="L304"></a>m.isReq = m.uint8() != 0;
    <a id="L305"></a>m.rpcNumber = m.uint32();
    <a id="L306"></a>m.gotHeader = m.status == OK; <span class="comment">// signal that header parsed successfully</span>
    <a id="L307"></a>if m.gotHeader &amp;&amp; !m.isReq {
        <a id="L308"></a>status := Errno(m.uint32());
        <a id="L309"></a>m.gotHeader = m.status == OK; <span class="comment">// still ok?</span>
        <a id="L310"></a>if m.gotHeader {
            <a id="L311"></a>m.status = status
        <a id="L312"></a>}
    <a id="L313"></a>}
<a id="L314"></a>}

<a id="L316"></a>func (m *msg) packHeader() {
    <a id="L317"></a>m.wuint32(m.protocol);
    <a id="L318"></a>m.wuint64(m.requestId);
    <a id="L319"></a>if m.isReq {
        <a id="L320"></a>m.wuint8(1)
    <a id="L321"></a>} else {
        <a id="L322"></a>m.wuint8(0)
    <a id="L323"></a>}
    <a id="L324"></a>m.wuint32(m.rpcNumber);
    <a id="L325"></a>if !m.isReq {
        <a id="L326"></a>m.wuint32(uint32(m.status))
    <a id="L327"></a>}
<a id="L328"></a>}

<a id="L330"></a>func (m *msg) unpackValues(v []interface{}) {
    <a id="L331"></a>for i := range v {
        <a id="L332"></a>t := m.uint8();
        <a id="L333"></a>m.fmt += string(t);
        <a id="L334"></a>switch t {
        <a id="L335"></a>default:
            <a id="L336"></a>if m.status == OK {
                <a id="L337"></a>m.status = ErrBadArgType
            <a id="L338"></a>}
            <a id="L339"></a>return;
        <a id="L340"></a>case &#39;b&#39;: <span class="comment">// bool[1]</span>
            <a id="L341"></a>v[i] = m.uint8() &gt; 0
        <a id="L342"></a>case &#39;C&#39;: <span class="comment">// char array</span>
            <a id="L343"></a>v[i] = m.bytes(int(m.uint32()))
        <a id="L344"></a>case &#39;d&#39;: <span class="comment">// double</span>
            <a id="L345"></a>v[i] = math.Float64frombits(m.uint64())
        <a id="L346"></a>case &#39;D&#39;: <span class="comment">// double array</span>
            <a id="L347"></a>a := make([]float64, int(m.uint32()));
            <a id="L348"></a>for j := range a {
                <a id="L349"></a>a[j] = math.Float64frombits(m.uint64())
            <a id="L350"></a>}
            <a id="L351"></a>v[i] = a;
        <a id="L352"></a>case &#39;h&#39;: <span class="comment">// file descriptor (handle)</span>
            <a id="L353"></a>if len(m.rdesc) == 0 {
                <a id="L354"></a>if m.status == OK {
                    <a id="L355"></a>m.status = ErrBadArgType
                <a id="L356"></a>}
                <a id="L357"></a>return;
            <a id="L358"></a>}
            <a id="L359"></a>v[i] = int(m.rdesc[0]);
            <a id="L360"></a>m.rdesc = m.rdesc[1:len(m.rdesc)];
        <a id="L361"></a>case &#39;i&#39;: <span class="comment">// int</span>
            <a id="L362"></a>v[i] = int32(m.uint32())
        <a id="L363"></a>case &#39;I&#39;: <span class="comment">// int array</span>
            <a id="L364"></a>a := make([]int32, int(m.uint32()));
            <a id="L365"></a>for j := range a {
                <a id="L366"></a>a[j] = int32(m.uint32())
            <a id="L367"></a>}
            <a id="L368"></a>v[i] = a;
        <a id="L369"></a>case &#39;s&#39;: <span class="comment">// string</span>
            <a id="L370"></a>v[i] = string(m.bytes(int(m.uint32())))
        <a id="L371"></a>}
    <a id="L372"></a>}
<a id="L373"></a>}

<a id="L375"></a>func (m *msg) packValues(v []interface{}) {
    <a id="L376"></a>for i := range v {
        <a id="L377"></a>switch x := v[i].(type) {
        <a id="L378"></a>default:
            <a id="L379"></a>if m.status == OK {
                <a id="L380"></a>m.status = ErrInternalError
            <a id="L381"></a>}
            <a id="L382"></a>return;
        <a id="L383"></a>case bool:
            <a id="L384"></a>m.wuint8(&#39;b&#39;);
            <a id="L385"></a>if x {
                <a id="L386"></a>m.wuint8(1)
            <a id="L387"></a>} else {
                <a id="L388"></a>m.wuint8(0)
            <a id="L389"></a>}
        <a id="L390"></a>case []byte:
            <a id="L391"></a>m.wuint8(&#39;C&#39;);
            <a id="L392"></a>m.wuint32(uint32(len(x)));
            <a id="L393"></a>m.wbytes(x);
        <a id="L394"></a>case float64:
            <a id="L395"></a>m.wuint8(&#39;d&#39;);
            <a id="L396"></a>m.wuint64(math.Float64bits(x));
        <a id="L397"></a>case []float64:
            <a id="L398"></a>m.wuint8(&#39;D&#39;);
            <a id="L399"></a>m.wuint32(uint32(len(x)));
            <a id="L400"></a>for _, f := range x {
                <a id="L401"></a>m.wuint64(math.Float64bits(f))
            <a id="L402"></a>}
        <a id="L403"></a>case int32:
            <a id="L404"></a>m.wuint8(&#39;i&#39;);
            <a id="L405"></a>m.wuint32(uint32(x));
        <a id="L406"></a>case []int32:
            <a id="L407"></a>m.wuint8(&#39;I&#39;);
            <a id="L408"></a>m.wuint32(uint32(len(x)));
            <a id="L409"></a>for _, i := range x {
                <a id="L410"></a>m.wuint32(uint32(i))
            <a id="L411"></a>}
        <a id="L412"></a>case string:
            <a id="L413"></a>m.wuint8(&#39;s&#39;);
            <a id="L414"></a>m.wuint32(uint32(len(x)));
            <a id="L415"></a>m.wstring(x);
        <a id="L416"></a>}
    <a id="L417"></a>}
<a id="L418"></a>}

<a id="L420"></a>func (m *msg) unpackRequest() {
    <a id="L421"></a>m.status = OK;
    <a id="L422"></a>if m.unpackHeader(); m.status != OK {
        <a id="L423"></a>return
    <a id="L424"></a>}
    <a id="L425"></a>if m.protocol != protocol || !m.isReq {
        <a id="L426"></a>m.status = ErrProtocolMismatch;
        <a id="L427"></a>return;
    <a id="L428"></a>}

    <a id="L430"></a><span class="comment">// type-tagged argument values</span>
    <a id="L431"></a>m.Arg = make([]interface{}, m.uint32());
    <a id="L432"></a>m.unpackValues(m.Arg);
    <a id="L433"></a>if m.status != OK {
        <a id="L434"></a>return
    <a id="L435"></a>}

    <a id="L437"></a><span class="comment">// type-tagged expected return sizes.</span>
    <a id="L438"></a><span class="comment">// fill in zero values for each return value</span>
    <a id="L439"></a><span class="comment">// and save sizes.</span>
    <a id="L440"></a>m.fmt += &#34;:&#34;;
    <a id="L441"></a>m.Ret = make([]interface{}, m.uint32());
    <a id="L442"></a>m.Size = make([]int, len(m.Ret));
    <a id="L443"></a>for i := range m.Ret {
        <a id="L444"></a>t := m.uint8();
        <a id="L445"></a>m.fmt += string(t);
        <a id="L446"></a>switch t {
        <a id="L447"></a>default:
            <a id="L448"></a>if m.status == OK {
                <a id="L449"></a>m.status = ErrBadArgType
            <a id="L450"></a>}
            <a id="L451"></a>return;
        <a id="L452"></a>case &#39;b&#39;: <span class="comment">// bool[1]</span>
            <a id="L453"></a>m.Ret[i] = false
        <a id="L454"></a>case &#39;C&#39;: <span class="comment">// char array</span>
            <a id="L455"></a>m.Size[i] = int(m.uint32());
            <a id="L456"></a>m.Ret[i] = []byte(nil);
        <a id="L457"></a>case &#39;d&#39;: <span class="comment">// double</span>
            <a id="L458"></a>m.Ret[i] = float64(0)
        <a id="L459"></a>case &#39;D&#39;: <span class="comment">// double array</span>
            <a id="L460"></a>m.Size[i] = int(m.uint32());
            <a id="L461"></a>m.Ret[i] = []float64(nil);
        <a id="L462"></a>case &#39;h&#39;: <span class="comment">// file descriptor (handle)</span>
            <a id="L463"></a>m.Ret[i] = int(-1)
        <a id="L464"></a>case &#39;i&#39;: <span class="comment">// int</span>
            <a id="L465"></a>m.Ret[i] = int32(0)
        <a id="L466"></a>case &#39;I&#39;: <span class="comment">// int array</span>
            <a id="L467"></a>m.Size[i] = int(m.uint32());
            <a id="L468"></a>m.Ret[i] = []int32(nil);
        <a id="L469"></a>case &#39;s&#39;: <span class="comment">// string</span>
            <a id="L470"></a>m.Size[i] = int(m.uint32());
            <a id="L471"></a>m.Ret[i] = &#34;&#34;;
        <a id="L472"></a>}
    <a id="L473"></a>}
<a id="L474"></a>}

<a id="L476"></a>func (m *msg) packRequest() {
    <a id="L477"></a>m.packHeader();
    <a id="L478"></a>m.wuint32(uint32(len(m.Arg)));
    <a id="L479"></a>m.packValues(m.Arg);
    <a id="L480"></a>m.wuint32(uint32(len(m.Ret)));
    <a id="L481"></a>for i, v := range m.Ret {
        <a id="L482"></a>switch x := v.(type) {
        <a id="L483"></a>case bool:
            <a id="L484"></a>m.wuint8(&#39;b&#39;)
        <a id="L485"></a>case []byte:
            <a id="L486"></a>m.wuint8(&#39;C&#39;);
            <a id="L487"></a>m.wuint32(uint32(m.Size[i]));
        <a id="L488"></a>case float64:
            <a id="L489"></a>m.wuint8(&#39;d&#39;)
        <a id="L490"></a>case []float64:
            <a id="L491"></a>m.wuint8(&#39;D&#39;);
            <a id="L492"></a>m.wuint32(uint32(m.Size[i]));
        <a id="L493"></a>case int:
            <a id="L494"></a>m.wuint8(&#39;h&#39;)
        <a id="L495"></a>case int32:
            <a id="L496"></a>m.wuint8(&#39;i&#39;)
        <a id="L497"></a>case []int32:
            <a id="L498"></a>m.wuint8(&#39;I&#39;);
            <a id="L499"></a>m.wuint32(uint32(m.Size[i]));
        <a id="L500"></a>case string:
            <a id="L501"></a>m.wuint8(&#39;s&#39;);
            <a id="L502"></a>m.wuint32(uint32(m.Size[i]));
        <a id="L503"></a>}
    <a id="L504"></a>}
<a id="L505"></a>}

<a id="L507"></a>func (m *msg) unpackResponse() {
    <a id="L508"></a>m.status = OK;
    <a id="L509"></a>if m.unpackHeader(); m.status != OK {
        <a id="L510"></a>return
    <a id="L511"></a>}
    <a id="L512"></a>if m.protocol != protocol || m.isReq {
        <a id="L513"></a>m.status = ErrProtocolMismatch;
        <a id="L514"></a>return;
    <a id="L515"></a>}

    <a id="L517"></a><span class="comment">// type-tagged return values</span>
    <a id="L518"></a>m.fmt = &#34;&#34;;
    <a id="L519"></a>m.Ret = make([]interface{}, m.uint32());
    <a id="L520"></a>m.unpackValues(m.Ret);
<a id="L521"></a>}

<a id="L523"></a>func (m *msg) packResponse() {
    <a id="L524"></a>m.packHeader();
    <a id="L525"></a>m.wuint32(uint32(len(m.Ret)));
    <a id="L526"></a>m.packValues(m.Ret);
<a id="L527"></a>}
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
