<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/hashmap.c</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/hashmap.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;runtime.h&#34;
#include &#34;hashmap.h&#34;
#include &#34;type.h&#34;

/* Return a pointer to the struct/union of type &#34;type&#34;
   whose &#34;field&#34; field is addressed by pointer &#34;p&#34;. */


struct hash {	   /* a hash table; initialize with hash_init() */
	uint32 count;	  /* elements in table - must be first */

	uint8 datasize;   /* amount of data to store in entry */
	uint8 max_power;  /* max power of 2 to create sub-tables */
	uint8 max_probes; /* max entries to probe before rehashing */
	int32 changes;	      /* inc&#39;ed whenever a subtable is created/grown */
	hash_hash_t (*data_hash) (uint32, void *a);  /* return hash of *a */
	uint32 (*data_eq) (uint32, void *a, void *b);   /* return whether *a == *b */
	void (*data_del) (uint32, void *arg, void *data);  /* invoked on deletion */
	struct hash_subtable *st;    /* first-level table */

	uint32	keysize;
	uint32	valsize;
	uint32	datavo;

	// three sets of offsets: the digit counts how many
	// of key, value are passed as inputs:
	//	0 = func() (key, value)
	//	1 = func(key) (value)
	//	2 = func(key, value)
	uint32	ko0;
	uint32	vo0;
	uint32	ko1;
	uint32	vo1;
	uint32	po1;
	uint32	ko2;
	uint32	vo2;
	uint32	po2;
	Alg*	keyalg;
	Alg*	valalg;
};

struct hash_entry {
	hash_hash_t hash;     /* hash value of data */
	byte data[1];	 /* user data has &#34;datasize&#34; bytes */
};

struct hash_subtable {
	uint8 power;	 /* bits used to index this table */
	uint8 used;	  /* bits in hash used before reaching this table */
	uint8 datasize;      /* bytes of client data in an entry */
	uint8 max_probes;    /* max number of probes when searching */
	int16 limit_bytes;	   /* max_probes * (datasize+sizeof (hash_hash_t)) */
	struct hash_entry *end;      /* points just past end of entry[] */
	struct hash_entry entry[1];  /* 2**power+max_probes-1 elements of elemsize bytes */
};

#define HASH_DATA_EQ(h,x,y) ((*h-&gt;data_eq) (h-&gt;keysize, (x), (y)))

#define HASH_REHASH 0x2       /* an internal flag */
/* the number of bits used is stored in the flags word too */
#define HASH_USED(x)      ((x) &gt;&gt; 2)
#define HASH_MAKE_USED(x) ((x) &lt;&lt; 2)

#define HASH_LOW	6
#define HASH_ONE	(((hash_hash_t)1) &lt;&lt; HASH_LOW)
#define HASH_MASK       (HASH_ONE - 1)
#define HASH_ADJUST(x)  (((x) &lt; HASH_ONE) &lt;&lt; HASH_LOW)

#define HASH_BITS       (sizeof (hash_hash_t) * 8)

#define HASH_SUBHASH    HASH_MASK
#define HASH_NIL	0
#define HASH_NIL_MEMSET 0

#define HASH_OFFSET(base, byte_offset) \
	  ((struct hash_entry *) (((byte *) (base)) + (byte_offset)))


/* return a hash layer with 2**power empty entries */
static struct hash_subtable *
hash_subtable_new (struct hash *h, int32 power, int32 used)
{
	int32 elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);
	int32 bytes = elemsize &lt;&lt; power;
	struct hash_subtable *st;
	int32 limit_bytes = h-&gt;max_probes * elemsize;
	int32 max_probes = h-&gt;max_probes;

	if (bytes &lt; limit_bytes) {
		limit_bytes = bytes;
		max_probes = 1 &lt;&lt; power;
	}
	bytes += limit_bytes - elemsize;
	st = malloc (offsetof (struct hash_subtable, entry[0]) + bytes);
	st-&gt;power = power;
	st-&gt;used = used;
	st-&gt;datasize = h-&gt;datasize;
	st-&gt;max_probes = max_probes;
	st-&gt;limit_bytes = limit_bytes;
	st-&gt;end = HASH_OFFSET (st-&gt;entry, bytes);
	memset (st-&gt;entry, HASH_NIL_MEMSET, bytes);
	return (st);
}

static void
init_sizes (int64 hint, int32 *init_power, int32 *max_power)
{
	int32 log = 0;
	int32 i;

	for (i = 32; i != 0; i &gt;&gt;= 1) {
		if ((hint &gt;&gt; (log + i)) != 0) {
			log += i;
		}
	}
	log += 1 + (((hint &lt;&lt; 3) &gt;&gt; log) &gt;= 11);  /* round up for utilization */
	if (log &lt;= 14) {
		*init_power = log;
	} else {
		*init_power = 12;
	}
	*max_power = 12;
}

