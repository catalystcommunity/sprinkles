/* tslint:disable */
/* eslint-disable */
/**
 * sprinkles/v1/api.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1OptionOverride } from './V1OptionOverride';
import {
    V1OptionOverrideFromJSON,
    V1OptionOverrideFromJSONTyped,
    V1OptionOverrideToJSON,
} from './V1OptionOverride';

/**
 * 
 * @export
 * @interface V1UpsertOptionOverridesRequest
 */
export interface V1UpsertOptionOverridesRequest {
    /**
     * 
     * @type {Array<V1OptionOverride>}
     * @memberof V1UpsertOptionOverridesRequest
     */
    optionOverrides?: Array<V1OptionOverride>;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1UpsertOptionOverridesRequest
     */
    fields?: Array<string>;
}

/**
 * Check if a given object implements the V1UpsertOptionOverridesRequest interface.
 */
export function instanceOfV1UpsertOptionOverridesRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1UpsertOptionOverridesRequestFromJSON(json: any): V1UpsertOptionOverridesRequest {
    return V1UpsertOptionOverridesRequestFromJSONTyped(json, false);
}

export function V1UpsertOptionOverridesRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1UpsertOptionOverridesRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'optionOverrides': !exists(json, 'optionOverrides') ? undefined : ((json['optionOverrides'] as Array<any>).map(V1OptionOverrideFromJSON)),
        'fields': !exists(json, 'fields') ? undefined : json['fields'],
    };
}

export function V1UpsertOptionOverridesRequestToJSON(value?: V1UpsertOptionOverridesRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'optionOverrides': value.optionOverrides === undefined ? undefined : ((value.optionOverrides as Array<any>).map(V1OptionOverrideToJSON)),
        'fields': value.fields,
    };
}

