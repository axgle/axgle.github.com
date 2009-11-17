<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/dwarf/const.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/debug/dwarf/const.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Constants</span>

<a id="L7"></a>package dwarf

<a id="L9"></a>import &#34;strconv&#34;

<a id="L11"></a><span class="comment">// An Attr identifies the attribute type in a DWARF Entry&#39;s Field.</span>
<a id="L12"></a>type Attr uint32

<a id="L14"></a>const (
    <a id="L15"></a>AttrSibling        Attr = 0x01;
    <a id="L16"></a>AttrLocation       Attr = 0x02;
    <a id="L17"></a>AttrName           Attr = 0x03;
    <a id="L18"></a>AttrOrdering       Attr = 0x09;
    <a id="L19"></a>AttrByteSize       Attr = 0x0B;
    <a id="L20"></a>AttrBitOffset      Attr = 0x0C;
    <a id="L21"></a>AttrBitSize        Attr = 0x0D;
    <a id="L22"></a>AttrStmtList       Attr = 0x10;
    <a id="L23"></a>AttrLowpc          Attr = 0x11;
    <a id="L24"></a>AttrHighpc         Attr = 0x12;
    <a id="L25"></a>AttrLanguage       Attr = 0x13;
    <a id="L26"></a>AttrDiscr          Attr = 0x15;
    <a id="L27"></a>AttrDiscrValue     Attr = 0x16;
    <a id="L28"></a>AttrVisibility     Attr = 0x17;
    <a id="L29"></a>AttrImport         Attr = 0x18;
    <a id="L30"></a>AttrStringLength   Attr = 0x19;
    <a id="L31"></a>AttrCommonRef      Attr = 0x1A;
    <a id="L32"></a>AttrCompDir        Attr = 0x1B;
    <a id="L33"></a>AttrConstValue     Attr = 0x1C;
    <a id="L34"></a>AttrContainingType Attr = 0x1D;
    <a id="L35"></a>AttrDefaultValue   Attr = 0x1E;
    <a id="L36"></a>AttrInline         Attr = 0x20;
    <a id="L37"></a>AttrIsOptional     Attr = 0x21;
    <a id="L38"></a>AttrLowerBound     Attr = 0x22;
    <a id="L39"></a>AttrProducer       Attr = 0x25;
    <a id="L40"></a>AttrPrototyped     Attr = 0x27;
    <a id="L41"></a>AttrReturnAddr     Attr = 0x2A;
    <a id="L42"></a>AttrStartScope     Attr = 0x2C;
    <a id="L43"></a>AttrStrideSize     Attr = 0x2E;
    <a id="L44"></a>AttrUpperBound     Attr = 0x2F;
    <a id="L45"></a>AttrAbstractOrigin Attr = 0x31;
    <a id="L46"></a>AttrAccessibility  Attr = 0x32;
    <a id="L47"></a>AttrAddrClass      Attr = 0x33;
    <a id="L48"></a>AttrArtificial     Attr = 0x34;
    <a id="L49"></a>AttrBaseTypes      Attr = 0x35;
    <a id="L50"></a>AttrCalling        Attr = 0x36;
    <a id="L51"></a>AttrCount          Attr = 0x37;
    <a id="L52"></a>AttrDataMemberLoc  Attr = 0x38;
    <a id="L53"></a>AttrDeclColumn     Attr = 0x39;
    <a id="L54"></a>AttrDeclFile       Attr = 0x3A;
    <a id="L55"></a>AttrDeclLine       Attr = 0x3B;
    <a id="L56"></a>AttrDeclaration    Attr = 0x3C;
    <a id="L57"></a>AttrDiscrList      Attr = 0x3D;
    <a id="L58"></a>AttrEncoding       Attr = 0x3E;
    <a id="L59"></a>AttrExternal       Attr = 0x3F;
    <a id="L60"></a>AttrFrameBase      Attr = 0x40;
    <a id="L61"></a>AttrFriend         Attr = 0x41;
    <a id="L62"></a>AttrIdentifierCase Attr = 0x42;
    <a id="L63"></a>AttrMacroInfo      Attr = 0x43;
    <a id="L64"></a>AttrNamelistItem   Attr = 0x44;
    <a id="L65"></a>AttrPriority       Attr = 0x45;
    <a id="L66"></a>AttrSegment        Attr = 0x46;
    <a id="L67"></a>AttrSpecification  Attr = 0x47;
    <a id="L68"></a>AttrStaticLink     Attr = 0x48;
    <a id="L69"></a>AttrType           Attr = 0x49;
    <a id="L70"></a>AttrUseLocation    Attr = 0x4A;
    <a id="L71"></a>AttrVarParam       Attr = 0x4B;
    <a id="L72"></a>AttrVirtuality     Attr = 0x4C;
    <a id="L73"></a>AttrVtableElemLoc  Attr = 0x4D;
    <a id="L74"></a>AttrAllocated      Attr = 0x4E;
    <a id="L75"></a>AttrAssociated     Attr = 0x4F;
    <a id="L76"></a>AttrDataLocation   Attr = 0x50;
    <a id="L77"></a>AttrStride         Attr = 0x51;
    <a id="L78"></a>AttrEntrypc        Attr = 0x52;
    <a id="L79"></a>AttrUseUTF8        Attr = 0x53;
    <a id="L80"></a>AttrExtension      Attr = 0x54;
    <a id="L81"></a>AttrRanges         Attr = 0x55;
    <a id="L82"></a>AttrTrampoline     Attr = 0x56;
    <a id="L83"></a>AttrCallColumn     Attr = 0x57;
    <a id="L84"></a>AttrCallFile       Attr = 0x58;
    <a id="L85"></a>AttrCallLine       Attr = 0x59;
    <a id="L86"></a>AttrDescription    Attr = 0x5A;
<a id="L87"></a>)