static void
hash_init (struct hash *h,
		int32 datasize,
		hash_hash_t (*data_hash) (uint32, void *),
		uint32 (*data_eq) (uint32, void *, void *),
		void (*data_del) (uint32, void *, void *),
		int64 hint)
{
	int32 init_power;
	int32 max_power;

	if(datasize &lt; sizeof (void *))
		datasize = sizeof (void *);
	datasize = rnd(datasize, sizeof (void *));
	init_sizes (hint, &amp;init_power, &amp;max_power);
	h-&gt;datasize = datasize;
	h-&gt;max_power = max_power;
	h-&gt;max_probes = 15;
	assert (h-&gt;datasize == datasize);
	assert (h-&gt;max_power == max_power);
	assert (sizeof (void *) &lt;= h-&gt;datasize || h-&gt;max_power == 255);
	h-&gt;count = 0;
	h-&gt;changes = 0;
	h-&gt;data_hash = data_hash;
	h-&gt;data_eq = data_eq;
	h-&gt;data_del = data_del;
	h-&gt;st = hash_subtable_new (h, init_power, 0);
}

static void
hash_remove_n (struct hash_subtable *st, struct hash_entry *dst_e, int32 n)
{
	int32 elemsize = st-&gt;datasize + offsetof (struct hash_entry, data[0]);
	struct hash_entry *src_e = HASH_OFFSET (dst_e, n * elemsize);
	struct hash_entry *end_e = st-&gt;end;
	int32 shift = HASH_BITS - (st-&gt;power + st-&gt;used);
	int32 index_mask = (((hash_hash_t)1) &lt;&lt; st-&gt;power) - 1;
	int32 dst_i = (((byte *) dst_e) - ((byte *) st-&gt;entry)) / elemsize;
	int32 src_i = dst_i + n;
	hash_hash_t hash;
	int32 skip;
	int32 bytes;

	while (dst_e != src_e) {
		if (src_e != end_e) {
			struct hash_entry *cp_e = src_e;
			int32 save_dst_i = dst_i;
			while (cp_e != end_e &amp;&amp; (hash = cp_e-&gt;hash) != HASH_NIL &amp;&amp;
			     ((hash &gt;&gt; shift) &amp; index_mask) &lt;= dst_i) {
				cp_e = HASH_OFFSET (cp_e, elemsize);
				dst_i++;
			}
			bytes = ((byte *) cp_e) - (byte *) src_e;
			memmove (dst_e, src_e, bytes);
			dst_e = HASH_OFFSET (dst_e, bytes);
			src_e = cp_e;
			src_i += dst_i - save_dst_i;
			if (src_e != end_e &amp;&amp; (hash = src_e-&gt;hash) != HASH_NIL) {
				skip = ((hash &gt;&gt; shift) &amp; index_mask) - dst_i;
			} else {
				skip = src_i - dst_i;
			}
		} else {
			skip = src_i - dst_i;
		}
		bytes = skip * elemsize;
		memset (dst_e, HASH_NIL_MEMSET, bytes);
		dst_e = HASH_OFFSET (dst_e, bytes);
		dst_i += skip;
	}
}

static int32
hash_insert_internal (struct hash_subtable **pst, int32 flags, hash_hash_t hash,
		struct hash *h, void *data, void **pres);

