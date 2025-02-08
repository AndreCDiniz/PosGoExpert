package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// TestEvent é uma estrutura de exemplo que implementa a interface EventInterface.
type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

// TestEventHandler é um handler de exemplo que implementa a interface EventHandlerInterface.
type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {}

// EventDispatcherTestSuite é uma estrutura que agrupa os testes do EventDispatcher.
// Ela utiliza a ferramenta suite do testify para organizar e executar os testes.
// Vantagens do uso do suite:
// - Facilita a organização dos testes em grupos coesos.
// - Permite a reutilização de código de setup e teardown.
// - Fornece métodos auxiliares para asserções.
type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

// SetupTest é um método que é executado antes de cada teste.
// Ele inicializa os campos da estrutura de teste.
// Este método é útil para evitar a repetição de código de inicialização em cada teste.
func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "test1", Payload: "test1"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

// TestSuite é a função que inicia a execução da suite de testes.
// A função suite.Run executa todos os métodos de teste definidos na estrutura EventDispatcherTestSuite.
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

// TestEventDispatcher_Register verifica se os handlers são registrados corretamente no EventDispatcher.
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	// Registra o primeiro handler e verifica se foi registrado corretamente
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Registra o segundo handler e verifica se foi registrado corretamente
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Registra o terceiro handler e verifica se foi registrado corretamente
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(3, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Verifica se os handlers registrados são os esperados
	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
	assert.Equal(suite.T(), &suite.handler3, suite.eventDispatcher.handlers[suite.event.GetName()][2])
}

// TestEventDispatcher_Register_WithSameHandler verifica se o mesmo handler não pode ser registrado duas vezes para o mesmo evento.
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	// Registra o handler e verifica se foi registrado corretamente
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Tenta registrar o mesmo handler novamente e verifica se o erro esperado é retornado
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

// TestEventDispatcher_Clear verifica se todos os handlers são removidos corretamente.
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Registra handlers para dois eventos diferentes
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	// Limpa todos os handlers e verifica se foram removidos corretamente
	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))
}

// TestEventDispatcher_Has verifica se o EventDispatcher possui handlers registrados para um evento específico.
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	// Registra handlers para um evento
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Verifica se os handlers registrados são os esperados
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

// MockHandler é um mock para o EventHandlerInterface, usado para verificar se o método Handle foi chamado corretamente.
type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

// TestEventDispatcher_Dispatcher verifica se os handlers são chamados corretamente quando um evento é despachado.
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatcher() {
	// Cria um mock handler e registra-o para um evento
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", &suite.event)
	suite.eventDispatcher.Register(suite.event.GetName(), eventHandler)

	// Despacha o evento e verifica se o handler foi chamado corretamente
	suite.eventDispatcher.Dispatch(&suite.event)
	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}