<a id="L89"></a>var attrNames = [...]string{
    <a id="L90"></a>AttrSibling: &#34;Sibling&#34;,
    <a id="L91"></a>AttrLocation: &#34;Location&#34;,
    <a id="L92"></a>AttrName: &#34;Name&#34;,
    <a id="L93"></a>AttrOrdering: &#34;Ordering&#34;,
    <a id="L94"></a>AttrByteSize: &#34;ByteSize&#34;,
    <a id="L95"></a>AttrBitOffset: &#34;BitOffset&#34;,
    <a id="L96"></a>AttrBitSize: &#34;BitSize&#34;,
    <a id="L97"></a>AttrStmtList: &#34;StmtList&#34;,
    <a id="L98"></a>AttrLowpc: &#34;Lowpc&#34;,
    <a id="L99"></a>AttrHighpc: &#34;Highpc&#34;,
    <a id="L100"></a>AttrLanguage: &#34;Language&#34;,
    <a id="L101"></a>AttrDiscr: &#34;Discr&#34;,
    <a id="L102"></a>AttrDiscrValue: &#34;DiscrValue&#34;,
    <a id="L103"></a>AttrVisibility: &#34;Visibility&#34;,
    <a id="L104"></a>AttrImport: &#34;Import&#34;,
    <a id="L105"></a>AttrStringLength: &#34;StringLength&#34;,
    <a id="L106"></a>AttrCommonRef: &#34;CommonRef&#34;,
    <a id="L107"></a>AttrCompDir: &#34;CompDir&#34;,
    <a id="L108"></a>AttrConstValue: &#34;ConstValue&#34;,
    <a id="L109"></a>AttrContainingType: &#34;ContainingType&#34;,
    <a id="L110"></a>AttrDefaultValue: &#34;DefaultValue&#34;,
    <a id="L111"></a>AttrInline: &#34;Inline&#34;,
    <a id="L112"></a>AttrIsOptional: &#34;IsOptional&#34;,
    <a id="L113"></a>AttrLowerBound: &#34;LowerBound&#34;,
    <a id="L114"></a>AttrProducer: &#34;Producer&#34;,
    <a id="L115"></a>AttrPrototyped: &#34;Prototyped&#34;,
    <a id="L116"></a>AttrReturnAddr: &#34;ReturnAddr&#34;,
    <a id="L117"></a>AttrStartScope: &#34;StartScope&#34;,
    <a id="L118"></a>AttrStrideSize: &#34;StrideSize&#34;,
    <a id="L119"></a>AttrUpperBound: &#34;UpperBound&#34;,
    <a id="L120"></a>AttrAbstractOrigin: &#34;AbstractOrigin&#34;,
    <a id="L121"></a>AttrAccessibility: &#34;Accessibility&#34;,
    <a id="L122"></a>AttrAddrClass: &#34;AddrClass&#34;,
    <a id="L123"></a>AttrArtificial: &#34;Artificial&#34;,
    <a id="L124"></a>AttrBaseTypes: &#34;BaseTypes&#34;,
    <a id="L125"></a>AttrCalling: &#34;Calling&#34;,
    <a id="L126"></a>AttrCount: &#34;Count&#34;,
    <a id="L127"></a>AttrDataMemberLoc: &#34;DataMemberLoc&#34;,
    <a id="L128"></a>AttrDeclColumn: &#34;DeclColumn&#34;,
    <a id="L129"></a>AttrDeclFile: &#34;DeclFile&#34;,
    <a id="L130"></a>AttrDeclLine: &#34;DeclLine&#34;,
    <a id="L131"></a>AttrDeclaration: &#34;Declaration&#34;,
    <a id="L132"></a>AttrDiscrList: &#34;DiscrList&#34;,
    <a id="L133"></a>AttrEncoding: &#34;Encoding&#34;,
    <a id="L134"></a>AttrExternal: &#34;External&#34;,
    <a id="L135"></a>AttrFrameBase: &#34;FrameBase&#34;,
    <a id="L136"></a>AttrFriend: &#34;Friend&#34;,
    <a id="L137"></a>AttrIdentifierCase: &#34;IdentifierCase&#34;,
    <a id="L138"></a>AttrMacroInfo: &#34;MacroInfo&#34;,
    <a id="L139"></a>AttrNamelistItem: &#34;NamelistItem&#34;,
    <a id="L140"></a>AttrPriority: &#34;Priority&#34;,
    <a id="L141"></a>AttrSegment: &#34;Segment&#34;,
    <a id="L142"></a>AttrSpecification: &#34;Specification&#34;,
    <a id="L143"></a>AttrStaticLink: &#34;StaticLink&#34;,
    <a id="L144"></a>AttrType: &#34;Type&#34;,
    <a id="L145"></a>AttrUseLocation: &#34;UseLocation&#34;,
    <a id="L146"></a>AttrVarParam: &#34;VarParam&#34;,
    <a id="L147"></a>AttrVirtuality: &#34;Virtuality&#34;,
    <a id="L148"></a>AttrVtableElemLoc: &#34;VtableElemLoc&#34;,
    <a id="L149"></a>AttrAllocated: &#34;Allocated&#34;,
    <a id="L150"></a>AttrAssociated: &#34;Associated&#34;,
    <a id="L151"></a>AttrDataLocation: &#34;DataLocation&#34;,
    <a id="L152"></a>AttrStride: &#34;Stride&#34;,
    <a id="L153"></a>AttrEntrypc: &#34;Entrypc&#34;,
    <a id="L154"></a>AttrUseUTF8: &#34;UseUTF8&#34;,
    <a id="L155"></a>AttrExtension: &#34;Extension&#34;,
    <a id="L156"></a>AttrRanges: &#34;Ranges&#34;,
    <a id="L157"></a>AttrTrampoline: &#34;Trampoline&#34;,
    <a id="L158"></a>AttrCallColumn: &#34;CallColumn&#34;,
    <a id="L159"></a>AttrCallFile: &#34;CallFile&#34;,
    <a id="L160"></a>AttrCallLine: &#34;CallLine&#34;,
    <a id="L161"></a>AttrDescription: &#34;Description&#34;,
<a id="L162"></a>}

