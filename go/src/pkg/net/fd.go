<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/fd.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/fd.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// TODO(rsc): All the prints in this file should go to standard error.</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;once&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;sync&#34;;
    <a id="L13"></a>&#34;syscall&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// Network file descriptor.</span>
<a id="L17"></a>type netFD struct {
    <a id="L18"></a><span class="comment">// immutable until Close</span>
    <a id="L19"></a>fd     int;
    <a id="L20"></a>family int;
    <a id="L21"></a>proto  int;
    <a id="L22"></a>file   *os.File;
    <a id="L23"></a>cr     chan *netFD;
    <a id="L24"></a>cw     chan *netFD;
    <a id="L25"></a>net    string;
    <a id="L26"></a>laddr  Addr;
    <a id="L27"></a>raddr  Addr;

    <a id="L29"></a><span class="comment">// owned by client</span>
    <a id="L30"></a>rdeadline_delta int64;
    <a id="L31"></a>rdeadline       int64;
    <a id="L32"></a>rio             sync.Mutex;
    <a id="L33"></a>wdeadline_delta int64;
    <a id="L34"></a>wdeadline       int64;
    <a id="L35"></a>wio             sync.Mutex;

    <a id="L37"></a><span class="comment">// owned by fd wait server</span>
    <a id="L38"></a>ncr, ncw int;
<a id="L39"></a>}

<a id="L41"></a><span class="comment">// A pollServer helps FDs determine when to retry a non-blocking</span>
<a id="L42"></a><span class="comment">// read or write after they get EAGAIN.  When an FD needs to wait,</span>
<a id="L43"></a><span class="comment">// send the fd on s.cr (for a read) or s.cw (for a write) to pass the</span>
<a id="L44"></a><span class="comment">// request to the poll server.  Then receive on fd.cr/fd.cw.</span>
<a id="L45"></a><span class="comment">// When the pollServer finds that i/o on FD should be possible</span>
<a id="L46"></a><span class="comment">// again, it will send fd on fd.cr/fd.cw to wake any waiting processes.</span>
<a id="L47"></a><span class="comment">// This protocol is implemented as s.WaitRead() and s.WaitWrite().</span>
<a id="L48"></a><span class="comment">//</span>
<a id="L49"></a><span class="comment">// There is one subtlety: when sending on s.cr/s.cw, the</span>
<a id="L50"></a><span class="comment">// poll server is probably in a system call, waiting for an fd</span>
<a id="L51"></a><span class="comment">// to become ready.  It&#39;s not looking at the request channels.</span>
<a id="L52"></a><span class="comment">// To resolve this, the poll server waits not just on the FDs it has</span>
<a id="L53"></a><span class="comment">// been given but also its own pipe.  After sending on the</span>
<a id="L54"></a><span class="comment">// buffered channel s.cr/s.cw, WaitRead/WaitWrite writes a</span>
<a id="L55"></a><span class="comment">// byte to the pipe, causing the pollServer&#39;s poll system call to</span>
<a id="L56"></a><span class="comment">// return.  In response to the pipe being readable, the pollServer</span>
<a id="L57"></a><span class="comment">// re-polls its request channels.</span>
<a id="L58"></a><span class="comment">//</span>
<a id="L59"></a><span class="comment">// Note that the ordering is &#34;send request&#34; and then &#34;wake up server&#34;.</span>
<a id="L60"></a><span class="comment">// If the operations were reversed, there would be a race: the poll</span>
<a id="L61"></a><span class="comment">// server might wake up and look at the request channel, see that it</span>
<a id="L62"></a><span class="comment">// was empty, and go back to sleep, all before the requester managed</span>
<a id="L63"></a><span class="comment">// to send the request.  Because the send must complete before the wakeup,</span>
<a id="L64"></a><span class="comment">// the request channel must be buffered.  A buffer of size 1 is sufficient</span>
<a id="L65"></a><span class="comment">// for any request load.  If many processes are trying to submit requests,</span>
<a id="L66"></a><span class="comment">// one will succeed, the pollServer will read the request, and then the</span>
<a id="L67"></a><span class="comment">// channel will be empty for the next process&#39;s request.  A larger buffer</span>
<a id="L68"></a><span class="comment">// might help batch requests.</span>

