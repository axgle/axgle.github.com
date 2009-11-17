<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/sync/mutex.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/sync/mutex.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The sync package provides basic synchronization primitives</span>
<a id="L6"></a><span class="comment">// such as mutual exclusion locks.  These are intended for use</span>
<a id="L7"></a><span class="comment">// by low-level library routines.  Higher-level synchronization</span>
<a id="L8"></a><span class="comment">// is better done via channels and communication.</span>
<a id="L9"></a>package sync

<a id="L11"></a>import &#34;runtime&#34;

<a id="L13"></a>func cas(val *uint32, old, new uint32) bool

<a id="L15"></a><span class="comment">// A Mutex is a mutual exclusion lock.</span>
<a id="L16"></a><span class="comment">// Mutexes can be created as part of other structures;</span>
<a id="L17"></a><span class="comment">// the zero value for a Mutex is an unlocked mutex.</span>
<a id="L18"></a>type Mutex struct {
    <a id="L19"></a>key  uint32;
    <a id="L20"></a>sema uint32;
<a id="L21"></a>}

<a id="L23"></a>func xadd(val *uint32, delta int32) (new uint32) {
    <a id="L24"></a>for {
        <a id="L25"></a>v := *val;
        <a id="L26"></a>nv := v + uint32(delta);
        <a id="L27"></a>if cas(val, v, nv) {
            <a id="L28"></a>return nv
        <a id="L29"></a>}
    <a id="L30"></a>}
    <a id="L31"></a>panic(&#34;unreached&#34;);
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// Lock locks m.</span>
<a id="L35"></a><span class="comment">// If the lock is already in use, the calling goroutine</span>
<a id="L36"></a><span class="comment">// blocks until the mutex is available.</span>
<a id="L37"></a>func (m *Mutex) Lock() {
    <a id="L38"></a>if xadd(&amp;m.key, 1) == 1 {
        <a id="L39"></a><span class="comment">// changed from 0 to 1; we hold lock</span>
        <a id="L40"></a>return
    <a id="L41"></a>}
    <a id="L42"></a>runtime.Semacquire(&amp;m.sema);
<a id="L43"></a>}

<a id="L45"></a><span class="comment">// Unlock unlocks m.</span>
<a id="L46"></a><span class="comment">// It is a run-time error if m is not locked on entry to Unlock.</span>
<a id="L47"></a><span class="comment">//</span>
<a id="L48"></a><span class="comment">// A locked Mutex is not associated with a particular goroutine.</span>
<a id="L49"></a><span class="comment">// It is allowed for one goroutine to lock a Mutex and then</span>
<a id="L50"></a><span class="comment">// arrange for another goroutine to unlock it.</span>
<a id="L51"></a>func (m *Mutex) Unlock() {
    <a id="L52"></a>if xadd(&amp;m.key, -1) == 0 {
        <a id="L53"></a><span class="comment">// changed from 1 to 0; no contention</span>
        <a id="L54"></a>return
    <a id="L55"></a>}
    <a id="L56"></a>runtime.Semrelease(&amp;m.sema);
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// Stub implementation of r/w locks.</span>
<a id="L60"></a><span class="comment">// This satisfies the semantics but</span>
<a id="L61"></a><span class="comment">// is not terribly efficient.</span>

<a id="L63"></a><span class="comment">// The next comment goes in the BUGS section of the document,</span>
<a id="L64"></a><span class="comment">// in its own paragraph, without the (rsc) tag.</span>

<a id="L66"></a><span class="comment">// BUG(rsc): RWMutex does not (yet) allow multiple readers;</span>
<a id="L67"></a><span class="comment">// instead it behaves as if RLock and RUnlock were Lock and Unlock.</span>

<a id="L69"></a><span class="comment">// An RWMutex is a reader/writer mutual exclusion lock.</span>
<a id="L70"></a><span class="comment">// The lock can be held by an arbitrary number of readers</span>
<a id="L71"></a><span class="comment">// or a single writer.</span>
<a id="L72"></a><span class="comment">// RWMutexes can be created as part of other</span>
<a id="L73"></a><span class="comment">// structures; the zero value for a RWMutex is</span>
<a id="L74"></a><span class="comment">// an unlocked mutex.</span>
<a id="L75"></a>type RWMutex struct {
    <a id="L76"></a>m Mutex;
<a id="L77"></a>}

<a id="L79"></a><span class="comment">// RLock locks rw for reading.</span>
<a id="L80"></a><span class="comment">// If the lock is already locked for writing or there is a writer already waiting</span>
<a id="L81"></a><span class="comment">// to acquire the lock, RLock blocks until the writer has released the lock.</span>
<a id="L82"></a>func (rw *RWMutex) RLock() { rw.m.Lock() }

<a id="L84"></a><span class="comment">// RUnlock undoes a single RLock call;</span>
<a id="L85"></a><span class="comment">// it does not affect other simultaneous readers.</span>
<a id="L86"></a><span class="comment">// It is a run-time error if rw is not locked for reading</span>
<a id="L87"></a><span class="comment">// on entry to RUnlock.</span>
<a id="L88"></a>func (rw *RWMutex) RUnlock() { rw.m.Unlock() }

<a id="L90"></a><span class="comment">// Lock locks rw for writing.</span>
<a id="L91"></a><span class="comment">// If the lock is already locked for reading or writing,</span>
<a id="L92"></a><span class="comment">// Lock blocks until the lock is available.</span>
<a id="L93"></a><span class="comment">// To ensure that the lock eventually becomes available,</span>
<a id="L94"></a><span class="comment">// a blocked Lock call excludes new readers from acquiring</span>
<a id="L95"></a><span class="comment">// the lock.</span>
<a id="L96"></a>func (rw *RWMutex) Lock() { rw.m.Lock() }

<a id="L98"></a><span class="comment">// Unlock unlocks rw for writing.</span>
<a id="L99"></a><span class="comment">// It is a run-time error if rw is not locked for writing</span>
<a id="L100"></a><span class="comment">// on entry to Unlock.</span>
<a id="L101"></a><span class="comment">//</span>
<a id="L102"></a><span class="comment">// Like for Mutexes,</span>
<a id="L103"></a><span class="comment">// a locked RWMutex is not associated with a particular goroutine.</span>
<a id="L104"></a><span class="comment">// It is allowed for one goroutine to RLock (Lock) an RWMutex and then</span>
<a id="L105"></a><span class="comment">// arrange for another goroutine to RUnlock (Unlock) it.</span>
<a id="L106"></a>func (rw *RWMutex) Unlock() { rw.m.Unlock() }
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
