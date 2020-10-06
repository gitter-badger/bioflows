Introduction
############

.. image:: ../_static/images/logo.gif

BioFlows is a fully integrated cluster-enabled and container-enabled Bioinformatics Pipeline
Engine built in Golang..

Overview
========

BioFlows is a distributed pipeline framework for expressing , designing and running scalable reproducible and distributed computational bioinformatics workflows in cloud containers.

BioFlows Framework consists of software tools and cloud microservices that communicate together to achieve a highly distributed , highly coordinated and fault tolerant environment to run parallel bioinformatics pipelines onto cloud containers and cloud servers.

BioFlows also has BioFlows Description Language (BDL) which is an imperative and declarative standard for describing and expressing computational bioinformatics tools and pipelines, BDL is flexible , easy to use and a human readable language that enables researchers to design reproducible and scalable computational pipelines.

The language is based entirely on Yet Another Markup Language (YAML).

Design Goals
============

Portability
^^^^^^^^^^^
BioFlows Framework enables researchers to run massively parallel bioinformatics pipelines on multiple Cloud Platforms , Operating Systems and different Cloud Containers.

Extensibility
^^^^^^^^^^^^^

BioFlows Framework is highly Extensible Framework which allows researchers to extend its functionalities by providing alternative implementation to its different layered components.

Moreover, BioFlows has very rich set of APIs and documentations which allows the community to develop software tools that take advantage of BioFlows Framework.

Fault Tolerance
^^^^^^^^^^^^^^^

BioFlows Framework ensures a seamless and highly consistent execution environment for bioinformatics tools. Tools execution in BioFlows is managed by two highly coordinated microservices.

The Cluster Manager and JobManager , Cluster Manager ensures that JobManager is always running on the cluster Node while JobManager ensures the integrity of the running job.

if the job failed, the Job Manager restarts the job locally and if it failed for n times which defaults to 3 times, the cluster manager will distribute this particular job onto another cluster Node which satisfies the current job computing specifications or wait until available suitable resource becomes free to use.

Reproducibility
^^^^^^^^^^^^^^^

BioFlows Framework enables researchers to write self-contained computational pipelines that are sufficient to run on different Environments with zero configuration because it contains all execution and configuration parameters without any modification, thus, The framework allows different institutions to rapidly reproduce results.

Shareability
^^^^^^^^^^^^

BioFlows Pipelines and Tools could be shared online on GitHub or public servers. BioFlows Command Line could use a remotely shared pipeline or tool definition file, download the file and run it.

BioFlows also has BioFlows Hub which is a purposely-built sharing platform for the scientific community to share their pipeline definition files and allow others to use these pipelines to reproduce their research easily.