<a id="L70"></a>type pollServer struct {
    <a id="L71"></a>cr, cw   chan *netFD; <span class="comment">// buffered &gt;= 1</span>
    <a id="L72"></a>pr, pw   *os.File;
    <a id="L73"></a>pending  map[int]*netFD;
    <a id="L74"></a>poll     *pollster; <span class="comment">// low-level OS hooks</span>
    <a id="L75"></a>deadline int64;     <span class="comment">// next deadline (nsec since 1970)</span>
<a id="L76"></a>}

<a id="L78"></a>func newPollServer() (s *pollServer, err os.Error) {
    <a id="L79"></a>s = new(pollServer);
    <a id="L80"></a>s.cr = make(chan *netFD, 1);
    <a id="L81"></a>s.cw = make(chan *netFD, 1);
    <a id="L82"></a>if s.pr, s.pw, err = os.Pipe(); err != nil {
        <a id="L83"></a>return nil, err
    <a id="L84"></a>}
    <a id="L85"></a>var e int;
    <a id="L86"></a>if e = syscall.SetNonblock(s.pr.Fd(), true); e != 0 {
    <a id="L87"></a>Errno:
        <a id="L88"></a>err = &amp;os.PathError{&#34;setnonblock&#34;, s.pr.Name(), os.Errno(e)};
    <a id="L89"></a>Error:
        <a id="L90"></a>s.pr.Close();
        <a id="L91"></a>s.pw.Close();
        <a id="L92"></a>return nil, err;
    <a id="L93"></a>}
    <a id="L94"></a>if e = syscall.SetNonblock(s.pw.Fd(), true); e != 0 {
        <a id="L95"></a>goto Errno
    <a id="L96"></a>}
    <a id="L97"></a>if s.poll, err = newpollster(); err != nil {
        <a id="L98"></a>goto Error
    <a id="L99"></a>}
    <a id="L100"></a>if err = s.poll.AddFD(s.pr.Fd(), &#39;r&#39;, true); err != nil {
        <a id="L101"></a>s.poll.Close();
        <a id="L102"></a>goto Error;
    <a id="L103"></a>}
    <a id="L104"></a>s.pending = make(map[int]*netFD);
    <a id="L105"></a>go s.Run();
    <a id="L106"></a>return s, nil;
<a id="L107"></a>}

<a id="L109"></a>func (s *pollServer) AddFD(fd *netFD, mode int) {
    <a id="L110"></a><span class="comment">// TODO(rsc): This check handles a race between</span>
    <a id="L111"></a><span class="comment">// one goroutine reading and another one closing,</span>
    <a id="L112"></a><span class="comment">// but it doesn&#39;t solve the race completely:</span>
    <a id="L113"></a><span class="comment">// it still could happen that one goroutine closes</span>
    <a id="L114"></a><span class="comment">// but we read fd.fd before it does, and then</span>
    <a id="L115"></a><span class="comment">// another goroutine creates a new open file with</span>
    <a id="L116"></a><span class="comment">// that fd, which we&#39;d now be referring to.</span>
    <a id="L117"></a><span class="comment">// The fix is probably to send the Close call</span>
    <a id="L118"></a><span class="comment">// through the poll server too, except that</span>
    <a id="L119"></a><span class="comment">// not all Reads and Writes go through the poll</span>
    <a id="L120"></a><span class="comment">// server even now.</span>
    <a id="L121"></a>intfd := fd.fd;
    <a id="L122"></a>if intfd &lt; 0 {
        <a id="L123"></a><span class="comment">// fd closed underfoot</span>
        <a id="L124"></a>if mode == &#39;r&#39; {
            <a id="L125"></a>fd.cr &lt;- fd
        <a id="L126"></a>} else {
            <a id="L127"></a>fd.cw &lt;- fd
        <a id="L128"></a>}
        <a id="L129"></a>return;
    <a id="L130"></a>}
    <a id="L131"></a>if err := s.poll.AddFD(intfd, mode, false); err != nil {
        <a id="L132"></a>panicln(&#34;pollServer AddFD &#34;, intfd, &#34;: &#34;, err.String(), &#34;\n&#34;);
        <a id="L133"></a>return;
    <a id="L134"></a>}

    <a id="L136"></a>var t int64;
    <a id="L137"></a>key := intfd &lt;&lt; 1;
    <a id="L138"></a>if mode == &#39;r&#39; {
        <a id="L139"></a>fd.ncr++;
        <a id="L140"></a>t = fd.rdeadline;
    <a id="L141"></a>} else {
        <a id="L142"></a>fd.ncw++;
        <a id="L143"></a>key++;
        <a id="L144"></a>t = fd.wdeadline;
    <a id="L145"></a>}
    <a id="L146"></a>s.pending[key] = fd;
    <a id="L147"></a>if t &gt; 0 &amp;&amp; (s.deadline == 0 || t &lt; s.deadline) {
        <a id="L148"></a>s.deadline = t
    <a id="L149"></a>}
<a id="L150"></a>}

