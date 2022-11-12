import { Attribute, Event, Log } from "@cosmjs/stargate/build/logs"

export type RoadOperatorCreatedEvent = Event

export const getCreateRoadOperatorEvent = (
    log: Log,
): RoadOperatorCreatedEvent | undefined => {
    return log.events!.find(
        (event: Event) => event.type === "new-road-operator-created",
    )
}

export const getCreatedRoadOperatorId = (
    createdRoadOperatorEvent: RoadOperatorCreatedEvent,
): string => {
    return createdRoadOperatorEvent.attributes.find(
        (attribute: Attribute) => attribute.key === "road-operator-index",
    ).value
}