static void
hash_conv (struct hash *h,
		struct hash_subtable *st, int32 flags,
		hash_hash_t hash,
		struct hash_entry *e)
{
	int32 new_flags = (flags + HASH_MAKE_USED (st-&gt;power)) | HASH_REHASH;
	int32 shift = HASH_BITS - HASH_USED (new_flags);
	hash_hash_t prefix_mask = (-(hash_hash_t)1) &lt;&lt; shift;
	int32 elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);
	void *dummy_result;
	struct hash_entry *de;
	int32 index_mask = (1 &lt;&lt; st-&gt;power) - 1;
	hash_hash_t e_hash;
	struct hash_entry *pe = HASH_OFFSET (e, -elemsize);

	while (e != st-&gt;entry &amp;&amp; (e_hash = pe-&gt;hash) != HASH_NIL &amp;&amp; (e_hash &amp; HASH_MASK) != HASH_SUBHASH) {
		e = pe;
		pe = HASH_OFFSET (pe, -elemsize);
	}

	de = e;
	while (e != st-&gt;end &amp;&amp;
	    (e_hash = e-&gt;hash) != HASH_NIL &amp;&amp;
	    (e_hash &amp; HASH_MASK) != HASH_SUBHASH) {
		struct hash_entry *target_e = HASH_OFFSET (st-&gt;entry, ((e_hash &gt;&gt; shift) &amp; index_mask) * elemsize);
		struct hash_entry *ne = HASH_OFFSET (e, elemsize);
		hash_hash_t current = e_hash &amp; prefix_mask;
		if (de &lt; target_e) {
			memset (de, HASH_NIL_MEMSET, ((byte *) target_e) - (byte *) de);
			de = target_e;
		}
		if ((hash &amp; prefix_mask) == current ||
		   (ne != st-&gt;end &amp;&amp; (e_hash = ne-&gt;hash) != HASH_NIL &amp;&amp;
		   (e_hash &amp; prefix_mask) == current)) {
			struct hash_subtable *new_st = hash_subtable_new (h, 1, HASH_USED (new_flags));
			int32 rc = hash_insert_internal (&amp;new_st, new_flags, e-&gt;hash, h, e-&gt;data, &amp;dummy_result);
			assert (rc == 0);
			memcpy(dummy_result, e-&gt;data, h-&gt;datasize);
			e = ne;
			while (e != st-&gt;end &amp;&amp; (e_hash = e-&gt;hash) != HASH_NIL &amp;&amp; (e_hash &amp; prefix_mask) == current) {
				assert ((e_hash &amp; HASH_MASK) != HASH_SUBHASH);
				rc = hash_insert_internal (&amp;new_st, new_flags, e_hash, h, e-&gt;data, &amp;dummy_result);
				assert (rc == 0);
				memcpy(dummy_result, e-&gt;data, h-&gt;datasize);
				e = HASH_OFFSET (e, elemsize);
			}
			memset (de-&gt;data, HASH_NIL_MEMSET, h-&gt;datasize);
			*(struct hash_subtable **)de-&gt;data = new_st;
			de-&gt;hash = current | HASH_SUBHASH;
		} else {
			if (e != de) {
				memcpy (de, e, elemsize);
			}
			e = HASH_OFFSET (e, elemsize);
		}
		de = HASH_OFFSET (de, elemsize);
	}
	if (e != de) {
		hash_remove_n (st, de, (((byte *) e) - (byte *) de) / elemsize);
	}
}

static void
hash_grow (struct hash *h, struct hash_subtable **pst, int32 flags)
{
	struct hash_subtable *old_st = *pst;
	int32 elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);
	*pst = hash_subtable_new (h, old_st-&gt;power + 1, HASH_USED (flags));
	struct hash_entry *end_e = old_st-&gt;end;
	struct hash_entry *e;
	void *dummy_result;
	int32 used = 0;

	flags |= HASH_REHASH;
	for (e = old_st-&gt;entry; e != end_e; e = HASH_OFFSET (e, elemsize)) {
		hash_hash_t hash = e-&gt;hash;
		if (hash != HASH_NIL) {
			int32 rc = hash_insert_internal (pst, flags, e-&gt;hash, h, e-&gt;data, &amp;dummy_result);
			assert (rc == 0);
			memcpy(dummy_result, e-&gt;data, h-&gt;datasize);
			used++;
		}
	}
	free (old_st);
}

int32
hash_lookup (struct hash *h, void *data, void **pres)
{
	int32 elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);
	hash_hash_t hash = (*h-&gt;data_hash) (h-&gt;keysize, data) &amp; ~HASH_MASK;
	struct hash_subtable *st = h-&gt;st;
	int32 used = 0;
	hash_hash_t e_hash;
	struct hash_entry *e;
	struct hash_entry *end_e;

	hash += HASH_ADJUST (hash);
	for (;;) {
		int32 shift = HASH_BITS - (st-&gt;power + used);
		int32 index_mask = (1 &lt;&lt; st-&gt;power) - 1;
		int32 i = (hash &gt;&gt; shift) &amp; index_mask;	   /* i is the natural position of hash */

		e = HASH_OFFSET (st-&gt;entry, i * elemsize); /* e points to element i */
		e_hash = e-&gt;hash;
		if ((e_hash &amp; HASH_MASK) != HASH_SUBHASH) {      /* a subtable */
			break;
		}
		used += st-&gt;power;
		st = *(struct hash_subtable **)e-&gt;data;
	}
	end_e = HASH_OFFSET (e, st-&gt;limit_bytes);
	while (e != end_e &amp;&amp; (e_hash = e-&gt;hash) != HASH_NIL &amp;&amp; e_hash &lt; hash) {
		e = HASH_OFFSET (e, elemsize);
	}
	while (e != end_e &amp;&amp; ((e_hash = e-&gt;hash) ^ hash) &lt; HASH_SUBHASH) {
		if (HASH_DATA_EQ (h, data, e-&gt;data)) {    /* a match */
			*pres = e-&gt;data;
			return (1);
		}
		e = HASH_OFFSET (e, elemsize);
	}
	USED(e_hash);
	*pres = 0;
	return (0);
}