<a id="L164"></a>func (a Attr) String() string {
    <a id="L165"></a>if int(a) &lt; len(attrNames) {
        <a id="L166"></a>s := attrNames[a];
        <a id="L167"></a>if s != &#34;&#34; {
            <a id="L168"></a>return s
        <a id="L169"></a>}
    <a id="L170"></a>}
    <a id="L171"></a>return strconv.Itoa(int(a));
<a id="L172"></a>}

<a id="L174"></a>func (a Attr) GoString() string {
    <a id="L175"></a>if int(a) &lt; len(attrNames) {
        <a id="L176"></a>s := attrNames[a];
        <a id="L177"></a>if s != &#34;&#34; {
            <a id="L178"></a>return &#34;dwarf.Attr&#34; + s
        <a id="L179"></a>}
    <a id="L180"></a>}
    <a id="L181"></a>return &#34;dwarf.Attr(&#34; + strconv.Itoa64(int64(a)) + &#34;)&#34;;
<a id="L182"></a>}

<a id="L184"></a><span class="comment">// A format is a DWARF data encoding format.</span>
<a id="L185"></a>type format uint32

<a id="L187"></a>const (
    <a id="L188"></a><span class="comment">// value formats</span>
    <a id="L189"></a>formAddr        format = 0x01;
    <a id="L190"></a>formDwarfBlock2 format = 0x03;
    <a id="L191"></a>formDwarfBlock4 format = 0x04;
    <a id="L192"></a>formData2       format = 0x05;
    <a id="L193"></a>formData4       format = 0x06;
    <a id="L194"></a>formData8       format = 0x07;
    <a id="L195"></a>formString      format = 0x08;
    <a id="L196"></a>formDwarfBlock  format = 0x09;
    <a id="L197"></a>formDwarfBlock1 format = 0x0A;
    <a id="L198"></a>formData1       format = 0x0B;
    <a id="L199"></a>formFlag        format = 0x0C;
    <a id="L200"></a>formSdata       format = 0x0D;
    <a id="L201"></a>formStrp        format = 0x0E;
    <a id="L202"></a>formUdata       format = 0x0F;
    <a id="L203"></a>formRefAddr     format = 0x10;
    <a id="L204"></a>formRef1        format = 0x11;
    <a id="L205"></a>formRef2        format = 0x12;
    <a id="L206"></a>formRef4        format = 0x13;
    <a id="L207"></a>formRef8        format = 0x14;
    <a id="L208"></a>formRefUdata    format = 0x15;
    <a id="L209"></a>formIndirect    format = 0x16;
<a id="L210"></a>)

