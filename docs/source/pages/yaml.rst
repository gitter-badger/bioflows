BioFlows Definition Language
############################

BioFlows Definition Language is based on a declarative and structured markup language called (YAML: Yet Another Markup Language).
You express an individual tool or a complete pipeline/workflow by utilizing predefined set of attributes called "Directives".
These Directives allow you to take full control of tool definition and execution.
In BioFlows, The same directives are used to define both an individual tool or a whole complex pipeline with nested whole pipeline(s).

The following section enumerates all these directives in greater details....

BioFlows Directives
===================
.. csv-table:: BioFlows Tool/Pipeline Directives
   :file:  directives.csv
   :header-rows: 1
   :widths: 130, 170


Maintainer Directive
^^^^^^^^^^^^^^^^^^^^

Maintainer directive describes metadata information about the researcher or the bioinformatician who has written that tool or pipeline.
This person is considered to be in charge and support for this tool or pipeline. Users of the pipeline can use this information to communicate with him.

.. code-block:: yaml

   maintainer:
    username: xxxx
    fullname: xxxxx xxxxx
    email: xxx@xx.com

References Directive
^^^^^^^^^^^^^^^^^^^^

This is an optional list of references. Each reference is an object composed of nested directives.
this directive is used to include references to any scientific publications, papers,
posters and/or articles that might act as additional information sources for users
of this tool and/or pipeline.

You define references directive, as follows....

.. code-block:: yaml

   references:
    - name: "Name of your reference"
      description: "long or short snippet of description about this reference"
      website: http://www.yourreference-url.com
    - name: "Name of your reference"
      description: "long or short snippet of description about this reference"
      website: http://www.yourreference-url.com

Inputs Directive
^^^^^^^^^^^^^^^^

How to define Input Parameters (Inputs)
***************************************

Each separate tool or a tool in a bioinformatics pipeline requires some input(s) parameters
to work on and might or might not produce any output(s). Some Bioflows tools might act as decision steps
or state modifiers in a pipeline and hence these tools will only require some input(s) from previous step(s)
and will not produce any output(s). These tools should be shadowed having ``shadow=true`` in their definition.

In order to define input(s) for a tool or a pipeline, the following is an example inputs definition for a dummy tool..

.. code-block:: yaml

   inputs:
      - type: string
        displayname: The input directory for the command
        description: short or long description of the input file
        name: input_dir
        value: /your/original/dir/location
      - type: string
        displayname: The data directory where the rest of the required files reside
        description: short or long description of the data directory
        name: data_dir
        value: /your/data/dir



The type of the input parameter could be a ``string``, a ``file`` , a ``dir`` or it could be anything else.
It really does not matter the value of this type directive as long as the author of the tool knows how to use it
in either the scripts directive or the command directive.

Output(s) Directive
^^^^^^^^^^^^^^^^^^^

Output(s) directive defines a set of output parameter(s) which might be produced
by a tool during its execution. the outputs are the actual variables which could be utilized
by other downstream dependent tools in the pipeline. A tool might or might not produce any output(s).
Outputs directive follows the same definition markup as that of the inputs shown above.

.. code-block:: yaml

   outputs:
      - type: file
        displayname: "...."
        description: "...."
        name: output_file
        value: myfile.txt




Pipeline Definition Example(s)
==============================

Please use the following pipeline as an example to understand how to define the previously explained directives in the table above.

.. code-block:: yaml

    id: secondPipeline
    bioflowId: secondPipeline
    type: pipeline
    name: Second Pipeline
    description:
      -"This tool is the second pipeline"
      -"This tool is the second pipeline"
    website: http://hub.bioflows.io
    version: 1.0.0
    steps:
      - id: 1
        bioflowId: mytool1
        name: Generate
        inputs:
          - type: string
            displayname: The input directory for the command
            name: input_dir
            value: /home/snouto
        outputs:
          - type: file
            name: output_file
            value: myfile.txt
        command: ls -ll {{input_dir}} > {{self_dir}}/{{output_file}}
      - id: 2
        bioflowId: mytool2
        name: Move
        depends: 1
        description: "This is a tool that will list all linux directories"
        website: http://hub.bioflows.io
        inputs:
          - type: file
            displayname: The input file to move
            name: input_file
            value: "{{1.location}}/{{1.output_file}}"
          - type: dir
            name: dest_dir
            description: Destination Directory
            value: "{{self_dir}}/movedFile.txt"
        command: mv {{input_file}} {{dest_dir}}
      - id: 3
        name: count
        depends: 1,2
        command: wc -l {{2.dest_dir}} > {{self_dir}}/count.txt


Another Tool definition....

