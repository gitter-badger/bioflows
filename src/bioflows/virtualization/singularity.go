package virtualization

import "bytes"

type SingularityVirtualizationManager struct {

}

func (s *SingularityVirtualizationManager) ListContainers() []BioflowContainer{
	return nil
}

func (s *SingularityVirtualizationManager)  StartContainer(containerName string , imageName string,commands []string) (*bytes.Buffer,*bytes.Buffer,error) {
	return nil,nil,nil
}
func (s *SingularityVirtualizationManager) StopContainer (containerID string) error{

	return nil
}

func (s *SingularityVirtualizationManager) PullImage(imageURL string) (*bytes.Buffer , error){
	return nil , nil
}
