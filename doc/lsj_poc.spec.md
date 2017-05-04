Libre Science Journal
---------------------

## Proof of concept

This specification is a work in progress

Author: R. Díaz de León Plata <leon@elinter.net>

## Scenarios

The software is to be used to publish, review and collaborate in
scientific research. A typical user would write a scientific article
in its favorite text editor and use the lsj client to publish it, the
article *must* have metadata added, so other researcher can filter only
relevant articles and be able to review the ones where they got the
experience, an example could be a sociology article, sociology experts
should be able to review the theory, citations and overall arguments
of the papers while math enthusiasts and staticians should review the
statistics used to make those arguments.

Every contributor must first create his public key:

	$ lsj create-key
	passphrase:

The software should require a reasonable passphrase to secure the stored
private key.

When a contributor wants to publish an article, she must write the document
and save it as an HTML file, using the software (lsj) to publish it:

	$ lsj publish fishes.html
	<identifier>

The HTML file *must* contain the appropiate metadata, if it is to be found.

A user wanting to review articles should be able to use the software to search
for the latest published ones in the network, filtering by subject or other
criteria.

	$ lsj search tag:math.discrete
	<id> - <title>
	<id> - <title>
	<id> - <title>

Once the user has found a situable article, it must be retrieved for review.

	$ lsj fetch <id>

Then an article can be accepted or commented

	$ lsj approve <id>
	$ lsj comment <id> -m "Missig reference for X"

The comment must be appended to the article and automatically signed by the
author of the comment (this will be exploited by creating a cacophony, *must*
be improved.)

An author must be able to add revisions to her articles.

## Non-goals

+ Manage comments
+ Manage reputation
+ Be reasonably secure
+ Be able to browse revisions
+ Match revisions with comments

## Overview

Commenting and reviewing should be a feature of future releases but are out
of the scope of this proof of concept.

All the interaction should be cryptographycally signed, this will allow us
to, in later versions, create a reputation system and add review channels
and workflows.

The system should be peer-to-peer, distributed repository model so no one
individual has control over the publications (torrent? IPFS? block chain?)

