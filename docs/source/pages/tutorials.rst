Practical Tutorials
###################

In this section, we will put what we have learnt so far about Bioflows into action. We are going to write real-world bioinformatics pipelines in BioFlows.
We will start from the easiest example and try to make it harder as we move along these tutorials. If the tutorial was written somewhere else, the original website will be acknowledged.

How to install
^^^^^^^^^^^^^^

Before you can follow these practical tutorials, you need to have Bioflows installed onto your system. For this, we have prepared a complete
convenient bash script to automate the whole process for you.

Please follow the three simple steps below to have `bf` executable available in your linux box.


.. code-block:: bash

    $ curl -OL http://bioflows.github.io/bioflows.install.sh
    $ chmod a+x bioflows.install.sh
    $ ./bioflows.install.sh

After the above three steps complete successfully, you are ready to practice the following real bioinformatics tutorials.


System Configuration File
^^^^^^^^^^^^^^^^^^^^^^^^^

Bioflows require a system-wide YAML configuration file to be located at your home directory `~/` with name `.bf.yaml`.

The contents of this file are:

.. code-block:: yaml

    remote: false
    email:
      type: smtp
      host: smtp.gmail.com
      port: 587
      username: "yourusername@gmail.com"
      password: <Your Password>
      ssl: false
      tls: true


    cluster:
      address: 127.0.0.1
      port: 8500
      scheme: http


Please copy these contents into a file named `.bf.yaml` and place the file into your home directory `~/`.

BioFlows Tutorials
^^^^^^^^^^^^^^^^^^

.. toctree::
   :maxdepth: 2
   :caption: Practical Tutorials:

   tutorials/blast





