package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// EventDispatcher é uma estrutura que gerencia handlers de eventos.
// Ela mantém um mapa onde a chave é o nome do evento (string) e o valor é uma lista de handlers (EventHandlerInterface) associados a esse evento.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

// NewEventDispatcher cria uma nova instância de EventDispatcher.
// A função inicializa o campo handlers usando a função make para criar um mapa vazio.
// O make é usado para alocar e inicializar objetos do tipo slice, map ou chan. Neste caso, ele está criando um mapa vazio onde as chaves são strings
// e os valores são slices de EventHandlerInterface.
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch despacha um evento para todos os handlers registrados para esse evento.
// O método percorre todos os handlers associados ao nome do evento e chama o método Handle de cada handler.
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}

// Register adiciona um handler para um evento específico.
// O método tem um receiver (ed *EventDispatcher), que é uma referência à instância de EventDispatcher que está chamando o método.
// Isso permite que o método acesse e modifique os campos da instância.
func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// Verifica se o evento já está registrado
	if _, ok := ed.handlers[eventName]; ok {
		// Verifica se o handler já está registrado para o evento
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	// Adiciona o handler ao evento
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

// Clear remove todos os handlers registrados.
// O método recria o mapa de handlers, efetivamente removendo todos os handlers registrados anteriormente.
func (ed *EventDispatcher) Clear() error {
	ed.handlers = make(map[string][]EventHandlerInterface)
	return nil
}

// Has verifica se um handler específico está registrado para um evento.
// O método percorre todos os handlers associados ao nome do evento e verifica se o handler está presente.
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}