int32
hash_remove (struct hash *h, void *data, void *arg)
{
	int32 elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);
	hash_hash_t hash = (*h-&gt;data_hash) (h-&gt;keysize, data) &amp; ~HASH_MASK;
	struct hash_subtable *st = h-&gt;st;
	int32 used = 0;
	hash_hash_t e_hash;
	struct hash_entry *e;
	struct hash_entry *end_e;

	hash += HASH_ADJUST (hash);
	for (;;) {
		int32 shift = HASH_BITS - (st-&gt;power + used);
		int32 index_mask = (1 &lt;&lt; st-&gt;power) - 1;
		int32 i = (hash &gt;&gt; shift) &amp; index_mask;	   /* i is the natural position of hash */

		e = HASH_OFFSET (st-&gt;entry, i * elemsize); /* e points to element i */
		e_hash = e-&gt;hash;
		if ((e_hash &amp; HASH_MASK) != HASH_SUBHASH) {      /* a subtable */
			break;
		}
		used += st-&gt;power;
		st = *(struct hash_subtable **)e-&gt;data;
	}
	end_e = HASH_OFFSET (e, st-&gt;limit_bytes);
	while (e != end_e &amp;&amp; (e_hash = e-&gt;hash) != HASH_NIL &amp;&amp; e_hash &lt; hash) {
		e = HASH_OFFSET (e, elemsize);
	}
	while (e != end_e &amp;&amp; ((e_hash = e-&gt;hash) ^ hash) &lt; HASH_SUBHASH) {
		if (HASH_DATA_EQ (h, data, e-&gt;data)) {    /* a match */
			(*h-&gt;data_del) (h-&gt;keysize, arg, e-&gt;data);
			hash_remove_n (st, e, 1);
			h-&gt;count--;
			return (1);
		}
		e = HASH_OFFSET (e, elemsize);
	}
	USED(e_hash);
	return (0);
}

static int32
hash_insert_internal (struct hash_subtable **pst, int32 flags, hash_hash_t hash,
				 struct hash *h, void *data, void **pres)
{
	int32 elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);

	if ((flags &amp; HASH_REHASH) == 0) {
		hash += HASH_ADJUST (hash);
		hash &amp;= ~HASH_MASK;
	}
	for (;;) {
		struct hash_subtable *st = *pst;
		int32 shift = HASH_BITS - (st-&gt;power + HASH_USED (flags));
		int32 index_mask = (1 &lt;&lt; st-&gt;power) - 1;
		int32 i = (hash &gt;&gt; shift) &amp; index_mask;	   /* i is the natural position of hash */
		struct hash_entry *start_e =
			HASH_OFFSET (st-&gt;entry, i * elemsize);    /* start_e is the pointer to element i */
		struct hash_entry *e = start_e;		   /* e is going to range over [start_e, end_e) */
		struct hash_entry *end_e;
		hash_hash_t e_hash = e-&gt;hash;

		if ((e_hash &amp; HASH_MASK) == HASH_SUBHASH) {      /* a subtable */
			pst = (struct hash_subtable **) e-&gt;data;
			flags += HASH_MAKE_USED (st-&gt;power);
			continue;
		}
		end_e = HASH_OFFSET (start_e, st-&gt;limit_bytes);
		while (e != end_e &amp;&amp; (e_hash = e-&gt;hash) != HASH_NIL &amp;&amp; e_hash &lt; hash) {
			e = HASH_OFFSET (e, elemsize);
			i++;
		}
		if (e != end_e &amp;&amp; e_hash != HASH_NIL) {
			/* ins_e ranges over the elements that may match */
			struct hash_entry *ins_e = e;
			int32 ins_i = i;
			hash_hash_t ins_e_hash;
			while (ins_e != end_e &amp;&amp; ((e_hash = ins_e-&gt;hash) ^ hash) &lt; HASH_SUBHASH) {
				if (HASH_DATA_EQ (h, data, ins_e-&gt;data)) {    /* a match */
					*pres = ins_e-&gt;data;
					return (1);
				}
				assert (e_hash != hash || (flags &amp; HASH_REHASH) == 0);
				hash += (e_hash == hash);	   /* adjust hash if it collides */
				ins_e = HASH_OFFSET (ins_e, elemsize);
				ins_i++;
				if (e_hash &lt;= hash) {	       /* set e to insertion point */
					e = ins_e;
					i = ins_i;
				}
			}
			/* set ins_e to the insertion point for the new element */
			ins_e = e;
			ins_i = i;
			ins_e_hash = 0;
			/* move ins_e to point at the end of the contiguous block, but
			   stop if any element can&#39;t be moved by one up */
			while (ins_e != st-&gt;end &amp;&amp; (ins_e_hash = ins_e-&gt;hash) != HASH_NIL &amp;&amp;
			       ins_i + 1 - ((ins_e_hash &gt;&gt; shift) &amp; index_mask) &lt; st-&gt;max_probes &amp;&amp;
			       (ins_e_hash &amp; HASH_MASK) != HASH_SUBHASH) {
				ins_e = HASH_OFFSET (ins_e, elemsize);
				ins_i++;
			}
			if (e == end_e || ins_e == st-&gt;end || ins_e_hash != HASH_NIL) {
				e = end_e;    /* can&#39;t insert; must grow or convert to subtable */
			} else {	      /* make space for element */
				memmove (HASH_OFFSET (e, elemsize), e, ((byte *) ins_e) - (byte *) e);
			}
		}
		if (e != end_e) {
			e-&gt;hash = hash;
			*pres = e-&gt;data;
			return (0);
		}
		h-&gt;changes++;
		if (st-&gt;power &lt; h-&gt;max_power) {
			hash_grow (h, pst, flags);
		} else {
			hash_conv (h, st, flags, hash, start_e);
		}
	}
}

