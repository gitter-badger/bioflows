package main

func main(){
	//pipeline := &pipelines.BioPipeline{}
	//tool_in, err := os.Open("/home/snouto/projects/bioflows/scripts/pipeline.yaml")
	//
	//if err != nil {
	//	fmt.Printf("There was an error opening the tool file, %v\n",err)
	//	os.Exit(1)
	//}
	//mytool_content, err := ioutil.ReadAll(tool_in)
	//if err != nil {
	//	fmt.Printf("Error reading the contents of the tool , %v\n",err)
	//	os.Exit(1)
	//}
	//err = yaml.Unmarshal([]byte(mytool_content),pipeline)
	//if err != nil {
	//	//fmt.Println("There was a problem unmarshaling the current tool")
	//	fmt.Println(err.Error())
	//	return
	//}
	//b , err := pipelines.PreparePipeline(pipeline)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//g , err := pipelines.CreateGraph(b)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//gg, _ , err := pipelines.ToDotGraph(b,g)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Print(gg)
	//parents := g.SourceVertices()
	//for _, v := range parents {
	//	fmt.Println("Executing: ",v.ID)
	//	if v.Children.Size() > 0 {
	//		for _ , c := range v.Children.Values(){
	//			cv := c.(*dag.Vertex)
	//			fmt.Println("===Executing : ",cv.ID)
	//		}
	//	}
	//}

	//fmt.Println(pipeline.ToJson())
}

