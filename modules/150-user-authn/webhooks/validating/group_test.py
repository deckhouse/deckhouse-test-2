#!/usr/bin/env python3

# Copyright 2024 Flant JSC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import unittest
import json
import typing

from group import main
from deckhouse import hook, validations
from dotmap import DotMap

def _assert_validation(t: unittest.TestCase, v: validations.ValidationsCollector, allowed: bool, msg: typing.Tuple[str, ...] | str | None):
    t.assertEqual(len(v.data), 1)
    a = t.assertFalse
    if allowed:
        a = t.assertTrue
    a(v.data[0]["allowed"])
    if not msg is None:
        if isinstance(msg, str):
            t.assertEqual(len(v.data[0]["warnings"]), 1)
            t.assertEqual(v.data[0]["warnings"][0], msg)
        elif isinstance(msg, tuple):
            t.assertEqual(v.data[0]["warnings"], msg)
        else:
            t.fail("Incorrect msg type")


def assert_validation_allowed(t: unittest.TestCase, v: validations.ValidationsCollector, msg: typing.Tuple[str, ...] | str | None ):
    _assert_validation(t, v, True, msg)


def _prepare_update_binding_context(new_spec: dict) -> DotMap:
    binding_context_json = """
{
    "binding": "groups-unique.deckhouse.io",
    "review": {
        "request": {
            "uid": "8af60184-b30b-4b90-a33e-0c190f10e96d",
            "kind": {
                "group": "deckhouse.io",
                "version": "v1alpha1",
                "kind": "Group"
            },
            "resource": {
                "group": "deckhouse.io",
                "version": "v1alpha1",
                "resource": "groups"
            },
            "requestKind": {
                "group": "deckhouse.io",
                "version": "v1alpha1",
                "kind": "Group"
            },
            "requestResource": {
                "group": "deckhouse.io",
                "version": "v1alpha1",
                "resource": "groups"
            },
            "name": "candi-admins",
            "operation": "UPDATE",
            "userInfo": {
                "username": "kubernetes-admin",
                "groups": [
                    "system:masters",
                    "system:authenticated"
                ]
            },
            "object": {
                "apiVersion": "deckhouse.io/v1alpha1",
                "kind": "Group",
                "metadata": {
                    "creationTimestamp": "2023-07-17T13:40:39Z",
                    "generation": 3,
                    "managedFields": [
                        {
                            "apiVersion": "deckhouse.io/v1alpha1",
                            "fieldsType": "FieldsV1",
                            "fieldsV1": {
                                "f:spec": {
                                    ".": {},
                                    "f:name": {}
                                }
                            },
                            "manager": "deckhouse-controller",
                            "operation": "Update",
                            "time": "2023-07-17T13:40:39Z"
                        },
                        {
                            "apiVersion": "deckhouse.io/v1alpha1",
                            "fieldsType": "FieldsV1",
                            "fieldsV1": {
                                "f:spec": {
                                    "f:members": {}
                                }
                            },
                            "manager": "kubectl-edit",
                            "operation": "Update",
                            "time": "2024-11-21T14:44:35Z"
                        }
                    ],
                    "name": "candi-admins",
                    "resourceVersion": "1184522270",
                    "uid": "7820c68b-6423-49f0-b97f-b7e314e23c0b"
                },
                "spec": {
                    "members": [
                        {
                            "kind": "User",
                            "name": "superadmin"
                        },
                        {
                            "kind": "Group",
                            "name": "none-exists-2"
                        }
                    ],
                    "name": "candi-admins"
                }
            },
            "oldObject": {
                "apiVersion": "deckhouse.io/v1alpha1",
                "kind": "Group",
                "metadata": {
                    "creationTimestamp": "2023-07-17T13:40:39Z",
                    "generation": 2,
                    "managedFields": [
                        {
                            "apiVersion": "deckhouse.io/v1alpha1",
                            "fieldsType": "FieldsV1",
                            "fieldsV1": {
                                "f:spec": {
                                    ".": {},
                                    "f:name": {}
                                }
                            },
                            "manager": "deckhouse-controller",
                            "operation": "Update",
                            "time": "2023-07-17T13:40:39Z"
                        },
                        {
                            "apiVersion": "deckhouse.io/v1alpha1",
                            "fieldsType": "FieldsV1",
                            "fieldsV1": {
                                "f:spec": {
                                    "f:members": {}
                                }
                            },
                            "manager": "kubectl-edit",
                            "operation": "Update",
                            "time": "2024-11-20T14:00:21Z"
                        }
                    ],
                    "name": "candi-admins",
                    "resourceVersion": "1184522270",
                    "uid": "7820c68b-6423-49f0-b97f-b7e314e23c0b"
                },
                "spec": {
                    "members": [
                        {
                            "kind": "User",
                            "name": "superadmin"
                        },
                        {
                            "kind": "Group",
                            "name": "none-exists"
                        }
                    ],
                    "name": "candi-admins"
                }
            },
            "dryRun": false,
            "options": {
                "kind": "UpdateOptions",
                "apiVersion": "meta.k8s.io/v1",
                "fieldManager": "kubectl-edit",
                "fieldValidation": "Strict"
            }
        }
    },
    "snapshots": {
        "groups": [
            {
                "filterResult": {
                    "groupName": "candi-admins",
                    "members": [
                        {
                            "kind": "User",
                            "name": "superadmin"
                        },
                        {
                            "kind": "Group",
                            "name": "none-exists"
                        }
                    ],
                    "name": "candi-admins"
                }
            },
            {
                "filterResult": {
                    "groupName": "crowd-supplier-calendar-ro",
                    "members": [
                        {
                            "kind": "User",
                            "name": "test"
                        }
                    ],
                    "name": "crowd-supplier-calendar-ro"
                }
            }
        ],
        "users": [
            {
                "filterResult": {
                    "userName": "superadmin"
                }
            },
            {
                "filterResult": {
                    "userName": "test"
                }
            }
        ]
    },
    "type": "Validating"
}
"""
    ctx_dict = json.loads(binding_context_json)
    ctx = DotMap(ctx_dict)
    ctx.review.request.object.spec = new_spec
    return ctx

