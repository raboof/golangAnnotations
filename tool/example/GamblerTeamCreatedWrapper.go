
// Generated automatically: do not edit manually

package example

import (
  "encoding/json"
  "log"
  "time"

  "code.google.com/p/go-uuid/uuid"
)

func (s *GamblerTeamCreated) Wrap() (*Envelope,error) {
    var err error
    envelope := new(Envelope)
    envelope.Uuid = uuid.New()
    envelope.SequenceNumber = 0 // Set later by event-store
    envelope.Timestamp = time.Now()
    envelope.AggregateName = "gambler"
    envelope.AggregateUid = s.GetUid()
    envelope.EventTypeName = "GamblerTeamCreated"
    blob, err := json.Marshal(s)
    if err != nil {
        log.Printf("Error marshalling GamblerTeamCreated payload %+v", err)
        return nil, err
    }
    envelope.EventData = string(blob)

    return envelope, nil
}

func IsGamblerTeamCreated(envelope *Envelope) bool {
    return envelope.EventTypeName == "GamblerTeamCreated"
}

func GetIfIsGamblerTeamCreated(envelop *Envelope) (*GamblerTeamCreated, bool) {
    if IsGamblerTeamCreated(envelop) == false {
        return nil, false
    }
    event := UnWrapGamblerTeamCreated(envelop)
    return event, true
}

func UnWrapGamblerTeamCreated(envelop *Envelope) (*GamblerTeamCreated,error) {
    if IsGamblerTeamCreated(envelop) == false {
        return nil
    }
    var event GamblerTeamCreated
    err := json.Unmarshal([]byte(envelop.EventData), &event)
    if err != nil {
        log.Printf("Error unmarshalling GamblerTeamCreated payload %+v", err)
        return nil, err
    }

    return &event, nil
}
