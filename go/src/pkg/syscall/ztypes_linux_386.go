<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/ztypes_linux_386.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/ztypes_linux_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// godefs -gsyscall -f-m32 types_linux.c</span>

<a id="L3"></a><span class="comment">// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L5"></a>package syscall

<a id="L7"></a><span class="comment">// Constants</span>
<a id="L8"></a>const (
    <a id="L9"></a>sizeofPtr           = 0x4;
    <a id="L10"></a>sizeofShort         = 0x2;
    <a id="L11"></a>sizeofInt           = 0x4;
    <a id="L12"></a>sizeofLong          = 0x4;
    <a id="L13"></a>sizeofLongLong      = 0x8;
    <a id="L14"></a>PathMax             = 0x1000;
    <a id="L15"></a>SizeofSockaddrInet4 = 0x10;
    <a id="L16"></a>SizeofSockaddrInet6 = 0x1c;
    <a id="L17"></a>SizeofSockaddrAny   = 0x70;
    <a id="L18"></a>SizeofSockaddrUnix  = 0x6e;
    <a id="L19"></a>SizeofLinger        = 0x8;
    <a id="L20"></a>SizeofMsghdr        = 0x1c;
    <a id="L21"></a>SizeofCmsghdr       = 0xc;
<a id="L22"></a>)

<a id="L24"></a><span class="comment">// Types</span>

<a id="L26"></a>type _C_short int16

<a id="L28"></a>type _C_int int32

<a id="L30"></a>type _C_long int32

<a id="L32"></a>type _C_long_long int64

<a id="L34"></a>type Timespec struct {
    <a id="L35"></a>Sec  int32;
    <a id="L36"></a>Nsec int32;
<a id="L37"></a>}

<a id="L39"></a>type Timeval struct {
    <a id="L40"></a>Sec  int32;
    <a id="L41"></a>Usec int32;
<a id="L42"></a>}

<a id="L44"></a>type Timex struct {
    <a id="L45"></a>Modes     uint32;
    <a id="L46"></a>Offset    int32;
    <a id="L47"></a>Freq      int32;
    <a id="L48"></a>Maxerror  int32;
    <a id="L49"></a>Esterror  int32;
    <a id="L50"></a>Status    int32;
    <a id="L51"></a>Constant  int32;
    <a id="L52"></a>Precision int32;
    <a id="L53"></a>Tolerance int32;
    <a id="L54"></a>Time      Timeval;
    <a id="L55"></a>Tick      int32;
    <a id="L56"></a>Ppsfreq   int32;
    <a id="L57"></a>Jitter    int32;
    <a id="L58"></a>Shift     int32;
    <a id="L59"></a>Stabil    int32;
    <a id="L60"></a>Jitcnt    int32;
    <a id="L61"></a>Calcnt    int32;
    <a id="L62"></a>Errcnt    int32;
    <a id="L63"></a>Stbcnt    int32;
    <a id="L64"></a>Pad0      int32;
    <a id="L65"></a>Pad1      int32;
    <a id="L66"></a>Pad2      int32;
    <a id="L67"></a>Pad3      int32;
    <a id="L68"></a>Pad4      int32;
    <a id="L69"></a>Pad5      int32;
    <a id="L70"></a>Pad6      int32;
    <a id="L71"></a>Pad7      int32;
    <a id="L72"></a>Pad8      int32;
    <a id="L73"></a>Pad9      int32;
    <a id="L74"></a>Pad10     int32;
    <a id="L75"></a>Pad11     int32;
<a id="L76"></a>}

<a id="L78"></a>type Time_t int32

<a id="L80"></a>type Tms struct {
    <a id="L81"></a>Utime  int32;
    <a id="L82"></a>Stime  int32;
    <a id="L83"></a>Cutime int32;
    <a id="L84"></a>Cstime int32;
<a id="L85"></a>}

<a id="L87"></a>type Utimbuf struct {
    <a id="L88"></a>Actime  int32;
    <a id="L89"></a>Modtime int32;
<a id="L90"></a>}