int32
hash_insert (struct hash *h, void *data, void **pres)
{
	int32 rc = hash_insert_internal (&amp;h-&gt;st, 0, (*h-&gt;data_hash) (h-&gt;keysize, data), h, data, pres);

	h-&gt;count += (rc == 0);    /* increment count if element didn&#39;t previously exist */
	return (rc);
}

uint32
hash_count (struct hash *h)
{
	return (h-&gt;count);
}

static void
iter_restart (struct hash_iter *it, struct hash_subtable *st, int32 used)
{
	int32 elemsize = it-&gt;elemsize;
	hash_hash_t last_hash = it-&gt;last_hash;
	struct hash_entry *e;
	hash_hash_t e_hash;
	struct hash_iter_sub *sub = &amp;it-&gt;subtable_state[it-&gt;i];
	struct hash_entry *end;

	for (;;) {
		int32 shift = HASH_BITS - (st-&gt;power + used);
		int32 index_mask = (1 &lt;&lt; st-&gt;power) - 1;
		int32 i = (last_hash &gt;&gt; shift) &amp; index_mask;

		end = st-&gt;end;
		e = HASH_OFFSET (st-&gt;entry, i * elemsize);
		sub-&gt;start = st-&gt;entry;
		sub-&gt;end = end;

		if ((e-&gt;hash &amp; HASH_MASK) != HASH_SUBHASH) {
			break;
		}
		sub-&gt;e = HASH_OFFSET (e, elemsize);
		sub = &amp;it-&gt;subtable_state[++(it-&gt;i)];
		used += st-&gt;power;
		st = *(struct hash_subtable **)e-&gt;data;
	}
	while (e != end &amp;&amp; ((e_hash = e-&gt;hash) == HASH_NIL || e_hash &lt;= last_hash)) {
		e = HASH_OFFSET (e, elemsize);
	}
	sub-&gt;e = e;
}

void *
hash_next (struct hash_iter *it)
{
	int32 elemsize = it-&gt;elemsize;
	struct hash_iter_sub *sub = &amp;it-&gt;subtable_state[it-&gt;i];
	struct hash_entry *e = sub-&gt;e;
	struct hash_entry *end = sub-&gt;end;
	hash_hash_t e_hash = 0;

	if (it-&gt;changes != it-&gt;h-&gt;changes) {    /* hash table&#39;s structure changed; recompute */
		it-&gt;changes = it-&gt;h-&gt;changes;
		it-&gt;i = 0;
		iter_restart (it, it-&gt;h-&gt;st, 0);
		sub = &amp;it-&gt;subtable_state[it-&gt;i];
		e = sub-&gt;e;
		end = sub-&gt;end;
	}
	if (e != sub-&gt;start &amp;&amp; it-&gt;last_hash != HASH_OFFSET (e, -elemsize)-&gt;hash) {
		struct hash_entry *start = HASH_OFFSET (e, -(elemsize * it-&gt;h-&gt;max_probes));
		struct hash_entry *pe = HASH_OFFSET (e, -elemsize);
		hash_hash_t last_hash = it-&gt;last_hash;
		if (start &lt; sub-&gt;start) {
			start = sub-&gt;start;
		}
		while (e != start &amp;&amp; ((e_hash = pe-&gt;hash) == HASH_NIL || last_hash &lt; e_hash)) {
			e = pe;
			pe = HASH_OFFSET (pe, -elemsize);
		}
		while (e != end &amp;&amp; ((e_hash = e-&gt;hash) == HASH_NIL || e_hash &lt;= last_hash)) {
			e = HASH_OFFSET (e, elemsize);
		}
	}

	for (;;) {
		while (e != end &amp;&amp; (e_hash = e-&gt;hash) == HASH_NIL) {
			e = HASH_OFFSET (e, elemsize);
		}
		if (e == end) {
			if (it-&gt;i == 0) {
				it-&gt;last_hash = HASH_OFFSET (e, -elemsize)-&gt;hash;
				sub-&gt;e = e;
				return (0);
			} else {
				it-&gt;i--;
				sub = &amp;it-&gt;subtable_state[it-&gt;i];
				e = sub-&gt;e;
				end = sub-&gt;end;
			}
		} else if ((e_hash &amp; HASH_MASK) != HASH_SUBHASH) {
			it-&gt;last_hash = e-&gt;hash;
			sub-&gt;e = HASH_OFFSET (e, elemsize);
			return (e-&gt;data);
		} else {
			struct hash_subtable *st =
				*(struct hash_subtable **)e-&gt;data;
			sub-&gt;e = HASH_OFFSET (e, elemsize);
			it-&gt;i++;
			assert (it-&gt;i &lt; sizeof (it-&gt;subtable_state) /
					sizeof (it-&gt;subtable_state[0]));
			sub = &amp;it-&gt;subtable_state[it-&gt;i];
			sub-&gt;e = e = st-&gt;entry;
			sub-&gt;start = st-&gt;entry;
			sub-&gt;end = end = st-&gt;end;
		}
	}
}