<a id="L212"></a><span class="comment">// A Tag is the classification (the type) of an Entry.</span>
<a id="L213"></a>type Tag uint32

<a id="L215"></a>const (
    <a id="L216"></a>TagArrayType              Tag = 0x01;
    <a id="L217"></a>TagClassType              Tag = 0x02;
    <a id="L218"></a>TagEntryPoint             Tag = 0x03;
    <a id="L219"></a>TagEnumerationType        Tag = 0x04;
    <a id="L220"></a>TagFormalParameter        Tag = 0x05;
    <a id="L221"></a>TagImportedDeclaration    Tag = 0x08;
    <a id="L222"></a>TagLabel                  Tag = 0x0A;
    <a id="L223"></a>TagLexDwarfBlock          Tag = 0x0B;
    <a id="L224"></a>TagMember                 Tag = 0x0D;
    <a id="L225"></a>TagPointerType            Tag = 0x0F;
    <a id="L226"></a>TagReferenceType          Tag = 0x10;
    <a id="L227"></a>TagCompileUnit            Tag = 0x11;
    <a id="L228"></a>TagStringType             Tag = 0x12;
    <a id="L229"></a>TagStructType             Tag = 0x13;
    <a id="L230"></a>TagSubroutineType         Tag = 0x15;
    <a id="L231"></a>TagTypedef                Tag = 0x16;
    <a id="L232"></a>TagUnionType              Tag = 0x17;
    <a id="L233"></a>TagUnspecifiedParameters  Tag = 0x18;
    <a id="L234"></a>TagVariant                Tag = 0x19;
    <a id="L235"></a>TagCommonDwarfBlock       Tag = 0x1A;
    <a id="L236"></a>TagCommonInclusion        Tag = 0x1B;
    <a id="L237"></a>TagInheritance            Tag = 0x1C;
    <a id="L238"></a>TagInlinedSubroutine      Tag = 0x1D;
    <a id="L239"></a>TagModule                 Tag = 0x1E;
    <a id="L240"></a>TagPtrToMemberType        Tag = 0x1F;
    <a id="L241"></a>TagSetType                Tag = 0x20;
    <a id="L242"></a>TagSubrangeType           Tag = 0x21;
    <a id="L243"></a>TagWithStmt               Tag = 0x22;
    <a id="L244"></a>TagAccessDeclaration      Tag = 0x23;
    <a id="L245"></a>TagBaseType               Tag = 0x24;
    <a id="L246"></a>TagCatchDwarfBlock        Tag = 0x25;
    <a id="L247"></a>TagConstType              Tag = 0x26;
    <a id="L248"></a>TagConstant               Tag = 0x27;
    <a id="L249"></a>TagEnumerator             Tag = 0x28;
    <a id="L250"></a>TagFileType               Tag = 0x29;
    <a id="L251"></a>TagFriend                 Tag = 0x2A;
    <a id="L252"></a>TagNamelist               Tag = 0x2B;
    <a id="L253"></a>TagNamelistItem           Tag = 0x2C;
    <a id="L254"></a>TagPackedType             Tag = 0x2D;
    <a id="L255"></a>TagSubprogram             Tag = 0x2E;
    <a id="L256"></a>TagTemplateTypeParameter  Tag = 0x2F;
    <a id="L257"></a>TagTemplateValueParameter Tag = 0x30;
    <a id="L258"></a>TagThrownType             Tag = 0x31;
    <a id="L259"></a>TagTryDwarfBlock          Tag = 0x32;
    <a id="L260"></a>TagVariantPart            Tag = 0x33;
    <a id="L261"></a>TagVariable               Tag = 0x34;
    <a id="L262"></a>TagVolatileType           Tag = 0x35;
    <a id="L263"></a>TagDwarfProcedure         Tag = 0x36;
    <a id="L264"></a>TagRestrictType           Tag = 0x37;
    <a id="L265"></a>TagInterfaceType          Tag = 0x38;
    <a id="L266"></a>TagNamespace              Tag = 0x39;
    <a id="L267"></a>TagImportedModule         Tag = 0x3A;
    <a id="L268"></a>TagUnspecifiedType        Tag = 0x3B;
    <a id="L269"></a>TagPartialUnit            Tag = 0x3C;
    <a id="L270"></a>TagImportedUnit           Tag = 0x3D;
    <a id="L271"></a>TagMutableType            Tag = 0x3E;
<a id="L272"></a>)