<a id="L92"></a>type Rusage struct {
    <a id="L93"></a>Utime    Timeval;
    <a id="L94"></a>Stime    Timeval;
    <a id="L95"></a>Maxrss   int32;
    <a id="L96"></a>Ixrss    int32;
    <a id="L97"></a>Idrss    int32;
    <a id="L98"></a>Isrss    int32;
    <a id="L99"></a>Minflt   int32;
    <a id="L100"></a>Majflt   int32;
    <a id="L101"></a>Nswap    int32;
    <a id="L102"></a>Inblock  int32;
    <a id="L103"></a>Oublock  int32;
    <a id="L104"></a>Msgsnd   int32;
    <a id="L105"></a>Msgrcv   int32;
    <a id="L106"></a>Nsignals int32;
    <a id="L107"></a>Nvcsw    int32;
    <a id="L108"></a>Nivcsw   int32;
<a id="L109"></a>}

<a id="L111"></a>type Rlimit struct {
    <a id="L112"></a>Cur uint64;
    <a id="L113"></a>Max uint64;
<a id="L114"></a>}

<a id="L116"></a>type _Gid_t uint32

<a id="L118"></a>type Stat_t struct {
    <a id="L119"></a>Dev       uint64;
    <a id="L120"></a>X__pad1   uint16;
    <a id="L121"></a>Pad0      [2]byte;
    <a id="L122"></a>X__st_ino uint32;
    <a id="L123"></a>Mode      uint32;
    <a id="L124"></a>Nlink     uint32;
    <a id="L125"></a>Uid       uint32;
    <a id="L126"></a>Gid       uint32;
    <a id="L127"></a>Rdev      uint64;
    <a id="L128"></a>X__pad2   uint16;
    <a id="L129"></a>Pad1      [2]byte;
    <a id="L130"></a>Size      int64;
    <a id="L131"></a>Blksize   int32;
    <a id="L132"></a>Blocks    int64;
    <a id="L133"></a>Atim      Timespec;
    <a id="L134"></a>Mtim      Timespec;
    <a id="L135"></a>Ctim      Timespec;
    <a id="L136"></a>Ino       uint64;
<a id="L137"></a>}

<a id="L139"></a>type Statfs_t struct {
    <a id="L140"></a>Type    int32;
    <a id="L141"></a>Bsize   int32;
    <a id="L142"></a>Blocks  uint64;
    <a id="L143"></a>Bfree   uint64;
    <a id="L144"></a>Bavail  uint64;
    <a id="L145"></a>Files   uint64;
    <a id="L146"></a>Ffree   uint64;
    <a id="L147"></a>Fsid    [8]byte; <span class="comment">/* __fsid_t */</span>
    <a id="L148"></a>Namelen int32;
    <a id="L149"></a>Frsize  int32;
    <a id="L150"></a>Spare   [5]int32;
<a id="L151"></a>}

<a id="L153"></a>type Dirent struct {
    <a id="L154"></a>Ino    uint64;
    <a id="L155"></a>Off    int64;
    <a id="L156"></a>Reclen uint16;
    <a id="L157"></a>Type   uint8;
    <a id="L158"></a>Name   [256]int8;
    <a id="L159"></a>Pad0   [1]byte;
<a id="L160"></a>}

<a id="L162"></a>type RawSockaddrInet4 struct {
    <a id="L163"></a>Family uint16;
    <a id="L164"></a>Port   uint16;
    <a id="L165"></a>Addr   [4]byte; <span class="comment">/* in_addr */</span>
    <a id="L166"></a>Zero   [8]uint8;
<a id="L167"></a>}

<a id="L169"></a>type RawSockaddrInet6 struct {
    <a id="L170"></a>Family   uint16;
    <a id="L171"></a>Port     uint16;
    <a id="L172"></a>Flowinfo uint32;
    <a id="L173"></a>Addr     [16]byte; <span class="comment">/* in6_addr */</span>
    <a id="L174"></a>Scope_id uint32;
<a id="L175"></a>}

<a id="L177"></a>type RawSockaddrUnix struct {
    <a id="L178"></a>Family uint16;
    <a id="L179"></a>Path   [108]int8;
<a id="L180"></a>}

<a id="L182"></a>type RawSockaddr struct {
    <a id="L183"></a>Family uint16;
    <a id="L184"></a>Data   [14]int8;
<a id="L185"></a>}

<a id="L187"></a>type RawSockaddrAny struct {
    <a id="L188"></a>Addr RawSockaddr;
    <a id="L189"></a>Pad  [96]int8;
<a id="L190"></a>}