<a id="L152"></a>func (s *pollServer) LookupFD(fd int, mode int) *netFD {
    <a id="L153"></a>key := fd &lt;&lt; 1;
    <a id="L154"></a>if mode == &#39;w&#39; {
        <a id="L155"></a>key++
    <a id="L156"></a>}
    <a id="L157"></a>netfd, ok := s.pending[key];
    <a id="L158"></a>if !ok {
        <a id="L159"></a>return nil
    <a id="L160"></a>}
    <a id="L161"></a>s.pending[key] = nil, false;
    <a id="L162"></a>return netfd;
<a id="L163"></a>}

<a id="L165"></a>func (s *pollServer) WakeFD(fd *netFD, mode int) {
    <a id="L166"></a>if mode == &#39;r&#39; {
        <a id="L167"></a>for fd.ncr &gt; 0 {
            <a id="L168"></a>fd.ncr--;
            <a id="L169"></a>fd.cr &lt;- fd;
        <a id="L170"></a>}
    <a id="L171"></a>} else {
        <a id="L172"></a>for fd.ncw &gt; 0 {
            <a id="L173"></a>fd.ncw--;
            <a id="L174"></a>fd.cw &lt;- fd;
        <a id="L175"></a>}
    <a id="L176"></a>}
<a id="L177"></a>}

<a id="L179"></a>func (s *pollServer) Now() int64 {
    <a id="L180"></a>sec, nsec, err := os.Time();
    <a id="L181"></a>if err != nil {
        <a id="L182"></a>panic(&#34;net: os.Time: &#34;, err.String())
    <a id="L183"></a>}
    <a id="L184"></a>nsec += sec * 1e9;
    <a id="L185"></a>return nsec;
<a id="L186"></a>}

<a id="L188"></a>func (s *pollServer) CheckDeadlines() {
    <a id="L189"></a>now := s.Now();
    <a id="L190"></a><span class="comment">// TODO(rsc): This will need to be handled more efficiently,</span>
    <a id="L191"></a><span class="comment">// probably with a heap indexed by wakeup time.</span>

    <a id="L193"></a>var next_deadline int64;
    <a id="L194"></a>for key, fd := range s.pending {
        <a id="L195"></a>var t int64;
        <a id="L196"></a>var mode int;
        <a id="L197"></a>if key&amp;1 == 0 {
            <a id="L198"></a>mode = &#39;r&#39;
        <a id="L199"></a>} else {
            <a id="L200"></a>mode = &#39;w&#39;
        <a id="L201"></a>}
        <a id="L202"></a>if mode == &#39;r&#39; {
            <a id="L203"></a>t = fd.rdeadline
        <a id="L204"></a>} else {
            <a id="L205"></a>t = fd.wdeadline
        <a id="L206"></a>}
        <a id="L207"></a>if t &gt; 0 {
            <a id="L208"></a>if t &lt;= now {
                <a id="L209"></a>s.pending[key] = nil, false;
                <a id="L210"></a>if mode == &#39;r&#39; {
                    <a id="L211"></a>s.poll.DelFD(fd.fd, mode);
                    <a id="L212"></a>fd.rdeadline = -1;
                <a id="L213"></a>} else {
                    <a id="L214"></a>s.poll.DelFD(fd.fd, mode);
                    <a id="L215"></a>fd.wdeadline = -1;
                <a id="L216"></a>}
                <a id="L217"></a>s.WakeFD(fd, mode);
            <a id="L218"></a>} else if next_deadline == 0 || t &lt; next_deadline {
                <a id="L219"></a>next_deadline = t
            <a id="L220"></a>}
        <a id="L221"></a>}
    <a id="L222"></a>}
    <a id="L223"></a>s.deadline = next_deadline;
