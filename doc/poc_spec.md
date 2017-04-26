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

## Non-goals

+ Commenting on an article
+ Adding Comments
+ Manage reputation
+ Be computationaly fast

## Overview

Commenting and reviewing should be a feature of future releases but are out
of the scope of this proof of concept.

All the interaction should be cryptographycally signed, this will allow us
to, in later versions, create a reputation system and add review channels
and workflows.

The system should be peer-to-peer, distributed repository model so no one
individual has control over the publications (torrent? IPFS? block chain?)