void
hash_iter_init (struct hash *h, struct hash_iter *it)
{
	it-&gt;elemsize = h-&gt;datasize + offsetof (struct hash_entry, data[0]);
	it-&gt;changes = h-&gt;changes;
	it-&gt;i = 0;
	it-&gt;h = h;
	it-&gt;last_hash = 0;
	it-&gt;subtable_state[0].e = h-&gt;st-&gt;entry;
	it-&gt;subtable_state[0].start = h-&gt;st-&gt;entry;
	it-&gt;subtable_state[0].end = h-&gt;st-&gt;end;
}

static void
clean_st (struct hash_subtable *st, int32 *slots, int32 *used)
{
	int32 elemsize = st-&gt;datasize + offsetof (struct hash_entry, data[0]);
	struct hash_entry *e = st-&gt;entry;
	struct hash_entry *end = st-&gt;end;
	int32 lslots = (((byte *) end) - (byte *) e) / elemsize;
	int32 lused = 0;

	while (e != end) {
		hash_hash_t hash = e-&gt;hash;
		if ((hash &amp; HASH_MASK) == HASH_SUBHASH) {
			clean_st (*(struct hash_subtable **)e-&gt;data, slots, used);
		} else {
			lused += (hash != HASH_NIL);
		}
		e = HASH_OFFSET (e, elemsize);
	}
	free (st);
	*slots += lslots;
	*used += lused;
}

void
hash_destroy (struct hash *h)
{
	int32 slots = 0;
	int32 used = 0;

	clean_st (h-&gt;st, &amp;slots, &amp;used);
	free (h);
}

static void
hash_visit_internal (struct hash_subtable *st,
		int32 used, int32 level,
		void (*data_visit) (void *arg, int32 level, void *data),
		void *arg)
{
	int32 elemsize = st-&gt;datasize + offsetof (struct hash_entry, data[0]);
	struct hash_entry *e = st-&gt;entry;
	int32 shift = HASH_BITS - (used + st-&gt;power);
	int32 i = 0;

	while (e != st-&gt;end) {
		int32 index = ((e-&gt;hash &gt;&gt; (shift - 1)) &gt;&gt; 1) &amp; ((1 &lt;&lt; st-&gt;power) - 1);
		if ((e-&gt;hash &amp; HASH_MASK) == HASH_SUBHASH) {
			  (*data_visit) (arg, level, e-&gt;data);
			  hash_visit_internal (*(struct hash_subtable **)e-&gt;data,
				used + st-&gt;power, level + 1, data_visit, arg);
		} else {
			  (*data_visit) (arg, level, e-&gt;data);
		}
		if (e-&gt;hash != HASH_NIL) {
			  assert (i &lt; index + st-&gt;max_probes);
			  assert (index &lt;= i);
		}
		e = HASH_OFFSET (e, elemsize);
		i++;
	}
}

void
hash_visit (struct hash *h, void (*data_visit) (void *arg, int32 level, void *data), void *arg)
{
	hash_visit_internal (h-&gt;st, 0, 0, data_visit, arg);
}

//
/// interfaces to go runtime
//

static void
donothing(uint32 s, void *a, void *b)
{
	USED(s);
	USED(a);
	USED(b);
}

static	int32	debug	= 0;