.. code-block:: yaml

    id: pipeline1
    bioflowId: pipeline123
    type: pipeline
    name: my pipeline
    description:
      -"this tool will list directories"
      -"this tool will list all linux directories for a given input directory parameter"
    website: http://www.google.com
    version: 1.0.0
    icon: here you can place the base64 encoded string value of an icon in png format
    maintainer:
      -fullname: XXXXXXXX
      email: XXXXXXX@gmail.com
      username: XXXXXXX

    references:
      - name: "name of the reference"
        description: "long or short snippet of a description goes here"
        website: http://www.ncbi.nlm.gov.nl

    steps:
      - id: 1
        bioflowId: xdir3525
        name: 1
        description: this is a tool that will list all linux directories for a given input directory parameter
        discussions:
          - this tool will list directories
          - this tool will list all linux directories for a given input directory parameter
        website: http://hub.bioflows.io
        version: 1.0.0
        icon: here you can place base64 encoded string value of an icon in png format
        # shadow property indicates that this tool will have no output, it exists in a pipeline perhaps to modify some pipeline config param values or act
        # as a decision tool
        shadow: false
        maintainer:
          -fullname: XXXXXXXX
          email: xxxx@XXX.com
          username: xxxx

        references:
          - name: "name of the reference"
            description: "long or short snippet of a description goes here"
            website: http://www.ncbi.nlm.gov.nl

        inputs:
          - type: string
            displayname: the input directory for the command
            name: input_dir
            description: long or short description about the parameter goes here
            value: /your/dir/location
        scripts:
          - type: js
            before: true
            order: 1
            code: >
              self.input_dir = "/your/dir/location";
          - type: js
            order: 2
            before: true
            code: >
              self.input_dir = "/your/other/dir";
        # this tool has no outputs
        command: ls -ll {{input_dir}}

      - id: 2
        bioflowId: xdir3526
        name: 2
        description: this is a tool that will list all linux directories for a given input directory parameter
        discussions:
          - this tool will list directories
          - this tool will list all linux directories for a given input directory parameter
        website: http://hub.bioflows.io
        version: 1.0.0
        icon: here you can place base64 encoded string value of an icon in png format
        # shadow property indicates that this tool will have no output, it exists in a pipeline perhaps to modify some pipeline config param values or act
        # as a decision tool
        shadow: false
        maintainer:
          -fullname: XXXXXXXXX
          email: xx@xx.com
          username: XXXX

        references:
          - name: "name of the reference"
            description: "long or short snippet of a description goes here"
            website: http://www.ncbi.nlm.gov.nl
        notification:
          to: xx@xx.com
          title: "Step 2 has finished"
          body: "Step 2 has finished"


        inputs:
          - type: string
            displayname: the input directory for the command
            name: input_dir
            description: long or short description about the parameter goes here
            value: /your/dir
        scripts:
          - type: js
            before: true
            order: 2
            code: >
              self.input_dir = "/your/dir/location";
          - type: js
            order: 1
            before: true
            code: >
              self.input_dir = "/your/dir/location";
        # this tool has no outputs
        command: ls -ll {{input_dir}} > myfile.txt

      - id: 3
        bioflowId: xdir3528
        depends: 1,2
        name: 3
        description: this is a tool that will list all linux directories for a given input directory parameter
        discussions:
          - this tool will list directories
          - this tool will list all linux directories for a given input directory parameter
        website: http://hub.bioflows.io
        version: 1.0.0
        icon: here you can place base64 encoded string value of an icon in png format
        # shadow property indicates that this tool will have no output, it exists in a pipeline perhaps to modify some pipeline config param values or act
        # as a decision tool
        shadow: false
        maintainer:
          -fullname: XXXXXXXXX
          email: xx@xx.com
          username: XXXX

        references:
          - name: "name of the reference"
            description: "long or short snippet of a description goes here"
            website: http://www.ncbi.nlm.gov.nl

        inputs:
          - type: string
            displayname: the input directory for the command
            name: input_dir
            description: long or short description about the parameter goes here
            value: /your/dir
        scripts:
          - type: js
            before: true
            order: 1
            code: >
              self.input_dir = "/your/dir/location";
          - type: js
            order: 2
            before: true
            code: >
              self.input_dir = "/your/dir/location";
        # this tool has no outputs
        command: ls -ll {{input_dir}}

      - id: 5
        bioflowId: xdir3529
        depends: 1,3
        name: 5
        description: this is a tool that will list all linux directories for a given input directory parameter
        discussions:
          - this tool will list directories
          - this tool will list all linux directories for a given input directory parameter
        website: http://hub.bioflows.io
        version: 1.0.0
        icon: here you can place base64 encoded string value of an icon in png format
        # shadow property indicates that this tool will have no output, it exists in a pipeline perhaps to modify some pipeline config param values or act
        # as a decision tool
        shadow: false
        maintainer:
          -fullname: XXXXXXXXX
          email: xx@xx.com
          username: XXXX

        references:
          - name: "name of the reference"
            description: "long or short snippet of a description goes here"
            website: http://www.ncbi.nlm.gov.nl

        inputs:
          - type: string
            displayname: the input directory for the command
            name: input_dir
            description: long or short description about the parameter goes here
            value: /your/dir
        scripts:
          - type: js
            before: true
            order: 1
            code: >
              self.input_dir = "/your/dir/location";
          - type: js
            order: 2
            before: true
            code: >
              self.input_dir = "/your/dir/location";
        # this tool has no outputs
        command: ls -ll {{input_dir}}

      - id: 4
        bioflowId: xdir3529
        depends: 3,5
        name: 4
        description: this is a tool that will list all linux directories for a given input directory parameter
        discussions:
          - this tool will list directories
          - this tool will list all linux directories for a given input directory parameter
        website: http://hub.bioflows.io
        version: 1.0.0
        icon: here you can place base64 encoded string value of an icon in png format
        # shadow property indicates that this tool will have no output, it exists in a pipeline perhaps to modify some pipeline config param values or act
        # as a decision tool
        shadow: false
        maintainer:
          -fullname: XXXXXXXXX
          email: xx@xx.com
          username: XXXX

        references:
          - name: "name of the reference"
            description: "long or short snippet of a description goes here"
            website: http://www.ncbi.nlm.gov.nl

        inputs:
          - type: string
            displayname: the input directory for the command
            name: input_dir
            description: long or short description about the parameter goes here
            value: /your/dir
        scripts:
          - type: js
            before: true
            order: 1
            code: >
              self.input_dir = "/your/dir/location";
          - type: js
            order: 2
            before: true
            code: >
              self.input_dir = "/your/dir/location";
        # this tool has no outputs
        command: ls -ll {{input_dir}}