<a id="L224"></a>}

<a id="L226"></a>func (s *pollServer) Run() {
    <a id="L227"></a>var scratch [100]byte;
    <a id="L228"></a>for {
        <a id="L229"></a>var t = s.deadline;
        <a id="L230"></a>if t &gt; 0 {
            <a id="L231"></a>t = t - s.Now();
            <a id="L232"></a>if t &lt; 0 {
                <a id="L233"></a>s.CheckDeadlines();
                <a id="L234"></a>continue;
            <a id="L235"></a>}
        <a id="L236"></a>}
        <a id="L237"></a>fd, mode, err := s.poll.WaitFD(t);
        <a id="L238"></a>if err != nil {
            <a id="L239"></a>print(&#34;pollServer WaitFD: &#34;, err.String(), &#34;\n&#34;);
            <a id="L240"></a>return;
        <a id="L241"></a>}
        <a id="L242"></a>if fd &lt; 0 {
            <a id="L243"></a><span class="comment">// Timeout happened.</span>
            <a id="L244"></a>s.CheckDeadlines();
            <a id="L245"></a>continue;
        <a id="L246"></a>}
        <a id="L247"></a>if fd == s.pr.Fd() {
            <a id="L248"></a><span class="comment">// Drain our wakeup pipe.</span>
            <a id="L249"></a>for nn, _ := s.pr.Read(&amp;scratch); nn &gt; 0; {
                <a id="L250"></a>nn, _ = s.pr.Read(&amp;scratch)
            <a id="L251"></a>}

            <a id="L253"></a><span class="comment">// Read from channels</span>
            <a id="L254"></a>for fd, ok := &lt;-s.cr; ok; fd, ok = &lt;-s.cr {
                <a id="L255"></a>s.AddFD(fd, &#39;r&#39;)
            <a id="L256"></a>}
            <a id="L257"></a>for fd, ok := &lt;-s.cw; ok; fd, ok = &lt;-s.cw {
                <a id="L258"></a>s.AddFD(fd, &#39;w&#39;)
            <a id="L259"></a>}
        <a id="L260"></a>} else {
            <a id="L261"></a>netfd := s.LookupFD(fd, mode);
            <a id="L262"></a>if netfd == nil {
                <a id="L263"></a>print(&#34;pollServer: unexpected wakeup for fd=&#34;, netfd, &#34; mode=&#34;, string(mode), &#34;\n&#34;);
                <a id="L264"></a>continue;
            <a id="L265"></a>}
            <a id="L266"></a>s.WakeFD(netfd, mode);
        <a id="L267"></a>}
    <a id="L268"></a>}
<a id="L269"></a>}

<a id="L271"></a>var wakeupbuf [1]byte

<a id="L273"></a>func (s *pollServer) Wakeup() { s.pw.Write(&amp;wakeupbuf) }

<a id="L275"></a>func (s *pollServer) WaitRead(fd *netFD) {
    <a id="L276"></a>s.cr &lt;- fd;
    <a id="L277"></a>s.Wakeup();
    <a id="L278"></a>&lt;-fd.cr;
<a id="L279"></a>}

<a id="L281"></a>func (s *pollServer) WaitWrite(fd *netFD) {
    <a id="L282"></a>s.cw &lt;- fd;
    <a id="L283"></a>s.Wakeup();
    <a id="L284"></a>&lt;-fd.cw;
<a id="L285"></a>}


<a id="L288"></a><span class="comment">// Network FD methods.</span>
<a id="L289"></a><span class="comment">// All the network FDs use a single pollServer.</span>

<a id="L291"></a>var pollserver *pollServer

<a id="L293"></a>func startServer() {
    <a id="L294"></a>p, err := newPollServer();
    <a id="L295"></a>if err != nil {
        <a id="L296"></a>print(&#34;Start pollServer: &#34;, err.String(), &#34;\n&#34;)
    <a id="L297"></a>}
    <a id="L298"></a>pollserver = p;
<a id="L299"></a>}

<a id="L301"></a>func newFD(fd, family, proto int, net string, laddr, raddr Addr) (f *netFD, err os.Error) {
    <a id="L302"></a>once.Do(startServer);
    <a id="L303"></a>if e := syscall.SetNonblock(fd, true); e != 0 {
        <a id="L304"></a>return nil, &amp;OpError{&#34;setnonblock&#34;, net, laddr, os.Errno(e)}
    <a id="L305"></a>}
    <a id="L306"></a>f = &amp;netFD{
        <a id="L307"></a>fd: fd,
        <a id="L308"></a>family: family,
        <a id="L309"></a>proto: proto,
        <a id="L310"></a>net: net,
        <a id="L311"></a>laddr: laddr,
        <a id="L312"></a>raddr: raddr,
    <a id="L313"></a>};
    <a id="L314"></a>var ls, rs string;
    <a id="L315"></a>if laddr != nil {
        <a id="L316"></a>ls = laddr.String()
    <a id="L317"></a>}
    <a id="L318"></a>if raddr != nil {
        <a id="L319"></a>rs = raddr.String()
    <a id="L320"></a>}
    <a id="L321"></a>f.file = os.NewFile(fd, net+&#34;:&#34;+ls+&#34;-&gt;&#34;+rs);
    <a id="L322"></a>f.cr = make(chan *netFD, 1);
    <a id="L323"></a>f.cw = make(chan *netFD, 1);
    <a id="L324"></a>return f, nil;
<a id="L325"></a>}

<a id="L327"></a>func isEAGAIN(e os.Error) bool {
    <a id="L328"></a>if e1, ok := e.(*os.PathError); ok {
        <a id="L329"></a>return e1.Error == os.EAGAIN
    <a id="L330"></a>}
    <a id="L331"></a>return e == os.EAGAIN;
<a id="L332"></a>}

<a id="L334"></a>func (fd *netFD) Close() os.Error {
    <a id="L335"></a>if fd == nil || fd.file == nil {
        <a id="L336"></a>return os.EINVAL
    <a id="L337"></a>}

    <a id="L339"></a><span class="comment">// In case the user has set linger,</span>
    <a id="L340"></a><span class="comment">// switch to blocking mode so the close blocks.</span>
    <a id="L341"></a><span class="comment">// As long as this doesn&#39;t happen often,</span>
    <a id="L342"></a><span class="comment">// we can handle the extra OS processes.</span>
    <a id="L343"></a><span class="comment">// Otherwise we&#39;ll need to use the pollserver</span>
    <a id="L344"></a><span class="comment">// for Close too.  Sigh.</span>
    <a id="L345"></a>syscall.SetNonblock(fd.file.Fd(), false);

    <a id="L347"></a>e := fd.file.Close();
    <a id="L348"></a>fd.file = nil;
    <a id="L349"></a>fd.fd = -1;
    <a id="L350"></a>return e;
<a id="L351"></a>}

<a id="L353"></a>func (fd *netFD) Read(p []byte) (n int, err os.Error) {
    <a id="L354"></a>if fd == nil || fd.file == nil {
        <a id="L355"></a>return 0, os.EINVAL
    <a id="L356"></a>}
    <a id="L357"></a>fd.rio.Lock();
    <a id="L358"></a>defer fd.rio.Unlock();
    <a id="L359"></a>if fd.rdeadline_delta &gt; 0 {
        <a id="L360"></a>fd.rdeadline = pollserver.Now() + fd.rdeadline_delta
    <a id="L361"></a>} else {
        <a id="L362"></a>fd.rdeadline = 0
    <a id="L363"></a>}
    <a id="L364"></a>for {
        <a id="L365"></a>n, err = fd.file.Read(p);
        <a id="L366"></a>if isEAGAIN(err) &amp;&amp; fd.rdeadline &gt;= 0 {
            <a id="L367"></a>pollserver.WaitRead(fd);
            <a id="L368"></a>continue;
        <a id="L369"></a>}
        <a id="L370"></a>break;
    <a id="L371"></a>}
    <a id="L372"></a>return;
<a id="L373"></a>}

<a id="L375"></a>func (fd *netFD) Write(p []byte) (n int, err os.Error) {
    <a id="L376"></a>if fd == nil || fd.file == nil {
        <a id="L377"></a>return 0, os.EINVAL
    <a id="L378"></a>}
    <a id="L379"></a>fd.wio.Lock();
    <a id="L380"></a>defer fd.wio.Unlock();
    <a id="L381"></a>if fd.wdeadline_delta &gt; 0 {
        <a id="L382"></a>fd.wdeadline = pollserver.Now() + fd.wdeadline_delta
    <a id="L383"></a>} else {
        <a id="L384"></a>fd.wdeadline = 0
    <a id="L385"></a>}
    <a id="L386"></a>err = nil;
    <a id="L387"></a>nn := 0;
    <a id="L388"></a>for nn &lt; len(p) {
        <a id="L389"></a>n, err = fd.file.Write(p[nn:len(p)]);
        <a id="L390"></a>if n &gt; 0 {
            <a id="L391"></a>nn += n
        <a id="L392"></a>}
        <a id="L393"></a>if nn == len(p) {
            <a id="L394"></a>break
        <a id="L395"></a>}
        <a id="L396"></a>if isEAGAIN(err) &amp;&amp; fd.wdeadline &gt;= 0 {
            <a id="L397"></a>pollserver.WaitWrite(fd);
            <a id="L398"></a>continue;
        <a id="L399"></a>}
        <a id="L400"></a>if n == 0 || err != nil {
            <a id="L401"></a>break
        <a id="L402"></a>}
    <a id="L403"></a>}
    <a id="L404"></a>return nn, err;
<a id="L405"></a>}

<a id="L407"></a>func (fd *netFD) accept(toAddr func(syscall.Sockaddr) Addr) (nfd *netFD, err os.Error) {
    <a id="L408"></a>if fd == nil || fd.file == nil {
        <a id="L409"></a>return nil, os.EINVAL
    <a id="L410"></a>}

    <a id="L412"></a><span class="comment">// See ../syscall/exec.go for description of ForkLock.</span>
    <a id="L413"></a><span class="comment">// It is okay to hold the lock across syscall.Accept</span>
    <a id="L414"></a><span class="comment">// because we have put fd.fd into non-blocking mode.</span>
    <a id="L415"></a>syscall.ForkLock.RLock();
    <a id="L416"></a>var s, e int;
    <a id="L417"></a>var sa syscall.Sockaddr;
    <a id="L418"></a>for {
        <a id="L419"></a>s, sa, e = syscall.Accept(fd.fd);
        <a id="L420"></a>if e != syscall.EAGAIN {
            <a id="L421"></a>break
        <a id="L422"></a>}
        <a id="L423"></a>syscall.ForkLock.RUnlock();
        <a id="L424"></a>pollserver.WaitRead(fd);
        <a id="L425"></a>syscall.ForkLock.RLock();
    <a id="L426"></a>}
    <a id="L427"></a>if e != 0 {
        <a id="L428"></a>syscall.ForkLock.RUnlock();
        <a id="L429"></a>return nil, &amp;OpError{&#34;accept&#34;, fd.net, fd.laddr, os.Errno(e)};
    <a id="L430"></a>}
    <a id="L431"></a>syscall.CloseOnExec(s);
    <a id="L432"></a>syscall.ForkLock.RUnlock();

    <a id="L434"></a>if nfd, err = newFD(s, fd.family, fd.proto, fd.net, fd.laddr, toAddr(sa)); err != nil {
        <a id="L435"></a>syscall.Close(s);
        <a id="L436"></a>return nil, err;
    <a id="L437"></a>}
    <a id="L438"></a>return nfd, nil;
<a id="L439"></a>}
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
