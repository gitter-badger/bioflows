Writing New Pipelines and Tools
###############################

In this section, we will explore how to compose new BioFlows individual tools and pipelines using the directives we learnt about
in the previous section of this document.

Bioflows praises the concept of sharing computational analyses over the internet. Subsequently, Computational analyses in BioFlows
could be constructed either as a single individual tool or a complete pipeline with various steps. The author of Bioflows tools or
pipelines could share his/her tool or pipeline online through various means including but not limited to ``GitHub``, ``Bitbucket``,
publicly available HTTP server(s) or through BioFlows Hub Platform (**coming soon**).

Each single tool should perform a single function or a compound function through utilizing linux shell capabilities (i.e: pipes).
On contrary, a single pipeline is composed of one or more computational steps that may or may not depend on each others. Each step
in a pipeline could be a single tool or a whole nested pipeline. In general, all scientific computing pipelines and bioinformatics
pipelines in particular could be efficiently represented as a Directed Acyclic Graph (``DAG``). **Directed** means that the flow of
analysis has direction, it starts from one or more vertices and moves along the graph till the end. While, **Acyclic** means that no
vertex in the graph has any self-reference back to itself.

In this section, we will explain how to author new tools and pipelines using simple well-known linux commands just to allow beginners
to practice the art of writing new tools and focus more on the roles of each directive. In the next Practical Tutorials section,
we will explore writing real-world bioinformatics tools and pipelines and we will look at how to set up an environment and execute them
in more greater details.


Tool/Pipeline Definition
************************

Let's build together a tool that would list directories on a linux server....

.. code-block:: yaml

    id: listDir
    type: tool
    name: list_directories
    description: "this tool will list directories on a linux server"


You start by giving your tool an ``ID`` which should not contain any spaces and it should be unique in the whole pipeline. While
``type`` can have a value of either a "tool" or a "pipeline" depending upon whether the following definition file represents a tool
or a whole pipeline. ``name`` represents the name of the current tool, you should replace spaces by underscores, because the name of
the tool or a pipeline will be the name of its output directory. The best practices is to use a short name and replace spaces with underscores.

.. note::
    ``name`` directive in a tool or a pipeline should be short and have underscores instead of spaces..

The description is an optional field, but it is better if you can provide a textual description of your tool and the purpose of it.

Adding Command section
**********************

Now, we need to add the most important directive in a tool, which is the command. This tool is going to list directories in a linux server.

so we should write a shell command to list directories. Let's do it.

.. code-block:: yaml

    id: listDir
    type: tool
    name: list_directories
    description: "this tool will list directories on a linux server"
    command: ls -ll <Your directory location>

Now, we have added the command directive with a shell command to list directories on a linux server, ``-ll`` switch indicates that
we need to list directories using the long list format instead of the common shorter format. Afterwards, we specified the location
of the directory we need to list... Right !!!!

Apparently, hard-coded file system will only work with your local linux box, but it won't run with other people because definitely
they have a different file system layout. Subsequently, we need to externalize the file system location as an input parameter to our
tool to make it reusable across different environments. Let's do this now....

.. code-block:: yaml

    id: listDir
    type: tool
    name: list_directories
    description: "this tool will list directories on a linux server"
    inputs:
        - type: dir
          name: input_dir
          displayname: Input directory
          description: "Input directory to list its contents"
          value: /your/directory/location

    command: ls -ll {{input_dir}}


What we have done here, is that we have moved the file system location given to ``ls`` linux tool as an input parameter and used
a placeholder variable to dynamically mention the value of that parameter in the command using ``Mustache`` templating expression.

.. note::
    BioFlows fully supports Internal Templating engine called Mustache which facilitates dynamic placeholders for common parameters and variables in the file definition file.

Now any user with Bioflows can run the tool giving ``input_dir`` as an input to ``bf`` executable program and bioflows will override
the default value given in ``value`` sub-directive in the definition of the parameter.

When this tool run, there will be a folder named after the tool name concatenated with the id of the tool in the output directory of
this run, with a log file containing the output of the ``ls`` linux tool. because ``ls`` outputs its contents to the standard output
which is caught automatically by ``bf`` executable and written into a file with ``.logs`` extension in the output directory of this tool.

