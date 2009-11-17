<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/ztypes_darwin_386.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/ztypes_darwin_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// godefs -gsyscall -f-m32 types_darwin.c</span>

<a id="L3"></a><span class="comment">// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L5"></a>package syscall

<a id="L7"></a><span class="comment">// Constants</span>
<a id="L8"></a>const (
    <a id="L9"></a>sizeofPtr           = 0x4;
    <a id="L10"></a>sizeofShort         = 0x2;
    <a id="L11"></a>sizeofInt           = 0x4;
    <a id="L12"></a>sizeofLong          = 0x4;
    <a id="L13"></a>sizeofLongLong      = 0x8;
    <a id="L14"></a>O_CLOEXEC           = 0;
    <a id="L15"></a>SizeofSockaddrInet4 = 0x10;
    <a id="L16"></a>SizeofSockaddrInet6 = 0x1c;
    <a id="L17"></a>SizeofSockaddrAny   = 0x6c;
    <a id="L18"></a>SizeofSockaddrUnix  = 0x6a;
    <a id="L19"></a>SizeofLinger        = 0x8;
    <a id="L20"></a>SizeofMsghdr        = 0x1c;
    <a id="L21"></a>SizeofCmsghdr       = 0xc;
    <a id="L22"></a>PTRACE_TRACEME      = 0;
    <a id="L23"></a>PTRACE_CONT         = 0x7;
    <a id="L24"></a>PTRACE_KILL         = 0x8;
<a id="L25"></a>)

<a id="L27"></a><span class="comment">// Types</span>

<a id="L29"></a>type _C_short int16

<a id="L31"></a>type _C_int int32

<a id="L33"></a>type _C_long int32

<a id="L35"></a>type _C_long_long int64

<a id="L37"></a>type Timespec struct {
    <a id="L38"></a>Sec  int32;
    <a id="L39"></a>Nsec int32;
<a id="L40"></a>}

<a id="L42"></a>type Timeval struct {
    <a id="L43"></a>Sec  int32;
    <a id="L44"></a>Usec int32;
<a id="L45"></a>}

<a id="L47"></a>type Rusage struct {
    <a id="L48"></a>Utime    Timeval;
    <a id="L49"></a>Stime    Timeval;
    <a id="L50"></a>Maxrss   int32;
    <a id="L51"></a>Ixrss    int32;
    <a id="L52"></a>Idrss    int32;
    <a id="L53"></a>Isrss    int32;
    <a id="L54"></a>Minflt   int32;
    <a id="L55"></a>Majflt   int32;
    <a id="L56"></a>Nswap    int32;
    <a id="L57"></a>Inblock  int32;
    <a id="L58"></a>Oublock  int32;
    <a id="L59"></a>Msgsnd   int32;
    <a id="L60"></a>Msgrcv   int32;
    <a id="L61"></a>Nsignals int32;
    <a id="L62"></a>Nvcsw    int32;
    <a id="L63"></a>Nivcsw   int32;
<a id="L64"></a>}

<a id="L66"></a>type Rlimit struct {
    <a id="L67"></a>Cur uint64;
    <a id="L68"></a>Max uint64;
<a id="L69"></a>}

<a id="L71"></a>type _Gid_t uint32

<a id="L73"></a>type Stat_t struct {
    <a id="L74"></a>Dev           int32;
    <a id="L75"></a>Mode          uint16;
    <a id="L76"></a>Nlink         uint16;
    <a id="L77"></a>Ino           uint64;
    <a id="L78"></a>Uid           uint32;
    <a id="L79"></a>Gid           uint32;
    <a id="L80"></a>Rdev          int32;
    <a id="L81"></a>Atimespec     Timespec;
    <a id="L82"></a>Mtimespec     Timespec;
    <a id="L83"></a>Ctimespec     Timespec;
    <a id="L84"></a>Birthtimespec Timespec;
    <a id="L85"></a>Size          int64;
    <a id="L86"></a>Blocks        int64;
    <a id="L87"></a>Blksize       int32;
    <a id="L88"></a>Flags         uint32;
    <a id="L89"></a>Gen           uint32;
    <a id="L90"></a>Lspare        int32;
    <a id="L91"></a>Qspare        [2]int64;
<a id="L92"></a>}

<a id="L94"></a>type Statfs_t struct {
    <a id="L95"></a>Bsize       uint32;
    <a id="L96"></a>Iosize      int32;
    <a id="L97"></a>Blocks      uint64;
    <a id="L98"></a>Bfree       uint64;
    <a id="L99"></a>Bavail      uint64;
    <a id="L100"></a>Files       uint64;
    <a id="L101"></a>Ffree       uint64;
    <a id="L102"></a>Fsid        [8]byte; <span class="comment">/* fsid */</span>
    <a id="L103"></a>Owner       uint32;
    <a id="L104"></a>Type        uint32;
    <a id="L105"></a>Flags       uint32;
    <a id="L106"></a>Fssubtype   uint32;
    <a id="L107"></a>Fstypename  [16]int8;
    <a id="L108"></a>Mntonname   [1024]int8;
    <a id="L109"></a>Mntfromname [1024]int8;
    <a id="L110"></a>Reserved    [8]uint32;
<a id="L111"></a>}

