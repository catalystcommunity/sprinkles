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
/**
 * 
 * @export
 * @interface V1DeleteResponse
 */
export interface V1DeleteResponse {
    /**
     * 
     * @type {Array<string>}
     * @memberof V1DeleteResponse
     */
    errors?: Array<string>;
}

/**
 * Check if a given object implements the V1DeleteResponse interface.
 */
export function instanceOfV1DeleteResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1DeleteResponseFromJSON(json: any): V1DeleteResponse {
    return V1DeleteResponseFromJSONTyped(json, false);
}

export function V1DeleteResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1DeleteResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'errors': !exists(json, 'errors') ? undefined : json['errors'],
    };
}

export function V1DeleteResponseToJSON(value?: V1DeleteResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'errors': value.errors,
    };
}