<a id="L274"></a>var tagNames = [...]string{
    <a id="L275"></a>TagArrayType: &#34;ArrayType&#34;,
    <a id="L276"></a>TagClassType: &#34;ClassType&#34;,
    <a id="L277"></a>TagEntryPoint: &#34;EntryPoint&#34;,
    <a id="L278"></a>TagEnumerationType: &#34;EnumerationType&#34;,
    <a id="L279"></a>TagFormalParameter: &#34;FormalParameter&#34;,
    <a id="L280"></a>TagImportedDeclaration: &#34;ImportedDeclaration&#34;,
    <a id="L281"></a>TagLabel: &#34;Label&#34;,
    <a id="L282"></a>TagLexDwarfBlock: &#34;LexDwarfBlock&#34;,
    <a id="L283"></a>TagMember: &#34;Member&#34;,
    <a id="L284"></a>TagPointerType: &#34;PointerType&#34;,
    <a id="L285"></a>TagReferenceType: &#34;ReferenceType&#34;,
    <a id="L286"></a>TagCompileUnit: &#34;CompileUnit&#34;,
    <a id="L287"></a>TagStringType: &#34;StringType&#34;,
    <a id="L288"></a>TagStructType: &#34;StructType&#34;,
    <a id="L289"></a>TagSubroutineType: &#34;SubroutineType&#34;,
    <a id="L290"></a>TagTypedef: &#34;Typedef&#34;,
    <a id="L291"></a>TagUnionType: &#34;UnionType&#34;,
    <a id="L292"></a>TagUnspecifiedParameters: &#34;UnspecifiedParameters&#34;,
    <a id="L293"></a>TagVariant: &#34;Variant&#34;,
    <a id="L294"></a>TagCommonDwarfBlock: &#34;CommonDwarfBlock&#34;,
    <a id="L295"></a>TagCommonInclusion: &#34;CommonInclusion&#34;,
    <a id="L296"></a>TagInheritance: &#34;Inheritance&#34;,
    <a id="L297"></a>TagInlinedSubroutine: &#34;InlinedSubroutine&#34;,
    <a id="L298"></a>TagModule: &#34;Module&#34;,
    <a id="L299"></a>TagPtrToMemberType: &#34;PtrToMemberType&#34;,
    <a id="L300"></a>TagSetType: &#34;SetType&#34;,
    <a id="L301"></a>TagSubrangeType: &#34;SubrangeType&#34;,
    <a id="L302"></a>TagWithStmt: &#34;WithStmt&#34;,
    <a id="L303"></a>TagAccessDeclaration: &#34;AccessDeclaration&#34;,
    <a id="L304"></a>TagBaseType: &#34;BaseType&#34;,
    <a id="L305"></a>TagCatchDwarfBlock: &#34;CatchDwarfBlock&#34;,
    <a id="L306"></a>TagConstType: &#34;ConstType&#34;,
    <a id="L307"></a>TagConstant: &#34;Constant&#34;,
    <a id="L308"></a>TagEnumerator: &#34;Enumerator&#34;,
    <a id="L309"></a>TagFileType: &#34;FileType&#34;,
    <a id="L310"></a>TagFriend: &#34;Friend&#34;,
    <a id="L311"></a>TagNamelist: &#34;Namelist&#34;,
    <a id="L312"></a>TagNamelistItem: &#34;NamelistItem&#34;,
    <a id="L313"></a>TagPackedType: &#34;PackedType&#34;,
    <a id="L314"></a>TagSubprogram: &#34;Subprogram&#34;,
    <a id="L315"></a>TagTemplateTypeParameter: &#34;TemplateTypeParameter&#34;,
    <a id="L316"></a>TagTemplateValueParameter: &#34;TemplateValueParameter&#34;,
    <a id="L317"></a>TagThrownType: &#34;ThrownType&#34;,
    <a id="L318"></a>TagTryDwarfBlock: &#34;TryDwarfBlock&#34;,
    <a id="L319"></a>TagVariantPart: &#34;VariantPart&#34;,
    <a id="L320"></a>TagVariable: &#34;Variable&#34;,
    <a id="L321"></a>TagVolatileType: &#34;VolatileType&#34;,
    <a id="L322"></a>TagDwarfProcedure: &#34;DwarfProcedure&#34;,
    <a id="L323"></a>TagRestrictType: &#34;RestrictType&#34;,
    <a id="L324"></a>TagInterfaceType: &#34;InterfaceType&#34;,
    <a id="L325"></a>TagNamespace: &#34;Namespace&#34;,
    <a id="L326"></a>TagImportedModule: &#34;ImportedModule&#34;,
    <a id="L327"></a>TagUnspecifiedType: &#34;UnspecifiedType&#34;,
    <a id="L328"></a>TagPartialUnit: &#34;PartialUnit&#34;,
    <a id="L329"></a>TagImportedUnit: &#34;ImportedUnit&#34;,
    <a id="L330"></a>TagMutableType: &#34;MutableType&#34;,
<a id="L331"></a>}