<a id="L113"></a>type Flock_t struct {
    <a id="L114"></a>Start  int64;
    <a id="L115"></a>Len    int64;
    <a id="L116"></a>Pid    int32;
    <a id="L117"></a>Type   int16;
    <a id="L118"></a>Whence int16;
<a id="L119"></a>}

<a id="L121"></a>type Fstore_t struct {
    <a id="L122"></a>Flags      uint32;
    <a id="L123"></a>Posmode    int32;
    <a id="L124"></a>Offset     int64;
    <a id="L125"></a>Length     int64;
    <a id="L126"></a>Bytesalloc int64;
<a id="L127"></a>}

<a id="L129"></a>type Radvisory_t struct {
    <a id="L130"></a>Offset int64;
    <a id="L131"></a>Count  int32;
<a id="L132"></a>}

<a id="L134"></a>type Fbootstraptransfer_t struct {
    <a id="L135"></a>Offset int64;
    <a id="L136"></a>Length uint32;
    <a id="L137"></a>Buffer *byte;
<a id="L138"></a>}

<a id="L140"></a>type Log2phys_t struct {
    <a id="L141"></a>Flags       uint32;
    <a id="L142"></a>Contigbytes int64;
    <a id="L143"></a>Devoffset   int64;
<a id="L144"></a>}

<a id="L146"></a>type Dirent struct {
    <a id="L147"></a>Ino     uint64;
    <a id="L148"></a>Seekoff uint64;
    <a id="L149"></a>Reclen  uint16;
    <a id="L150"></a>Namlen  uint16;
    <a id="L151"></a>Type    uint8;
    <a id="L152"></a>Name    [1024]int8;
    <a id="L153"></a>Pad0    [3]byte;
<a id="L154"></a>}

<a id="L156"></a>type RawSockaddrInet4 struct {
    <a id="L157"></a>Len    uint8;
    <a id="L158"></a>Family uint8;
    <a id="L159"></a>Port   uint16;
    <a id="L160"></a>Addr   [4]byte; <span class="comment">/* in_addr */</span>
    <a id="L161"></a>Zero   [8]int8;
<a id="L162"></a>}

<a id="L164"></a>type RawSockaddrInet6 struct {
    <a id="L165"></a>Len      uint8;
    <a id="L166"></a>Family   uint8;
    <a id="L167"></a>Port     uint16;
    <a id="L168"></a>Flowinfo uint32;
    <a id="L169"></a>Addr     [16]byte; <span class="comment">/* in6_addr */</span>
    <a id="L170"></a>Scope_id uint32;
<a id="L171"></a>}

<a id="L173"></a>type RawSockaddrUnix struct {
    <a id="L174"></a>Len    uint8;
    <a id="L175"></a>Family uint8;
    <a id="L176"></a>Path   [104]int8;
<a id="L177"></a>}

<a id="L179"></a>type RawSockaddr struct {
    <a id="L180"></a>Len    uint8;
    <a id="L181"></a>Family uint8;
    <a id="L182"></a>Data   [14]int8;
<a id="L183"></a>}

<a id="L185"></a>type RawSockaddrAny struct {
    <a id="L186"></a>Addr RawSockaddr;
    <a id="L187"></a>Pad  [92]int8;
<a id="L188"></a>}

<a id="L190"></a>type _Socklen uint32

<a id="L192"></a>type Linger struct {
    <a id="L193"></a>Onoff  int32;
    <a id="L194"></a>Linger int32;
<a id="L195"></a>}

<a id="L197"></a>type Iovec struct {
    <a id="L198"></a>Base *byte;
    <a id="L199"></a>Len  uint32;
<a id="L200"></a>}

<a id="L202"></a>type Msghdr struct {
    <a id="L203"></a>Name       *byte;
    <a id="L204"></a>Namelen    uint32;
    <a id="L205"></a>Iov        *Iovec;
    <a id="L206"></a>Iovlen     int32;
    <a id="L207"></a>Control    *byte;
    <a id="L208"></a>Controllen uint32;
    <a id="L209"></a>Flags      int32;
<a id="L210"></a>}

<a id="L212"></a>type Cmsghdr struct {
    <a id="L213"></a>Len   uint32;
    <a id="L214"></a>Level int32;
    <a id="L215"></a>Type  int32;
<a id="L216"></a>}

<a id="L218"></a>type Kevent_t struct {
    <a id="L219"></a>Ident  uint32;
    <a id="L220"></a>Filter int16;
    <a id="L221"></a>Flags  uint16;
    <a id="L222"></a>Fflags uint32;
    <a id="L223"></a>Data   int32;
    <a id="L224"></a>Udata  *byte;
<a id="L225"></a>}

<a id="L227"></a>type FdSet struct {
    <a id="L228"></a>Bits [32]int32;
<a id="L229"></a>}
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