<a id="L192"></a>type _Socklen uint32

<a id="L194"></a>type Linger struct {
    <a id="L195"></a>Onoff  int32;
    <a id="L196"></a>Linger int32;
<a id="L197"></a>}

<a id="L199"></a>type Iovec struct {
    <a id="L200"></a>Base *byte;
    <a id="L201"></a>Len  uint32;
<a id="L202"></a>}

<a id="L204"></a>type Msghdr struct {
    <a id="L205"></a>Name       *byte;
    <a id="L206"></a>Namelen    uint32;
    <a id="L207"></a>Iov        *Iovec;
    <a id="L208"></a>Iovlen     uint32;
    <a id="L209"></a>Control    *byte;
    <a id="L210"></a>Controllen uint32;
    <a id="L211"></a>Flags      int32;
<a id="L212"></a>}

<a id="L214"></a>type Cmsghdr struct {
    <a id="L215"></a>Len   uint32;
    <a id="L216"></a>Level int32;
    <a id="L217"></a>Type  int32;
<a id="L218"></a>}

<a id="L220"></a>type PtraceRegs struct {
    <a id="L221"></a>Ebx      int32;
    <a id="L222"></a>Ecx      int32;
    <a id="L223"></a>Edx      int32;
    <a id="L224"></a>Esi      int32;
    <a id="L225"></a>Edi      int32;
    <a id="L226"></a>Ebp      int32;
    <a id="L227"></a>Eax      int32;
    <a id="L228"></a>Ds       uint16;
    <a id="L229"></a>X__ds    uint16;
    <a id="L230"></a>Es       uint16;
    <a id="L231"></a>X__es    uint16;
    <a id="L232"></a>Fs       uint16;
    <a id="L233"></a>X__fs    uint16;
    <a id="L234"></a>Gs       uint16;
    <a id="L235"></a>X__gs    uint16;
    <a id="L236"></a>Orig_eax int32;
    <a id="L237"></a>Eip      int32;
    <a id="L238"></a>Cs       uint16;
    <a id="L239"></a>X__cs    uint16;
    <a id="L240"></a>Eflags   int32;
    <a id="L241"></a>Esp      int32;
    <a id="L242"></a>Ss       uint16;
    <a id="L243"></a>X__ss    uint16;
<a id="L244"></a>}

<a id="L246"></a>type FdSet struct {
    <a id="L247"></a>Bits [32]int32;
<a id="L248"></a>}

<a id="L250"></a>type Sysinfo_t struct {
    <a id="L251"></a>Uptime    int32;
    <a id="L252"></a>Loads     [3]uint32;
    <a id="L253"></a>Totalram  uint32;
    <a id="L254"></a>Freeram   uint32;
    <a id="L255"></a>Sharedram uint32;
    <a id="L256"></a>Bufferram uint32;
    <a id="L257"></a>Totalswap uint32;
    <a id="L258"></a>Freeswap  uint32;
    <a id="L259"></a>Procs     uint16;
    <a id="L260"></a>Pad       uint16;
    <a id="L261"></a>Totalhigh uint32;
    <a id="L262"></a>Freehigh  uint32;
    <a id="L263"></a>Unit      uint32;
    <a id="L264"></a>X_f       [8]int8;
<a id="L265"></a>}

<a id="L267"></a>type Utsname struct {
    <a id="L268"></a>Sysname    [65]int8;
    <a id="L269"></a>Nodename   [65]int8;
    <a id="L270"></a>Release    [65]int8;
    <a id="L271"></a>Version    [65]int8;
    <a id="L272"></a>Machine    [65]int8;
    <a id="L273"></a>Domainname [65]int8;
<a id="L274"></a>}

<a id="L276"></a>type Ustat_t struct {
    <a id="L277"></a>Tfree  int32;
    <a id="L278"></a>Tinode uint32;
    <a id="L279"></a>Fname  [6]int8;
    <a id="L280"></a>Fpack  [6]int8;
<a id="L281"></a>}

<a id="L283"></a>type EpollEvent struct {
    <a id="L284"></a>Events uint32;
    <a id="L285"></a>Fd     int32;
    <a id="L286"></a>Pad    int32;
<a id="L287"></a>}
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