// makemap(key, val *Type, hint uint32) (hmap *map[any]any);
Hmap*
makemap(Type *key, Type *val, uint32 hint)
{
	Hmap *h;
	int32 keyalg, valalg, keysize, valsize;

	keyalg = key-&gt;alg;
	valalg = val-&gt;alg;
	keysize = key-&gt;size;
	valsize = val-&gt;size;

	if(keyalg &gt;= nelem(algarray) || algarray[keyalg].hash == nohash) {
		printf(&#34;map(keyalg=%d)\n&#34;, keyalg);
		throw(&#34;runtime·makemap: unsupported map key type&#34;);
	}

	if(valalg &gt;= nelem(algarray)) {
		printf(&#34;map(valalg=%d)\n&#34;, valalg);
		throw(&#34;runtime·makemap: unsupported map value type&#34;);
	}

	h = mal(sizeof(*h));

	// align value inside data so that mark-sweep gc can find it.
	// might remove in the future and just assume datavo == keysize.
	h-&gt;datavo = keysize;
	if(valsize &gt;= sizeof(void*))
		h-&gt;datavo = rnd(keysize, sizeof(void*));

	hash_init(h, h-&gt;datavo+valsize,
		algarray[keyalg].hash,
		algarray[keyalg].equal,
		donothing,
		hint);

	h-&gt;keysize = keysize;
	h-&gt;valsize = valsize;
	h-&gt;keyalg = &amp;algarray[keyalg];
	h-&gt;valalg = &amp;algarray[valalg];

	// these calculations are compiler dependent.
	// figure out offsets of map call arguments.

	// func() (key, val)
	h-&gt;ko0 = rnd(sizeof(h), Structrnd);
	h-&gt;vo0 = rnd(h-&gt;ko0+keysize, val-&gt;align);

	// func(key) (val[, pres])
	h-&gt;ko1 = rnd(sizeof(h), key-&gt;align);
	h-&gt;vo1 = rnd(h-&gt;ko1+keysize, Structrnd);
	h-&gt;po1 = rnd(h-&gt;vo1+valsize, 1);

	// func(key, val[, pres])
	h-&gt;ko2 = rnd(sizeof(h), key-&gt;align);
	h-&gt;vo2 = rnd(h-&gt;ko2+keysize, val-&gt;align);
	h-&gt;po2 = rnd(h-&gt;vo2+valsize, 1);

	if(debug) {
		printf(&#34;makemap: map=%p; keysize=%d; valsize=%d; keyalg=%d; valalg=%d; offsets=%d,%d; %d,%d,%d; %d,%d,%d\n&#34;,
			h, keysize, valsize, keyalg, valalg, h-&gt;ko0, h-&gt;vo0, h-&gt;ko1, h-&gt;vo1, h-&gt;po1, h-&gt;ko2, h-&gt;vo2, h-&gt;po2);
	}

	return h;
}

// makemap(key, val *Type, hint uint32) (hmap *map[any]any);
void
runtime·makemap(Type *key, Type *val, uint32 hint, Hmap *ret)
{
	ret = makemap(key, val, hint);
	FLUSH(&amp;ret);
}

void
mapaccess(Hmap *h, byte *ak, byte *av, bool *pres)
{
	byte *res;

	res = nil;
	if(hash_lookup(h, ak, (void**)&amp;res)) {
		*pres = true;
		h-&gt;valalg-&gt;copy(h-&gt;valsize, av, res+h-&gt;datavo);
	} else {
		*pres = false;
		h-&gt;valalg-&gt;copy(h-&gt;valsize, av, nil);
	}
}

// mapaccess1(hmap *map[any]any, key any) (val any);
void
runtime·mapaccess1(Hmap *h, ...)
{
	byte *ak, *av;
	bool pres;

	ak = (byte*)&amp;h + h-&gt;ko1;
	av = (byte*)&amp;h + h-&gt;vo1;

	mapaccess(h, ak, av, &amp;pres);
	if(!pres)
		throw(&#34;runtime·mapaccess1: key not in map&#34;);

	if(debug) {
		prints(&#34;runtime·mapaccess1: map=&#34;);
		runtime·printpointer(h);
		prints(&#34;; key=&#34;);
		h-&gt;keyalg-&gt;print(h-&gt;keysize, ak);
		prints(&#34;; val=&#34;);
		h-&gt;valalg-&gt;print(h-&gt;valsize, av);
		prints(&#34;; pres=&#34;);
		runtime·printbool(pres);
		prints(&#34;\n&#34;);
	}
}

// mapaccess2(hmap *map[any]any, key any) (val any, pres bool);
void
runtime·mapaccess2(Hmap *h, ...)
{
	byte *ak, *av, *ap;

	ak = (byte*)&amp;h + h-&gt;ko1;
	av = (byte*)&amp;h + h-&gt;vo1;
	ap = (byte*)&amp;h + h-&gt;po1;

	mapaccess(h, ak, av, ap);

	if(debug) {
		prints(&#34;runtime·mapaccess2: map=&#34;);
		runtime·printpointer(h);
		prints(&#34;; key=&#34;);
		h-&gt;keyalg-&gt;print(h-&gt;keysize, ak);
		prints(&#34;; val=&#34;);
		h-&gt;valalg-&gt;print(h-&gt;valsize, av);
		prints(&#34;; pres=&#34;);
		runtime·printbool(*ap);
		prints(&#34;\n&#34;);
	}
}

void
mapassign(Hmap *h, byte *ak, byte *av)
{
	byte *res;
	int32 hit;

	res = nil;
	if(av == nil) {
		hash_remove(h, ak, (void**)&amp;res);
		return;
	}

	hit = hash_insert(h, ak, (void**)&amp;res);
	h-&gt;keyalg-&gt;copy(h-&gt;keysize, res, ak);
	h-&gt;valalg-&gt;copy(h-&gt;valsize, res+h-&gt;datavo, av);

	if(debug) {
		prints(&#34;mapassign: map=&#34;);
		runtime·printpointer(h);
		prints(&#34;; key=&#34;);
		h-&gt;keyalg-&gt;print(h-&gt;keysize, ak);
		prints(&#34;; val=&#34;);
		h-&gt;valalg-&gt;print(h-&gt;valsize, av);
		prints(&#34;; hit=&#34;);
		runtime·printint(hit);
		prints(&#34;; res=&#34;);
		runtime·printpointer(res);
		prints(&#34;\n&#34;);
	}
}

// mapassign1(hmap *map[any]any, key any, val any);
void
runtime·mapassign1(Hmap *h, ...)
{
	byte *ak, *av;

	ak = (byte*)&amp;h + h-&gt;ko2;
	av = (byte*)&amp;h + h-&gt;vo2;

	mapassign(h, ak, av);
}

// mapassign2(hmap *map[any]any, key any, val any, pres bool);
void
runtime·mapassign2(Hmap *h, ...)
{
	byte *ak, *av, *ap;

	ak = (byte*)&amp;h + h-&gt;ko2;
	av = (byte*)&amp;h + h-&gt;vo2;
	ap = (byte*)&amp;h + h-&gt;po2;

	if(*ap == false)
		av = nil;	// delete

	mapassign(h, ak, av);

	if(debug) {
		prints(&#34;mapassign2: map=&#34;);
		runtime·printpointer(h);
		prints(&#34;; key=&#34;);
		h-&gt;keyalg-&gt;print(h-&gt;keysize, ak);
		prints(&#34;\n&#34;);
	}
}

// mapiterinit(hmap *map[any]any, hiter *any);
void
runtime·mapiterinit(Hmap *h, struct hash_iter *it)
{
	if(h == nil) {
		it-&gt;data = nil;
		return;
	}
	hash_iter_init(h, it);
	it-&gt;data = hash_next(it);
	if(debug) {
		prints(&#34;runtime·mapiterinit: map=&#34;);
		runtime·printpointer(h);
		prints(&#34;; iter=&#34;);
		runtime·printpointer(it);
		prints(&#34;; data=&#34;);
		runtime·printpointer(it-&gt;data);
		prints(&#34;\n&#34;);
	}
}

struct hash_iter*
mapiterinit(Hmap *h)
{
	struct hash_iter *it;

	it = mal(sizeof *it);
	runtime·mapiterinit(h, it);
	return it;
}

// mapiternext(hiter *any);
void
runtime·mapiternext(struct hash_iter *it)
{
	it-&gt;data = hash_next(it);
	if(debug) {
		prints(&#34;runtime·mapiternext: iter=&#34;);
		runtime·printpointer(it);
		prints(&#34;; data=&#34;);
		runtime·printpointer(it-&gt;data);
		prints(&#34;\n&#34;);
	}
}

void
mapiternext(struct hash_iter *it)
{
	runtime·mapiternext(it);
}

// mapiter1(hiter *any) (key any);
void
runtime·mapiter1(struct hash_iter *it, ...)
{
	Hmap *h;
	byte *ak, *res;

	h = it-&gt;h;
	ak = (byte*)&amp;it + h-&gt;ko0;

	res = it-&gt;data;
	if(res == nil)
		throw(&#34;runtime·mapiter2: key:val nil pointer&#34;);

	h-&gt;keyalg-&gt;copy(h-&gt;keysize, ak, res);

	if(debug) {
		prints(&#34;mapiter2: iter=&#34;);
		runtime·printpointer(it);
		prints(&#34;; map=&#34;);
		runtime·printpointer(h);
		prints(&#34;\n&#34;);
	}
}

bool
mapiterkey(struct hash_iter *it, void *ak)
{
	Hmap *h;
	byte *res;

	h = it-&gt;h;
	res = it-&gt;data;
	if(res == nil)
		return false;
	h-&gt;keyalg-&gt;copy(h-&gt;keysize, ak, res);
	return true;
}

// mapiter2(hiter *any) (key any, val any);
void
runtime·mapiter2(struct hash_iter *it, ...)
{
	Hmap *h;
	byte *ak, *av, *res;

	h = it-&gt;h;
	ak = (byte*)&amp;it + h-&gt;ko0;
	av = (byte*)&amp;it + h-&gt;vo0;

	res = it-&gt;data;
	if(res == nil)
		throw(&#34;runtime·mapiter2: key:val nil pointer&#34;);

	h-&gt;keyalg-&gt;copy(h-&gt;keysize, ak, res);
	h-&gt;valalg-&gt;copy(h-&gt;valsize, av, res+h-&gt;datavo);

	if(debug) {
		prints(&#34;mapiter2: iter=&#34;);
		runtime·printpointer(it);
		prints(&#34;; map=&#34;);
		runtime·printpointer(h);
		prints(&#34;\n&#34;);
	}
}
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
