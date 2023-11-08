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
 * @interface V1OptionOverrides
 */
export interface V1OptionOverrides {
    /**
     * 
     * @type {Array<V1OptionOverride>}
     * @memberof V1OptionOverrides
     */
    optionOverrides?: Array<V1OptionOverride>;
}

/**
 * Check if a given object implements the V1OptionOverrides interface.
 */
export function instanceOfV1OptionOverrides(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1OptionOverridesFromJSON(json: any): V1OptionOverrides {
    return V1OptionOverridesFromJSONTyped(json, false);
}

export function V1OptionOverridesFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1OptionOverrides {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'optionOverrides': !exists(json, 'optionOverrides') ? undefined : ((json['optionOverrides'] as Array<any>).map(V1OptionOverrideFromJSON)),
    };
}

export function V1OptionOverridesToJSON(value?: V1OptionOverrides | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'optionOverrides': value.optionOverrides === undefined ? undefined : ((value.optionOverrides as Array<any>).map(V1OptionOverrideToJSON)),
    };
}