<a id="L333"></a>func (t Tag) String() string {
    <a id="L334"></a>if int(t) &lt; len(tagNames) {
        <a id="L335"></a>s := tagNames[t];
        <a id="L336"></a>if s != &#34;&#34; {
            <a id="L337"></a>return s
        <a id="L338"></a>}
    <a id="L339"></a>}
    <a id="L340"></a>return strconv.Itoa(int(t));
<a id="L341"></a>}

<a id="L343"></a>func (t Tag) GoString() string {
    <a id="L344"></a>if int(t) &lt; len(tagNames) {
        <a id="L345"></a>s := tagNames[t];
        <a id="L346"></a>if s != &#34;&#34; {
            <a id="L347"></a>return &#34;dwarf.Tag&#34; + s
        <a id="L348"></a>}
    <a id="L349"></a>}
    <a id="L350"></a>return &#34;dwarf.Tag(&#34; + strconv.Itoa64(int64(t)) + &#34;)&#34;;
<a id="L351"></a>}

<a id="L353"></a><span class="comment">// Location expression operators.</span>
<a id="L354"></a><span class="comment">// The debug info encodes value locations like 8(R3)</span>
<a id="L355"></a><span class="comment">// as a sequence of these op codes.</span>
<a id="L356"></a><span class="comment">// This package does not implement full expressions;</span>
<a id="L357"></a><span class="comment">// the opPlusUconst operator is expected by the type parser.</span>
<a id="L358"></a>const (
    <a id="L359"></a>opAddr       = 0x03; <span class="comment">/* 1 op, const addr */</span>
    <a id="L360"></a>opDeref      = 0x06;
    <a id="L361"></a>opConst1u    = 0x08; <span class="comment">/* 1 op, 1 byte const */</span>
    <a id="L362"></a>opConst1s    = 0x09; <span class="comment">/*	&#34; signed */</span>
    <a id="L363"></a>opConst2u    = 0x0A; <span class="comment">/* 1 op, 2 byte const  */</span>
    <a id="L364"></a>opConst2s    = 0x0B; <span class="comment">/*	&#34; signed */</span>
    <a id="L365"></a>opConst4u    = 0x0C; <span class="comment">/* 1 op, 4 byte const */</span>
    <a id="L366"></a>opConst4s    = 0x0D; <span class="comment">/*	&#34; signed */</span>
    <a id="L367"></a>opConst8u    = 0x0E; <span class="comment">/* 1 op, 8 byte const */</span>
    <a id="L368"></a>opConst8s    = 0x0F; <span class="comment">/*	&#34; signed */</span>
    <a id="L369"></a>opConstu     = 0x10; <span class="comment">/* 1 op, LEB128 const */</span>
    <a id="L370"></a>opConsts     = 0x11; <span class="comment">/*	&#34; signed */</span>
    <a id="L371"></a>opDup        = 0x12;
    <a id="L372"></a>opDrop       = 0x13;
    <a id="L373"></a>opOver       = 0x14;
    <a id="L374"></a>opPick       = 0x15; <span class="comment">/* 1 op, 1 byte stack index */</span>
    <a id="L375"></a>opSwap       = 0x16;
    <a id="L376"></a>opRot        = 0x17;
    <a id="L377"></a>opXderef     = 0x18;
    <a id="L378"></a>opAbs        = 0x19;
    <a id="L379"></a>opAnd        = 0x1A;
    <a id="L380"></a>opDiv        = 0x1B;
    <a id="L381"></a>opMinus      = 0x1C;
    <a id="L382"></a>opMod        = 0x1D;
    <a id="L383"></a>opMul        = 0x1E;
    <a id="L384"></a>opNeg        = 0x1F;
    <a id="L385"></a>opNot        = 0x20;
    <a id="L386"></a>opOr         = 0x21;
    <a id="L387"></a>opPlus       = 0x22;
    <a id="L388"></a>opPlusUconst = 0x23; <span class="comment">/* 1 op, ULEB128 addend */</span>
    <a id="L389"></a>opShl        = 0x24;
    <a id="L390"></a>opShr        = 0x25;
    <a id="L391"></a>opShra       = 0x26;
    <a id="L392"></a>opXor        = 0x27;
    <a id="L393"></a>opSkip       = 0x2F; <span class="comment">/* 1 op, signed 2-byte constant */</span>
    <a id="L394"></a>opBra        = 0x28; <span class="comment">/* 1 op, signed 2-byte constant */</span>
    <a id="L395"></a>opEq         = 0x29;
    <a id="L396"></a>opGe         = 0x2A;
    <a id="L397"></a>opGt         = 0x2B;
    <a id="L398"></a>opLe         = 0x2C;
    <a id="L399"></a>opLt         = 0x2D;
    <a id="L400"></a>opNe         = 0x2E;
    <a id="L401"></a>opLit0       = 0x30;
    <a id="L402"></a><span class="comment">/* OpLitN = OpLit0 + N for N = 0..31 */</span>
    <a id="L403"></a>opReg0 = 0x50;
    <a id="L404"></a><span class="comment">/* OpRegN = OpReg0 + N for N = 0..31 */</span>
    <a id="L405"></a>opBreg0 = 0x70; <span class="comment">/* 1 op, signed LEB128 constant */</span>
    <a id="L406"></a><span class="comment">/* OpBregN = OpBreg0 + N for N = 0..31 */</span>
    <a id="L407"></a>opRegx       = 0x90; <span class="comment">/* 1 op, ULEB128 register */</span>
    <a id="L408"></a>opFbreg      = 0x91; <span class="comment">/* 1 op, SLEB128 offset */</span>
    <a id="L409"></a>opBregx      = 0x92; <span class="comment">/* 2 op, ULEB128 reg; SLEB128 off */</span>
    <a id="L410"></a>opPiece      = 0x93; <span class="comment">/* 1 op, ULEB128 size of piece */</span>
    <a id="L411"></a>opDerefSize  = 0x94; <span class="comment">/* 1-byte size of data retrieved */</span>
    <a id="L412"></a>opXderefSize = 0x95; <span class="comment">/* 1-byte size of data retrieved */</span>
    <a id="L413"></a>opNop        = 0x96;
    <a id="L414"></a><span class="comment">/* next four new in Dwarf v3 */</span>
    <a id="L415"></a>opPushObjAddr = 0x97;
    <a id="L416"></a>opCall2       = 0x98; <span class="comment">/* 2-byte offset of DIE */</span>
    <a id="L417"></a>opCall4       = 0x99; <span class="comment">/* 4-byte offset of DIE */</span>
    <a id="L418"></a>opCallRef     = 0x9A; <span class="comment">/* 4- or 8- byte offset of DIE */</span>
    <a id="L419"></a><span class="comment">/* 0xE0-0xFF reserved for user-specific */</span>
<a id="L420"></a>)

<a id="L422"></a><span class="comment">// Basic type encodings -- the value for AttrEncoding in a TagBaseType Entry.</span>
<a id="L423"></a>const (
    <a id="L424"></a>encAddress        = 0x01;
    <a id="L425"></a>encBoolean        = 0x02;
    <a id="L426"></a>encComplexFloat   = 0x03;
    <a id="L427"></a>encFloat          = 0x04;
    <a id="L428"></a>encSigned         = 0x05;
    <a id="L429"></a>encSignedChar     = 0x06;
    <a id="L430"></a>encUnsigned       = 0x07;
    <a id="L431"></a>encUnsignedChar   = 0x08;
    <a id="L432"></a>encImaginaryFloat = 0x09;
<a id="L433"></a>)
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