At this point, we have a tool that list directories of a given file system and outputs the contents of this directory to the standard output
 we also made the tool reusable by externalizing the ``input_dir`` parameter so that other users can take advantage of this by passing the parameter to the pipeline during execution.

But still our tool is of limited use, because the contents of the input directory has been written to the overall tool standard output
which may or may not contain other textual output data from the tool itself. This prevents us from performing any further downstream analysis
on the output of ``ls``. In real world cases, this data might be a structured or semi-structured formatted data that we need to further work on it
so mixing it with the tool outputs will hinder any downstream parsing or further processing on it....

So a better strategy would be to direct the output of ``ls`` tool to an output parameter and save it as a file in the same directory of the running tool
So how can we do this ? Let's see...

.. code-block:: yaml

    id: listDir
    type: tool
    name: list_directories
    description: "this tool will list directories on a linux server"
    outputs:
        - type: file
          name: output_file
          value: "{{self_dir}}/ls_output.txt"

    inputs:
        - type: dir
          name: input_dir
          displayname: Input directory
          description: "Input directory to list its contents"
          value: /your/directory/location

    command: "ls -ll {{input_dir}} > {{output_file}}"

We have defined an output parameter with a type of ``file`` named ``ls_output.txt``. please note that ``{{self_dir}}`` is an implicit variable
given to you by bioflows which contains the fully qualified path of the current tool output directory. For more information about
all other implicit variables, please take a look at **Implicit Variables** section of this documentation.

We have also directed the outputs of the tool to another file called ``output_file``

.. note::
    Please note that it is strictly recommended to define your tool output parameters of type ``file``, ``dir`` with fully qualified paths in order to allow these parameters to be referenced directly in downstream dependent steps without referencing that tool output directory with the output file name or directory each time


Now your tool looks really great, it is reusable and can run anywhere with ``bf`` executable. Now let's add some metadata about the author
of this tool and a website where other people could visit who are interested to read more about you or your research.

.. code-block:: yaml

    id: listDir
    type: tool
    name: list_directories
    description: "this tool will list directories on a linux server"
    discussions:
      - this tool will list directories
      - this tool will list all linux directories for a given input directory parameter
    website: http://john.university.com
    version: 1.0.0
    maintainer:
        - fullname: Dr. John Doe
          email: john.doe@university.com
    outputs:
        - type: file
          name: output_file
          value: "{{self_dir}}/ls_output.txt"
    inputs:
        - type: dir
          name: input_dir
          displayname: Input directory
          description: "Input directory to list its contents"
          value: /your/directory/location

    command: "ls -ll {{input_dir}} > {{output_file}}"


Share Your Tool
***************

Now, you can share and reuse this tool with other researchers in your field, or simply, you can create a GitHub or BitBucket account
and put your tool definition file in there so others can use it and mention you and your research in their publications.

Reuse Your Tool
***************

Now assume that you or other researchers want to use your previously published tool and incorporate it into their pipeline.

Let's do just this....

.. code-block:: yaml

    id: countwords
    name: countwords
    type: pipeline
    description: this pipeline will list the contents of a specific directory and save that to a file and count the lines in this file.
    steps:
        - id: listDir
          url: https://raw.githubusercontent.com/mfawzysami/bioflows/master/scripts/tool.yaml

        - id: countstep
          name: countstep
          depends: listDir
          command: "wc -l {{self.listDir.output_file}}"


As you can see, when we wanted to mention that tool, we only used the ``raw`` url of GitHub of this tool and we didn't define anything
else because the current tool in our pipeline will inherit all the directives from the remote tool.

Furthermore, we defined a new step called ``countstep`` which basically ``depends`` on ``listDir`` step
and we have mentioned the listDir's Output file using ``{{self.listDir.output_file}}`` in one shot. because we have created ``output_file``
as a fully qualified file path, but instead if the ``output_file`` was only a file name without a full directory path, we could have mentioned it differently like this

.. code-block:: yaml

    command: "wc -l {{self.listDir.location}}/{{self.listDir.output_file}}"


which is a bit verbose and error prone especially in more complex pipeline definition....

.. note::
    Please note that, **depends** is one of the most important directives which without it will corrupt the directed acyclic graph and the order of other tools in the graph which might have unfavorable processing consequences. So please make sure, to mention **depends** in a dependent step mentioning the ``ID`` of the parent tool.
    If your current step depends on multiple previous steps, you can mention them as comma separated list of IDs.





















