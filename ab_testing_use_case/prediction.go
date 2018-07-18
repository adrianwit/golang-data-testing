package viant



type Prediction struct {
	Value1  float64
	Value2  float64
	ValueN  float64
	Static1 string
	Static2 string
	StaticN string
}







type Service interface {
	Predict(request *PredicationRequest) *PredicationResponse
}

type PredicationRequest struct {
	Input1 interface{}
	Input2 interface{}
	InputN interface{}
}

type PredicationResponse struct {
	*Prediction
	ExperimentID int
}











type service struct{}

func (s *service) Predict(request *PredicationRequest) *PredicationResponse {
	var response = &PredicationResponse{
		ExperimentID: randInt(1, 2),
		Prediction: &Prediction{
			Static1:"value1",
			Static2:"value2",
			StaticN:"valueN",
		},
	}
	switch response.ExperimentID {
	case 1:
		response.Value1 = randFloat(0.11, 0.44)
		response.Value2 = randFloat(0.43, 0.84)
		response.ValueN = randFloat(0.01, 0.3)

	case 2:
		response.Value1 = randFloat(0.22, 0.5)
		response.Value2 = randFloat(0.44, 0.70)
		response.ValueN = randFloat(0.2, 0.4)
	}
	return response
}


func New() Service{
	return &service{}
}