class TestGroupValidationWebhook(unittest.TestCase):
    def test_update_group_with_new_group_member_none_exists_group(self):
        ctx = _prepare_update_binding_context({
            "members": [
                {
                    "kind": "User",
                    "name": "superadmin"
                },
                {
                    "kind": "Group",
                    "name": "none-exists-2"
                }
            ],
            "name": "candi-admins"
        })
        out = hook.testrun(main, [ctx])
        assert_validation_allowed(self, out.validations, 'groups.deckhouse.io "none-exists-2" not exist')

    def test_update_group_with_new_group_member_none_exists_user(self):
        ctx = _prepare_update_binding_context({
            "members": [
                {
                    "kind": "User",
                    "name": "superadmin"
                },
                {
                    "kind": "User",
                    "name": "not-exists"
                }
            ],
            "name": "candi-admins"
        })
        out = hook.testrun(main, [ctx])
        assert_validation_allowed(self, out.validations, 'users.deckhouse.io "not-exists" not exist')

    def test_update_group_with_new_group_member_none_exists_user_and_group(self):
        ctx = _prepare_update_binding_context({
            "members": [
                {
                    "kind": "User",
                    "name": "superadmin"
                },
                {
                    "kind": "Group",
                    "name": "none-exists-2"
                },
                {
                    "kind": "User",
                    "name": "not-exists"
                }
            ],
            "name": "candi-admins"
        })
        out = hook.testrun(main, [ctx])
        assert_validation_allowed(self, out.validations, (
            'groups.deckhouse.io "none-exists-2" not exist',
            'users.deckhouse.io "not-exists" not exist'
        ))

if __name__ == '__main__':
    unittest.main()
